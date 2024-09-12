package services

import (
	"Blog/internal/db"
	"Blog/pkg/logger"
	"context"
	"time"

	"github.com/google/uuid"
)

// BlogService implements the BlogService interface
type BlogService struct {
	queries     *db.Queries
	LastFetch   time.Time
	CachedPosts []db.Post
}

// NewBlogService creates a new blog service
func NewBlogService(queries *db.Queries) *BlogService {
	return &BlogService{queries: queries}
}

func (s *BlogService) GetAllPosts() ([]db.Post, error) {
	// Get the current year and week number
	year, week := time.Now().ISOWeek()

	// Check if the cached posts are from the current week
	yearLF, weekLF := s.LastFetch.ISOWeek()
	if yearLF == year && weekLF == week {
		return s.CachedPosts, nil
	}

	// Fetch new posts from the database
	posts, err := s.queries.ListPosts(context.Background(), 100)
	if err != nil {
		logger.Error("Error getting all posts: %e", err)
		return nil, err
	}

	// Update cache
	s.CachedPosts = posts
	s.LastFetch = time.Now()

	return posts, nil
}

func (s *BlogService) GetPostByID(id uuid.UUID) (*db.Post, error) {
	// First check if the post is in the cached posts
	for _, post := range s.CachedPosts {
		if post.ID == id {
			return &post, nil // Return a pointer to the cached post
		}
	}

	// If not found in cache, fetch from the database
	post, err := s.queries.GetPost(context.Background(), id.String())
	if err != nil {
		logger.Error("Error getting post by ID: %v", err)
		return nil, err
	}

	// Update the cache with this newly fetched post
	s.CachedPosts = append(s.CachedPosts, post)
	return &post, nil
}
