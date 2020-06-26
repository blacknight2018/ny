package school

import "ny/db"

type school struct {
	Id         int64  `json:"id";gorm:"column:id;unique_index;PRIMARY_KEY;"`
	SchoolName string `json:"school_name";gorm:"column:school_name;"`
}

func (s *school) TableName() string {
	return "school"
}

// 更新非空字段内容
func (s *school) update() bool {
	err := db.GetDB().Model(s).Update(s).Error
	return err == nil
}

// 返回所有学校
func querySchool() (bool, []school) {
	var s []school
	err := db.GetDB().Model(s).Find(&s).Error
	if err != nil {
		return false, nil
	}
	return true, s
}
