package helpers

import (
	"encoding/json"
	"strconv"
	"strings"
)

func MustParseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func MustParseBool(s string) bool {
	return !StrSliceIn(strings.ToLower(s), []string{
		"",
		"0",
		"false",
		"f",
	})
}

func MustJsonMarshal(v interface{}) []byte {
	if b, e := json.Marshal(v); e == nil {
		return b
	} else {
		panic(e)
	}
}
func MustJsonMarshalString(v interface{}) string {
	return string(MustJsonMarshal(v))
}
