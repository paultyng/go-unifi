package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	version "github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func TestLatestUnifiVersion(t *testing.T) {
	respData := map[string]interface{}{
		"downloads": []map[string]interface{}{
			{
				"architecture":     "",
				"build":            "",
				"category__name":   "User Guides",
				"category__slug":   "user-guides",
				"changelog":        "",
				"date_published":   "2018-03-05",
				"description":      "",
				"featured":         true,
				"file_path":        "/downloads/guides/UniFi/UniFi_Controller_V5_UG.pdf",
				"filename":         "UniFi_Controller_V5_UG.pdf",
				"id":               898,
				"mib":              "",
				"name":             "UniFi® Controller v5 User Guide",
				"products":         "UAP|UAP-AC-EDU|UAP–AC–IW|UAP–AC–IW–PRO|UAP-AC-LITE|UAP-AC-LR|UAP-AC-M|UAP-AC-M-PRO|UAP-AC-PRO|UAP-HD|UAP-IW|UAP-LR|UAP-nanoHD|UAP-PRO|UAP-SHD|UAP‑XG|UAS-XG|UC-CK|US-16-150W|US‑16‑XG|US-24|US-24-250W|US-24-500W|US-48|US-48-500W|US-48-750W|US-8|US-8-150W|US-8-60W|USG|USG-PRO-4|USG-XG-8|US-XG-6POE|UWB‑XG|UWB‑XG‑BK",
				"rank":             1,
				"revision_history": "",
				"sdk__id":          nil,
				"size":             nil,
				"slug":             "unifi-controller-v5-user-guide",
				"thumbnail":        "https://prd-www-cdn.ubnt.com/media/images/download/user-guides/UniFi_Controller_V5_UG.jpg",
				"thumbnail_retina": "https://prd-www-cdn.ubnt.com/media/images/download/user-guides/UniFi_Controller_V5_UG-2x.jpg",
				"version":          "",
			},
			{
				"architecture":     "",
				"build":            "",
				"category__name":   "Software",
				"category__slug":   "software",
				"changelog":        "https://community.ui.com/releases/UniFi-Network-Controller-6-0-22/910ceffc-f0e9-4518-86c1-df5eeee34695",
				"date_published":   "2020-09-17",
				"description":      "UniFi Network Controller 6.0.22 for Debian/Ubuntu Linux and UniFi Cloud Key.",
				"featured":         false,
				"file_path":        "/downloads/unifi/6.0.22/unifi_sysvinit_all.deb",
				"filename":         "unifi_sysvinit_all.deb",
				"id":               2607,
				"mib":              "",
				"name":             "UniFi Network Controller 6.0.22 for Debian/Ubuntu Linux and UniFi Cloud Key",
				"products":         "UAP|UAP-AC-EDU|UAP–AC–IW|UAP–AC–IW–PRO|UAP-AC-LITE|UAP-AC-LR|UAP-AC-M|UAP-AC-M-PRO|UAP-AC-PRO|UAP-BeaconHD|UAP-FlexHD|UAP-HD|UAP-IW|UAP-IW-HD|UAP-LR|UAP-nanoHD|UAP-Outdoor|UAP-Outdoor+|UAP-Outdoor5|UAP-PRO|UAP-SHD|UAP‑XG|UAS-XG|UBB|UC-CK|UCK-G2|UCK-G2-PLUS|US-16-150W|US‑16‑XG|US-24|US-24-250W|US-24-500W|US-48|US-48-500W|US-48-750W|US-8|US-8-150W|US-8-60W|USG|USG-PRO-4|USG-XG-8|US-L2-24-POE|US-L2-48-POE|USW-16-POE|USW-24-POE|USW-48-POE|USW-Flex|USW-Flex-Mini|USW-Industrial|USW-Lite-16-POE|USW-Pro-24-POE|USW-Pro-48-POE|US-XG-6POE|UWB‑XG|UWB‑XG‑BK",
				"rank":             350,
				"revision_history": "",
				"sdk__id":          nil,
				"size":             nil,
				"slug":             "unifi-network-controller-6022-debianubuntu-linux-and-unifi-cloud-key",
				"thumbnail":        nil,
				"thumbnail_retina": nil,
				"version":          "6.0.22",
			},
			{
				"architecture":     "",
				"build":            "",
				"category__name":   "Software",
				"category__slug":   "software",
				"changelog":        "https://community.ui.com/releases/0cffd3ed-7429-4529-9a20-9fead78ebf66",
				"date_published":   "2021-03-25",
				"description":      "UniFi Network Controller 6.1.71 for Debian/Ubuntu Linux and UniFi Cloud Key.",
				"featured":         false,
				"file_path":        "/downloads/unifi/6.1.71/unifi_sysvinit_all.deb",
				"filename":         "unifi_sysvinit_all.deb",
				"id":               2777,
				"mib":              "",
				"name":             "UniFi Network Controller 6.1.71 for Debian/Ubuntu Linux and UniFi Cloud Key",
				"products":         "UAP|UAP-AC-EDU|UAP–AC–IW|UAP–AC–IW–PRO|UAP-AC-LITE|UAP-AC-LR|UAP-AC-M|UAP-AC-M-PRO|UAP-AC-PRO|UAP-BeaconHD|UAP-FlexHD|UAP-HD|UAP-IW|UAP-IW-HD|UAP-LR|UAP-nanoHD|UAP-Outdoor|UAP-Outdoor+|UAP-Outdoor5|UAP-PRO|UAP-SHD|UAP‑XG|UAS-XG|UBB|UC-CK|UCK-G2|UCK-G2-PLUS|US-16-150W|US‑16‑XG|US-24|US-24-250W|US-24-500W|US-48|US-48-500W|US-48-750W|US-8|US-8-150W|US-8-60W|USG|USG-PRO-4|USG-XG-8|US-L2-24-POE|US-L2-48-POE|USP-RPS|USW-16-POE|USW-24|USW-24-POE|USW-48|USW-48-POE|USW-Aggregation|USW-Flex|USW-Flex-Mini|USW-Industrial|USW-Lite-16-POE|USW-Lite-8-PoE|USW-Pro-24|USW-Pro-24-POE|USW-Pro-48|USW-Pro-48-POE|US-XG-6POE|UWB‑XG|UWB‑XG‑BK",
				"rank":             423,
				"revision_history": "",
				"sdk__id":          nil,
				"size":             nil,
				"slug":             "unifi-network-controller-6171-debianubuntu-linux-and-unifi-cloud-key",
				"thumbnail":        nil,
				"thumbnail_retina": nil,
				"version":          "6.1.71",
			},
			{
				"architecture":     "",
				"build":            "",
				"category__name":   "Software",
				"category__slug":   "software",
				"changelog":        "https://community.ui.com/releases/0dfcbc77-8a4f-4e20-bb93-07bbb0237e3a",
				"date_published":   "2021-06-21",
				"description":      "UniFi Network Application 6.2.26 for Debian/Ubuntu Linux and UniFi Cloud Key.",
				"featured":         true,
				"file_path":        "/downloads/unifi/6.2.26/unifi_sysvinit_all.deb",
				"filename":         "unifi_sysvinit_all.deb",
				"id":               2840,
				"mib":              "",
				"name":             "UniFi Network Application 6.2.26 for Debian/Ubuntu Linux and UniFi Cloud Key",
				"products":         "UAP|UAP-AC-EDU|UAP–AC–IW|UAP–AC–IW–PRO|UAP-AC-LITE|UAP-AC-LR|UAP-AC-M|UAP-AC-M-PRO|UAP-AC-PRO|UAP-BeaconHD|UAP-FlexHD|UAP-HD|UAP-IW|UAP-IW-HD|UAP-LR|UAP-nanoHD|UAP-Outdoor|UAP-Outdoor+|UAP-Outdoor5|UAP-PRO|UAP-SHD|UAP‑XG|UAS-XG|UBB|UC-CK|UCK-G2|UCK-G2-PLUS|UDM|UDM-Pro|US-16-150W|US‑16‑XG|US-24|US-24-250W|US-24-500W|US-48|US-48-500W|US-48-750W|US-8|US-8-150W|US-8-60W|USG|USG-PRO-4|USG-XG-8|US-L2-24-POE|US-L2-48-POE|USP-RPS|USW-16-POE|USW-24|USW-24-POE|USW-48|USW-48-POE|USW-Aggregation|USW-Enterprise-24-PoE|USW-Flex|USW-Flex-Mini|USW-Industrial|USW-Lite-16-POE|USW-Lite-8-PoE|USW-Pro-24|USW-Pro-24-POE|USW-Pro-48|USW-Pro-48-POE|USW-Pro-Aggregation|US-XG-6POE|UWB‑XG|UWB‑XG‑BK",
				"rank":             440,
				"revision_history": "",
				"sdk__id":          nil,
				"size":             nil,
				"slug":             "unifi-network-application-6226-debianubuntu-linux-and-unifi-cloud-key",
				"thumbnail":        nil,
				"thumbnail_retina": nil,
				"version":          "6.2.26",
			},
			{
				"architecture":     "",
				"build":            "",
				"category__name":   "Firmware",
				"category__slug":   "firmware",
				"changelog":        "https://community.ui.com/releases/a98a71d1-ce1e-4823-a1d2-4a5fa3d642b9",
				"date_published":   "2021-07-14",
				"description":      "UniFi firmware 5.60.9 for U6-Lite",
				"featured":         true,
				"file_path":        "/downloads/unifi/firmware/UAL6/5.60.9.12980/BZ.mt7621_5.60.9+12980.210702.0701.bin",
				"filename":         "BZ.mt7621_5.60.9+12980.210702.0701.bin",
				"id":               2847,
				"mib":              "",
				"name":             "UniFi firmware 5.60.9 for U6-Lite",
				"products":         "U6-Lite",
				"rank":             444,
				"revision_history": "",
				"sdk__id":          nil,
				"size":             nil,
				"slug":             "unifi-firmware-5609-u6-lite",
				"thumbnail":        nil,
				"thumbnail_retina": nil,
				"version":          "5.60.9",
			},
		},
		"products": []map[string]interface{}{},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, []string{"XMLHttpRequest"}, req.Header["X-Requested-With"])

		resp, err := json.Marshal(respData)
		assert.Nil(t, err)

		_, err = rw.Write(resp)
		assert.Nil(t, err)
	}))
	defer server.Close()

	expected, err := version.NewVersion("6.2.26")
	assert.Nil(t, err)

	uiDownloadUrl = server.URL
	actual, err := latestUnifiVersion()
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
