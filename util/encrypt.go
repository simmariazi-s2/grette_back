package util

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func DecodeBase64(data string) string {
	result, _ := base64.StdEncoding.DecodeString(data)

	// if err != nil {
	// 	return ""
	// }
	return fmt.Sprintf("%x", result)
}

func errInvalidHash() error {
	return errors.New("올바른 해시 형태가 아닙니다")
}

func errInvalidParam() error {
	return errors.New("올바른 매개변수가 아닙니다")
}

func EncodeScrypt(password string) (string, error) {
	salt, err := GenerateRandomBytes(32)
	if err != nil {
		return "", err
	}

	dk, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)

	if err != nil {
		return "", err
	}
	params := log2(16384)<<16 | 8<<8 | 1
	encodeParam := fmt.Sprintf("%x", params)
	fmt.Printf("%x, %x\n", salt, dk)
	fmt.Printf("%s, %s", EncodeBase64(string(salt)), EncodeBase64(string(dk)))

	// 나중에 이걸로 바꾸고 싶음
	// return fmt.Sprintf("%d$%d$%d$%s$%s", 16384, 8, 1, EncodeBase64(string(salt)), EncodeBase64(string(dk))), nil
	return fmt.Sprintf("$s0$%s$%s$%s", string(encodeParam), EncodeBase64(string(salt)), EncodeBase64(string(dk))), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func CompareHashAndPassword(hash string, password string) bool {
	params, salt, dk, err := decodeHash(hash)
	if err != nil {
		return false
	}

	other, err := scrypt.Key([]byte(password), salt, params.N, params.R, params.P, params.DKLen)
	if err != nil {
		return false
	}

	if subtle.ConstantTimeCompare(dk, other) == 1 {
		return true
	}

	return false
}

func decodeHash(hash string) (param, []byte, []byte, error) {
	var err error
	var resultParam param
	vals := strings.Split(hash, "$")

	if len(vals) != 5 || vals[1] != "s0" {
		return param{}, nil, nil, errInvalidHash()
	}
	hashParam, err := strconv.ParseUint(vals[2], 16, 32)

	if err != nil {
		return resultParam, nil, nil, errInvalidHash()
	}

	resultParam.N = int(math.Pow(2, float64(hashParam>>16&0xffff)))
	resultParam.R = int(hashParam >> 8 & 0xff)
	resultParam.P = int(hashParam & 0xff)

	// resultParam.P == 0는 상황에 따라 가능 함
	if resultParam.N == 0 || resultParam.R == 0 {
		return resultParam, nil, nil, errInvalidHash()
	}

	salt, err := hex.DecodeString(DecodeBase64(vals[3]))
	if err != nil {
		return resultParam, nil, nil, errInvalidHash()
	}
	resultParam.SaltLen = len(salt)

	dk, err := hex.DecodeString(DecodeBase64(vals[4]))
	if err != nil {
		return resultParam, nil, nil, errInvalidHash()
	}
	resultParam.DKLen = len(dk)

	if err := resultParam.Validate(); err != nil {
		return resultParam, nil, nil, err
	}

	return resultParam, salt, dk, nil
}

func log2(n int) int {
	log := 0
	if (n & 0xffff0000) != 0 {
		n = n >> 16
		log = 16
	}
	if n >= 256 {
		n = n >> 8
		log += 8
	}
	if n >= 16 {
		n = n >> 4
		log += 4
	}
	if n >= 4 {
		n = n >> 2
		log += 2
	}
	return log + (n >> 1)
}

type param struct {
	N       int
	R       int
	P       int
	SaltLen int
	DKLen   int
}

func (p *param) Validate() error {
	if p.N > (1<<31-1) || p.N <= 1 || p.N%2 != 0 {
		return errInvalidParam()
	}

	if p.R < 1 || p.R > (1<<31-1) {
		return errInvalidParam()
	}

	if p.P < 1 || p.P > (1<<31-1) {
		return errInvalidParam()
	}

	if uint64(p.R)*uint64(p.P) >= 1<<30 || p.R > (1<<31-1)/128/p.P || p.R > (1<<31-1)/256 || p.N > (1<<31-1)/128/p.R {
		return errInvalidParam()
	}

	if p.SaltLen < 8 || p.SaltLen > (1<<31-1) {
		return errInvalidParam()
	}

	if p.DKLen < 16 || p.DKLen > (1<<31-1) {
		return errInvalidParam()
	}

	return nil
}
