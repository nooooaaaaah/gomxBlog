package blog

import (
	"Blog/pkg/db"
	"Blog/pkg/logger"
	"time"

	"github.com/edgedb/edgedb-go"
)

// BlogService implements the BlogService interface
type BlogService struct {
	lastFetch   time.Time
	cachedPosts []db.Post
}

// NewBlogService creates a new blog service
func NewBlogService() *BlogService {
	return &BlogService{}
}

func (s *BlogService) getAllPosts() ([]db.Post, error) {
	// Get the current year and week number
	year, week := time.Now().ISOWeek()

	// Check if the cached posts are from the current week
	yearLF, weekLF := s.lastFetch.ISOWeek()
	if yearLF == year && weekLF == week {
		return s.cachedPosts, nil
	}

	// Fetch new posts from the database
	posts, err := db.GetPosts()
	if err != nil {
		logger.LogError.Println("Error getting all posts: ", err)
		return nil, err
	}

	// Update cache
	s.cachedPosts = posts
	s.lastFetch = time.Now()

	return posts, nil
}

func (s *BlogService) getPostByID(id edgedb.UUID) (*db.Post, error) {
	// First check if the post is in the cached posts
	for _, post := range s.cachedPosts {
		if post.Id == id {
			return &post, nil // Return a pointer to the cached post
		}
	}

	// If not found in cache, fetch from the database
	post, err := db.GetPostByID(id)
	if err != nil {
		logger.LogError.Println("Error getting post by ID:", err)
		return nil, err
	}

	// Optionally, you could update the cache with this newly fetched post
	// This step depends on your caching strategy; if you wish to keep the cache strictly once per week,
	// you might skip updating the cache here.
	return post, nil
}
