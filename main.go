package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		var dat map[string]interface{}

		if err := json.Unmarshal([]byte(input), &dat); err != nil {
			panic(err)
		}

		if b, err := json.MarshalIndent(dat, "", "  "); err != nil {
			panic(err)
		} else {
			b = append(b, '\n')
			os.Stdout.Write(b)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
