package producer

import (
	"context"

	"github.com/swagchat/chat-api/config"
	scpb "github.com/swagchat/protobuf/protoc-gen-go"
)

type provider interface {
	PublishMessage(*scpb.EventData) error
}

func Provider(ctx context.Context) provider {
	cfg := config.Config()

	var p provider
	switch cfg.Producer.Provider {
	case "direct":
		p = &directProvider{
			ctx: ctx,
		}
	case "nsq":
		p = &nsqProvider{
			ctx: ctx,
		}
	case "kafka":
		p = &kafkaProvider{
			ctx: ctx,
		}
	default:
		p = &noopProvider{
			ctx: ctx,
		}
	}

	return p
}
