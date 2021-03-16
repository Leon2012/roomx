package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gotomicro/ego-component/eredis"
	"roomx/model"
	"roomx/components/logrus"
)

const ROOMX_SORTEDSET_PREFIX = "roomx.messages.sortedset"
const ROOMX_HASH_PREFIX = "roomx.message.hash"
const ROOMX_MESSAGES_MAX = 20000
const ROOMX_MESSAGES_PAGE_NUM = 5


func MessagesNext(rds *eredis.Component, uid, rid, currId int32, currMessage *model.Message) (model.Messages, int32, error) {
	var (
		mMessages, messages model.Messages
		xMessage            *model.Message
		err                 error
		nextId              int32
		pageNum             int64 = ROOMX_MESSAGES_PAGE_NUM + 1
		key                 string
		rank, start, stop         int64
		ctx                 = context.Background()
		items                []string
		item                 string
		total int64
	)
	key = genSortedsetKey(rid)
	//第一次返回数据
	if currId == 0 {
		if total, err = rds.ZCard(ctx, key); err != nil {
			return messages, nextId, err
		}
		start = total - pageNum
		stop = total

		if items, err = rds.ZRange(ctx, key, start, stop); err != nil {
			return messages, nextId, err
		}
		mMessages, err = model.NewMessagesWithBase64String(items)
		if len(mMessages) > int(ROOMX_MESSAGES_PAGE_NUM) {
			messages = mMessages[0:ROOMX_MESSAGES_PAGE_NUM]
			xMessage = mMessages[len(mMessages)-1]
			nextId = xMessage.Id
		} else {
			messages = mMessages
			xMessage = mMessages[len(mMessages)-1]
			nextId = xMessage.Id
		}
	} else {
		item, err = currMessage.ToBase64String()
		logrus.Info("current message : %s", item)
		if err != nil {
			return messages, nextId, err
		}
		rank, err = rds.Client().ZRank(ctx, key, item).Result()
		if err != nil {
			return messages, nextId, err
		}
		logrus.Info("get message rank : %d", rank)
		start = rank
		stop = rank + ROOMX_MESSAGES_PAGE_NUM
		if items, err = rds.ZRange(ctx, key, start, stop); err != nil {
			return messages, nextId, err
		}
		mMessages, err = model.NewMessagesWithBase64String(items)
		if len(mMessages) > int(ROOMX_MESSAGES_PAGE_NUM) {
			messages = mMessages[0:ROOMX_MESSAGES_PAGE_NUM]
			xMessage = mMessages[len(mMessages)-1]
			nextId = xMessage.Id
		} else {
			messages = mMessages
			xMessage = mMessages[len(mMessages)-1]
			nextId = xMessage.Id
		}
	}

	return messages, nextId, err
}

func MessageAdd(rds *eredis.Component, message *model.Message) (int64, error) {
	var (
		data                 string
		err                  error
		key                  string
		val                  *redis.Z
		id, cnt, start, stop int64
		ctx                  = context.Background()
	)
	data, err = message.ToBase64String()
	if err != nil {
		return 0, err
	}
	key = genSortedsetKey(message.Rid)
	val = &redis.Z{
		Score:  float64(message.Id),
		Member: data,
	}
	cnt, err = rds.ZCard(ctx, key)
	if err != nil {
		return 0, err
	}
	if cnt >= ROOMX_MESSAGES_MAX {
		start = 0
		stop = cnt - ROOMX_MESSAGES_MAX
		cnt, err = rds.ZRemRangeByRank(ctx, key, start, stop)
		if err != nil {
			return 0, err
		}
	}

	id, err = rds.ZAdd(ctx, key, val)
	if err != nil {
		return 0, err
	}

	return id, err
}

func genSortedsetKey(rid int32) string {
	var (
		key string
	)
	key = fmt.Sprintf("%s.%d", ROOMX_SORTEDSET_PREFIX, rid)
	return key
}

func genHashKey(rid int32) string {
	var (
		key string
	)
	key = fmt.Sprintf("%s.%d", ROOMX_HASH_PREFIX, rid)
	return key
}
