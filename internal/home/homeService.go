package home

import (
	"Blog/pkg/github"
	"Blog/pkg/logger"
	"time"
)

type HomeService struct {
	lastFetch time.Time
	cachedGh  GhInfoCache
}

type GhInfoCache struct {
	GhPro       *github.GitHubProfile
	PinnedRepos []github.Repo
}

func NewHomeService() *HomeService {
	return &HomeService{}
}

func (s *HomeService) ghProfile() (*github.GitHubProfile, error) {
	login := "nooooaaaaah"
	ghPro, err := github.GetGitHubProfile(login)
	if err != nil {
		logger.LogError.Println("couldn't get gh profile: ", err)
		return nil, err
	}
	return ghPro, nil
}

func (s *HomeService) pinnedPosts() ([]github.Repo, error) {
	login := "nooooaaaaah"
	pinnedPosts, err := github.GetPinnedRepos(login)
	if err != nil {
		logger.LogError.Println("couldn't get pinned repos: ", err)
		return nil, err
	}
	return pinnedPosts, nil
}

func (s *HomeService) GetCachedGhInfo() (*GhInfoCache, error) {
	// Check if the cached data is still valid
	if time.Since(s.lastFetch) < 10*time.Minute {
		return &s.cachedGh, nil
	}

	// Fetch new data if the cache is stale
	ghPro, err := s.ghProfile()
	if err != nil {
		return nil, err
	}

	pinnedRepos, err := s.pinnedPosts()
	if err != nil {
		return nil, err
	}

	// Update the cache
	s.cachedGh = GhInfoCache{
		GhPro:       ghPro,
		PinnedRepos: pinnedRepos,
	}
	s.lastFetch = time.Now()

	return &s.cachedGh, nil
}