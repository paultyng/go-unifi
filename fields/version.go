package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-version"
)

func latestUnifiVersion() (*version.Version, *url.URL, error) {
	url, err := url.Parse(firmwareUpdateApi)
	if err != nil {
		return nil, nil, err
	}

	query := url.Query()
	query.Add("filter", firmwareUpdateApiFilter("channel", releaseChannel))
	query.Add("filter", firmwareUpdateApiFilter("product", unifiControllerProduct))
	url.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var respData firmwareUpdateApiResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, nil, err
	}

	for _, firmware := range respData.Embedded.Firmware {
		if firmware.Platform != debianPlatform {
			continue
		}

		return firmware.Version.Core(), firmware.Links.Data.Href, nil
	}

	return nil, nil, nil
}
