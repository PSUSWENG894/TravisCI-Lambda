package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

func main() {

	//the open source repo api url prefix
	urlPrefix := "https://api.travis-ci.org/"

	//the "%2F" is important - don't replace it with a /
	repoSlug := "PSUSWENG894%2FGroup3/"

	//this is the authorization token for TravisCI to allow the connection to the API
	apiToken := "mU_3AeG5xd6aw_6rL1e9aA"

	//triggers a build of the master branch
	build("master", "Request to build from API with GO", apiToken)

	//lists the last five builds
	apiGet(urlPrefix + "repo/"+ repoSlug+ "builds?limit=5", apiToken)

}

//generic function to make a Travis CI get API call with no body
func apiGet(url string, apiToken string){

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Travis-API-Version", "3")
	req.Header.Set("Authorization", "token " + apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	head, _ :=json.MarshalIndent(resp.Header, "", "  ")

	//this reads the response body so it will be pretty printed below
	body, _ := ioutil.ReadAll(resp.Body)
	var prettyBody bytes.Buffer
	json.Indent(&prettyBody, body,"","  ")

	//this dumps the response out to the console
	fmt.Println("Response Status: ", resp.Status)
	fmt.Println("Response Headers: ", string(head))
	fmt.Println("Response Body: ", string(prettyBody.Bytes()))
}

//trigger a TravisCI build of the passed branch with the passed message will show
//up on the build in TravisCI
func build(buildBranch string , message string, apiToken string){

	//structs are use to construct jsom objects
	//the string to the right of the variables defines what the json key will be (otherwise it is the variable name)
	type ReqBody struct {
		Branch string `json:"branch"`
		Message string `json:"message"`
	}

	type ReqStruct struct {
		ReqBody ReqBody `json:"request"`
	}

	rBody := ReqStruct{
		ReqBody{
			buildBranch,
			message,
		},
	}

	jsonRequest, err := json.Marshal(rBody)

	/* Example of the marshalled json
	{
 		"request": {
 			"message": "Override the commit message: this is an api request",
 			"branch":"master",
		}
	}

	 */

	 //build the POST request
	req, _ := http.NewRequest("POST", "https://api.travis-ci.org/repo/PSUSWENG894%2FGroup3/requests", bytes.NewReader(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Travis-API-Version", "3")
	req.Header.Set("Authorization", "token " + apiToken)

	//execute the request
	client := &http.Client{}
	resp, err := client.Do(req)

	//handle errors
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//dump the output of the body to the console
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

//This is the equivalent command to the above written in curl
/*
curl -s -X POST \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-H "Travis-API-Version: 3" \
-H "Authorization: token mU_3AeG5xd6aw_6rL1e9aA" \
-d '{"request": {"message": "Some message", "branch":"master"}}'\
https://api.travis-ci.org/repo/PSUSWENG894%2FGroup3/requests
*/
