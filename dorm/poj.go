package dorm

import "ny/db"

//dorm

type dorm struct {
	Id            int64   `json:"id";gorm:"column:id;unique_index;PRIMARY_KEY;"`
	SchoolId      int64   `json:"school_id";gorm:"column:school_id;"`
	DormName      string  `json:"dorm_name";gorm:"column:dorm_name;"`
	DormLongitude float32 `json:"dorm_longitude";gorm:"column:dorm_longitude;"`
	DormLatitude  float32 `json:"dorm_latitude";gorm:"column:dorm_latitude;"`
}

func (d *dorm) TableName() string {
	return "dorm"
}

// 更新非空字段内容
func (d *dorm) update() bool {
	err := db.GetDB().Model(d).Update(d).Error
	return err == nil
}

// 根据主键信息填充其他字段
func (d *dorm) queryById() bool {
	err := db.GetDB().Model(d).Where("id = ?", d.Id).First(d).Error
	return err == nil
}

// 获取第一条记录
func (d *dorm) queryFirst() bool {
	err := db.GetDB().Model(d).First(d).Error
	return err == nil
}

// 查询shcool_id下的dorm
func queryDormList(schoolId int64) (bool, []dorm) {
	var r []dorm
	err := db.GetDB().Where("school_id = ?", schoolId).Find(&r).Error
	return err == nil, r
}
