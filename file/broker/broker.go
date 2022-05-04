package main

import (
	"file/messages"
	"fmt"
	"log"
	"net"
)

func handleClient(msgHandler *messages.MessageHandler, directory string) {

	for {
		wrapper, _ := msgHandler.Receive()
		switch msg := wrapper.Msg.(type) {
		case *messages.Wrapper_ClimateData:
			handleClimateData(msg)
			continue
		case nil:
			log.Println("Received an empty message, terminating client")
			msgHandler.Close()
			return
		default:
			log.Printf("Unexpected message type: %T", msg)
		}
	}
}

func handleClimateData(msg *messages.Wrapper_ClimateData) {
	timestamp := msg.ClimateData.UTC_TIMESTAMP
	airTemperature := msg.ClimateData.AIR_TEMPERATURE

	log.Printf("Timestamp: %T", timestamp)
	log.Printf("Air Temperature: %T", airTemperature)
}

func main() {
	port := "20111"
	fmt.Println("Listening at port: " + port)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	for {
		if conn, err := listener.Accept(); err == nil {
			msgHandler := messages.NewMessageHandler(conn)
			defer msgHandler.Close()
			go handleClient(msgHandler, "/")
		}
	}
}
