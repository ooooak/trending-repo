package db

import (
	"encoding/json"
	"io/ioutil"

	"../types"
)

// FileName !
const FileName = "output.json"

// Write !
func Write(data types.Records) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(FileName, json, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Read !
func Read() (types.Records, error) {
	var r types.Records
	body, err := ioutil.ReadFile(FileName)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
