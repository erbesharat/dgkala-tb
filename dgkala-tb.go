package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/mamal72/dgkala"
)

func checkerr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func task() {
	godotenv.Load()
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	offers, err := dgkala.SpecialOffers()
	checkerr(err)
	for _, item := range offers {
		tweetText := item.ProductTitleFa + "\n" + "قیمت: " + strconv.Itoa(int(item.Price)) + "ریال"
		_, _, err := client.Statuses.Update(tweetText, nil)
		checkerr(err)
	}

}

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Days().Do(task)
	<-s.Start()
}
