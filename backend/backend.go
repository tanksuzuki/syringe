package backend

import (
	"errors"
	"fmt"

	"github.com/tanksuzuki/syringe/backend/env"
	"github.com/tanksuzuki/syringe/backend/json"
	"github.com/tanksuzuki/syringe/backend/toml"
)

func GetKeyValueFromString(backend, str string) (map[string]interface{}, error) {
	switch backend {
	case "env":
		// env is not supported the stdin from the pipe.
		return nil, nil
	case "json":
		keyValue, err := json.GetKeyValueFromString(str)
		if err != nil {
			return nil, err
		}
		return keyValue, nil
	case "toml":
		keyValue, err := toml.GetKeyValueFromString(str)
		if err != nil {
			return nil, err
		}
		return keyValue, nil
	default:
		return nil, errors.New(fmt.Sprintf("'%s' is not a valid backend type.", backend))
	}
}

func GetKeyValueFromBackends(backend string, args []string) (map[string]interface{}, error) {
	switch backend {
	case "env":
		if len(args) > 0 {
			return nil, errors.New("'env' is not supported the arguments of backend.")
		}
		return env.GetKeyValue(), nil
	case "json":
		return json.GetKeyValueFromFiles(args)
	case "toml":
		return toml.GetKeyValueFromFiles(args)
	default:
		return nil, errors.New(fmt.Sprintf("'%s' is not a valid backend type.", backend))
	}
}
