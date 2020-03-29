package ctlib

import (
	"bytes"
	"strings"

	ctconfig "github.com/hashicorp/consul-template/config"
	ctman "github.com/hashicorp/consul-template/manager"
)

// Execute populates a Consul Template template string, equivalent to running the command:
// `consul-template -template=<file containing the template string> -dry -once`
func Execute(template string) (*string, error) {

	// overrides for never touching the filesystem
	createDestDirs := false
	errMissingKey := true
	destination := "dummydest"

	// feed our template string into a default template config, with above overrides
	templateConfig := ctconfig.DefaultTemplateConfig()
	templateConfig.Contents = &template
	templateConfig.CreateDestDirs = &createDestDirs
	templateConfig.ErrMissingKey = &errMissingKey
	templateConfig.Destination = &destination
	templateConfigs := ctconfig.TemplateConfigs([]*ctconfig.TemplateConfig{templateConfig})

	// default config, except it runs to completion and has the template argument injected
	conf := ctconfig.DefaultConfig()
	conf.Once = true
	conf.Templates = &templateConfigs

	// default runner, except it writes output to a byte buffer in memory
	runner, err := ctman.NewRunner(conf, true)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	runner.SetOutStream(&buf)

	// start the runner, which closes by itself, and read any errors from its error channel
	runner.Start()
	close(runner.ErrCh)
	err = <-runner.ErrCh
	if err != nil {
		return nil, err
	}

	// sanitize and return the rendered output string
	rendered := buf.String()
	rendered = strings.Replace(rendered, "> dummydest\n", "", 1)
	return &rendered, nil
}
