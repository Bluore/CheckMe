package change

import (
	"encoding/json"

	"gorm.io/datatypes"
)

func ChangeData(cov *map[string]map[string]interface{}, req datatypes.JSON, app string) (datatypes.JSON, error) {

	conf, ok := (*cov)[app]
	if !ok {
		return req, nil
	}

	var data map[string]interface{}
	err := json.Unmarshal(req, &data)
	if err != nil {
		return nil, err
	}

	for key, val := range conf {
		data[key] = val
	}

	if val, ok := data["is_hide_music"]; ok || val == true {
		delete(data, "music_name")
	}

	res, err := MapToJSON(data)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func MapToJSON(m map[string]interface{}) (datatypes.JSON, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(b), nil
}
