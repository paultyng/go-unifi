package unifi

import (
	"context"
	"fmt"
)

type SysInfo struct {
	Timezone                 string `json:"timezone"`
	Build                    string `json:"build"`
	Version                  string `json:"version"`
	SimulatorDemoModeEnabled bool   `json:"simulator_demo_mode_enabled,omitempty"`
}

func (c *Client) Sysinfo(ctx context.Context, id string) (*SysInfo, error) {
	var respBody struct {
		Meta meta      `json:"meta"`
		Data []SysInfo `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/stat/sysinfo", id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	return &respBody.Data[0], nil
}
