package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swierq/golang/internal/loadek"
	"github.com/swierq/golang/pkg/webapp"
)

func TestNewApp(t *testing.T) {
	app, err := newApp()
	assert.NotNil(t, err)
	assert.Nil(t, app)

	webConfig := webapp.NewConfig()
	//TODO: use option functions to set the port and log level
	webapp := webapp.NewApp(webConfig)
	loadekConfig, _ := loadek.NewConfig(loadek.WithCpuLoadMi(10), loadek.WithMemLoadMb(10))
	loadek := loadek.NewApp(loadek.WithConfig(loadekConfig))
	app, err = newApp(withLoadek(loadek), withWebApp(webapp))
	assert.Nil(t, err)
	assert.NotNil(t, app)
}
