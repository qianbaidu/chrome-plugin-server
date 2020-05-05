package util

import (
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"math/big"
	"strings"
)

const (
	SUCCESS        = 200
	ACCEPTED       = 202
	ERROR          = 501
	FAILED         = 502
	PARAM_ERROR    = 400
	AUTH_ERROR     = 403
	NO_AUTH        = 405
	LOGIN_TIME_OUT = 401
)

var ReturnMessage = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	FAILED:         "failed",
	PARAM_ERROR:    "parameter error",
	AUTH_ERROR:     "auth error",
	LOGIN_TIME_OUT: "login time out",
	NO_AUTH:        "Insufficient authority",
}

var CategoryMap = map[int]string{
	1: "常用",
	2: "稍后阅读",
	3: "工作",
	4: "学习",
	5: "生活",
	6: "娱乐",
	7: "其他",
	8: "资源",
}

func HtmlDecode(str string) string {
	str = strings.Replace(str, "<", "&lt;", -1)
	str = strings.Replace(str, ">", "&gt", -1)
	return str
}

func CheckUrl(url string) bool {
	if !strings.HasPrefix(url,"http") {
		return false
	}
	return true
}

func GetMysqlDns()  string {
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbUser := beego.AppConfig.String("dbUser")
	dbPassword := beego.AppConfig.String("dbPassword")
	dbName := beego.AppConfig.String("dbName")
	dns := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	return dns
}

func GetCategoryName(categoryId int) string {
	if c, ok := CategoryMap[categoryId]; ok {
		return c
	} else {
		log.Infof("分类id %d 暂未添加", categoryId)
		return "未知分类"
	}
}

func GetMsgByCode(code int) string {
	return ReturnMessage[code]
}

func JsonMsg(msg interface{}, code int, data interface{}) (returnData map[string]interface{}) {
	errorMsg := fmt.Sprintf("%s", msg)
	if len(errorMsg) < 1 {
		errorMsg = GetMsgByCode(code)
	}
	returnData = map[string]interface{}{"code": code, "msg": errorMsg, "data": data}
	return returnData
}

func GetUserSalt() (string, error) {
	return RandomChars(10)
}

func RandomChars(n int) (string, error) {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	randomInt := func(max *big.Int) (int, error) {
		r, err := rand.Int(rand.Reader, max)
		if err != nil {
			return 0, err
		}

		return int(r.Int64()), nil
	}

	buffer := make([]byte, n)
	max := big.NewInt(int64(len(alphanum)))
	for i := 0; i < n; i++ {
		index, err := randomInt(max)
		if err != nil {
			return "", err
		}

		buffer[i] = alphanum[index]
	}

	return string(buffer), nil
}
