package main

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
	SenderChat  *Chat      `json:"chat"`
	Animation   *Animation `json:"animation"`
}
type Update struct {
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type UpdateArrayFromResponse struct {
	Array []Update `json:"result"`
}

func (anim *Animation) IsAnimation() bool {
	if anim != nil {
		return true
	} else {
		return false
	}
}

func (chat *Chat) DeleteGifs(message Message) bool {
	if message.Animation.IsAnimation() {

		return true
	} else {
		return false
	}

}
