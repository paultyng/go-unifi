package main

import (
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-version"
)

var uiDownloadUrl = "https://www.ui.com/download/?platform=unifi"

func latestUnifiVersion() (*version.Version, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uiDownloadUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Downloads []struct {
			Id           int    `json:"id"`
			CategorySlug string `json:"category__slug"`
			Filename     string `json:"filename"`
			Version      string `json:"version"`
		} `json:"downloads"`
	}
	var latestVersion *version.Version

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	for _, download := range respData.Downloads {
		if download.CategorySlug != "software" {
			continue
		}

		if download.Filename != "unifi_sysvinit_all.deb" {
			continue
		}

		downloadVersion, err := version.NewVersion(download.Version)
		if err != nil {
			// Skip this entry if the version isn't valid.
			continue
		}

		if latestVersion == nil || downloadVersion.GreaterThan(latestVersion) {
			latestVersion = downloadVersion
		}
	}

	return latestVersion, nil
}
