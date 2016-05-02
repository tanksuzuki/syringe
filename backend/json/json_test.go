package json

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetKeyValueFromString(t *testing.T) {
	m, err := GetKeyValueFromString("{\"test\":\"foo\"}")
	if !reflect.DeepEqual(m, map[string]interface{}{"test": "foo"}) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromString_Invalid(t *testing.T) {
	expected := "json: invalid character 'i' looking for beginning of value"
	if _, err := GetKeyValueFromString("invalid"); fmt.Sprintf("%s", err) != expected {
		t.Errorf("err: %s\n", err)
	}
}

func TestGetKeyValueFromFile(t *testing.T) {
	m, err := GetKeyValueFromFiles([]string{"../../test/backend/json/single.json"})
	if !reflect.DeepEqual(m, map[string]interface{}{"test": "foo"}) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromFile_Invalid(t *testing.T) {
	expected := "json: invalid character 'i' looking for beginning of value"
	if _, err := GetKeyValueFromFiles([]string{"../../test/backend/toml/invalid.toml"}); fmt.Sprintf("%s", err) != expected {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromFile_NotFound(t *testing.T) {
	expected := "open notfound: no such file or directory"
	if _, err := GetKeyValueFromFiles([]string{"notfound"}); fmt.Sprintf("%s", err) != expected {
		t.Fatalf("err: %s\n", err)
	}
}
