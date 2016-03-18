package config

import (
	"testing"
	"os"
)

func TestInit(t *testing.T) {
	configPath := os.Getenv("GOPATH") + string(os.PathSeparator) + "src/github.com/scofieldpeng/go-libs/config/test.ini"
	if err := Config.Init(configPath);err != nil {
		t.Fatal(err)
	}
}
