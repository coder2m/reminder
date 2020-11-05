/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 17:01
 */
package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

//连接信息
// 用户名 密码 ip:端口/虚拟机
//const MQURL = "amqp://admin:admin@127.0.0.1:5672/testMQ"

//rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

//创建结构体实例
func newRabbitMQ(queueName string, exchange string, key string, options *Options) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: fmt.Sprintf("amqp://%v:%v@%v:%v/%v", options.Username, options.Password, options.Host, options.Port, options.Name)}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

//创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string, options *Options) (rabbitmq *RabbitMQ, err error) {
	//创建RabbitMQ实例
	rabbitmq = newRabbitMQ(queueName, "", "", options)
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		return
	}
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return
	}
	return
}

//直接模式队列生产
func (r *RabbitMQ) PublishSimple(message []byte) (err error) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err = r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return
	}
	//调用channel 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	return
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple(f func([]byte) error, stopCh <-chan struct{}) (err error) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return
	}

	//接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return
	}

	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			_ = f(d.Body)
		}
	}()
	<-stopCh
	return
}

//订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string, options *Options) (rabbitmq *RabbitMQ, err error) {
	//创建RabbitMQ实例
	rabbitmq = newRabbitMQ("", exchangeName, "", options)
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		return
	}
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return
	}
	return
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message []byte) (err error) {
	//1.尝试创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	return err
}

//订阅模式消费端代码
func (r *RabbitMQ) RecieveSub(f func([]byte) error, stopCh <-chan struct{}) (err error) {
	//1.试探性创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil)

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	go func() {
		for d := range messges {
			_ = f(d.Body)
		}
	}()

	<-stopCh

	return err
}

//路由模式
//创建RabbitMQ实例
func NewRabbitMQRouting(exchangeName string, routingKey string, options *Options) (rabbitmq *RabbitMQ, err error) {
	//创建RabbitMQ实例
	rabbitmq = newRabbitMQ("", exchangeName, routingKey, options)
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		return
	}
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return
	}
	return
}

//路由模式发送消息
func (r *RabbitMQ) PublishRouting(message []byte) (err error) {
	//1.尝试创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//要改成direct
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return
	}

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		//要设置
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	return
}

//路由模式接受消息
func (r *RabbitMQ) RecieveRouting(f func([]byte) error, stopCh <-chan struct{}) (err error) {
	//1.试探性创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return
	}

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//需要绑定key
		r.Key,
		r.Exchange,
		false,
		nil)

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range messges {
			_ = f(d.Body)
		}
	}()

	<-stopCh

	return
}

//话题模式
//创建RabbitMQ实例
func NewRabbitMQTopic(exchangeName string, routingKey string, options *Options) (rabbitmq *RabbitMQ, err error) {
	//创建RabbitMQ实例
	rabbitmq = newRabbitMQ("", exchangeName, routingKey, options)
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		return
	}
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return
	}
	return
}

//话题模式发送消息
func (r *RabbitMQ) PublishTopic(message []byte) (err error) {
	//1.尝试创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//要改成topic
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return
	}

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		//要设置
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	return
}

//话题模式接受消息
//要注意key,规则
//其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
//匹配 myxy99.* 表示匹配 myxy99.hello, 但是myxy99.hello.one需要用myxy99.#才能匹配到
func (r *RabbitMQ) RecieveTopic(f func([]byte) error, stopCh <-chan struct{}) (err error) {
	//1.试探性创建交换机
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return
	}

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		r.Key,
		r.Exchange,
		false,
		nil)

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	go func() {
		for d := range messges {
			_ = f(d.Body)
		}
	}()
	<-stopCh
	return
}
