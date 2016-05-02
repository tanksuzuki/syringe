package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestNoArgument(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "\x1b[31mERROR\x1b[0m: syringe requires a minimum of 1 argument. Please see 'syringe --help'.\n"

	if s := c.run(nil); s != 1 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", errStream) != expected {
		t.Fatalf("stderr: %s\n", errStream)
	}
}

func TestDebugFlag(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "\x1b[36mDEBUG\x1b[0m: "

	c.run(strings.Split("--debug", " "))
	if !strings.Contains(fmt.Sprintf("%s", outStream), expected) {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestShortHelpFlag(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "syringe [options] <template> [<backend>...]"

	if s := c.run(strings.Split("-h", " ")); s != 1 {
		t.Fatalf("exit code: %d\n", s)
	}
	if !strings.Contains(fmt.Sprintf("%s", outStream), expected) {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestLongHelpFlag(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "syringe [options] <template> [<backend>...]"

	if s := c.run(strings.Split("--help", " ")); s != 1 {
		t.Fatalf("exit code: %d\n", s)
	}
	if !strings.Contains(fmt.Sprintf("%s", outStream), expected) {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestVersionFlag(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "syringe version " + VERSION + "\n"

	if s := c.run(strings.Split("--version", " ")); s != 1 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", outStream) != expected {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestImportKeyValueFromPipe(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	inStream.Write([]byte("test = \"foo\""))
	expected := "test: foo\n"

	if s := c.run(strings.Split("test/template/single.txt", " ")); s != 0 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", outStream) != expected {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestImportKeyValueFromFile(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "test: foo\n"

	if s := c.run(strings.Split("test/template/single.txt test/backend/toml/single.toml", " ")); s != 0 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", outStream) != expected {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestImportKeyValueFromFile_NotFound(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "\x1b[31mERROR\x1b[0m: open notfound: no such file or directory\n"

	if s := c.run(strings.Split("notfound", " ")); s != 1 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", errStream) != expected {
		t.Fatalf("stderr: %s\n", errStream)
	}
}

func TestImportKeyValueFromFlag(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "test: foo\n"

	if s := c.run(strings.Split("test/template/single.txt -v test:foo", " ")); s != 0 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", outStream) != expected {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestOverrideKeyValueFromPipeAndFile(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	inStream.Write([]byte("test = \"bar\""))
	expected := "test: foo\n"

	if s := c.run(strings.Split("test/template/single.txt test/backend/toml/single.toml", " ")); s != 0 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", outStream) != expected {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestOverrideKeyValueFromFileAndFlag(t *testing.T) {
	inStream, outStream, errStream := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	c := &cli{inStream: inStream, outStream: outStream, errStream: errStream}
	expected := "test: bar\n"

	if s := c.run(strings.Split("test/template/single.txt test/backend/toml/single.toml -v test:bar", " ")); s != 0 {
		t.Fatalf("exit code: %d\n", s)
	}
	if fmt.Sprintf("%s", outStream) != expected {
		t.Fatalf("stdout: %s\n", outStream)
	}
}

func TestParseArgs(t *testing.T) {
	flag, args, err := parseArgs(strings.Split("arg1 -b json arg2", " "))

	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	if flag.Backend != "json" {
		t.Fatalf("flag: %+v\n", flag)
	}
	if !reflect.DeepEqual(args, []string{"arg1", "arg2"}) {
		t.Fatalf("args: %+v\n", args)
	}
}

func TestParseArgsWithEnv(t *testing.T) {
	envBackend := setTestEnv("SY_BACKEND", "test_backend")
	defer envBackend()
	envDebug := setTestEnv("SY_DEBUG", "true")
	defer envDebug()
	envDelimLeft := setTestEnv("SY_DELIML", "test_deliml")
	defer envDelimLeft()
	envDelimRight := setTestEnv("SY_DELIMR", "test_delimr")
	defer envDelimRight()

	flag, _, err := parseArgs(nil)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	if flag.Backend != "test_backend" {
		t.Fatalf("flag: %+v\n", flag)
	}
	if !flag.Debug {
		t.Fatalf("flag: %+v\n", flag)
	}
	if flag.DelimLeft != "test_deliml" {
		t.Fatalf("flag: %+v\n", flag)
	}
	if flag.DelimRight != "test_delimr" {
		t.Fatalf("flag: %+v\n", flag)
	}
}

func TestParseArgsWithUnknownFlag(t *testing.T) {
	if _, _, err := parseArgs(strings.Split("--unknown", " ")); err == nil {
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
