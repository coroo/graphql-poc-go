package utils

import (
	"net/http"
	"time"
	"bytes"
	"log"
	// "io/ioutil"
	"encoding/json"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func PostJson(url string, target interface{}, postBody []byte) error {
	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	// ==== Activate it only for debug only ====
	// body, err := ioutil.ReadAll(resp)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// sb := string(body)
	// log.Printf(sb)

    return json.NewDecoder(resp.Body).Decode(target)
}