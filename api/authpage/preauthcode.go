package authpage

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/gin-gonic/gin"
)

type getPreAuthCodeReq struct {
	ComponentAppid string `wx:"component_appid"`
}

type getPreAuthCodeResp struct {
	PreAuthCode string `wx:"pre_auth_code"`
}

func getPreAuthCodeHandler(c *gin.Context) {
	req := getPreAuthCodeReq{
		ComponentAppid: wxbase.GetAppid(),
	}
	log.Infof("Appid",wxbase.GetAppid())
	log.Infof("getpro-auth-code",req)
	_, body, err := wx.PostWxJsonWithComponentToken("/cgi-bin/component/api_create_preauthcode", "", req)
	if err != nil {
		log.Errorf("1111111 err, %v", errno.ErrSystemError.WithData(err.Error()))
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	log.Infof("PostWxJsonWithComponentToken",err)
	var resp getPreAuthCodeResp
	if err := wx.WxJson.Unmarshal(body, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{
		"preAuthCode": resp.PreAuthCode,
	}))
}
