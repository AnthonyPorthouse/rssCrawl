package main

import (
	"fmt"
	"github.com/AnthonyPorthouse/rssCrawl/feeds"
	"log"
	"time"
)

func main() {
	startTime := time.Now()

	feedURLs := []string{
		"https://cad-comic.com/feed/",
		"https://girlswithslingshots.com/rss.php",
		"http://feeds.feedburner.com/LookingForGroup?format=xml",
		"http://www.girlgeniusonline.com/ggmain.rss",
		"http://www.questionablecontent.net/QCRSS.xml",
	}

	for _, url := range feedURLs {
		feed, err := feeds.GetFeed(url)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Got Feed: %s\n", feed.Title)
		fmt.Printf("Total Items: %d\n", len(feed.Items))
	}

	fmt.Printf("Time Taken: %s\n", time.Since(startTime))
}
