package main

import (
	"JwsAlgsTest/alg"
	"JwsAlgsTest/config"
	"JwsAlgsTest/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DmitriyVTitov/size"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	rsaKeys   model.RSA
	ecdsaKeys model.ECDSA
)

func init() {
	rsaKeys = *alg.GenRSAKeys()
	ecdsaKeys = *alg.GenECDSAKeys()
}

func main() {
	jsonString := config.JsonString

	rs256TokenStr := createAndMeasure(jsonString, jwt.SigningMethodRS256)
	verifyAndMeasure(rs256TokenStr, jwt.SigningMethodRS256)

	es256TokenStr := createAndMeasure(jsonString, jwt.SigningMethodES256)
	verifyAndMeasure(es256TokenStr, jwt.SigningMethodES256)
}

func createAndMeasure(jsonStr string, alg jwt.SigningMethod) (tokenStr string) {
	fmt.Printf("Length of claims: %d.\n", len(jsonStr))
	fmt.Printf("Size of claims: %d.\n", size.Of(jsonStr))
	start := time.Now()
	tokenStr = createToken(jsonStr, alg)
	elapsed := time.Since(start)
	fmt.Printf("--%s start--\n", alg.Alg())
	fmt.Println(tokenStr)
	fmt.Printf("--%s end--\n", alg.Alg())

	fmt.Printf("Length of %s token: %d.\n", alg.Alg(), len(tokenStr))
	fmt.Printf("Size of %s token: %d.\n", alg.Alg(), size.Of(tokenStr))

	log.Printf("%s sign took %s\n", alg.Alg(), elapsed)
	return
}

func verifyAndMeasure(tokenStr string, alg jwt.SigningMethod) {
	start := time.Now()
	token := verifyToken(tokenStr, alg)
	elapsed := time.Since(start)
	fmt.Printf("%s valid: %t \n", alg.Alg(), token.Valid)
	log.Printf("%s verify took %s", alg.Alg(), elapsed)
}

func createToken(jsonStr string, alg jwt.SigningMethod) (tokenStr string) {

	token := jwt.NewWithClaims(alg, getClaims(jsonStr))
	uuid := uuid.New()
	key := uuid.String()
	delete(token.Header, "typ")
	token.Header["kid"] = key

	err := error(nil)
	switch alg {
	case jwt.SigningMethodES256:
		if tokenStr, err = token.SignedString(ecdsaKeys.PrivateKey); err != nil {
			log.Printf("Signing Error -> es256\n")
			os.Exit(1)
		}
		break
	case jwt.SigningMethodRS256:
		if tokenStr, err = token.SignedString(rsaKeys.PrivateKey); err != nil {
			log.Printf("Signing Error -> rs256\n")
			os.Exit(1)
		}
		break
	}
	return
}

func getClaims(jsonStr string) (claims jwt.MapClaims) {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &result)
	claims = result
	return
}

func verifyToken(tokenStr string, alg jwt.SigningMethod) (token *jwt.Token) {
	err := error(nil)
	switch alg {
	case jwt.SigningMethodES256:
		if token, err = jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return ecdsaKeys.PublicKey, nil
		}); err != nil {
			log.Printf("Parse token Error -> es256\n %s", err)
			os.Exit(1)
		}
	case jwt.SigningMethodRS256:
		if token, err = jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return rsaKeys.PublicKey, nil
		}); err != nil {
			log.Printf("Parse token Error -> rs256\n %s", err)
			os.Exit(1)
		}
	}
	return
}
