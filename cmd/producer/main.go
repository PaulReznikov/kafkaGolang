package main

//import (
//	"fmt"
//	"github.com/sirupsen/logrus"
//	k "kafkaGolangProducer/internal/kafka"
//)
//
//const (
//	topic = "my-topic"
//)
//
//var address = []string{"localhost:9092", "localhost:9093", "localhost:9094"}
//
//func main() {
//	p, err := k.NewProducer(address)
//	if err != nil {
//		logrus.Fatal(err)
//	}
//
//	for i := 0; i < 100; i++ {
//		msg := fmt.Sprintf("Kafka message %d", i)
//		if err = p.Produce(msg, topic); err != nil {
//			logrus.Error(err)
//		}
//	}
//
//}

/*
#include <stdio.h>
void hello() { printf("Hello from C!\n"); }
*/
import "C"

func main() {
	C.hello()
}
