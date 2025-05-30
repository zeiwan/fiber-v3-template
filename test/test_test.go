package test

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func createPassword(plaintext, salt string) string {
	// 第一步：计算 plaintext + salt 的 MD5
	innerHash := md5.Sum([]byte(plaintext + salt))
	innerHashStr := hex.EncodeToString(innerHash[:])

	// 第二步：计算 salt + innerHash 的 MD5
	finalHash := md5.Sum([]byte(salt + innerHashStr))
	return hex.EncodeToString(finalHash[:])
}

func TestName(t *testing.T) {
	println(createPassword("mogen123", "43dc"))

}
