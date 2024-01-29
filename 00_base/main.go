package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	//ctx, cancel = context.WithTimeout(10 * time.Second)
	//defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	return nil
}
