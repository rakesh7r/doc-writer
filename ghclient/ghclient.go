package ghclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GitHubClient struct {
	httpClient *http.Client
	token      string
}

func WithGitHubToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, "githubToken", token)
}

func NewGitHubClient(ctx context.Context) *GitHubClient {
	return &GitHubClient{
		httpClient: &http.Client{Timeout: 15 * time.Second},
		token:      ctx.Value("githubToken").(string),
	}
}

func (c *GitHubClient) newRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	return req, nil
}

type Repo struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	DefaultBranch string `json:"default_branch"`
	Language      string `json:"language"`
}

func (c *GitHubClient) GetRepo(owner, repo string) (*Repo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := c.newRequest("GET", url)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("github api error: %s", resp.Status)
	}

	var r Repo
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

type TreeItem struct {
	Path string `json:"path"`
	Type string `json:"type"` // "blob" or "tree"
}

type TreeResponse struct {
	Tree []TreeItem `json:"tree"`
}

func (c *GitHubClient) GetRepoTree(owner, repo, branch string) ([]TreeItem, error) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1",
		owner, repo, branch,
	)

	req, _ := c.newRequest("GET", url)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tr TreeResponse
	json.NewDecoder(resp.Body).Decode(&tr)
	return tr.Tree, nil
}

type ContentResponse struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

func (c *GitHubClient) GetFile(owner, repo, path string) (string, error) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/%s",
		owner, repo, path,
	)

	req, _ := c.newRequest("GET", url)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cr ContentResponse
	json.NewDecoder(resp.Body).Decode(&cr)

	if cr.Encoding != "base64" {
		return "", fmt.Errorf("unsupported encoding")
	}

	decoded, err := base64.StdEncoding.DecodeString(
		strings.ReplaceAll(cr.Content, "\n", ""),
	)
	return string(decoded), err
}
