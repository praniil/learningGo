package golangthings

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sha256() {
	data := "test string"

	hash := sha256.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum([]byte("hi"))         //argument to Sum is to append to existing type		//returns inbytes
	hashString := hex.EncodeToString(byteSlice) //hexadecimal encoding
	fmt.Println(data)
	fmt.Printf("%x\n", byteSlice)
	fmt.Println(hashString)
}
