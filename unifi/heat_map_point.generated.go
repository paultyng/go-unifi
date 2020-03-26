// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"fmt"
)

// just to fix compile issues with the import
var (
	_ fmt.Formatter
	_ context.Context
)

type HeatMapPoint struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	DownloadSpeed float64 `json:"download_speed,omitempty"`
	HeatmapID     string  `json:"heatmap_id"`
	UploadSpeed   float64 `json:"upload_speed,omitempty"`
	X             float64 `json:"x,omitempty"`
	Y             float64 `json:"y,omitempty"`
}

func (c *Client) listHeatMapPoint(ctx context.Context, site string) ([]HeatMapPoint, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []HeatMapPoint `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/heatmappoint", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *Client) getHeatMapPoint(ctx context.Context, site, id string) (*HeatMapPoint, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []HeatMapPoint `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/heatmappoint/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) deleteHeatMapPoint(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/heatmappoint/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createHeatMapPoint(ctx context.Context, site string, d *HeatMapPoint) (*HeatMapPoint, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []HeatMapPoint `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/heatmappoint", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *Client) updateHeatMapPoint(ctx context.Context, site string, d *HeatMapPoint) (*HeatMapPoint, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []HeatMapPoint `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/heatmappoint/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
