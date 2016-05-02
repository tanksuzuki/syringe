package toml

import (
	"errors"
	"fmt"
	t "github.com/BurntSushi/toml"
	"io/ioutil"
)

func GetKeyValueFromString(str string) (map[string]interface{}, error) {
	var keyValue map[string]interface{}
	_, err := t.Decode(str, &keyValue)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("toml: %s", err))
	}
	return keyValue, nil
}

func GetKeyValueFromFiles(paths []string) (map[string]interface{}, error) {
	keyValue := map[string]interface{}{}

	for _, p := range paths {
		m := map[string]interface{}{}
		b, err := ioutil.ReadFile(p)
		if err != nil {
			return nil, err
		}
		_, err = t.Decode(string(b), &m)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("toml: %s", err))
		}
		for key, _ := range m {
			keyValue[key] = m[key]
		}
	}

	return keyValue, nil
}
