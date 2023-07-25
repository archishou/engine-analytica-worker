package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"net/http"
	"os"
)

type WorkerReadyResponse struct {
	RepoURL    string `json:"repoUrl"`
	BaseBranch string `json:"baseBranch"`
	DevBranch  string `json:"devBranch"`
	BatchSize  string `json:"batchSize"`
}

var tmpDirectory = "tmp"
var cutechessBinaryDir = tmpDirectory + "/cutechess-binaries"
var repoDirs = tmpDirectory + "/repos"
var baseDir = repoDirs + "/base"
var devDir = repoDirs + "/dev"

func fetchCutechessBinaries() {
	if _, err := os.Stat(cutechessBinaryDir); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("[INFO] Fetching cutechess binaries.")
			_, _ = git.PlainClone(cutechessBinaryDir, false, &git.CloneOptions{
				URL:      "https://github.com/archishou/cutechess-binaries.git",
				Progress: os.Stdout,
			})
		} else {
			fmt.Println("Failed to fetch cutchess-cli-binaries", err)
		}
	}
}

func workerReady(instanceUrl string) (WorkerReadyResponse, error) {
	isReadyRequest := instanceUrl + "/worker-ready"
	fmt.Println("[INFO] Fetching workload.")
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
	fetchCutechessBinaries()
	workerResponse, err := workerReady(url)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = git.PlainClone(baseDir, false, &git.CloneOptions{
		URL:           workerResponse.RepoURL,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(workerResponse.BaseBranch),
	})

	_, _ = git.PlainClone(devDir, false, &git.CloneOptions{
		URL:           workerResponse.RepoURL,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(workerResponse.DevBranch),
	})
}
