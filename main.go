package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	outputLines, err := run(os.Args[1:])
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, "Error: ", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	println(strings.Join(outputLines, "\n"))
}

func run(args []string) ([]string, error) {
	if len(args) != 1 {
		return []string{}, errors.New("usage: github-activity <username>")
	}
	events, err := fetchApi(args[0])
	if err != nil {
		return []string{}, err
	}
	lines := make([]string, 0, len(events))
	for _, event := range events {
		lines = append(lines, "- "+event.HumanString())
	}
	return lines, nil
}

func fetchApi(username string) ([]GithubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	response, err := http.Get(url)
	if err != nil {
		return []GithubEvent{}, errors.Join(errors.New("GitHub API error"), err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return []GithubEvent{}, fmt.Errorf("GitHub API error: %d %s", response.StatusCode, response.Status)
	}
	var events []GithubEvent
	err = json.NewDecoder(response.Body).Decode(&events)
	if err != nil {
		return nil, errors.Join(errors.New("json decode error"), err)
	}
	return events, nil
}
