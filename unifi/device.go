package unifi

import (
	"context"
	"fmt"
)

const (
	DEVICE_STATE_UNKNOWN           = 0
	DEVICE_STATE_CONNECTED         = 1
	DEVICE_STATE_PENDING           = 2
	DEVICE_STATE_FIRMWARE_MISMATCH = 3
	DEVICE_STATE_UPGRADING         = 4
	DEVICE_STATE_PROVISIONING      = 5
	DEVICE_STATE_HEARTBEAT_MISSED  = 6
	DEVICE_STATE_ADOPTING          = 7
	DEVICE_STATE_DELETING          = 8
	DEVICE_STATE_INFORM_ERROR      = 9
	DEVICE_STATE_ADOPT_FAILED      = 10
	DEVICE_STATE_ISOLATED          = 11
)

func (c *Client) ListDevice(ctx context.Context, site string) ([]Device, error) {
	return c.listDevice(ctx, site)
}

func (c *Client) GetDeviceByMAC(ctx context.Context, site, mac string) (*Device, error) {
	return c.getDevice(ctx, site, mac)
}

func (c *Client) DeleteDevice(ctx context.Context, site, id string) error {
	return c.deleteDevice(ctx, site, id)
}

func (c *Client) CreateDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	return c.createDevice(ctx, site, d)
}

func (c *Client) UpdateDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	return c.updateDevice(ctx, site, d)
}

func (c *Client) GetDevice(ctx context.Context, site, id string) (*Device, error) {
	devices, err := c.ListDevice(ctx, site)

	if err != nil {
		return nil, err
	}

	for _, d := range devices {
		if d.ID == id {
			return &d, nil
		}
	}

	return nil, &NotFoundError{}
}

func (c *Client) AdoptDevice(ctx context.Context, site, mac string) error {
	reqBody := struct {
		Cmd string `json:"cmd"`
		MAC string `json:"mac"`
	}{
		Cmd: "adopt",
		MAC: mac,
	}

	var respBody struct {
		Meta meta `json:"meta"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/cmd/devmgr", site), reqBody, &respBody)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ForgetDevice(ctx context.Context, site, mac string) error {
	reqBody := struct {
		Cmd  string   `json:"cmd"`
		MACs []string `json:"macs"`
	}{
		Cmd:  "delete-device",
		MACs: []string{mac},
	}

	var respBody struct {
		Meta meta     `json:"meta"`
		Data []Device `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/cmd/sitemgr", site), reqBody, &respBody)
	if err != nil {
		return err
	}

	return nil
}
