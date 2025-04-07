package urlshort

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func parseYAML(data []byte) (parsedYAML []map[string]string, err error) {
	err = yaml.Unmarshal(data, &parsedYAML)
	if err != nil {
		return nil, err
	}
	return parsedYAML, nil
}

func parseJSON(data []byte) (parsedJSON []map[string]string, err error) {
	err = json.Unmarshal(data, &parsedJSON)
	if err != nil {
		return nil, err
	}
	return parsedJSON, nil
}
