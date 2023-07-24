package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"net/http"
	"os"
)

type WorkerReadyResponse struct {
	RepoURL    string `json:"repoUrl"`
	BaseBranch string `json:"baseBranch"`
	BatchSize  string `json:"batchSize"`
}

func workerReady(instanceUrl string) (WorkerReadyResponse, error) {
	isReadyRequest := instanceUrl + "/worker-ready"
	res, err := http.Get(isReadyRequest)

	if err != nil {
		return WorkerReadyResponse{}, err
	}

	response := WorkerReadyResponse{}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)

	if err != nil {
		return WorkerReadyResponse{}, err
	}

	return response, nil
}

func main() {
	url := "http://127.0.0.1:65123"
	workerResponse, err := workerReady(url)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = git.PlainClone("tmp", false, &git.CloneOptions{
		URL: workerResponse.RepoURL,
		//RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress: os.Stdout,
	})
}
