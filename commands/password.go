package commands

import (
	"crypto/rand"
	"fmt"
	"github.com/etng/bbx/helpers"
)

const ReadableAlphaNumMap = "23456789abcdefghjkmnpqrstuvwxyz"

// EncodeReadableAlphaNumForRand 可以把 二进制数据encode成 ReadableAlphaNum
// 但是此处不能decode,这个东西暂时仅能用于rand.(输入参数应该使用类似于MD5的东西)
func EncodeReadableAlphaNumForRand(b []byte) string {
	outBytes := make([]byte, len(b)/2)
	mapLen := len(ReadableAlphaNumMap)
	for i := 0; i < len(outBytes); i++ {
		outBytes[i] = ReadableAlphaNumMap[(int(b[2*i])*256+int(b[2*i+1]))%(mapLen)]
	}
	return string(outBytes)
}
func NewPassword(length int) string {
	src := make([]byte, length*2)
	if n, e := rand.Read(src); e != nil || n == 0 {
		fmt.Println("rand read fail", e, n)
	}
	return EncodeReadableAlphaNumForRand(src)
}

func init() {
	allCommands = append(allCommands, Command{
		Name:      "password",
		Desc:      "random readable string to act as tag or password",
		AliasList: []string{"passwd"},
		Handler: func(args ...string) {
			length := 6
			if len(args) > 0 {
				length = helpers.MustParseInt(args[0])
			}
			fmt.Println(NewPassword(length))
		},
		Weight: 0,
	})
}
