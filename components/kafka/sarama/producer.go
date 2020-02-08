package main


import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，默认设置8个分区
	config.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{}
	msg.Topic = `nginx_log`
	msg.Value = sarama.StringEncoder("this is a good test1")
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close err, ", err)
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)

	if err != nil {
		fmt.Println("send message failed, ", err)
		return
	}
	fmt.Printf("分区ID:%v, offset:%v \n", pid, offset)
}
