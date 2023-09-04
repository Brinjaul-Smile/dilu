package apis

import (
	"dilu/common/codes"
	"dilu/modules/sys/service"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type CaptchaApi struct {
	base.BaseApi
}

// GenerateCaptchaHandler 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 登陆
// @Success 200 {object} base.Resp{data=map[string]string} "{"code": 200, "data": [...]}"
// @Router /api/v1/captcha [get]
func (e CaptchaApi) GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := service.DriverDigitFunc()
	if err != nil {
		core.Log.Error("生成验证码失败", zap.Error(err))
		e.Err(c, errs.Err(codes.CaptchaErr, e.GetReqId(c), err))
		return
	}
	m := map[string]string{
		"data": b64s,
		"id":   id,
	}
	e.Ok(c, m)
}