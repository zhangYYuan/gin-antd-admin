package business

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/busModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
)


type CreateTreatTypeParams struct {
	Name        string `json:"name"`
	Describe string `json:"describe"`
}

type DeleteTreatTypeParams struct {
	ID uint `json:"id"`
}

// @Tags TreatType
// @Summary 创建TreatType
func Create(c *gin.Context) {
	var treatType busModel.TreatType
	_ = c.BindJSON(&treatType)
	err := treatType.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("TreatType创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "TreatType创建成功", gin.H{})
	}
}

func Delete(c *gin.Context) {
	var treatType busModel.TreatType
	_ = c.BindJSON(&treatType)
	err := treatType.Delete()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("TreatType删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "TreatType删除成功", gin.H{})
	}
}

type GetByIdStruct struct {
	Id float64 `json:"id"`
}

func GetById(c *gin.Context) {
	var idInfo GetByIdStruct
	_ = c.BindJSON(&idInfo)
	err, api := new(busModel.TreatType).GetById(idInfo.Id)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("TreatType获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "TreatType获取数据成功", gin.H{
			"api": api,
		})

	}
}

func Update(c *gin.Context) {
	var treatType busModel.TreatType
	_ = c.BindJSON(&treatType)
	err := treatType.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("TreatType修改数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "TreatType修改数据成功", gin.H{})
	}
}

func GetAlls(c *gin.Context) {
	err, apis := new(busModel.TreatType).GetAlls()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取TreatType数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取TreatType数据成功", gin.H{
			"apis": apis,
		})
	}
}

func GetPageList(c *gin.Context) {
	type searchParams struct {
		busModel.TreatType
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	fmt.Println("searchParams json: ", sp.PageInfo)
	err, list, total := sp.TreatType.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     list,
			"total":    total,
			"page":     sp.PageInfo.Page,
			"pageSize": sp.PageInfo.PageSize,
		})

	}
}
