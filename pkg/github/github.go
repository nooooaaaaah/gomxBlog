package github

import (
	"Blog/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

// GitHubProfile represents a user's GitHub profile.
type GitHubProfile struct {
	Login            string `json:"login"`
	AvatarURL        string `json:"avatar_url"`
	PublicReposCount int    `json:"public_repos"`
}

// Repo represents a GitHub repository.
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
}

// GetGitHubProfile fetches a user's GitHub profile.
func GetGitHubProfile(username string) (*GitHubProfile, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GitHub profile: %v", err)
	}
	defer resp.Body.Close()

	var profile GitHubProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("failed to decode GitHub profile JSON: %v", err)
	}
	logger.LogInfo.Printf("Fetched GitHub profile: %+v\n", profile)
	return &profile, nil
}

// GetPinnedRepos fetches a user's pinned GitHub repositories.
func GetPinnedRepos(username string) (*[]Repo, error) {
	url := fmt.Sprintf("https://gh-pinned-repos.now.sh/?username=%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pinned repositories: %v", err)
	}
	defer resp.Body.Close()

	var repos []Repo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, fmt.Errorf("failed to decode pinned repositories JSON: %v", err)
	}
	return &repos, nil
}
