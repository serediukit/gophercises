package decoder

import (
	"encoding/json"

	"cyoa/entity"
)

func Decode(data []byte) (entity.Stories, error) {
	stories := make(entity.Stories)
	err := json.Unmarshal(data, &stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}
