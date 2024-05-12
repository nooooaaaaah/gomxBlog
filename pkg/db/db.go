package db

import (
	"Blog/pkg/logger"
	"context"
	"log"

	edgedb "github.com/edgedb/edgedb-go"
)

var Client *edgedb.Client

func InitEdgeDB() {
	var err error
	// Assuming using DSN or environment variables to handle connection specifics
	Client, err = edgedb.CreateClient(context.Background(), edgedb.Options{})
	if err != nil {
		log.Fatalf("Failed to create EdgeDB client: %v", err)
	}
}

func CloseEdgeDB() {
	err := Client.Close()
	if err != nil {
		log.Printf("Failed to close EdgeDB client: %v", err)
	}
}

func GetPostByID(id edgedb.UUID) (*Post, error) {
	var post Post
	query := "SELECT Post { id, title, content, description, link, published_on } FILTER .id = <uuid>$id;"
	err := Client.QuerySingle(context.Background(), query, &post, id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPosts() ([]Post, error) {
	var posts []Post
	query := "Select Post { id, title, content, description, link, published_on }"
	err := Client.Query(context.Background(), query, &posts)
	if err != nil {
		logger.LogError.Println("getting posts from db failed, ", err)
		return nil, err
	}
	date, _ := posts[0].PublishedOn.Get()
	logger.LogInfo.Printf("GetPosts results: %v, %v", posts, date)
	return posts, nil
}
