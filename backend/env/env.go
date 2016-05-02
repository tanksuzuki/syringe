package env

import (
	"os"
	"strings"
)

func GetKeyValue() map[string]interface{} {
	keyValue := map[string]interface{}{}

	for _, e := range os.Environ() {
		i := strings.Index(e, "=")
		keyValue[e[0:i]] = e[i+1:]
	}

	return keyValue
}
