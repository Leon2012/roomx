package mysql

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"roomx/model"
	"roomx/sendsrv/pkg/invoker"
)

func MessageCreate(db *gorm.DB, data *model.Message) (err error) {
	if err = db.Create(data).Error; err != nil {
		invoker.Logger.Error("create user error", zap.Error(err))
		return
	}
	return
}

func MessageGet(db *gorm.DB, id int32) (message *model.Message, err error) {
	var m model.Message
	result := db.First(&m, id)
	if result.Error != nil {
		return nil, result.Error
	}
	message = &m
	return
}
