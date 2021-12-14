package gin_use

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// app secret
	appSecret = "123456"

	// sign expire date (unit is second)
	expireTime = "3600"
)

func SetUp() gin.HandlerFunc {
	// check sign
	return func(c *gin.Context) {
		isPass, err := verifySign(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "系统验签错误",
			})
			c.Abort()
			return
		}

		if !isPass {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "验签错误",
			})
			c.Abort()
		}

		c.Next()
	}
}

//verifySign sing verify
func verifySign(c *gin.Context) (res bool, err error) {
	// 1. get sign content
	appKey, timestamp, sign, body, err := getSignNeedContent(c)
	if err != nil {
		return false, err
	}

	// 2. sign and back
	thisSign := md5V(appKey + appSecret + timestamp + body)
	return thisSign == sign, nil
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// get sign need content
func getSignNeedContent(c *gin.Context) (appKey, timestamp, sign, reqBody string, err error) {
	// 1. get app key
	appKey = strings.Join(c.Request.Header["Appkey"], "")
	if appKey == "" {
		return "", "", "", "", errors.New("app key is empty")
	}

	// 2. get timestamp
	timestamp = strings.Join(c.Request.Header["Timestamp"], "")
	if err = timestampCheck(timestamp); err != nil {
		return "", "", "", "", err
	}

	// 3. get sign
	sign = strings.Join(c.Request.Header["Sign"], "")
	if sign == "" {
		return "", "", "", "", errors.New("sign is empty")
	}

	// 4. get request body
	reqBytesBody, err := getReqBody(c.Request.Body)
	if err != nil {
		return "", "", "", "", err
	}
	reqBody = string(reqBytesBody)

	// 5. get nonce value
	// value equal md5(timestamp+rand(0,1000))
	if err = nonceCheck(strings.Join(c.Request.Header["Nonce"], "")); err != nil {
		return "", "", "", "", err
	}

	return
}

func getReqBody(in io.ReadCloser) (res []byte, err error) {
	res, err = ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	return
}

func timestampCheck(timestampStr string) (err error) {
	if timestampStr == "" {
		return errors.New("timestamp is empty")
	}

	// check timestamp
	timestamp := time.Now().Unix()
	exp, _ := strconv.ParseInt(expireTime, 10, 64)
	tsInt, _ := strconv.ParseInt(timestampStr, 10, 64)
	if tsInt > timestamp || timestamp-tsInt >= exp {
		return errors.New("timestamp is over time limit")
	}

	return
}

func nonceCheck(nonce string) (err error) {
	if nonce == "" {
		return errors.New("nonce value is empty")
	}

	// todo check if it exists in redis

	return
}
