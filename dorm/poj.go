package dorm

import "ny/db"

//dorm

type dorm struct {
	Id            int     `json:"id";gorm:"column:id;unique_index;PRIMARY_KEY;"`
	SchoolId      int     `json:"school_id";gorm:"column:school_id;"`
	DormName      string  `json:"dorm_name";gorm:"column:dorm_name;"`
	DormLongitude float32 `json:"dorm_longitude";gorm:"column:dorm_longitude;"`
	DormLatitude  float32 `json:"dorm_latitude";gorm:"column:dorm_latitude;"`
}

func (d *dorm) TableName() string {
	return "dorm"
}

// 更新非空字段内容
func (d *dorm) Update() bool {
	err := db.GetDB().Model(d).Update(d).Error
	return err == nil
}

// 根据主键信息填充其他字段
func (d *dorm) QueryById() bool {
	err := db.GetDB().Model(d).Where("id = ?", d.Id).First(d).Error
	return err == nil
}

// 获取第一条记录
func (d *dorm) QueryFirst() bool {
	err := db.GetDB().Model(d).First(d).Error
	return err == nil
}

// 查询shcool_id下的dorm
func QueryDormList(schoolId int) (bool, []dorm) {
	var r []dorm
	err := db.GetDB().Where("school_id = ?", schoolId).Find(&r).Error
	return err == nil, r
}
