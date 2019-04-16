package thirdplatform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xiaofengeee/go-common/logger"
)

type FbIdsForBusinessRst struct {
	Data []struct {
		Id  string `json:"id"`
		App struct {
			Name      string `json:"name"`
			NameSpace string `json:"namespace"`
			Id        string `json:"id"`
		} `json:"app"`
	} `json:"data"`
}

//fb 检测用户
func Fb_CheckLogin(fbUrl, appId, token string) (string, error) { //根据token再取账号
	strUrl := fmt.Sprintf("%s/me/ids_for_business?access_token=%s", fbUrl, token)
	res, err := http.Get(strUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	logger.Info("FB_CheckLogin ids_for_business", string(buf))

	var stBus FbIdsForBusinessRst
	if err := json.Unmarshal(buf, &stBus); err != nil {
		return "", err
	}

	userId := ""
	for _, v := range stBus.Data {
		if v.App.Id == appId {
			userId = v.Id
			break
		}
	}

	if userId == "" {
		return "", fmt.Errorf("no user")
	}

	return userId, nil
}
