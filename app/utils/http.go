package utils

import (
	"net/http"
	"time"
	"bytes"
	"log"
	"io/ioutil"
	"errors"
	"strconv"
	
	"encoding/json"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	client := http.Client{}
	req , err := http.NewRequest("GET", url, nil)
    if err != nil {
		return err
    }
	req.Header = http.Header{
		"Authorization": []string{"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjJmMjYyMWQwNzE4MGFlODEyZDE4OTNiMzk4YmU0NDJiZTBjMzQ4ODVlODc5MGU0MTk0MjhlMTEwODI1MzY2YjkzNGM3NTA1Yzk4ZDgwZjk4In0.eyJhdWQiOiIyIiwianRpIjoiMmYyNjIxZDA3MTgwYWU4MTJkMTg5M2IzOThiZTQ0MmJlMGMzNDg4NWU4NzkwZTQxOTQyOGUxMTA4MjUzNjZiOTM0Yzc1MDVjOThkODBmOTgiLCJpYXQiOjE2MzQ5MTQyMzQsIm5iZiI6MTYzNDkxNDIzNCwiZXhwIjoxNjY2NDUwMjM0LCJzdWIiOiIyMzg0Iiwic2NvcGVzIjpbXX0.cY_a9s-Qe2N1A6B13qxwKEIcx6jBtUA2f7grqoB750ZvtOQng8hY33jyEtmV-JXYbO_s7m5JqeRrTMBbXbia8JtmPfclbdJvLwWset6EDo-9f6v0buMbqe9vNxnxcbq2_1Vw8ivo0xru_CfNbPLPV4-O7DBZONZVx3_rf-qOEOjwW-nwb2Dql7igpzg_UuwwY1N9Mn4hfxHw2efacXfHqqY18DXrBOoKeOCuCXIM3vlAkeygPTXz_oXaasgEd9d7fyg0yC-SeK-vEaf47czdhiEcM_JlQZom2YFa-sRFOL5c3t9ERcG_Ta2XhWqlXw9R_VU5IDS3KNv_8QQPVUP0mH-7QVS843MXpoAnj9KoSkLeKjTMTfwkM60qifz2EVPKPJGzc5qIuIDBnFE1tYIEcAUFm0GqNLFz3VCnwBPSF95-5nebSyguGXhNc-WbX8Kup5Zpn-SXKpHZuZiH9vYDkpzdoUZnTwsqIvTYPEaVBVf-v_a_xLZrkQNGOA_8JUaw7Wd06duo8ShOi7UDKe9BxZWyrdScW0KnaYY_UJ1mByH_-lPxV6h8uUafmJwe5fs2cc05-I0YtbGsPd4Qh2_j5_msDds6n2VpQOUN9vmf9dur8PER64bHMelVsnDCWbKzZBY3Krc3Q8hN9bhFaepR_tkxc1TsGD5DRdotsfRNt2c"},
	}
	res , err := client.Do(req)
    if err != nil {
		return err
    }
    defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		response := res.Body
		// == FOR CHECK THE ERROR: import "io/ioutil", "fmt"
		// bodyBytes, _ := ioutil.ReadAll(res.Body)
		// fmt.Println(string(bodyBytes))
	
		return json.NewDecoder(response).Decode(target)
	} else {
		data := res.Body
		body, err := ioutil.ReadAll(data)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		return errors.New(strconv.Itoa(res.StatusCode) + ": " + sb + " for " + url)
	}
}

func PostJson(url string, target interface{}, postBody []byte, token string) error {
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Add("Authorization","bearer " + token) 
	if err != nil {
		return err
	}
    resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// ==== Activate it only for debug only; import "io/ioutil" ====
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// sb := string(body)
		// log.Printf(sb)
	
		return json.NewDecoder(resp.Body).Decode(target)
	} else {
		data := resp.Body
		body, err := ioutil.ReadAll(data)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		return errors.New(strconv.Itoa(resp.StatusCode) + ": " + sb + " for " + url)
	}
}

func UpdateJson(url string, target interface{}, postBody []byte, token string) error {
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("UPDATE", url, requestBody)
	req.Header.Add("Authorization","bearer " + token) 
	if err != nil {
		return err
	}
    resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// ==== Activate it only for debug only; import "io/ioutil" ====
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// sb := string(body)
		// log.Printf(sb)
	
		return json.NewDecoder(resp.Body).Decode(target)
	} else {
		data := resp.Body
		body, err := ioutil.ReadAll(data)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		return errors.New(strconv.Itoa(resp.StatusCode) + ": " + sb + " for " + url)
	}
}

func DeleteJson(url string, target interface{}, postBody []byte, token string) error {
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("DELETE", url, requestBody)
	req.Header.Add("Authorization","bearer "+token) 
	if err != nil {
		return err
	}
    resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// ==== Activate it only for debug only; import "io/ioutil" ====
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// sb := string(body)
		// log.Printf(sb)
	
		return json.NewDecoder(resp.Body).Decode(target)
	} else {
		data := resp.Body
		body, err := ioutil.ReadAll(data)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		return errors.New(strconv.Itoa(resp.StatusCode) + ": " + sb + " for " + url)
	}
}