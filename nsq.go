package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
)


func req(){
	
}

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func testNSQ() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9
		fmt.Errorf()
		for _, var := range var {
			
		}

		//建立多个连接
		/* 	for i := 0; i < 5; i++ {
			consumer, err := nsq.NewConsumer("topicname", "zj", config) //topic， channel， config
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD("127.0.0.1:4150")
			if nil != err {
				fmt.Println("err", err)
				return
			}
		} */

		for i := 0; i < 5; i++ {
			consumer, err := nsq.NewConsumer("nsqtest", "nsq", config) //topic， channel， config
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD("127.0.0.1:4150")
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}

	}()

	waiter.Wait()
}
func main() {
	testNSQ()

}

//may be given multiple times
