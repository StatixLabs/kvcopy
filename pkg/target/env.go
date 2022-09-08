package target

import (
	"strings"
)

func Env(values map[string]string) string {
	var output strings.Builder
	for key, value := range values {
		output.WriteString(key + "=" + value + "\n")
	}
	return output.String()
}
