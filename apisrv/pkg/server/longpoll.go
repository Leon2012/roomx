package server

import (
	context2 "context"
	"math/rand"
	"net/http"
	"roomx/apisrv/pkg/invoker"
	"roomx/components/longpoll"
	"roomx/model"
	xmessage "roomx/proto/message"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/server"
	"roomx/apisrv/pkg/client"
)

var (
	Longpoll *longpoll.Component
	Nextseq  int32 = 0
)

func InitLongpoll() server.Server {
	InitGin()
	Longpoll = longpoll.Load("server.longpoll").Build(Gin)
	InitRouters()
	go RecvMessagesLoop()
	return Longpoll
}

func InitRouters() {
	Gin.GET("/basic", BasicExampleHomepage)
	Gin.GET("/recv", func(context *gin.Context) {
		seqStr := context.Query("next_seq")
		if seqStr != "" {
			seq, _ := strconv.ParseInt(seqStr, 10, 32)
			Nextseq = int32(seq)
		}
		invoker.Logger.Infof("next seq :%d", Nextseq)
		Longpoll.LongpollManager().SubscriptionHandler(context.Writer, context.Request)
	})
	Gin.POST("/send", func(context *gin.Context) {
		var (
			msg      model.Message
			err      error
			request  *xmessage.MessageSendReq
			response *xmessage.MessageSendResp
			ctx      context2.Context = context2.Background()
		)
		if err = context.ShouldBindJSON(&msg); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		request = &xmessage.MessageSendReq{
			Uid:     msg.Uid,
			Rid:     msg.Rid,
			Type:    msg.Type,
			Content: msg.Content,
			Extra:   msg.Extra,
		}

		response, err = client.SendClient.Send(ctx, request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"code": 0, "seq": response.Mid})
		return
	})
}

func RecvMessagesLoop() {
	var (
		ctx     context2.Context         = context2.Background()
		request *xmessage.MessageRecvReq = &xmessage.MessageRecvReq{
			Uid: 10,
			Rid: 10,
			Seq: Nextseq,
		}
		response *xmessage.MessageRecvResp
		err      error
	)
	for {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		request.Seq = Nextseq
		response, err = client.RecvClient.Recv(ctx, request)
		if err == nil {
			Longpoll.LongpollManager().Publish("farm", response)
		}
	}
}
