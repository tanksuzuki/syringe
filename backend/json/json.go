package json

import (
	j "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func GetKeyValueFromString(str string) (map[string]interface{}, error) {
	var keyValue map[string]interface{}
	err := j.Unmarshal([]byte(str), &keyValue)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("json: %s", err))
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
		err = j.Unmarshal(b, &m)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("json: %s", err))
		}
		for key, _ := range m {
			keyValue[key] = m[key]
		}
	}

	return keyValue, nil
}
