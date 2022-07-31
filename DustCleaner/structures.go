package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const botToken string = "5484587003:AAGAUGUB0_A6cnvLEBOAlDn_yYIkefteJH8"
const botApi string = "https://api.telegram.org/bot"

var botUrl string = botApi + botToken

func main() {
	for {
		updates, err := getUpdates()
		if err != nil {
			log.Println("Error", err.Error())
		}
		fmt.Println(updates)
	}

}
func getUpdates() ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
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
