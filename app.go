package main

import "log"

func main() {
	_, err := LoadTcpConfig()
	if err != nil {
		log.Printf("Failed")
		return
	}

	log.Printf("Success")
}
