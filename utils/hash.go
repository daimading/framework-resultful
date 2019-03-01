package utils

import (
	"github.com/json-iterator/go"
	"time"

	"crypto/sha256"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
)

const KSaltLen = 10
const KEmailCaptchaLen = 6

var (
	digitTable1    = "123456789"
	digitTable2    = "1234567890"
	digitTable1Len = len(digitTable1)
	digitTable2Len = len(digitTable2)
)

func GenSalt() string {
	return string(RandStringK(KSaltLen, RandKindAll))
}
func GenEmailCaptcha() string {
	return string(RandStringK(KEmailCaptchaLen, RandKindNum))
}

const kTokenValidateTime = 1000 * time.Second

func GenTokenVt() (token string, vt time.Time) {
	uuid := uuid.NewV4()
	return uuid.String(), time.Now().Add(kTokenValidateTime)
}

func MarshalMap(m map[string]interface{}) ([]byte, error) {
	bs, err := jsoniter.Marshal(m)
	if err != nil {
		logrus.Error("marshal json errors and reason \n", err.Error())
		return nil, err
	}

	return bs, nil
}

func GenHashedPassword(password, salt string) string {
	h := sha256.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	return hex.EncodeToString(h.Sum(nil))
}
