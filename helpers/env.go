package helpers

import (
	"os"
)

func GetIntEnv(name string) int {
	s := os.Getenv(name)
	if s == "" {
		return 0
	}
	return MustParseInt(s)
}

func GetBoolEnv(name string) bool {
	s := os.Getenv(name)
	return MustParseBool(s)
}
