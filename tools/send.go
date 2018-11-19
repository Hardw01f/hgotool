package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	name := "Go"
	text := "Hello from go"
	channel := "alert"

	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + text + `"}`

	req, err := http.NewRequest(
		"POST",
		"https://hooks.slack.com/services/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",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(resp)
	defer resp.Body.Close()
}
