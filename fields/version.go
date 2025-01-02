package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-version"
	"net/http"
	"net/url"
)

const (
	LatestVersionMarker = "latest"
	baseDownloadUrl     = "https://dl.ui.com/unifi/%s/unifi_sysvinit_all.deb"
)

type UnifiVersion struct {
	Version     *version.Version
	DownloadUrl *url.URL
}

func NewUnifiVersion(unifiVersion *version.Version, downloadUrl *url.URL) *UnifiVersion {
	return &UnifiVersion{
		Version:     unifiVersion,
		DownloadUrl: downloadUrl,
	}
}

func latestUnifiVersion() (*UnifiVersion, error) {
	url, err := url.Parse(firmwareUpdateApi)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	query.Add("filter", firmwareUpdateApiFilter("channel", releaseChannel))
	query.Add("filter", firmwareUpdateApiFilter("product", unifiControllerProduct))
	url.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url.String(), nil)
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
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	for _, firmware := range respData.Embedded.Firmware {
		if firmware.Platform != debianPlatform {
			continue
		}
		return NewUnifiVersion(firmware.Version.Core(), firmware.Links.Data.Href), nil
	}

	return nil, nil
}

func determineUnifiVersion(versionMarker string) (*UnifiVersion, error) {
	if versionMarker == LatestVersionMarker {
		return latestUnifiVersion()
	} else {
		unifiVersion, err := version.NewVersion(versionMarker)
		if err != nil {
			return nil, err
		}
		downloadUrl := fmt.Sprintf(baseDownloadUrl, unifiVersion)
		unifiDownloadUrl, err := url.Parse(downloadUrl)
		if err != nil {
			return nil, err
		}
		return NewUnifiVersion(unifiVersion, unifiDownloadUrl), nil
	}
}
