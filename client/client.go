package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net"
)

type Data struct {
	Data string
	Hash string
}

func main() {
	conn, err := net.Dial("tcp", "localhost:4221")

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	str := "Main Data"

	h := sha256.New()
	h.Write([]byte(str))

	hash := h.Sum(nil)
	// data := Data{
	// 	Data: str,
	// 	Hash: string(h.Sum(nil)),
	// }

	// json_data, err := json.Marshal(data)

	// if err != nil {
	// 	return
	// }
	var data string

	fmt.Print("Data: ")
	fmt.Scan(&data)

	data_main := make([]byte, len(data)+32)

	var version byte = 1

	data_main[0] = version

	copy(data_main[1:], data)

	log.Println(len(hash))
	println(hash)
	copy(data_main[len(data)+1:], hash)

	log.Println(len(data_main))

	_, err = conn.Write([]byte(data_main))

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	defer conn.Close()
}
