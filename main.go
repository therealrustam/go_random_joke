package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://uselessfacts.jsph.pl/random.json?language=en"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}
	fmt.Println(data["text"])

}
