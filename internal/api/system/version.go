package system

import (
	api "Open_IM/pkg/base_info"
	"Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// wgt版本
func WgtVersion(c *gin.Context) {
	var params api.WgtVersionReq
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	info, err := im_mysql_model.GetNewWgtVersion(params.AppId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": map[string]string{
			"app_id":  params.AppId,
			"version": "",
			"url":     "",
			"remarks": "",
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": map[string]string{
		"app_id":  params.AppId,
		"version": info.Version,
		"url":     info.Url,
		"remarks": info.Remarks,
	}})
	return
}
