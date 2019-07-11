package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

//提供HR系统 RSA非对称加密算法

// var publicKey = []byte(`
// -----BEGIN PUBLIC KEY-----
// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrGh1sc5AKD1EQ8WdA1iWF4m7w
// XtO6WoS7Dtfd0Jm2ud+LKBQ+e7R6YIXnwfEKB/4Jm+jNtCi7/Zrx5gtEpUuVAyrE
// o5+qr5al5KibeJq3xyI/626IBsDMFX5o3WOoXceTF7+lgi6r+OuokqFJgpeh7YAN
// XQ8Y8mn8ucw+Ly+LbQIDAQAB
// -----END PUBLIC KEY-----
// `)

// var privateKey = []byte(`
// -----BEGIN RSA PRIVATE KEY-----
// MIICXQIBAAKBgQCrGh1sc5AKD1EQ8WdA1iWF4m7wXtO6WoS7Dtfd0Jm2ud+LKBQ+
// e7R6YIXnwfEKB/4Jm+jNtCi7/Zrx5gtEpUuVAyrEo5+qr5al5KibeJq3xyI/626I
// BsDMFX5o3WOoXceTF7+lgi6r+OuokqFJgpeh7YANXQ8Y8mn8ucw+Ly+LbQIDAQAB
// AoGAGgoxbC3yP/WwyrlSk4WD1Gpvo9lqs7PO+4D4zWNP4YVMRitlWVUOVImYF3tm
// qbYprWCy/4tpn6KrECGImXvmkplXPxd4x3W+haZftx3VjTwh5fvT9yHp4swXxN+h
// LMItDdIOWS4U6wVJa77Dy7VfK303LZrPLqnxkf4oEywp5YECQQDZOz1WD7nOqOiy
// AlwDhfeLTmArN0f+gV6RLrxMp2XRqC2DN5nMq5O5BVVMK9LBgArNqYfxWYuMa3K2
// qliRDPPxAkEAyaNWq/fDvjpK9TgztqsHIiG+cUQpWI759zt5qHNA+QF4L43dtAVZ
// zBR/uam1jnRuM6K0ZCSZo2ITiqapmk8bPQJAEd9d3IbOssIS4xJun5uWElAQeX3C
// 3p2mOiuuMmBTcDx2AiXA8aXsMXzO18WDQYhXWzRniuPjJ1pvxbeeMdDvAQJBAMDh
// uZAJEzrOAlQurfFICyvQQZ+Rx0dKhbzFLOxBS96mVDSRLYn+MFbzKPcOa3lY0O4d
// 7xd4l2td7zmLkePlVjUCQQCY8VuIfKc0+AWvPnktKXbx9bBdJZSDginZM5cu7pdx
// W0uB9KZoLqgbGLIvWrLyA6SBqo87Q1j1//wFgLP+A2Gn
// -----END RSA PRIVATE KEY-----
// `)

//加密公钥
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCkGMFymAvXLH8nYtvSZqSxX9bRyp
X37bL/J1BoZsQpueVrntaYYfIJfGm/AxkcNcwB9fjAzXcDQxQlyL51ofZ0YjeqcmBX/
Zqire0GEHWdLkWJ4/iYCfomISRmPABXEmHA5mZB1fBP0Ag89Bz2+7pa+mfafMSh2e4i
GIQ8NdBbQIDAQAB
-----END PUBLIC KEY-----
`)

//解密私钥
var privateKey = []byte(`
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAMKQYwXKYC9csfyd
&#xD;&#xA;i29JmpLFf1tHKlfftsv8nUGhmxCm55Wue1phh8gl8ab8DGRw1zAH1+
MDNdwNDFCX&#xD;&#xA;IvnWh9nRiN6pyYFf9mqKt7QYQdZ0uRYnj+JgJ+iYhJGY
8AFcSYcDmZkHV8E/QCDz&#xD;&#xA;0HPb7ulr6Z9p8xKHZ7iIYhDw10FtAgMBAA
ECgYEAn0Dl/Jhg0IOUIEyoE+hwQFCt&#xD;&#xA;5P3EN/civadA5LatoRyslEUk
LJ+GL5p3SRIn5pLCYEsbN3KqRDrd6J09ALjLqlwK&#xD;&#xA;ZKXbi4hsPgTSdd
/bWiAIdf+hqr8vWGvCHPIVWnkQcNBTLlKtFc7JNgIKzbCZRxQn&#xD;&#xA;WtUP
HH+k8hc+Ob6OZwECQQD5gNWCB24n0ZTI6xUYJS+woLAbXlKGX8+ANAg4X24m&#xD
;&#xA;qkWxI6PmKVZej8aUMiE6YtNVtdDtWrBYAH5V9IG+tPytAkEAx6FVMOahA4
aS/NNS&#xD;&#xA;paR8OByWqtN+VRP+X9p8Ytb2sz+P0CdE14ThKhARK428ToYZ
PTkrs66cNm5CGB8g&#xD;&#xA;dyUvwQJBAJCx5aCGHJ0dD1NB+jbJghHF7rvAhM
2HDPiFtGq09VWZE9e6GpglSwCG&#xD;&#xA;ExzowZpxq6weSC8OlAxFJP9GUGQ/
4/UCQDk5+3ToODINiudlIOUREPb44wwXUrjK&#xD;&#xA;4XnS5SNkYhYiW3SdPT
PXCMEJGBL3L4sHEAcn82ov3OIRm2rUyXa+N0ECQCBcQwku&#xD;&#xA;3IIin7n4
mJqPue4Y80I4WoYYrfjLOD96D0YCQFz9q4sLsu6WwzY1bNuzxqyqRdLN&#xD;&#x
A;x2lkD+DZuKz5BTY=
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// func main() {
// 	data, _ := RsaEncrypt([]byte("test dataΩ......"))
// 	fmt.Println(base64.StdEncoding.EncodeToString(data))
// 	origData, _ := RsaDecrypt(data)
// 	fmt.Println(string(origData))
// }
