package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	fmt.Println("hello", workerResponse.RepoURL)
	fmt.Println("hello", workerResponse.BaseBranch)
	fmt.Println("hello", workerResponse.BatchSize)
}
