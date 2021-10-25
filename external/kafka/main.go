package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

func producerDemo() {
	config := sarama.NewConfig()
	// ack 确保消息写入 kafka 并落盘，保证数据不会丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 设置数据落地的分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//
	config.Producer.Return.Successes = true

	/**
	实例化一个生产者
	sarama.NewSyncProducer() 	同步生产者
	sarama.NewAsyncProducer()	异步生产者
	*/
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)

	if err != nil {
		fmt.Println("producer create err:", err)
		return
	}
	defer client.Close()

	/**
	构建一条 kafka 消息
	*/
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	// 执行发送
	pid, offset, err := client.SendMessage(msg) // pid 数据落盘的分区id，offset 写入后的偏移量
	if err != nil {
		fmt.Println("producer send msg err:", err)
		return
	}
	fmt.Printf("pic = %v, offset = %v\n", pid, offset)
}

func consumerDemo() {
	consumer, err := sarama.NewConsumer(strings.Split("127.0.0.1:9092", ","), nil)
	if err != nil {
		fmt.Printf("failed to start consumer: %v", err)
		return
	}
	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition: %d, offset: %d, key: %s, value: %s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}

	time.Sleep(time.Hour)
	consumer.Close()
}

func main() {
	// 生产 demo
	producerDemo()

	// 消费 demo
	consumerDemo()
}
