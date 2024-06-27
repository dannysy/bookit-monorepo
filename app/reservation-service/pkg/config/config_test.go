package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldLoadNewConf(t *testing.T) {
	temp, err := os.CreateTemp("", "config-*.yml")
	defer func(temp *os.File) {
		err := temp.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(temp)
	if err != nil {
		t.Fatal(err)
	}
	cfgFile, err := os.ReadFile("../../etc/config.yml.template")
	if err != nil {
		t.Fatal(err)
	}
	_, err = temp.Write(cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	err = temp.Sync()
	if err != nil {
		t.Fatal(err)
	}
	os.Args = []string{"test", "-config", temp.Name()}
	cfg := New()
	assert.NotEmpty(t, cfg.Database.URL, "url should not be empty")
}
