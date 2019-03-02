package oauth2config

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestConfigInterface(t *testing.T) {
	c := &oauth2.Config{}
	var i interface{} = c

	if _, ok := i.(Config); ok != true {
		t.Errorf("Given config doesn't implement Config")
	}
}

func TestFakeConfig(t *testing.T) {
	c := &FakeConfig{}
	var i interface{} = c

	if _, ok := i.(Config); ok != true {
		t.Errorf("FakeConfig doesn't implement Config")
	}
}
