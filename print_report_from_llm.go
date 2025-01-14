package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func printReportFromLLM(apiKey string, hosts []string) {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	prompt := getPrompt(hosts)
	res, err := model.GenerateContent(ctx, genai.Text(prompt))

	fmt.Println()
	if err != nil {
		fmt.Println("failed to generate report using LLM ", err)
		return
	}

	printReportFromLLMHelper(res)
}

func printReportFromLLMHelper(res *genai.GenerateContentResponse) {
	for _, candidate := range res.Candidates {
		for _, part := range candidate.Content.Parts {
			fmt.Println(part)
		}
	}
}

func getPrompt(hosts []string) string {
	promptStart := `Hi, I am making this request from my web crawler app. I will 
	provide the top URLs I found while crawling. Can you generate 
	a report with a short description of each URL (at most 10 lines for each and 
	preferably at least 3 lines)? Only provide title for the report and the report with the 
	form ' idx num - URL \n description text ' and nothing else. The URLs are the following. `

	return promptStart + strings.Join(hosts, ",")
}
