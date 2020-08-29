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

type ScheduleTask struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Action                  string `json:"action,omitempty"` // stream|upgrade
	AdditionalSoundsEnabled bool   `json:"additional_sounds_enabled"`
	BroadcastgroupID        string `json:"broadcastgroup_id"`
	CronExpr                string `json:"cron_expr,omitempty"`
	ExecuteOnlyOnce         bool   `json:"execute_only_once"`
	MediafileID             string `json:"mediafile_id"`
	Name                    string `json:"name,omitempty"`
	SampleFilename          string `json:"sample_filename,omitempty"`
	StreamType              string `json:"stream_type,omitempty"` // media|sample
	UpgradeTargets          []struct {
		MAC string `json:"mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	} `json:"upgrade_targets,omitempty"`
}

func (c *Client) listScheduleTask(ctx context.Context, site string) ([]ScheduleTask, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []ScheduleTask `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/scheduletask", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *Client) getScheduleTask(ctx context.Context, site, id string) (*ScheduleTask, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []ScheduleTask `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/scheduletask/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) deleteScheduleTask(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/scheduletask/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createScheduleTask(ctx context.Context, site string, d *ScheduleTask) (*ScheduleTask, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []ScheduleTask `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/scheduletask", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *Client) updateScheduleTask(ctx context.Context, site string, d *ScheduleTask) (*ScheduleTask, error) {
	var respBody struct {
		Meta meta           `json:"meta"`
		Data []ScheduleTask `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/scheduletask/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}
