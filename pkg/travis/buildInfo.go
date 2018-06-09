package travis

import (
	"time"
	"encoding/json"
	"strconv"
	"io/ioutil"
)

//TravisCI response body from the build info API call
type BuildsResponse struct {
	Type           string `json:"@type"`
	Href           string `json:"@href"`
	Representation string `json:"@representation"`
	Pagination     struct {
		Limit   int  `json:"limit"`
		Offset  int  `json:"offset"`
		Count   int  `json:"count"`
		IsFirst bool `json:"is_first"`
		IsLast  bool `json:"is_last"`
		Next    struct {
			Href   string `json:"@href"`
			Offset int    `json:"offset"`
			Limit  int    `json:"limit"`
		} `json:"next"`
		Prev  interface{} `json:"prev"`
		First struct {
			Href   string `json:"@href"`
			Offset int    `json:"offset"`
			Limit  int    `json:"limit"`
		} `json:"first"`
		Last struct {
			Href   string `json:"@href"`
			Offset int    `json:"offset"`
			Limit  int    `json:"limit"`
		} `json:"last"`
	} `json:"@pagination"`
	Builds []struct {
		Type           string `json:"@type"`
		Href           string `json:"@href"`
		Representation string `json:"@representation"`
		Permissions    struct {
			Read    bool `json:"read"`
			Cancel  bool `json:"cancel"`
			Restart bool `json:"restart"`
		} `json:"@permissions"`
		ID                int         `json:"id"`
		Number            string      `json:"number"`
		State             string      `json:"state"`
		Duration          int         `json:"duration"`
		EventType         string      `json:"event_type"`
		PreviousState     string      `json:"previous_state"`
		PullRequestTitle  interface{} `json:"pull_request_title"`
		PullRequestNumber interface{} `json:"pull_request_number"`
		StartedAt         time.Time   `json:"started_at"`
		FinishedAt        time.Time   `json:"finished_at"`
		Private           bool        `json:"private"`
		Repository        struct {
			Type           string `json:"@type"`
			Href           string `json:"@href"`
			Representation string `json:"@representation"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Slug           string `json:"slug"`
		} `json:"repository"`
		Branch struct {
			Type           string `json:"@type"`
			Href           string `json:"@href"`
			Representation string `json:"@representation"`
			Name           string `json:"name"`
		} `json:"branch"`
		Tag    interface{} `json:"tag"`
		Commit struct {
			Type           string    `json:"@type"`
			Representation string    `json:"@representation"`
			ID             int       `json:"id"`
			Sha            string    `json:"sha"`
			Ref            string    `json:"ref"`
			Message        string    `json:"message"`
			CompareURL     string    `json:"compare_url"`
			CommittedAt    time.Time `json:"committed_at"`
		} `json:"commit"`
		Jobs []struct {
			Type           string `json:"@type"`
			Href           string `json:"@href"`
			Representation string `json:"@representation"`
			ID             int    `json:"id"`
		} `json:"jobs"`
		Stages    []interface{} `json:"stages"`
		CreatedBy struct {
			Type           string `json:"@type"`
			Href           string `json:"@href"`
			Representation string `json:"@representation"`
			ID             int    `json:"id"`
			Login          string `json:"login"`
		} `json:"created_by"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"builds"`
}

func BuildsInfo(limit int, client *Client) (BuildsResponse){
	url := client.baseURL+ "/repo/"+ client.repoSlug + "builds?limit=" + strconv.Itoa(limit)

	response := ApiGet(url, client)

	body, _ := ioutil.ReadAll(response.Body)
	var responseStruct BuildsResponse
	err := json.Unmarshal(body, &responseStruct)

	defer response.Body.Close()

	if(err != nil) {
		panic(err)
	}

	DumpResponse(response, body)

	return responseStruct
}




