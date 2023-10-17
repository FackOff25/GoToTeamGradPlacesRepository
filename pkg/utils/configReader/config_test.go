package configReader

import (
	"testing"
)

func TestConfigReader(t *testing.T) {
	_, err := NewConfig("../../../configs/config.toml")
	if err != nil {
		t.Error()
	}
}

func TestConfigReaderNegative(t *testing.T) {
	_, err := NewConfig("")
	if err == nil {
		t.Error()
	}
}
