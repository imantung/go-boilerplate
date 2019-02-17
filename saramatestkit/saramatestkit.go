package saramatestkit

import (
	"bou.ke/monkey"
	"github.com/Shopify/sarama"
)

func PatchNewClient(client sarama.Client, err error) *monkey.PatchGuard {
	return monkey.Patch(sarama.NewClient, func(addrs []string, conf *sarama.Config) (sarama.Client, error) {
		return client, err
	})
}

func PatchNewSyncProducer(producer sarama.SyncProducer, err error) *monkey.PatchGuard {
	return monkey.Patch(sarama.NewSyncProducer, func(addrs []string, config *sarama.Config) (sarama.SyncProducer, error) {
		return producer, err
	})
}
