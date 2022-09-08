package source

import (
	"fmt"
	"os"
	"strings"
)

func File(path string) map[string]string {
	contents, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	entries := strings.Split(string(contents), "\n")
	output := make(map[string]string)
	for _, e := range entries {
		fmt.Println(e)
		parts := strings.Split(e, "=")
		if len(parts) == 2 {
			output[parts[0]] = parts[1]
		}
	}
	return output
}
