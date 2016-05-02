package template

import (
	"fmt"
	"strings"
	"testing"
)

func TestMerge(t *testing.T) {
	s, err := Merge("((.test))", "((", "))", map[string]interface{}{"test": "foo"})
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	if s != "foo" {
		t.Fatalf("s: %s\n", s)
	}
}

func TestMerge_Invalid(t *testing.T) {
	expected := "template: {{:1: unexpected unclosed action in command"
	_, err := Merge("{{", "{{", "}}", nil)
	if fmt.Sprintf("%s", err) != expected {
		t.Fatalf("err: %s\n", err)
	}
}

func TestTrim(t *testing.T) {
	templateString := `{{ $test := "foo" }}
{{/*comment*/}}
{{ $test }}`
	if s := trim(templateString, "{{", "}}"); strings.Contains("s", "\n") {
		t.Fatalf("s: %s\n", s)
	}
}

func TestToNetworkWithLength(t *testing.T) {
	if n := toNetwork("192.168.1.1", "24"); n != "192.168.1.0" {
		t.Fatalf("s: %s\n", n)
	}
}

func TestToNetworkWithMask(t *testing.T) {
	if n := toNetwork("192.168.1.1", "255.255.255.0"); n != "192.168.1.0" {
		t.Fatalf("s: %s\n", n)
	}
}

func TestToNetworkWithInvalidLength(t *testing.T) {
	if n := toNetwork("192.168.1.1", "33"); n != "Invalid prefix length" {
		t.Fatalf("s: %s\n", n)
	}
}

func TestToNetworkWithInvalidMask(t *testing.T) {
	if n := toNetwork("192.168.1.1", "255.255.255.255.0"); n != "Invalid mask" {
		t.Fatalf("s: %s\n", n)
	}
}

func TestToNetworkWithInvalidIP(t *testing.T) {
	if n := toNetwork("192.168.1.1.1", "255.255.255.0"); n != "Invalid IP address" {
		t.Fatalf("s: %s\n", n)
	}
}

func TestToPrefixLen(t *testing.T) {
	if p := toPrefixLen("255.255.255.0"); p != "24" {
		t.Fatalf("p: %s\n", p)
	}
}

func TestToPrefixLenInvalid(t *testing.T) {
	if p := toPrefixLen("255.255.255.255.0"); p != "Invalid mask" {
		t.Fatalf("p: %s\n", p)
	}
}

func TestToSubnetMask(t *testing.T) {
	if m := toSubnetMask("24"); m != "255.255.255.0" {
		t.Fatalf("m: %s\n", m)
	}
}

func TestToSubnetMaskInvalid(t *testing.T) {
	if m := toSubnetMask("33"); m != "Invalid prefix length" {
		t.Fatalf("m: %s\n", m)
	}
}

func TestExec(t *testing.T) {
	if e := exec("echo"); e != "\n" {
		t.Fatalf("e: %s\n", e)
	}
}

func TestExecWithArg(t *testing.T) {
	if e := exec("echo test"); e != "test\n" {
		t.Fatalf("e: %s\n", e)
	}
}

func TestExecInvalid(t *testing.T) {
	if e := exec(""); e != "Invalid command" {
		t.Fatalf("e: %s\n", e)
	}
}
