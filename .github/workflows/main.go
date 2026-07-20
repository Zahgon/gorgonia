package main

import (
	"io"
	"os"
)

const (
	latestGo   = "1.16.x"
	previousGo = "1.15.x"
)

func main() {
	workflowLinux, err := os.OpenFile("runner-github-ubuntu-amd64.yml", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer workflowLinux.Close()
	workflowMac, err := os.OpenFile("runner-github-macos-amd64.yml", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer workflowMac.Close()
	workflowSelf, err := os.OpenFile("runner-self-hosted.yml", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer workflowSelf.Close()

	err = generateWorkflow(workflowLinux, "Build and Tests on Linux/amd64", "Linux/amd64", "ubuntu-latest", map[string]bool{
		"none": false,
		"avx":  false,
		"sse":  false,
	}, true)
	if err != nil {
		panic(err)
	}
	err = generateWorkflow(workflowMac, "Build and Tests on MacOS/amd64", "MacOS/amd64", "macos-latest", map[string]bool{
		"none": false,
		"avx":  false,
		"sse":  false,
	}, false)
	if err != nil {
		panic(err)
	}
	err = generateWorkflow(workflowSelf, "Build and Tests on Self-Hosted (arm)", "Self-Hosted", "self-hosted", map[string]bool{
		"none": false,
	}, false)
	if err != nil {
		panic(err)
	}

}

func generateWorkflow(w io.Writer, workflowName, runnerName, runsOn string, tags map[string]bool, withRace bool) error {
	_ = "STUB: not implemented"
	return nil
}

func mapToList(m map[string]bool) string { _ = "STUB: not implemented"; return "" }

func hasExperimental(m map[string]bool) bool { _ = "STUB: not implemented"; return false }
