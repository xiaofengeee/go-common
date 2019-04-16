package thirdplatform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WxaAuthRet struct {
	Errcode      int64  `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

type WxUserInfoRet struct {
	Errcode  int64  `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Openid   string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Province string `json:"province"`
	City     string `json:"city"`
	Country  string `json:"country"`
	HeadUrl  string `json:"headimgurl"`
}

//WX_GetRetRedirectOpenId 获取微信openid
func Wx_GetRetRedirectOpenId(appId, appSec, code string) (*WxaAuthRet, error) {
	urlUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appId, appSec, code)

	res, err := http.Get(urlUrl)
	if err != nil {
		return nil, fmt.Errorf("登陆异常")
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("登陆读取异常")
	}

	var stRet WxaAuthRet
	if err := json.Unmarshal(buf, &stRet); err != nil {
		return nil, fmt.Errorf("登陆解析异常")
	}

	if stRet.Errcode != 0 {
		return nil, fmt.Errorf("登陆失败：%s", stRet.Errmsg)
	}

	return &stRet, nil
}

//读取信息
func Wx_GetBaseInfo(openid, code string) (*WxUserInfoRet, error) {
	urlUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", code, openid)

	res, err := http.Get(urlUrl)
	if err != nil {
		return nil, fmt.Errorf("获取信息异常")
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取信息异常")
	}

	var stRet WxUserInfoRet
	if err := json.Unmarshal(buf, &stRet); err != nil {
		return nil, fmt.Errorf("读取信息解析异常")
	}

	if stRet.Errcode != 0 {
		return nil, fmt.Errorf("读取信息失败：%s", stRet.Errmsg)
	}

	return &stRet, nil
}

func Wx_Login(code string) (*WxaAuthRet, *WxUserInfoRet, error) {
	authInfo, err := WX_GetRetRedirectOpenId(code)
	if err != nil {
		return nil, nil, err
	}

	userInfo, err := Wx_GetBaseInfo(authInfo.Openid, authInfo.AccessToken)
	if err != nil {
		return nil, nil, err
	}

	return authInfo, userInfo, nil
}
