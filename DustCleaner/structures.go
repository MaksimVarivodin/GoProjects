package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const botToken string = "5496691862:AAEiqlkAbjfz2RzLftJPFs6eCR3EoruO2wQ"
const botApi string = "https://api.telegram.org/bot"

var botUrl string = botApi + botToken

func main() {
	iter := 0
	/*
	 getting all of the previous messages together(if any)
	 and deleting all of the gif's
	*/
	// here is getting all of the info in the bot
	oldUpd, offset, err1 := getOffset()
	if err1 != nil {
		log.Println("Error", err1.Error())
	}
	// here we're deleting all gif's
	//(but they stay for some time at the server)
	iter = delInUpd(oldUpd, iter)
	/*
		 infinite cycle which gets updates and deletes
		all of the gif's recently sent
	*/
	for {
		// getting updates
		upd, err := getUpdates(offset)
		if err != nil {
			log.Println("Error", err.Error())
		}
		/*
			if the updates array didn't change,
			 we check again
		*/
		if len(oldUpd) == len(upd) {
			continue
		} else {
			oldUpd = upd
		}
		/*
			if the check says that array has bigger size
			we check most recent messages if any of them contain gif's
		*/
		iter = delInUpd(upd, iter)
		_, offset, err1 = getOffset()
		if err1 != nil {
			log.Println("Error", err1.Error())
		}
	}

}

func getUpdates(offset int) ([]Update, error) {

	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		fmt.Println("Error while getting, file: strustures.go,  line: 27", err.Error())
		fmt.Scanln()
		return nil, err
	}
	defer resp.Body.Close() // я так понимаю удаление переменной
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading, file: strustures.go,  line: 34", err.Error())
		fmt.Scanln()
		return nil, err
	}
	var arr UpdateArrayFromResponse
	err = json.Unmarshal(body, &arr)
	if err != nil {
		fmt.Println("Error while unmarshaling, file: strustures.go,  line: 41", err.Error())
		fmt.Scanln()
		return nil, err
	}
	return arr.Array, nil
}

func getOffset() ([]Update, int, error) {
	offset := 0
	upd, err := getUpdates(offset)
	if err != nil {
		return nil, 0, err
	}
	if len(upd) > 0 {
		offset = upd[len(upd)-1].Message.MessageId + 1
	} else {
		offset = 0
	}
	return upd, offset, nil
}

func delInUpd(arr []Update, iter int) int {
	for iter < len(arr) {
		var mes Message
		mes.Chat = new(Chat)
		mes.Chat.ChatId = arr[iter].Message.Chat.ChatId
		arr[iter].Message.DeleteGifs(botUrl)
		iter += 1
	}
	return iter
}
