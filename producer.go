package main

// 来自 https://www.cnblogs.com/haima/p/13953154.html
import (
	"fmt"

	"github.com/Shopify/sarama"
)

var Topic = "web_log" //主题名称

// 基于sarama第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"182.92.234.24:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	//例子一发单个消息
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = Topic
	content := "this is a test log"
	send01(client, msg, content)

	//例子二发多个消息
	for _, word := range []string{"Welcome11", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		send01(client, msg, word)
	}
}

// 发消息
func send01(client sarama.SyncProducer, msg *sarama.ProducerMessage, content string) {
	msg.Value = sarama.StringEncoder(content)

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}
