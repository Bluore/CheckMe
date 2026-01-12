package judge

import "gorm.io/datatypes"

func IsJSONNull(j datatypes.JSON) bool {
	if len(j) == 0 || string(j) == "null" {
		return true
	} else {
		return false
	}
}
