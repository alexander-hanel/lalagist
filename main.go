package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v62/github"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"log"
	"os"
	"strings"
	"time"
)

func prompt() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	gitHubToken := os.Getenv("GITHUB_TOKEN")
	gistID := os.Getenv("GITHUB_GIST_ID")
	llmName := os.Getenv("LLM_NAME")
	model := os.Getenv("MODEL")
	avatar := os.Getenv("AVATAR")

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(gitHubToken)

	_, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// Rate.Limit should most likely be 5000 when authorized.
	log.Printf("Rate: %#v\n", resp.Rate)

	// If a Token Expiration has been set, it will be displayed.
	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	comments, _, err := client.Gists.ListComments(ctx, gistID, nil)
	if len(comments) == 0 {
		fmt.Printf("no comments. skipping")
		return
	}
	lastComment := comments[len(comments)-1]
	if !strings.HasPrefix(strings.ToLower(*lastComment.Body), llmName) {
		return
	}
	prompt := strings.TrimSpace(strings.TrimPrefix(*lastComment.Body, llmName+","))

	// select ollama model
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}

	// run the ollama context in the background
	ctx_llm := context.Background()
	var response strings.Builder
	response.WriteString("> ")
	response.WriteString(prompt)
	response.WriteString("\n")
	response.WriteString("\n")
	mdAvatar := fmt.Sprintf("![](%s)\n", avatar)
	response.WriteString(mdAvatar)

	completion, err := llms.GenerateFromSinglePrompt(ctx_llm, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	response.WriteString(completion)

	newComment := &github.GistComment{
		Body: github.String(response.String()),
	}
	_, _, err = client.Gists.CreateComment(ctx, gistID, newComment)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	ticker := time.NewTicker(60 * time.Second)

	for {
		select {
		case <-ticker.C:
			prompt()
		}
		defer ticker.Stop()
	}
}
