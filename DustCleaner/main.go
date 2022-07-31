package main

type User struct {
	userId    int    `json:"id"`
	isBot     bool   `json:"is_bot"`
	firstName string `json:"first_name"`
}
type Chat struct {
	chatId    int    `json:"id"`
	chatType  string `json:"type"`
	chatTitle string `json:"title"`
}
type Animation struct {
	fileId       string `json:"file_id"`
	fileUniqueId string `json:"file_unique_id"`
	width        int    `json:"width"`
	height       int    `json:"height"`
	duration     int    `json:"duration"`
}

type Message struct {
	messageId   int        `json:"message_id"`
	messageFrom *User      `json:"from"`
	senderChat  *Chat      `json:"chat"`
	animation   *Animation `json:"animation"`
}
type Update struct {
	updateId int      `json:"update_id"`
	message  *Message `json:"message"`
}

type UpdateArrayFromResponse struct {
	array []Update `json:"result"`
}

func (anim *Animation) isAnimation() bool {
	if anim != nil {
		return true
	} else {
		return false
	}
}

func (chat *Chat) deleteGifs(message Message) bool {
	if message.animation.isAnimation() {

		return true
	} else {
		return false
	}

}
