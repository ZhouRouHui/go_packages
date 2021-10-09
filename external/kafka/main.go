package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	// ack 确保消息写入 kafka 并落盘，保证数据不会丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 设置数据落地的分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//
	config.Producer.Return.Successes = true

	/**
	构建一条 kafka 消息
	*/
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	/**
	实例化一个生产者
	vsarama.NewSyncProducer() 	同步生产者
	sarama.NewAsyncProducer()	异步生产者
	*/
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)

	if err != nil {
		fmt.Println("producer create err:", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg) // pid 数据落盘的分区id，offset 写入后的偏移量
	if err != nil {
		fmt.Println("producer send msg err:", err)
		return
	}
	fmt.Printf("pic = %v, offset = %v\n", pid, offset)
}
