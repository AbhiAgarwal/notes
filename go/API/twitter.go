package main

import (
	"fmt"
	anaconda "github.com/ChimeraCoder/anaconda"
	"github.com/abhiagarwal/goconf"
	"net/url"
)

func getConfiguration() map[string]interface{} {
	var parsedData map[string]interface{}
	parsedData = make(map[string]interface{})
	parsedData = goconf.Parse("twitter.conf")
	return parsedData
}

func main() {
	var configuration map[string]interface{}
	configuration = make(map[string]interface{})
	configuration = getConfiguration()

	ConsumerKey, _ := configuration["ConsumerKey"].(string)
	ConsumerSecret, _ := configuration["ConsumerSecret"].(string)
	AccessToken, _ := configuration["AccessToken"].(string)
	AccessTokenSecret, _ := configuration["AccessTokenSecret"].(string)

	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

	v := url.Values{}
	v.Set("count", "100")

	searchResult, _ := api.GetSearch("nyu", v)
	for _, tweet := range searchResult {
		fmt.Println(tweet.Text)
	}
}
