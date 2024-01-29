package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	return f1(ctx)
}

func printCandinates(cs []*genai.Candidate) {
	for _, c := range cs {
		for _, p := range c.Content.Parts {
			fmt.Println(p)
		}
	}
}

func f1(ctx context.Context) error {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_APIKEY")))
	if err != nil {
		return err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	prompt := genai.Text("プログラミング言語Goのマスコットキャラクタは？")

	resp, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return err
	}

	printCandinates(resp.Candidates)

	return nil
}
