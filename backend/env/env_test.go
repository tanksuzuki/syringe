package env

import (
	"os"
	"testing"
)

func TestGetKeyValue(t *testing.T) {
	reset := setTestEnv("SYRINGE_TEST", "foo")
	defer reset()

	m := GetKeyValue()
	if m["SYRINGE_TEST"] != "foo" {
		t.Fatalf("map: %+v\n", m)
	}
}

func setTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}
