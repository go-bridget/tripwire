package main

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

// Check represents an individual command to run
type Check struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

// A check may have multiple results (per-line)
type CheckResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func start() error {
	checksJSON, err := ioutil.ReadFile("tripwire.json")
	if err != nil {
		return err
	}

	checks := []*Check{}
	if err := json.Unmarshal(checksJSON, &checks); err != nil {
		return err
	}

	if len(checks) == 0 {
		return errors.New("No checks to run, empty")
	}

	errCount := 0

	var (
		OK   = color.GreenString("✓ %s %s")
		FAIL = color.YellowString("✗ %s %s")

		errRunningCheck    = color.RedString("error running check %s")
		errDecodingResults = color.RedString("error decoding check results for %s")
	)

	for _, check := range checks {
		results := []*CheckResult{}
		output, err := exec.Command(check.Command, check.Arguments...).Output()
		if err != nil {
			log.WithError(err).Errorf(errRunningCheck, check.Command)
			errCount++
			continue
		}

		if err := json.Unmarshal(output, &results); err != nil {
			log.WithError(err).Errorf(errDecodingResults, check.Command)
			errCount++
			continue
		}

		log.Infof("Got %d results from check %s", len(results), check.Command)

		for _, result := range results {
			if result.Value == "OK" {
				log.Infof(OK, result.Key, result.Value)
			} else {
				log.Infof(FAIL, result.Key, result.Value)
			}
		}
	}

	if errCount > 0 {
		return errors.Errorf("Encountered %d errors, checks are failing", errCount)
	}

	return nil
}

func main() {
	if err := start(); err != nil {
		log.WithError(err).Fatal("exiting")
	}
}
