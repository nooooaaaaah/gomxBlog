package github

import (
	"Blog/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

// GraphQL query to fetch pinned repositories
const query = `
{
  user(login: "%s") {
    pinnedItems(first: 6, types: REPOSITORY) {
      totalCount
      edges {
        node {
          ... on Repository {
            name
            description
            stargazers {
              totalCount
            }
            forks {
              totalCount
            }
          }
        }
      }
    }
  }
}`

// GraphQLResponse represents the response from the GraphQL API
type GraphQLResponse struct {
	Data struct {
		User struct {
			PinnedItems struct {
				Edges []struct {
					Node struct {
						Name        string `json:"name"`
						Description string `json:"description"`
						Stargazers  struct {
							TotalCount int `json:"totalCount"`
						} `json:"stargazers"`
						Forks struct {
							TotalCount int `json:"totalCount"`
						} `json:"forks"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"pinnedItems"`
		} `json:"user"`
	} `json:"data"`
}

// GetGitHubProfile fetches a user's GitHub profile.
func GetGitHubProfile(username string) (*GitHubProfile, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GitHub profile: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var profile GitHubProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("failed to decode GitHub profile JSON: %v", err)
	}
	logger.LogInfo.Printf("Fetched GitHub profile: %+v\n", profile)
	return &profile, nil
}

// GetPinnedRepos fetches a user's pinned GitHub repositories.
func GetPinnedRepos(username string) ([]Repo, error) {
	url := "https://api.github.com/graphql"
	query := fmt.Sprintf(`
		{
		  user(login: "%s") {
		    pinnedItems(first: 6, types: REPOSITORY) {
		      totalCount
		      edges {
		        node {
		          ... on Repository {
		            name
		            description
		            stargazers {
		              totalCount
		            }
		            forks {
		              totalCount
		            }
		          }
		        }
		      }
		    }
		  }
		}
	`, username)
	reqBody := fmt.Sprintf(`{"query": %q}`, query)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, bodyString)
	}

	var graphqlResponse GraphQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&graphqlResponse); err != nil {
		return nil, fmt.Errorf("failed to decode GraphQL response JSON: %v", err)
	}

	var repos []Repo
	for _, edge := range graphqlResponse.Data.User.PinnedItems.Edges {
		repo := Repo{
			Name:        edge.Node.Name,
			Description: edge.Node.Description,
			Stars:       edge.Node.Stargazers.TotalCount,
			Forks:       edge.Node.Forks.TotalCount,
		}
		repos = append(repos, repo)
	}
	logger.LogInfo.Println(repos)
	return repos, nil
}
