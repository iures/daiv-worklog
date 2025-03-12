package main

import (
	"fmt"

	internal "daiv-worklog/internal"

	plug "github.com/iures/daivplug"
)

// WorklogPlugin implements the Plugin interface
type WorklogPlugin struct{
  WorklogPath string
}

// Plugin is exported as a symbol for the daiv plugin system to find
var Plugin plug.Plugin = &WorklogPlugin{}

// Name returns the unique identifier for this plugin
func (p *WorklogPlugin) Name() string {
	return "worklog"
}

// Manifest returns the plugin manifest
func (p *WorklogPlugin) Manifest() *plug.PluginManifest {
	return &plug.PluginManifest{
		ConfigKeys: []plug.ConfigKey{
			{
				Type:        plug.ConfigTypeString,
				Key:         "worklog.directory",
				Name:        "Worklog Directory",
				Description: "The directory containing the worklog files",
				Required:    true,
			},
		},
	}
}

// Initialize sets up the plugin with its configuration
func (p *WorklogPlugin) Initialize(settings map[string]interface{}) error {
	worklogPath := settings["worklog.directory"].(string)
	if worklogPath == "" {
		return fmt.Errorf("worklog.directory is required")
	}

	p.WorklogPath = worklogPath

	return nil
}

// Shutdown performs cleanup when the plugin is being disabled/removed
func (p *WorklogPlugin) Shutdown() error {
	// TODO: Clean up any resources
	return nil
}

// GetStandupContext implements the StandupPlugin interface
func (p *WorklogPlugin) GetStandupContext(timeRange plug.TimeRange) (plug.StandupContext, error) {
	fmt.Println("Getting standup context for worklog plugin")
	standupContext := internal.NewStandupContext(p.WorklogPath, timeRange)
	content, err := standupContext.Render()
	if err != nil {
		return plug.StandupContext{}, err
	}

	return plug.StandupContext{
		PluginName: p.Name(),
		Content:    content,
	}, nil
}
