package model

import (
	"encoding/base64"
	"encoding/json"
	"roomx/libs/bytesconv"
)

type Message struct {
	Id       int32  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Uid      int32  `json:"uid" binding:"required"`
	Rid      int32  `json:"rid" binding:"required"`
	Type     int32  `json:"type" binding:"required"`
	Content  string `gorm:"not null" json:"content" binding:"required"`
	Extra    string `json:"extra"`
	Dateline int64  `json:"dateline,omitempty" `
}

type Messages []*Message

func (t Message) TableName() string {
	return "messages"
}

func (t *Message) ToBytes() (body []byte, err error) {
	body, err = json.Marshal(t)
	return body, err
}

func (t *Message) ToString() (string, error) {
	var (
		bytes []byte
		err   error
	)

	if bytes, err = t.ToBytes(); err != nil {
		return "", err
	}
	return bytesconv.BytesToString(bytes), nil
}

func (t *Message) ToBase64String() (string, error) {
	var (
		bytes []byte
		err   error
	)

	if bytes, err = t.ToBytes(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func NewMessage(body []byte) (message *Message, err error) {
	var m Message
	if err = json.Unmarshal(body, &m); err != nil {
		return
	}
	message = &m
	return
}

func NewMessagesWithBase64String(items []string)  (messages Messages, err error){
	var (
		message *Message
		byte    []byte
	)
	for i := 0; i < len(items); i++ {
		item := items[i]
		byte, err = base64.StdEncoding.DecodeString(item)
		if err == nil {
			if message, err = NewMessage(byte); err == nil {
				messages = append(messages, message)
			}
		}

	}
	return
}

func NewMessages(items []string) (messages Messages, err error) {
	var (
		message *Message
		byte    []byte
	)
	for i := 0; i < len(items); i++ {
		byte = bytesconv.StringToBytes(items[i])
		if message, err = NewMessage(byte); err == nil {
			messages = append(messages, message)
		}
	}
	return
}
