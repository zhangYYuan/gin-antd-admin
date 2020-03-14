package busModel

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

/**
Column	指定列名
Type	指定列数据类型
Size	指定列大小, 默认值255
PRIMARY_KEY	将列指定为主键
UNIQUE	将列指定为唯一
DEFAULT	指定列默认值
PRECISION   指定列精度
NOT NULL    将列指定为非 NULL
AUTO_INCREMENT  指定列是否为自增类型
INDEX   创建具有或不带名称的索引, 如果多个索引同名则创建复合索引
UNIQUE_INDEX    和 INDEX 类似，只不过创建的是唯一索引
EMBEDDED    将结构设置为嵌入
EMBEDDED_PREFIX 设置嵌入结构的前缀
*/
type TreatType struct {
	gorm.Model
	BusModel
	Name     string `json:"description" gorm:"unique;column:name;type:varchar(255);not null"`
	Describe string `json:"describe" gorm:"type:text;not null"`
	Level int `json:"level" gorm:"type:tinyint;default 1;not null"`
	Region int `json:"region" gorm:"type:tinyint;default 1;not null"`
}

//新增
func (tt *TreatType) Create() (err error) {
	findOne := qmsql.DEFAULTDB.Where("name = ?", tt.Name).Find(&TreatType{}).Error
	if findOne == nil {
		return errors.New("存在相同name")
	} else {
		err = qmsql.DEFAULTDB.Create(tt).Error
	}
	return err
}

//删除
func (tt *TreatType) Delete() (err error) {
	err = qmsql.DEFAULTDB.Delete(tt).Error
	return err
}

//更新
func (tt *TreatType) Update() (err error) {
	var oldA TreatType
	err = qmsql.DEFAULTDB.Where("id = ?", tt.ID).First(&oldA).Error
	if err != nil {
		return err
	} else {
		err = qmsql.DEFAULTDB.Save(tt).Error
	}
	return err
}

//根据Id获取
func (tt *TreatType) GetById(id float64) (err error, treatType TreatType) {
	err = qmsql.DEFAULTDB.Where("id = ?", id).First(&treatType).Error
	return
}

// 获取所有api信息
func (tt *TreatType) GetAlls() (err error, treatTypes []TreatType) {
	err = qmsql.DEFAULTDB.Find(&treatTypes).Error
	return
}

// 分页获取数据  需要分页实现这个接口即可
func (tt *TreatType) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(tt, info)
	fmt.Println("GetInfoList total ------>> ",total)
	if err != nil {
		return
	} else {
		var resList []TreatType
		//Where("name LIKE ?", "%"+tt.Name+"%").
		err = qmsql.DEFAULTDB.Find(&resList).Count(&total).Error
		if err!=nil{
			return err, resList, total
		}else{
			//Where("name LIKE ?", "%"+tt.Name+"%")
			err = db.Order("id", true).Find(&resList).Error
		}
		return err, resList, total
	}
}
