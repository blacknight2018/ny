package school

import (
	"encoding/json"
	"ny/utils"
)

func getSchoolList() (bool, string) {
	ok, r := querySchool()
	bytes, err := json.Marshal(r)
	if !ok || err != nil {
		return false, utils.EmptyString
	}
	return true, string(bytes)
}
