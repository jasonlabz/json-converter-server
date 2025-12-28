package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlabz/potato/consts"

	base "github.com/jasonlabz/json-converter-server/common/ginx"
	"github.com/jasonlabz/json-converter-server/server/service/health_check"
)

// HealthCheck 健康检查
//
//	@Summary	健康检查
//	@Tags		健康检查
//	@Accept		json
//	@Produce	json
//	@Router		/health-check [get]
func HealthCheck(c *gin.Context) {
	status := health_check.GetService().DoCheck(c)
	base.JsonResult(c, consts.APIVersionV1, status, nil)
}
