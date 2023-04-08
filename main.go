package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Alg interface {
	
}

type RSA struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

type ECDSA struct {
	publicKey  *ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

var (
	rsaKeys   RSA
	ecdsaKeys ECDSA
)

const (
	jsonString = `{
		"sub": "member001",
		"aud": "kioxiaApp2",
		"nbf": 1661700103,
		"scope": [
		  "read",
		  "openid",
		  "write"
		],
		"iss": "http://192.168.0.100:30200",
		"exp": 1662304903,
		"iat": 1661700103,
		"device": {
		  "deviceId": 1,
		  "employee": null,
		  "name": "Win",
		  "firewall": true,
		  "antivirus": false,
		  "os": "Windows",
		  "osVersion": "21H2",
		  "osLatestUpdate": 1516239022,
		  "deviceType": "Computer",
		  "firstLoginTime": 1661700100159,
		  "riskAnalysis": "Unevaluated",
		  "deviceActivityList": [
			{
			  "deviceActivityId": 1,
			  "ipAddress": "0:0:0:0:0:0:0:1",
			  "activityStatus": "Latest",
			  "loginTime": 1661700100165
			}
		  ]
		}
	  }`
)

func init() {
	genRSAKeys()
	genECDSAKeys()
}

func main() {
	// token, _ := createToken(jwt.SigningMethodRS256)
	// fmt.Println(token)

	fmt.Print(getClaims(jsonString))
}

func genRSAKeys() {
	err := error(nil)
	if rsaKeys.privateKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	rsaKeys.publicKey = &rsaKeys.privateKey.PublicKey
}

func genECDSAKeys() {
	err := error(nil)
	if ecdsaKeys.privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}

	ecdsaKeys.publicKey = &ecdsaKeys.privateKey.PublicKey
}

func createToken(alg jwt.SigningMethod) (tokenStr string, err error) {
	token := jwt.NewWithClaims(alg, getClaims(jsonString))
	if(alg.Alg() == jwt.SigningMethodES256) 
	return token.SignedString(privateKey)
}

func getClaims(jsonStr string) (claims jwt.MapClaims) {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &result)
	claims = result
	return
}
