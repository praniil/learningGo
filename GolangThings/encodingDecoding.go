package golangthings

import (
	bs4 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

// we do encoding for like transmission of data so that it wont get corrupt in this format. Also done in case of binary or json so that it can be stored in database which actually supports text
type Person struct {
	Name  string
	Age   int
	Email string
}

func EncodingDecoding() {
	person := Person{
		Name:  "Pranil",
		Age:   20,
		Email: "pparajuli15@gamil.com",
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Encoding Json: ", err)
	}

	os.Stdout.Write(jsonData)

	data := "abc123!@#$%^&*()_+><>?:*"
	encodedData := bs4.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encodedData: ", encodedData);

	decodedData, err := bs4.StdEncoding.DecodeString(encodedData)
	if err != nil {
		fmt.Println("error decoding string: ", err)
	}
	fmt.Println("decodedData: ", string(decodedData))		//if i dont show the data in string it shows the ASCII values

}
