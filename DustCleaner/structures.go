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
	oldUpd, offset, err1 := getOffset()
	if err1 != nil {
		log.Println("Error", err1.Error())
	}
	iter = delInUpd(oldUpd, iter)
	for {

		upd, err := getUpdates(offset)
		if err != nil {
			log.Println("Error", err.Error())
		}
		if len(oldUpd) == len(upd) {
			continue
		} else {
			//iter = len(oldUpd) - 1
			oldUpd = upd
		}
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
	offset = upd[len(upd)-1].Message.MessageId + 1
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
