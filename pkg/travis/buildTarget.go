package travis

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
)

//TravisCI request body for the build API call
type RequestBody struct {
	Request struct {
		Branch  string `json:"branch"`
		Message string `json:"message"`
		Config  struct {
			Env struct {
				Matrix []string `json:"matrix"`
			} `json:"env"`
		} `json:"config"`
	} `json:"request"`
}

//TravisCI response body from the build API call
type ResponseBody struct {
	Type              string `json:"@type"`
	RemainingRequests int    `json:"remaining_requests"`
	Repository        struct {
		Type           string `json:"@type"`
		Href           string `json:"@href"`
		Representation string `json:"@representation"`
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Slug           string `json:"slug"`
	} `json:"repository"`
	Request struct {
		Repository struct {
			ID        int    `json:"id"`
			OwnerName string `json:"owner_name"`
			Name      string `json:"name"`
		} `json:"repository"`
		User struct {
			ID int `json:"id"`
		} `json:"user"`
		ID      int    `json:"id"`
		Message string `json:"message"`
		Branch  string `json:"branch"`
		Config  struct {
			Env struct {
				Matrix []string `json:"matrix"`
			} `json:"env"`
		} `json:"config"`
	} `json:"request"`
	ResourceType string `json:"resource_type"`
}

//trigger a TravisCI build
func Build(client *Client, buildBranch string , message string,  envVars []string) (ResponseBody){
	reqBody := RequestBody{}

	//set the variables
	reqBody.Request.Branch = buildBranch
	reqBody.Request.Message = message
	reqBody.Request.Config.Env.Matrix = envVars

	jsonRequest, _ := json.Marshal(reqBody)

	/* Example of the marshalled json
 {
    "request": {
        "branch": "master",
        "message": "Request to build from API with GO",
        "config": {
            "env": {
                "matrix": [
                    "UNITTESTS=ALL",
                    "INTTESTS=NONE"
                ]
            }
        }
    }
}
	 */

	//build the POST request
	req, _ := http.NewRequest("POST","https://api.travis-ci.org/repo/" + client.repoSlug + "requests" , bytes.NewReader(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Travis-API-Version", "3")
	req.Header.Set("Authorization", "token " + client.apiToken)

	//execute the request
	resp, err := client.httpClient.Do(req)

	//handle errors
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return handleBuildResponse(resp)
}

func handleBuildResponse(response *http.Response) (ResponseBody){
	body, _ := ioutil.ReadAll(response.Body)

	var responseBody ResponseBody
	err := json.Unmarshal(body, &responseBody)

	if(err != nil){
		panic(err)
	}

	return responseBody;
}

