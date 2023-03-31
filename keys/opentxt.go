package keys

import (
	"log"
	"os"
)

func Token() string {
	content, err := os.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func AccessKey() string {
	content, err := os.ReadFile("accesskey.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
