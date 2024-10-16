package core

import (
	"log"
	"os"
)

func MustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
