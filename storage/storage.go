package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/oliviaclyde/grinchmas/grinchmas"
)

const filename = "grinchmas.json"

func Load() error {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var savedEvents []grinchmas.Event
	err = json.Unmarshal(fileContents, &savedEvents)
	if err != nil {
		return err
	}

	grinchmas.SetEvents(savedEvents)

	return nil
}

func Save() error {
	eventList := grinchmas.ListEvents()

	eventListBytes, err := json.MarshalIndent(eventList, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, eventListBytes, 0775)
	if err != nil {
		return err
	}

	return nil
}
