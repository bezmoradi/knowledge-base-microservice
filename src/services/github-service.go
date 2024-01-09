package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Depado/bfchroma/v2"
	"github.com/google/go-github/github"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/oauth2"
)

func refillTipsCollection() {
	gitHubAccessToken := os.Getenv("GITHUB_TOKEN")
	gitHubUsername := os.Getenv("GITHUB_USERNAME")
	gitHubRepository := os.Getenv("GITHUB_REPO")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitHubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	fetchContentsRecursively(client, gitHubUsername, gitHubRepository, "")
}

func markdownToHTML(markdownContent string) string {
	htmlContent := blackfriday.Run(
		[]byte(markdownContent),
		blackfriday.WithRenderer(bfchroma.NewRenderer(bfchroma.Style("nord"))),
	)

	return string(htmlContent)
}

func getContent(file string) string {
	gitHubAccessToken := os.Getenv("GITHUB_TOKEN")
	gitHubUsername := os.Getenv("GITHUB_USERNAME")
	gitHubRepository := os.Getenv("GITHUB_REPO")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitHubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	fileContent, _, _, err := client.Repositories.GetContents(context.Background(), gitHubUsername, gitHubRepository, file, nil)
	if err != nil {
		fmt.Println("content not found")

		return ""
	}

	content, _ := fileContent.GetContent()

	return markdownToHTML(content)
}

func fetchContentsRecursively(client *github.Client, owner, repo, path string) {
	_, contents, _, _ := client.Repositories.GetContents(context.Background(), owner, repo, path, nil)

	for _, content := range contents {
		if !shouldBeExcluded(*content.Name) {
			if *content.Type == "file" {
				createTip(*content.Path)
			} else if *content.Type == "dir" {
				fetchContentsRecursively(client, owner, repo, *content.Path)
			}
		}
	}
}

func shouldBeExcluded(path string) bool {
	return strings.Contains(path, "_") || path == ".gitignore"
}
