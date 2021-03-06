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
	"time"
)

func readFile(filename string) {

	size := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		parseLine(scanner.Text())
		time.Sleep(100 * time.Millisecond)
		size = size + 1
	}

	fmt.Printf("file size is: %d", size)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func parseLine(line string) {
	words := strings.Fields(line)
	timestamp, _ := strconv.ParseInt(words[1], 10, 64)
	longitude, _ := strconv.ParseFloat(words[6], 32)
	latitude, _ := strconv.ParseFloat(words[7], 32)
	air_temperature, _ := strconv.ParseFloat(words[8], 32)
	precipitaion, _ := strconv.ParseFloat(words[9], 32)
	solar_radiation, _ := strconv.ParseFloat(words[10], 32)
	surface_temperature, _ := strconv.ParseFloat(words[12], 32)
	relative_humidity, _ := strconv.ParseInt(words[15], 10, 64)

	sendData(timestamp, float32(longitude), float32(latitude), float32(air_temperature), float32(precipitaion), float32(solar_radiation), float32(surface_temperature), relative_humidity)

}

func sendData(timestamp int64, longitude float32, latitude float32, air_temperature float32, precipitaion float32, solar_radiation float32,
	surface_temperature float32, relative_humidity int64) {
	nodename := "localhost"
	conn, err := net.Dial("tcp", nodename+":20111")
	if err != nil {
		log.Fatalln(err.Error())
	}
	msgHandler := messages.NewMessageHandler(conn)
	msg := messages.ClimateData{UTC_TIMESTAMP: timestamp, LONGITUDE: longitude, LATITUDE: latitude, AIR_TEMPERATURE: air_temperature,
		PRECIPITATION: precipitaion, SOLAR_RADIATION: solar_radiation, SURFACE_TEMPERATURE: surface_temperature, RELATIVE_HUMIDITY: relative_humidity}
	wrapper := &messages.Wrapper{
		Msg: &messages.Wrapper_ClimateData{ClimateData: &msg},
	}
	msgHandler.Send(wrapper)
}

func main() {
	readFile("../../data/test.txt")
	// parseLine("63892 20180101 0005 20171231 1805      5  -88.14   32.84    -1.0     0.0 -99999 0 -9999.0 U 0 -9999 0 -99.000 -9999.0  1212 0 -99.00 0")

}
