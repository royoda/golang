package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

const (
	PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	PEM_END = "\n-----END RSA PRIVATE KEY-----"
)

func main() {
	msg := "hello world"
	privateKey := `MIIEowIBAAKCAQEAjAVTQ6OYVxFO/3ndRzMwXVKpm+KYwDy6fR8lP/6irVMQ3f8NmzVA6MHHn+5GzKnF9RXCYvUPWSx1vQsnc+R8T0r8X4yJAc+3VHf0iJhSasN3lw37H0tACEUG+ecFXYi6bRU0K+y1kbceKWmr2g1zxK7BJYplCUWHZ8iV9s5WJhhMZ/O8vdGYQp6TgRXt/w3sfVtw68spwqO2pbvZD9AZFFcviRtcVb+svIkoxP8zKVS8nLLO4yfwntSGQh80zAapRyjr6wpMoPTP+myyBPkY5LRpZg/QSEuX8Oxl9XRYmv4q8VPplxqbESd80na8FTrV/vuwhNxC/wBGFeEXoZzcMwIDAQABAoIBAHtmEgxxN8e71FAoeQ/PmBcWy607FP2OR5gPg8eTRJDVvO64YahgVVULuj9DQhgKoxAsXgTSDfW7zgONufZT7g3/es1GUFRNWDdUBCkBfNjkbRet1ZutuL/Q/aXtqHfXEN0jq7fuQ9IofKdonnBsHJPAoy3Pet2h7gRT7X32OEwXHeTSc7koJt6sDDVrcaOKSMEiagCt60SdBj/3qtjSJ4CMIKus5YHyWvrsvNDxHAMS43z64zvjqLgMpMsy+f2FeJE/D4/rsXeAG6MPrHyURwUNInnUgXPGSpXePH6nCMv5Ed9VjdUh1SdBqv0YeUCeDRX137wzqHYV9+/ItGlMIfECgYEAxlSgqi5V52gv2EoN/GG0fXvCPmqUvg7z8w3MDwww1dXYlX6jaukkNG+yGrkZVwD0Yr4Gb0dCvOCVU0ymYCtPv8VDcauFkTtbpt6QJpvqNIDXjXPI33KMjWDLsUl7Bd8TRfxUcxhvi0QX1AiBqfZ8Amf9l1BgtL9eDXLZPaDSVAsCgYEAtLw4/FIT6CyisUMiBytXJ3GicCrYCcefPsccVc8IOVo3V3aWzgcw5gSDlfP46jiGHwR5rNbmP/ksbJx/bNL2T/DJcre1sdJ1gYwJ7xj4iX6BRny9YrN8N6TT4wjnqaOtdCEiLAYMH/F7kY0Sg2//huxVXlzi5c+gPoq2PfY0SXkCgYAkJsTqgGu0NK+T81u9R2oI/YgYrIDUbBFXmoo5q/Uy5ToHdf//uuOYk0uZx2exDkF9xjmSCyFFIILXnj8thHeS8zgp/iPopzM7pZINF5qHW3zvEc3pvQ2Vr7Exbb15AoKW3TnfiPK8JftZ8PsBqiZMofv9mu8lw0Q3m1mx/CCapQKBgBrXXUsz0VsIzfyvcK1J5X6pJKut4TnKmL7VCUNBJQKHgHKP3SOp87wlbXlmq0/1pw2SN5PQso9LrSpQL6h0yDnlE7XizKwlCmh76LnGppqc3Awg2GuBJnCJHSK5ntYVxdvaU93f5AsWbPiXz+kPSA68qI0EFHp21GCVby1SlD9ZAoGBAMQ1aSRxPPBkmlZTMMSYBzbxjP+qpcTTqcyC81uuJAJo9cXOJeN+3glVwGeohWM2YeJJD4oVTNCpXqtBvS81YsruB9E9MNu2CiE/Noc2u6lT0N9Mv27GCkeQjmRf5Rb9v17naixXZiPXy4wuN2H1lH2NEdfoUGAEWwBsTH+FMPA5`
	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	rng := rand.Reader
	message := []byte(msg)
	hashed := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rng, priKey, crypto.SHA256, hashed[:])
	signatureHex := base64.StdEncoding.EncodeToString(signature)
	fmt.Println(signatureHex)
}

func ParsePrivateKey(privateKey string)(*rsa.PrivateKey, error) {
	privateKey = FormatPrivateKey(privateKey)
	// 2、解码私钥字节，生成加密对象
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}

func FormatPrivateKey(privateKey string) string  {
	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
		privateKey = PEM_BEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEM_END) {
		privateKey = privateKey + PEM_END
	}
	return privateKey
}
