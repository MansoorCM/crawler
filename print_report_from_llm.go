package main

import (
	"fmt"
	"strings"

	"github.com/sgaunet/perplexity-go"
)

func printReportFromLLM(apiKey string, hosts []string) {

	client := perplexity.NewClient(apiKey)
	prompt := getPrompt(hosts)
	messages := getMessage(prompt)

	res, err := client.CreateCompletion(messages)

	fmt.Println()
	if err != nil {
		fmt.Println("failed to generate report using LLM ", err)
		return
	}

	fmt.Println(res.GetLastContent())
}

func getPrompt(hosts []string) string {
	promptStart := `hi, I am making this request from my web crawler app. I will 
	provide the hostnames of the top websites I found while crawling. can you provide 
	a report with a short description of each host (at most 10 lines for each). only 
	provide the report with the form ' idx num - hostname \n description text ' and 
	nothing else. The hostnames are the following `

	return promptStart + strings.Join(hosts, ",")
}

func getMessage(prompt string) []perplexity.Message {
	return []perplexity.Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}
}
