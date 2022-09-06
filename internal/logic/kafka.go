package logic

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type kafka struct {
	client sarama.SyncProducer
}

func Initkafka(conf string) *kafka {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Metadata.AllowAutoTopicCreation = true
	client, err := sarama.NewSyncProducer([]string{conf}, config)
	if err != nil {
		log.Fatalln("producer closed, err:", err)
		return nil
	}
	log.Print("connect kafka success:", conf)
	return &kafka{client: client}
}
func (ka *kafka) push(req httpreq, topic string) error {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	value, err := json.Marshal(req)
	if err != nil {
		log.Printf("marshal err")
		return err
	}
	msg.Value = sarama.ByteEncoder(value)
	log.Print("kafka push:", msg)
	_, _, err = ka.client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return err
	}
	return nil
}
func (ka *kafka) close() {
	ka.client.Close()
}
