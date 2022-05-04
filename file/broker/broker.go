package main

import (
	"file/messages"
	"fmt"
	"log"
	"net"
	"time"
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
			break
		default:
			log.Printf("Unexpected message type: %T", msg)
		}
		time.Sleep(1 * time.Second)
	}
}

func handleClimateData(msg *messages.Wrapper_ClimateData) {
	timestamp := msg.ClimateData.GetUTC_TIMESTAMP()
	airTemperature := msg.ClimateData.GetAIR_TEMPERATURE()
	solarRadiation := msg.ClimateData.GetSOLAR_RADIATION()
	log.Printf("Timestamp: %v", timestamp)
	log.Printf("Air Temperature: %v", airTemperature)
	log.Printf("Solar Radiation: %v", solarRadiation)
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
