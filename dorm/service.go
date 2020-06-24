package dorm

import (
	"encoding/json"
	"ny/utils"
)

func getDormList(schoolId int) (bool, string) {

	ok, r := QueryDormList(schoolId)
	if ok {
		bytes, err := json.Marshal(r)
		if err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}
