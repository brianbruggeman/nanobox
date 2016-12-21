package component

import (
	"encoding/json"

	"github.com/jcelliott/lumber"

	"github.com/nanobox-io/nanobox/models"
)

// PlanPayload returns a string for the user hook payload
func PlanPayload(component *models.Component) string {
	config, err := componentConfig(component)
	if err != nil {
		lumber.Error("hooks:componentConfig(): %s", err.Error())
		return "{}"
	}

	payload := map[string]interface{}{
		"config": config,
	}

	// marshal the payload into json
	b, err := json.Marshal(payload)
	if err != nil {
		lumber.Error("hooks:json.Marshal(): %s", err.Error())
		return "{}"
	}

	return string(b)
}
