package dorm

import (
	"encoding/json"
	"ny/utils"
)

func getDormList(schoolId int64) (bool, string) {

	ok, r := queryDormList(schoolId)
	if ok {
		bytes, err := json.Marshal(r)
		if err == nil {
			return true, string(bytes)
		}
	}
	return false, utils.EmptyString
}

func getDormIdList(schoolId int64) []int64 {
	var idList []int64
	ok, r := queryDormList(schoolId)
	if ok {
		for _, v := range r {
			idList = append(idList, v.Id)
		}
	}
	return idList
}
