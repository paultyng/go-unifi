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

type SpatialRecord struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Devices []struct {
		MAC      string `json:"mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
		Position struct {
			X float64 `json:"x,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
			Y float64 `json:"y,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
			Z float64 `json:"z,omitempty"` // (^([-]?[\d]+)$)|(^([-]?[\d]+[.]?[\d]+)$)
		} `json:"position,omitempty"`
	} `json:"devices,omitempty"`
	Name string `json:"name,omitempty"` // .{1,128}
}

func (c *Client) listSpatialRecord(ctx context.Context, site string) ([]SpatialRecord, error) {
	var respBody struct {
		Meta meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/spatialrecord", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *Client) getSpatialRecord(ctx context.Context, site, id string) (*SpatialRecord, error) {
	var respBody struct {
		Meta meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) deleteSpatialRecord(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createSpatialRecord(ctx context.Context, site string, d *SpatialRecord) (*SpatialRecord, error) {
	var respBody struct {
		Meta meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/spatialrecord", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *Client) updateSpatialRecord(ctx context.Context, site string, d *SpatialRecord) (*SpatialRecord, error) {
	var respBody struct {
		Meta meta            `json:"meta"`
		Data []SpatialRecord `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/spatialrecord/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
