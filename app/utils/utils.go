package utils

import (
	"net/http"
	"io/ioutil"
	"time"
	"fmt"
	"log"
	"encoding/json"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string) []Product {
    r, err := myClient.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    defer r.Body.Close()
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the responseData to type string

    var responseObject []Product
    json.Unmarshal(responseData, &responseObject)

    // for i := 0; i < len(responseObject); i++ {
    //     fmt.Println(responseObject[i])
    // }
	// sb := string(responseData)

    return responseObject
}