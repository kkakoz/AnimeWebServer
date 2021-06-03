package kafka_runner

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/pkg/msg"
	"red-bean-anime-server/pkg/kafkax"
)

// 计数消息消费方法
func NewCountConsumerRunner(countUsecase domain.ICountUsecase)  kafkax.ConsumerRunFunc{
	return func(message *sarama.ConsumerMessage) error {
		var countMsg = &msg.CountMsg{}
		err := json.Unmarshal(message.Value, countMsg)
		if err != nil {
			return err
		}
		switch countMsg.MsgType {
		case msg.MsgTypeLike:
			return countUsecase.UserLikeAnime(countMsg.UserId, countMsg.AnimeId, true)
		case msg.MsgTypeLikeCancel:
			return countUsecase.UserLikeAnime(countMsg.UserId, countMsg.AnimeId, false)
		case msg.MsgTypeUnlike:
			return countUsecase.UserUnLikeAnime(countMsg.UserId, countMsg.AnimeId, true)
		case msg.MsgTypeUnlikeCancel:
			return countUsecase.UserUnLikeAnime(countMsg.UserId, countMsg.AnimeId, true)
		case msg.MsgTypeAddView:
			return countUsecase.AddAnimeView(countMsg.AnimeId)
		case msg.MsgTypeCollect:
			return nil
		case msg.MsgTypeCollectCancel:
			return nil
		default:
			return nil
		}
	}
}

var CountConsumerSet = wire.NewSet(NewCountConsumerRunner)