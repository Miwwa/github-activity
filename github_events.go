package main

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
	Payload   interface{} `json:"payload,omitempty"`
	Public    bool        `json:"public,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	Org       struct {
		Id         string `json:"id,omitempty"`
		Login      string `json:"login,omitempty"`
		GravatarId string `json:"gravatar_id,omitempty"`
		Url        string `json:"url,omitempty"`
		AvatarUrl  string `json:"avatar_url,omitempty"`
	} `json:"org"`
}

func (event *GithubEvent) HumanString() string {
	return event.Type
}
