package job

import (
	"context"

	"geim/internal/job/conf"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

// sarama 库中消费者组为一个接口 sarama.ConsumerGroup 所有实现该接口的类型都能当做消费者组使用。

// MyConsumerGroupHandler 实现 sarama.ConsumerGroup 接口，作为自定义ConsumerGroup
type MyConsumerGroupHandler struct {
	name  string
	count int64
}

// Setup 执行在 获得新 session 后 的第一步, 在 ConsumeClaim() 之前
func (MyConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup 执行在 session 结束前, 当所有 ConsumeClaim goroutines 都退出时
func (MyConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim 具体的消费逻辑
func (h MyConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("[consumer] name:%s topic:%q partition:%d offset:%d\n", h.name, msg.Topic, msg.Partition, msg.Offset)
		// 标记消息已被消费 内部会更新 consumer offset
		sess.MarkMessage(msg, "")
		go handleRoomMsg(msg.Value)
		h.count++
		if h.count%10000 == 0 {
			log.Printf("name:%s 消费数:%v\n", h.name, h.count)
		}
	}
	return nil
}

//"kafka-server:9092"
func ConsumerGroup(conf *conf.RoomKafka) {
	addr := conf.Addr
	group := conf.Group
	name := conf.Name
	topic := conf.Topic
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cg, err := sarama.NewConsumerGroup([]string{addr}, group, config)
	if err != nil {
		log.Fatal("NewConsumerGroup err: ", err)
	}
	defer cg.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler := MyConsumerGroupHandler{name: name}
		for {
			log.Println("running: ", name)
			/*
				![important]
				应该在一个无限循环中不停地调用 Consume()
				因为每次 Rebalance 后需要再次执行 Consume() 来恢复连接
				Consume 开始才发起 Join Group 请求 如果当前消费者加入后成为了 消费者组 leader,则还会进行 Rebalance 过程，从新分配
				组内每个消费组需要消费的 topic 和 partition，最后 Sync Group 后才开始消费
				具体信息见 https://github.com/lixd/kafka-go-example/issues/4
			*/
			err = cg.Consume(ctx, []string{topic}, handler)
			if err != nil {
				log.Println("Consume err: ", err)
			}
			// 如果 context 被 cancel 了，那么退出
			if ctx.Err() != nil {
				return
			}
		}
	}()
	wg.Wait()
}
