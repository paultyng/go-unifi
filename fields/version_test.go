package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLatestUnifiVersion(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	fwVersion := "v7.3.83+atag-7.3.83-19645"
	fwDownload := "https://fw-download.ubnt.com/data/unifi-controller/c31c-debian-7.3.83-c9249c913b91416693b869b9548850c3.deb"

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
							Href: "https://fw-download.ubnt.com/data/unifi-controller/edf8-document-7.3.83-2a6001087f794b3eb6e04dd262460457",
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
							Href: "https://fw-download.ubnt.com/data/unifi-controller/c0ce-windows-7.3.83-9d2d413d36ce4742a10d4351aac6f08d.exe",
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
		require.NoError(err)

		_, err = rw.Write(resp)
		require.NoError(err)
	}))
	defer server.Close()

	expected, err := version.NewVersion(fwVersion)
	require.NoError(err)

	firmwareUpdateApi = server.URL
	actual, err := latestUnifiVersion()
	require.NoError(err)

	assert.Equal(expected, actual)
}
