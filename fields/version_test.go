package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertLatestVersionUsingProvider(t *testing.T, provider func() (*UnifiVersion, error)) {
	t.Helper()
	assert := assert.New(t)
	require := require.New(t)

	fwVersion, err := version.NewVersion("7.3.83+atag-7.3.83-19645")
	require.NoError(err)

	fwDownload, err := url.Parse("https://fw-download.ubnt.com/data/unifi-controller/c31c-debian-7.3.83-c9249c913b91416693b869b9548850c3.deb")
	require.NoError(err)

	respData := firmwareUpdateApiResponse{
		Embedded: firmwareUpdateApiResponseEmbedded{
			Firmware: []firmwareUpdateApiResponseEmbeddedFirmware{
				{
					Channel:  releaseChannel,
					Created:  "2023-02-06T08:55:31+00:00",
					Id:       "c9249c91-3b91-4166-93b8-69b9548850c3",
					Platform: debianPlatform,
					Product:  unifiControllerProduct,
					Version:  fwVersion,
					Links: firmwareUpdateApiResponseEmbeddedFirmwareLinks{
						Data: firmwareUpdateApiResponseEmbeddedFirmwareDataLink{
							Href: fwDownload,
						},
					},
				},
				{
					Channel:  releaseChannel,
					Created:  "2023-02-06T08:51:36+00:00",
					Id:       "2a600108-7f79-4b3e-b6e0-4dd262460457",
					Platform: "document",
					Product:  unifiControllerProduct,
					Version:  fwVersion,
					Links: firmwareUpdateApiResponseEmbeddedFirmwareLinks{
						Data: firmwareUpdateApiResponseEmbeddedFirmwareDataLink{
							Href: nil,
						},
					},
				},
				{
					Channel:  releaseChannel,
					Created:  "2023-02-06T08:51:37+00:00",
					Id:       "9d2d413d-36ce-4742-a10d-4351aac6f08d",
					Platform: "windows",
					Product:  unifiControllerProduct,
					Version:  fwVersion,
					Links: firmwareUpdateApiResponseEmbeddedFirmwareLinks{
						Data: firmwareUpdateApiResponseEmbeddedFirmwareDataLink{
							Href: nil,
						},
					},
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		assert.Contains(query["filter"], firmwareUpdateApiFilter("channel", releaseChannel))
		assert.Contains(query["filter"], firmwareUpdateApiFilter("product", unifiControllerProduct))

		resp, err := json.Marshal(respData)
		assert.NoError(err)

		_, err = rw.Write(resp)
		assert.NoError(err)
	}))
	defer server.Close()

	firmwareUpdateApi = server.URL
	gotVersion, err := provider()
	require.NoError(err)

	assert.Equal(fwVersion.Core(), gotVersion.Version)
	assert.Equal(fwDownload, gotVersion.DownloadUrl)
}

func TestLatestUnifiVersion(t *testing.T) {
	t.Parallel()
	assertLatestVersionUsingProvider(t, func() (*UnifiVersion, error) {
		return latestUnifiVersion()
	})
}

func TestDetermineUnifiVersion_latest(t *testing.T) {
	t.Parallel()
	assertLatestVersionUsingProvider(t, func() (*UnifiVersion, error) {
		return determineUnifiVersion(LatestVersionMarker)
	})
}

func TestDetermineUnifiVersion_provided(t *testing.T) {
	t.Parallel()
	testCases := map[string]string{
		"7.3.83+atag-7.3.83-19645": "7.3.83",
		"7.3.83":                   "7.3.83",
		"7.3":                      "7.3.0",
		"7":                        "7.0.0",
	}

	for providedVersion, expectedVersion := range testCases {
		t.Run(providedVersion, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)
			require := require.New(t)

			unifiVersion, err := determineUnifiVersion(providedVersion)
			require.NoError(err)

			assert.Equal(expectedVersion, unifiVersion.Version.String())
			assert.Equal(fmt.Sprintf(baseDownloadUrl, expectedVersion), unifiVersion.DownloadUrl.String())
		})
	}
}

func TestDetermineUnifiVersion_invalid(t *testing.T) {
	t.Parallel()
	testCases := []string{
		"invalid",
		"-1",
		"",
	}
	assert := assert.New(t)

	for _, providedVersion := range testCases {
		t.Run(providedVersion, func(t *testing.T) {
			t.Parallel()
			_, err := determineUnifiVersion(providedVersion)
			assert.ErrorContains(err, providedVersion)
		})
	}
}
