package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
)

func main() {
	filename := "token"
	fp, _ := os.Open(filename)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	consumerKey := scanner.Text()
	scanner.Scan()
	consumerSecret := scanner.Text()
	scanner.Scan()
	accessToken := scanner.Text()
	scanner.Scan()
	accessSecret := scanner.Text()

	sample := `{"text":"hogehoge3"}`

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// http.Client will automatically authorize Requests
	client := config.Client(oauth1.NoContext, token)
	req, _ := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer([]byte(sample)))
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))
}
