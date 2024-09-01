package main

import (
	"encoding/json"
	"fmt"
)

type GithubEvent struct {
	Id    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Actor struct {
		Id           int    `json:"id,omitempty"`
		Login        string `json:"login,omitempty"`
		DisplayLogin string `json:"display_login,omitempty"`
		GravatarId   string `json:"gravatar_id,omitempty"`
		Url          string `json:"url,omitempty"`
		AvatarUrl    string `json:"avatar_url,omitempty"`
	} `json:"actor"`
	Repo struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"repo"`
	Payload   json.RawMessage `json:"payload,omitempty"`
	Public    bool            `json:"public,omitempty"`
	CreatedAt string          `json:"created_at,omitempty"`
	Org       struct {
		Id         int    `json:"id,omitempty"`
		Login      string `json:"login,omitempty"`
		GravatarId string `json:"gravatar_id,omitempty"`
		Url        string `json:"url,omitempty"`
		AvatarUrl  string `json:"avatar_url,omitempty"`
	} `json:"org"`
}

func (event *GithubEvent) HumanString() string {
	switch event.Type {
	case "CreateEvent":
		var payload CreateEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "CreateEvent invalid payload"
		}
		switch payload.RefType {
		case "branch", "tag":
			return fmt.Sprintf("%s '%s' created in repo %s", payload.RefType, payload.Ref, event.Repo.Name)
		case "repository":
			return fmt.Sprintf("%s '%s' created", payload.RefType, event.Repo.Name)
		}
	case "DeleteEvent":
		var payload DeleteEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "DeleteEvent invalid payload"
		}
		return fmt.Sprintf("%s '%s' deleted in repo %s", payload.RefType, payload.Ref, event.Repo.Name)
	case "CommitCommentEvent":
		return fmt.Sprintf("commit commented in repo %s", event.Repo.Name)
	case "ForkEvent":
		return fmt.Sprintf("created fork of %s", event.Repo.Name)
	case "GollumEvent":
		var payload GollumEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "GollumEvent invalid payload"
		}
		return fmt.Sprintf("%d wiki pages created in repo %s", len(payload.Pages), event.Repo.Name)
	case "IssueCommentEvent":
		var payload IssueCommentEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "IssueCommentEvent invalid payload"
		}
		return fmt.Sprintf("issue comment %s in repo %s", payload.Action, event.Repo.Name)
	case "IssuesEvent":
		var payload IssuesEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "IssueCommentEvent invalid payload"
		}
		return fmt.Sprintf("issue %s in repo %s", payload.Action, event.Repo.Name)
	case "MemberEvent":
		return fmt.Sprintf("new member added to repo %s", event.Repo.Name)
	case "PublicEvent":
		return fmt.Sprintf("repo %s is made public!", event.Repo.Name)
	case "PullRequestEvent":
		var payload PullRequestEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "PullRequestEvent invalid payload"
		}
		return fmt.Sprintf("pull request %s in repo %s", payload.Action, event.Repo.Name)
	case "PullRequestReviewEvent":
		var payload PullRequestReviewEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "PullRequestEvent invalid payload"
		}
		return fmt.Sprintf("pull request review %s in repo %s", event.Repo.Name)
	case "PullRequestReviewCommentEvent":
		var payload PullRequestReviewCommentEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "PullRequestEvent invalid payload"
		}
		return fmt.Sprintf("pull request review comment %s in repo %s", event.Repo.Name)
	case "PullRequestReviewThreadEvent":
		var payload PullRequestReviewThreadEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "PullRequestReviewThreadEvent invalid payload"
		}
		return fmt.Sprintf("Pull request review thread %s in repo %s", payload.Action, event.Repo.Name)
	case "PushEvent":
		var payload PushEvent
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return "PushEvent invalid payload"
		}
		return fmt.Sprintf("pushed %d commits to %s", payload.Size, event.Repo.Name)
	case "ReleaseEvent":
		return fmt.Sprintf("published a new release in %s", event.Repo.Name)
	case "SponsorshipEvent":
		return "SponsorshipEvent"
	case "WatchEvent":
		return fmt.Sprintf("starred repo %s", event.Repo.Name)
	}
	return "Unknown event"
}

// some struct below are incomplete because we don't need these data for current task

type PushEvent struct {
	RepositoryId int           `json:"repository_id,omitempty"`
	PushId       int           `json:"push_id,omitempty"`
	Size         int           `json:"size,omitempty"`
	DistinctSize int           `json:"distinct_size,omitempty"`
	Ref          string        `json:"ref,omitempty"`
	Head         string        `json:"head,omitempty"`
	Before       string        `json:"before,omitempty"`
	Commits      []interface{} `json:"commits,omitempty" json:"commits,omitempty"`
}

type PullRequestReviewThreadEvent struct {
	Action      string      `json:"action,omitempty"`
	PullRequest interface{} `json:"pull_request,omitempty"`
	Thread      interface{} `json:"thread,omitempty"`
}

type PullRequestEvent struct {
	Action      string        `json:"action,omitempty"`
	Number      int           `json:"number,omitempty"`
	Changes     []interface{} `json:"changes,omitempty"`
	PullRequest interface{}   `json:"pull_request,omitempty"`
	Reason      string        `json:"reason,omitempty"`
}

type CreateEvent struct {
	Ref          string `json:"ref,omitempty"`
	RefType      string `json:"ref_type,omitempty"`
	MasterBranch string `json:"master_branch,omitempty"`
	Description  string `json:"description,omitempty"`
	PusherType   string `json:"pusher_type,omitempty"`
}

type DeleteEvent struct {
	Ref     string `json:"ref,omitempty"`
	RefType string `json:"ref_type,omitempty"`
}

type GollumEvent struct {
	Pages []struct {
		PageName string `json:"page_name,omitempty"`
		Title    string `json:"title,omitempty"`
		Action   string `json:"action,omitempty"`
		Sha      string `json:"sha,omitempty"`
		HtmlUrl  string `json:"html_url,omitempty"`
	} `json:"pages,omitempty"`
}

type IssueCommentEvent struct {
	Action  string      `json:"action,omitempty"`
	Changes interface{} `json:"changes,omitempty"`
	Issue   interface{} `json:"issue,omitempty"`
	Comment interface{} `json:"comment,omitempty"`
}

type IssuesEvent struct {
	Action   string      `json:"action,omitempty"`
	Changes  interface{} `json:"changes,omitempty"`
	Issue    interface{} `json:"issue,omitempty"`
	Assignee interface{} `json:"assignee,omitempty"`
	Label    interface{} `json:"label,omitempty"`
}

type PullRequestReviewEvent struct {
	Action string `json:"action,omitempty"`
}

type PullRequestReviewCommentEvent struct {
	Action string `json:"action,omitempty"`
}
