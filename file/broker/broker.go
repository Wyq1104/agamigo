package main

import (
	"file/messages"
	"fmt"
	"log"
	"math/rand"
	"net"
)

type Item struct {
	airTemperature float32
	solarRadiation float32
	timestamp      int64
	priority       float64
	index          int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

var priorityQueue = make(PriorityQueue, 0)

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
		//time.Sleep(500 * time.Millisecond)
	}
}

func handleClimateData(msg *messages.Wrapper_ClimateData) {
	timestamp := msg.ClimateData.GetUTC_TIMESTAMP()
	airTemperature := msg.ClimateData.GetAIR_TEMPERATURE()
	solarRadiation := msg.ClimateData.GetSOLAR_RADIATION()
	log.Printf("Timestamp: %v", timestamp)
	log.Printf("Air Temperature: %v", airTemperature)
	log.Printf("Solar Radiation: %v", solarRadiation)

	doSample(airTemperature, solarRadiation, timestamp)
}

func doSample(airTemperature float32, solarRadiation float32, timestamp int64) {
	ran := rand.Float64()
	if len(priorityQueue) < 5 {
		priorityQueue.Push(&Item{airTemperature: airTemperature, solarRadiation: solarRadiation, timestamp: timestamp, priority: ran})
	} else {
		i := priorityQueue.Pop().(*Item)
		if i.priority < ran {
			priorityQueue.Push(&Item{airTemperature: airTemperature, solarRadiation: solarRadiation, timestamp: timestamp, priority: ran})
		} else {
			priorityQueue.Push(i)
		}
	}
	calculateSample()
}

func calculateSample() {

	log.Printf("Size of the PQ: %d\n", len(priorityQueue))

	var totalTemperature = 0.0
	var totalSolar = 0.0
	for _, item := range priorityQueue {
		totalTemperature += float64(item.airTemperature)
		totalSolar += float64(item.solarRadiation)
	}
	log.Printf("Aggregation Result:")
	log.Printf("Air Temperature: %v", float32(totalTemperature/float64(len(priorityQueue))))
	log.Printf("Solar Radiation: %v", float32(totalSolar/float64(len(priorityQueue))))
	sendToVisualize("Air Temperature", float32(totalTemperature/float64(len(priorityQueue))))
	sendToVisualize("Solar Radiation", float32(totalTemperature/float64(len(priorityQueue))))
}

func sendToVisualize(name string, value float32) {
	conn, err := net.Dial("tcp", "localhost:20112")
	if err != nil {
		log.Fatalln(err.Error())
	}
	msgHandler := messages.NewMessageHandler(conn)
	msg := messages.AggregateData{AGGREGATE_NAME: name, AGGREGATE_VALUE: value}
	msgHandler.Sendout(&msg)
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
