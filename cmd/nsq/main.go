package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	for {
		/* if err = producer.Publish("topicname", []byte("hello world"+strconv.Itoa(rand.Int()))); err != nil {
			fmt.Println(err)
		} */
		if err = producer.Publish("nsqtest", []byte(strconv.Itoa(rand.Int()))); err != nil {
			fmt.Println(err)
		}
	}
	//gracefully stop the producer.
	//producer.Stop()
}

//may be given multiple times
