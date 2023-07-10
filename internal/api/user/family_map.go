package user

import (
	"Open_IM/internal/api/common"
	api "Open_IM/pkg/base_info"
	imdb "Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportUserLocation(c *gin.Context) {
	var (
		params api.ReportUserLocationReq
	)

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusOK, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	reports := &imdb.FUserReports{
		UserId:    userId,
		Latitude:  params.Latitude,
		Longitude: params.Longitude,
		Battery:   int(params.Battery),
		Step:      int(params.Step),
	}
	// 保存用户数据
	err := imdb.CreateUserLocation(reports)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// todo 这里需要进行上报高德地图，进行围栏判断，走异步方式

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "errMsg": "上报成功"})
}

func GetGroupLocationList(c *gin.Context) {
	var (
		params api.GetGroupLocationListReq
	)

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusOK, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	_, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	// 1. 获取群成员列表
	group, err := imdb.GetGroupInfoByGroupID(params.GroupID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	if group.GroupType != 2 /* 家庭群*/ {
		c.JSON(http.StatusOK, gin.H{"errCode": 400, "errMsg": "该群不是家庭群"})
		return
	}

	// 2.获取所有的群成员
	members, err := imdb.GetGroupMemberListByGroupID(params.GroupID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	// userStr = 15,16,17,18
	var userStr string
	for num, member := range members {
		if num == len(members)-1 {
			userStr += member.UserID
			break
		}
		userStr += member.UserID + ","
	}

	// 3.查询群成员最后一条位置信息
	reports, err := imdb.GetUserLocationList(userStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//{
	//  "code": 200,
	//  "data": {
	//    "common_resp": {
	//      "code": 200,
	//      "msg": "设置成功"
	//    }
	//  }
	//}
	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": gin.H{"common_resp": gin.H{"code": 200, "msg": "设置成功"}, "list": reports}})
}
