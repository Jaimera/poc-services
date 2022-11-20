package broker

import (
	"encoding/json"
	"github.com/jaimera/poc-services/domain/contract"
	"github.com/jaimera/poc-services/domain/dto"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

type KafkaConnection struct {
	log       *logrus.Entry
	host      string
	topic     string
	partition int
	network   string
}

func NewKafkaConnection(
	log *logrus.Entry,
	host string,
	topic string,
	partition int,
	network string,
) KafkaConnection {
	return KafkaConnection{
		log:       log,
		host:      host,
		topic:     topic,
		partition: partition,
		network:   network,
	}
}

func (c KafkaConnection) Produce(ctx context.Context, port dto.PortDto) error {
	conn, err := kafka.DialLeader(context.Background(), c.network, c.host, c.topic, c.partition)
	if err != nil {
		c.log.Errorf("failed to dial leader: %v, host: %v", err, c.host)
		return err
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		c.log.Errorf("failed to set write deadline: %v", err)
		return err
	}

	bytes, err := json.Marshal(port)
	if err != nil {
		c.log.Errorf("failed to marshal port: %v", err)
		return err
	}

	_, err = conn.WriteMessages(
		kafka.Message{Value: bytes},
	)
	if err != nil {
		c.log.Errorf("failed to write messages: %v", err)
		return err
	}

	if err := conn.Close(); err != nil {
		c.log.Fatalf("failed to close writer: %v", err)
		return err
	}

	return nil
}

func (c KafkaConnection) Consume(ctx context.Context, portService contract.PortService) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{c.host},
		Topic:     c.topic,
		Partition: c.partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	atInterruption(func() { r.Close() })

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		var dto dto.PortDto
		err = json.Unmarshal(m.Value, &dto)
		if err != nil {
			c.log.Errorf("Failed to unmarshal port %v", err)
		}
		err = portService.Insert(ctx, dto)
		if err != nil {
			c.log.Errorf("Failed to insert port %v", err)
		}
		c.log.Infof("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		c.log.Fatalf("failed to close reader: %v", err)
		return err
	}

	return nil
}

func atInterruption(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fn()
		os.Exit(0)
	}()
}
