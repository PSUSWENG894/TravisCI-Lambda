package travis_test

import (
	"testing"
	"net/http"
	"../travis"

	"encoding/json"
)

func TestBuildInfo(t *testing.T) {
	teardown := setup()
	defer teardown()

	limit := 5

	url := "/repo/"+ repoSlug + "builds"

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := travis.BuildsResponse{Type: "Hello"};
		respJson,_ := json.Marshal(resp);

		w.Write(respJson)
	})

	buildsResponse := travis.BuildsInfo(limit, client)

	if(buildsResponse.Type != "Hello"){
		t.Fatal("Expected Hello - Got " + buildsResponse.Type)
	}
}

// ... other tests here