package pfx

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Pfx(password string) {
	// 从文件加载PFX文件
	pfxData, err := ioutil.ReadFile("PIX-HMG-CLIENTE.pfx")
	if err != nil {
		fmt.Println("读取PFX文件时发生错误:", err)
		return
	}

	// 提取证书和私钥
	cert, err := extractCertsAndKeyFromPFX(pfxData, password)
	if err != nil {
		fmt.Println("提取证书和私钥时发生错误:", err)
		return
	}

	// 创建TLS配置
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	// 创建带有自定义TLS配置的HTTP客户端
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	url := "https://api.qrcodes-h.sulcredi.coop.br/oauth/token"
	method := "POST"

	payload := strings.NewReader(`{
		"grant_type": "client_credentials",
		"client_id": "00011203699012079000138",
		"client_secret": "zY5YjhlNTktYmMwOC00NjBhLWJlNDAtY"
	}`)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// extractCertsAndKeyFromPFX 从PFX文件中提取证书和私钥
func extractCertsAndKeyFromPFX(pfxData []byte, password string) (tls.Certificate, error) {
	return tls.LoadX509KeyPair("PIX-HMG-CLIENTE.crt", "PIX-HMG-CLIENTE.key")
}
