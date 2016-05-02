package backend

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGetKeyValueFromString_Env(t *testing.T) {
	m, err := GetKeyValueFromString("env", "foo")
	if m != nil {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromString_Json(t *testing.T) {
	expected := map[string]interface{}{"test": "foo"}
	m, err := GetKeyValueFromString("json", "{\"test\":\"foo\"}")
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromString_Toml(t *testing.T) {
	expected := map[string]interface{}{"test": "foo"}
	m, err := GetKeyValueFromString("toml", "test = \"foo\"")
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromString_Invalid(t *testing.T) {
	expected := "'invalid' is not a valid backend type."
	m, err := GetKeyValueFromString("invalid", "")
	if m != nil {
		t.Fatalf("map: %+v\n", m)
	}
	if fmt.Sprintf("%s", err) != expected {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromBackends_Env(t *testing.T) {
	reset := setTestEnv("SYRINGE_TEST", "foo")
	defer reset()
	m, err := GetKeyValueFromBackends("env", nil)
	if m["SYRINGE_TEST"] != "foo" {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromBackends_EnvWithArgs(t *testing.T) {
	expected := "'env' is not supported the arguments of backend."
	_, err := GetKeyValueFromBackends("env", []string{"foo", "bar"})
	if fmt.Sprintf("%s", err) != expected {
		fmt.Printf("err: %s\n", err)
	}
}

func TestGetKeyValueFromBackends_Json(t *testing.T) {
	expected := map[string]interface{}{"test": "foo"}
	m, err := GetKeyValueFromBackends("json", []string{"../test/backend/json/single.json"})
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromBackends_Toml(t *testing.T) {
	expected := map[string]interface{}{"test": "foo"}
	m, err := GetKeyValueFromBackends("toml", []string{"../test/backend/toml/single.toml"})
	if !reflect.DeepEqual(m, expected) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromBackends_Invalid(t *testing.T) {
	expected := "'invalid' is not a valid backend type."
	m, err := GetKeyValueFromBackends("invalid", nil)
	if m != nil {
		t.Fatalf("map: %+v\n", m)
	}
	if fmt.Sprintf("%s", err) != expected {
		t.Fatalf("err: %s\n", err)
	}
}

func setTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}
