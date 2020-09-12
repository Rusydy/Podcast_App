package main

import (
	"Podcast_App/itunes"
	"log"
)

func main() {
	ias := itunes.NewItunesApiServices()

	res, err := ias.Search("Full Stack Radio")

	if err != nil {
		log.Fatalf("error while searching: %v", err)
	}

	for _, item := range res.Results {
		log.Println("------------------")
		log.Printf("Artist: %s", item.ArtistName)
		log.Printf("Podcast Name: %s", item.TrackName)
		log.Printf("Feed url: %s", item.FeedURL)
		log.Println("------------------")
	}
}
