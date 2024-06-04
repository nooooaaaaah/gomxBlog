package github

import (
	"Blog/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// GitHubProfile represents a user's GitHub profile.
type GitHubProfile struct {
	Login            string `json:"login"`
	AvatarURL        string `json:"avatar_url"`
	PublicReposCount int    `json:"public_repos"`
}

// License represents the license of a GitHub repository.
type License struct {
	Name string `json:"name"`
}

// Language represents a language used in a GitHub repository.
type Language struct {
	Name string `json:"name"`
}

// Repo represents a GitHub repository.
type Repo struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	License         License   `json:"licenseInfo"`
	OpenIssuesCount int       `json:"open_issues_count"`
	OpenPRCount     int       `json:"open_pr_count"`
	WatchersCount   int       `json:"watchers_count"`
	ForksCount      int       `json:"forks_count"`
	PrimaryLanguage Language  `json:"primary_language"`
	Languages       []Language
	HomepageURL     string `json:"homepage_url"`
	Topics          []string
}

// GraphQL query to fetch pinned repositories with additional information.
const queryTemplate = `
{
  user(login: "%s") {
    pinnedItems(first: 6, types: REPOSITORY) {
      edges {
        node {
          ... on Repository {
            name
            description
            createdAt
            updatedAt
            pushedAt
            licenseInfo {
              name
            }
            issues(states: OPEN) {
              totalCount
            }
            pullRequests(states: OPEN) {
              totalCount
            }
            watchers {
              totalCount
            }
            forks {
              totalCount
            }
            primaryLanguage {
              name
            }
            languages(first: 5) {
              edges {
                node {
                  name
                }
              }
            }
            homepageUrl
            repositoryTopics(first: 5) {
              edges {
                node {
                  topic {
                    name
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}`

// GraphQLResponse represents the response from the GraphQL API.
type GraphQLResponse struct {
	Data struct {
		User struct {
			PinnedItems struct {
				Edges []struct {
					Node struct {
						Name        string `json:"name"`
						Description string `json:"description"`
						CreatedAt   string `json:"createdAt"`
						UpdatedAt   string `json:"updatedAt"`
						PushedAt    string `json:"pushedAt"`
						LicenseInfo struct {
							Name string `json:"name"`
						} `json:"licenseInfo"`
						Issues struct {
							TotalCount int `json:"totalCount"`
						} `json:"issues"`
						PullRequests struct {
							TotalCount int `json:"totalCount"`
						} `json:"pullRequests"`
						Watchers struct {
							TotalCount int `json:"totalCount"`
						} `json:"watchers"`
						Forks struct {
							TotalCount int `json:"totalCount"`
						} `json:"forks"`
						PrimaryLanguage struct {
							Name string `json:"name"`
						} `json:"primaryLanguage"`
						Languages struct {
							Edges []struct {
								Node struct {
									Name string `json:"name"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"languages"`
						HomepageUrl      string `json:"homepageUrl"`
						RepositoryTopics struct {
							Edges []struct {
								Node struct {
									Topic struct {
										Name string `json:"name"`
									} `json:"topic"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"repositoryTopics"`
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

// GetPinnedRepos fetches a user's pinned GitHub repositories with additional information.
func GetPinnedRepos(username string) ([]Repo, error) {
	url := "https://api.github.com/graphql"
	query := fmt.Sprintf(queryTemplate, username)
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
			Name:            edge.Node.Name,
			Description:     edge.Node.Description,
			CreatedAt:       parseTime(edge.Node.CreatedAt),
			UpdatedAt:       parseTime(edge.Node.UpdatedAt),
			PushedAt:        parseTime(edge.Node.PushedAt),
			License:         License{Name: edge.Node.LicenseInfo.Name},
			OpenIssuesCount: edge.Node.Issues.TotalCount,
			OpenPRCount:     edge.Node.PullRequests.TotalCount,
			WatchersCount:   edge.Node.Watchers.TotalCount,
			ForksCount:      edge.Node.Forks.TotalCount,
			PrimaryLanguage: Language{Name: edge.Node.PrimaryLanguage.Name},
			HomepageURL:     edge.Node.HomepageUrl,
		}

		for _, langEdge := range edge.Node.Languages.Edges {
			repo.Languages = append(repo.Languages, Language{Name: langEdge.Node.Name})
		}

		for _, topicEdge := range edge.Node.RepositoryTopics.Edges {
			repo.Topics = append(repo.Topics, topicEdge.Node.Topic.Name)
		}

		repos = append(repos, repo)
	}

	return repos, nil
}

// parseTime is a helper function to parse time from the GitHub GraphQL API.
func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}
