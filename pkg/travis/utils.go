package travis

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
)

//generic function to make a Travis CI get API call with no body
func ApiGet(url string, client *Client) (response *http.Response){
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Travis-API-Version", "3")
	req.Header.Set("Authorization", "token " + client.apiToken)

//	client := &http.Client{}
	resp, err := client.httpClient.Do(req)

	if err != nil {
		panic(err)
	}

	return resp
}

//convenience method to dump the contents of an http response
func DumpResponse(response *http.Response, body []byte){
	head, _ :=json.MarshalIndent(response.Header, "", "  ")
	//body, _ := ioutil.ReadAll(response.Body)

	var prettyBody bytes.Buffer
	json.Indent(&prettyBody, body,"","  ")

	fmt.Println("Response Status: ", response.Status)
	fmt.Println("Response Headers: ", string(head))
	fmt.Println("Response Body: ", string(prettyBody.Bytes()))
}
