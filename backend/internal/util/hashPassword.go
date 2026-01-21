package util

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 生成 bcrypt 哈希
func HashPassword(plain string) (string, error) {
	// bcrypt 默认成本 10，如果需要更高安全可以改成 12
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword 校验明文密码和哈希是否匹配
func CheckPassword(hash, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
