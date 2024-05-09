package blog

// Post represents a blog post
type Post struct {
	ID    int
	Title string
	Body  string
}

// Service implements the BlogService interface
type Service struct {
	// Dependencies, e.g., a database client
}

// NewService creates a new blog service
func NewService() *Service {
	return &Service{}
}

// GetAllPosts returns all blog posts
func (s *Service) getAllPosts() ([]Post, error) {
	// Dummy data; replace with actual DB call
	return []Post{
		{ID: 1, Title: "First Post", Body: "This is the first post"},
		{ID: 2, Title: "Second Post", Body: "This is the second post"},
	}, nil
}

// GetPostByID returns a single blog post by ID
func (s *Service) getPostByID(id int) (Post, error) {
	// Dummy data; replace with actual DB call
	return Post{ID: id, Title: "Sample Post", Body: "This is a sample post"}, nil
}
