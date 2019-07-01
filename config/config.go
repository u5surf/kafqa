package config

import (
	"fmt"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Application struct {
	Producer
	Consumer
	Config
	Reporter
}

type Config struct {
	Environment string `default:"production"`
	DurationMs  int64  `split_words:"true" default:"10000"`
}

type Producer struct {
	Topic          string `default:"kafqa_test" envconfig:"KAFKA_TOPIC"`
	Concurrency    int    `default:"100"`
	TotalMessages  uint64 `split_words:"true" default:"10000"`
	KafkaBrokers   string `split_words:"true" required:"true"`
	FlushTimeoutMs int    `split_words:"true" default:"2000"`
}

type Consumer struct {
	// TODO: remove tags and load with split words while processing
	Topic         string `default:"kafqa_test" envconfig:"KAFKA_TOPIC"`
	Concurrency   int    `default:"20"`
	KafkaBrokers  string `split_words:"true" required:"true"`
	GroupID       string `split_words:"true" default:"kafqa_test_consumer"`
	OffsetReset   string `split_words:"true" default:"earliest"`
	PollTimeoutMs int64  `split_words:"true" default:"500"`
}

type Prometheus struct {
	Enabled bool `default:"false"`
	Port    int  `default:"9999"`
}

func (p Prometheus) BindPort() string {
	return fmt.Sprintf("0.0.0.0:%d", p.Port)
}

type Reporter struct {
	Prometheus
}

func App() Application {
	return application
}

func (a Application) RunDuration() time.Duration {
	return time.Duration(a.Config.DurationMs) * time.Millisecond
}

func (a Application) LogLevel() string {
	if a.DevEnvironment() {
		return "debug"
	}
	return "info"
}

func (a Application) DevEnvironment() bool {
	return a.Config.Environment == "development"
}

func (c Consumer) KafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		KafkaBootstrapServerKey: c.KafkaBrokers,
		ConsumerOffsetResetKey:  c.OffsetReset,
		ConsumerGroupIDKey:      c.GroupID,
	}
}

func (c Consumer) PollTimeout() time.Duration {
	return time.Duration(c.PollTimeoutMs) * time.Millisecond
}
