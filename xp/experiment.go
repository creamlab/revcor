package xp

import (
	"encoding/json"
	"io/ioutil"

	"github.com/creamlab/revcor/helpers"
)

type ExperimentSettings struct {
	Id             string `json:"id"`
	AdminPassword  string `json:"adminPassword"`
	AllowCreate    bool   `json:"allowCreate"`
	CreatePassword string `json:"createPassword"`
	TrialsPerBlock int    `json:"trialsPerBlock"`
	BlocksPerXp    int    `json:"blocksPerXp"`
	AddRepeatBlock bool   `json:"addRepeatBlock"`
}

// API
// Check if ids sent by client are valid (match a regex + configuration file exists)
func IsExperimentValid(experimentId string) bool {
	// check IDs format
	if !helpers.IsIdValid(experimentId) {
		return false
	}
	// check config exisis
	return helpers.PathExists("data/" + experimentId + "/config/settings.json")
}

func GetExperimentSettings(experimentId string) (e ExperimentSettings, err error) {
	settingsPath := "data/" + experimentId + "/config/settings.json"
	file, err := ioutil.ReadFile(settingsPath)
	if err != nil {
		return
	}

	e = ExperimentSettings{}
	if err = json.Unmarshal([]byte(file), &e); err != nil {
		return
	}
	e.Id = experimentId
	return
}

func GetSanitizedExperimentSettings(experimentId string) (e ExperimentSettings, err error) {
	e, err = GetExperimentSettings(experimentId)
	// sanitize
	e.AdminPassword = ""
	e.CreatePassword = ""
	return
}

func GetExperimentWordingRunString(experimentId string) (json string, err error) {
	wordingRunPath := "data/" + experimentId + "/config/wording.run.json"
	return helpers.ReadTrimJSON(wordingRunPath)
}

// no error is returned
func GetExperimentWordingNewMap(experimentId string) (m map[string]string) {
	wordingNewPath := "data/" + experimentId + "/config/wording.new.json"
	file, err := ioutil.ReadFile(wordingNewPath)
	if err != nil {
		return
	}

	json.Unmarshal([]byte(file), &m)
	return
}
