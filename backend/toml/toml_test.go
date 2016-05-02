package toml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetKeyValueFromString(t *testing.T) {
	m, err := GetKeyValueFromString("test = \"foo\"")
	if !reflect.DeepEqual(m, map[string]interface{}{"test": "foo"}) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromString_Invalid(t *testing.T) {
	expected := "toml: Near line 0 (last key parsed ''): Bare keys cannot contain '\\n'."
	if _, err := GetKeyValueFromString("invalid"); fmt.Sprintf("%s", err) != expected {
		t.Errorf("err: %s\n", err)
	}
}

func TestGetKeyValueFromFile(t *testing.T) {
	m, err := GetKeyValueFromFiles([]string{"../../test/backend/toml/single.toml"})
	if !reflect.DeepEqual(m, map[string]interface{}{"test": "foo"}) {
		t.Fatalf("map: %+v\n", m)
	}
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestGetKeyValueFromFile_Invalid(t *testing.T) {
	expected := "toml: Near line 0 (last key parsed ''): Bare keys cannot contain '\\n'."
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
