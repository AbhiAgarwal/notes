package main

import (
	"encoding/json"
	"fmt"
	anaconda "github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
)

type Configuration struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getConfiguration() Configuration {
	file, _ := os.Open("twitter-conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func main() {
	configuration := getConfiguration()
	anaconda.SetConsumerKey(configuration.ConsumerKey)
	anaconda.SetConsumerSecret(configuration.ConsumerSecret)
	api := anaconda.NewTwitterApi(configuration.AccessToken, configuration.AccessTokenSecret)

	v := url.Values{}
	v.Set("count", "100")

	searchResult, _ := api.GetSearch("nyu", v)
	for _, tweet := range searchResult {
		fmt.Println(tweet.Text)
	}
}
