package main

import (
	"bufio"
	"file/messages"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func parseLine(line string) {
	words := strings.Fields(line)

	timestamp, _ := strconv.Atoi(words[1])
	longitude, _ := strconv.ParseFloat(words[6], 32)
	latitude, _ := strconv.ParseFloat(words[7], 32)
	air_temperature, _ := strconv.ParseFloat(words[8], 32)
	precipitaion, _ := strconv.ParseFloat(words[9], 32)
	solar_radiation, _ := strconv.ParseFloat(words[10], 32)
	surface_temperature, _ := strconv.ParseFloat(words[12], 32)
	relative_humidity, _ := strconv.Atoi(words[15])

	sendData(timestamp, longitude, latitude, air_temperature, precipitaion, solar_radiation, surface_temperature, relative_humidity)

}

func sendData(timestamp int64, longitude float32, latitude float32, air_temperature float32, precipitaion float32, solar_radiation float32,
	surface_temperature float32, relative_humidity int64) {
	nodename := "localhost"
	conn, err := net.Dial("tcp", nodename+":20666")
	if err != nil {
		log.Fatalln(err.Error())
	}
	msgHandler := messages.NewMessageHandler(conn)
	msg := messages.ToMapper{UTC_TIMESTAMP: timestamp, LONGITUDE: longitude, LATITUDE: latitude, AIR_TEMPERATURE: air_temperature,
		PRECIPITATION: precipitaion, SOLAR_RADIATION: solar_radiation, SURFACE_TEMPERATURE: surface_temperature, RELATIVE_HUMIDITY: relative_humidity}
	wrapper := &messages.Wrapper{
		Msg: &messages.Wrapper_ToMapperMessage{ToMapperMessage: &msg},
	}
	msgHandler.Send(wrapper)
}

// int64 UTC_TIMESTAMP = 1;
//     float LONGITUDE = 2;
//     float LATITUDE = 3;
//     float AIR_TEMPERATURE = 4;
//     float PRECIPITATION = 5;
//     float SOLAR_RADIATION = 6;
//     float SURFACE_TEMPERATURE = 7;
//     int64 RELATIVE_HUMIDITY = 8;

func main() {
	// readFile("../../data/test.txt")
	parseLine("63892 20180101 0005 20171231 1805      5  -88.14   32.84    -1.0     0.0 -99999 0 -9999.0 U 0 -9999 0 -99.000 -9999.0  1212 0 -99.00 0")

}
