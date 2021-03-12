package consumer

import (
	"errors"
	"roomx/libs/queue"
	"roomx/model"
	"roomx/model/redis"
	"roomx/recvsrv/pkg/invoker"
)

type Handler struct {
	payload *queue.Payload
}

func (h *Handler) Id() string {
	return ""
}

func (h *Handler) Perform() error {
	var (
		body []byte
		err  error
		data *model.Message
	)
	body = h.payload.Message.Body
	if len(body) == 0 {
		err = errors.New("empty data")
	}
	if data, err = model.NewMessage(body); err != nil {
		return err
	}
	invoker.Logger.Info("recv content : " + data.Content)
	_, err = redis.MessageAdd(invoker.Redis, data)
	if err != nil {
		return err
	}
	return err
}

func (h *Handler) SetPayload(p *queue.Payload) {
	h.payload = p
}
