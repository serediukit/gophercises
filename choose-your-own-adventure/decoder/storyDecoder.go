package decoder

import (
	"encoding/json"

	"cyoa/entity"
)

func Decode(data []byte) (map[string]*entity.Arc, error) {
	stories := make(map[string]*entity.Arc)
	err := json.Unmarshal(data, &stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}
