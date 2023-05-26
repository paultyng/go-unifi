package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-version"
)

func latestUnifiVersion() (*version.Version, error) {
	url, err := url.Parse(firmwareUpdateApi)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	query.Add("filter", firmwareUpdateApiFilter("channel", releaseChannel))
	query.Add("filter", firmwareUpdateApiFilter("product", unifiControllerProduct))
	url.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData firmwareUpdateApiResponse
	var latestVersion *version.Version

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	for _, firmware := range respData.Embedded.Firmware {
		if firmware.Platform != debianPlatform {
			continue
		}

		latestVersion, err = version.NewVersion(firmware.Version)
		if err != nil {
			// Skip this entry if the version isn't valid.
			continue
		}

		break
	}

	return latestVersion, nil
}
