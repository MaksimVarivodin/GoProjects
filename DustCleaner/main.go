package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	UserId    int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
}

type Chat struct {
	ChatId    int    `json:"id"`
	ChatType  string `json:"type"`
	ChatTitle string `json:"title"`
}

type Animation struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Weight       int    `json:"height"`
	Duration     int    `json:"duration"`
}

type Message struct {
	MessageId   int        `json:"message_id"`
	MessageFrom *User      `json:"from"`
	Date        int        `json:"date"`
	Chat        *Chat      `json:"chat"`
	Text        string     `json:"text"`
	Animation   *Animation `json:"animation"`
}

type Update struct {
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type UpdateArrayFromResponse struct {
	Array []Update `json:"result"`
}

type InfoToDelete struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

type InfoToAdd struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

func (mess *Message) ToString() string {

	str := "MessageId: " + strconv.Itoa(mess.MessageId) + "\n"
	str += "MessageFrom: " + mess.MessageFrom.FirstName + "\n"
	i, err := strconv.ParseInt(strconv.Itoa(mess.Date), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)

	str += "Date: " + tm.Format("2006-01-02 15:04:05") + "\n"
	str += "Chat: " + mess.Chat.ChatTitle + "\n"

	str += "Text: " + mess.Text + "\n"
	str += "Animation: "
	if mess.Animation != nil {
		str += "true\n"
	} else {
		str += "false\n"
	}
	return str
}

func (anim *Animation) IsAnimation() bool {
	return anim != nil
}
func (mes *Message) SendMessage(botUrl string, text string) bool {
	var info InfoToAdd
	info.ChatId = mes.Chat.ChatId
	info.Text = text
	jsnfrmt, err := json.Marshal(info)
	if err != nil {
		return false
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(jsnfrmt))
	return err != nil
}
func (mes *Message) DeleteGifs(botUrl string) bool {
	if mes.Animation.IsAnimation() {
		mes.SendMessage(botUrl, "Deleted a message:\n"+mes.ToString())
		var delMssg InfoToDelete
		delMssg.ChatId = mes.Chat.ChatId
		delMssg.MessageId = mes.MessageId
		jsnfrmt, err := json.Marshal(delMssg)
		if err != nil {
			return false
		}
		_, err = http.Post(botUrl+"/deleteMessage", "application/json", bytes.NewBuffer(jsnfrmt))
		return err == nil
	} else {
		return false
	}
}
