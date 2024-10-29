package main

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net"
)

type Data struct {
	Data string
	Hash string
}

func main() {
	conn, err := net.Dial("tcp", "server:4221")

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	str := "Main Data"

	h := sha256.New()
	h.Write([]byte(str))

	data := Data{
		Data: str,
		Hash: string(h.Sum(nil)),
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		return
	}

	_, err = conn.Write([]byte(string(json_data)))

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	defer conn.Close()
}
