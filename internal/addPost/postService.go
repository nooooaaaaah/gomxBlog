package addpost

import (
	"Blog/pkg/db"
	"context"

	"github.com/google/uuid"
)

type PostService struct {
	queries db.Queries
}

func NewPostSerivce(queries db.Queries) *PostService {
	return &PostService{queries: queries}
}

func (p *PostService) MakePost(post db.Post) error {
	_, err := p.queries.CreatePost(context.Background(), db.CreatePostParams{
		ID:          uuid.New(),
		Title:       post.Title,
		Content:     post.Content,
		Description: post.Description,
		Link:        post.Link,
		PublishedOn: post.PublishedOn,
	})
	if err != nil {
		return err
	}
	return nil
}
