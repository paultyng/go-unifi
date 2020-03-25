package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	awsAlgorithm      = "AWS4-HMAC-SHA256"
	awsService        = "iotdevicegateway"
	awsV4Request      = "aws4_request"
	shortTimeFormat   = "20060102"
	signatureQueryKey = "X-Amz-Signature"
	timeFormat        = "20060102T150405Z"

	// emptyStringSHA256 is a SHA256 of an empty string
	emptyStringSHA256 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
)

// from AWS v4 signer.go
func hashSHA256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// from AWS v4 signer.go
func hmacSHA256(key []byte, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

// from AWS v4 signer.go
func formatShortTime(dt time.Time) string {
	return dt.UTC().Format(shortTimeFormat)
}

// from AWS v4 signer.go
func formatTime(dt time.Time) string {
	return dt.UTC().Format(timeFormat)
}

// from AWS v4 signer.go
func deriveSigningKey(region, service, secretKey string, dt time.Time) []byte {
	kDate := hmacSHA256([]byte("AWS4"+secretKey), []byte(formatShortTime(dt)))
	kRegion := hmacSHA256(kDate, []byte(region))
	kService := hmacSHA256(kRegion, []byte(service))
	signingKey := hmacSHA256(kService, []byte(awsV4Request))
	return signingKey
}

func signAWSMQTTWSURL(host, region, accessKeyID, secretAccessKey, sessionToken string) (string, error) {
	now := time.Now()
	datetime := formatTime(now)
	date := formatShortTime(now)

	mqttURL, err := url.Parse(fmt.Sprintf("wss://%s/mqtt", host))
	if err != nil {
		return "", err
	}

	credentialScope := fmt.Sprintf("%s/%s/%s/aws4_request", date, region, awsService)

	qs := mqttURL.Query()

	qs.Set("X-Amz-Algorithm", awsAlgorithm)
	qs.Set("X-Amz-Credential", fmt.Sprintf("%s/%s", accessKeyID, credentialScope))
	qs.Set("X-Amz-Date", datetime)
	qs.Set("X-Amz-SignedHeaders", "host")

	mqttURL.RawQuery = strings.Replace(qs.Encode(), "+", "%20", -1)

	canonicalHeaders := fmt.Sprintf("host:%s\n", host)

	canonicalString := fmt.Sprintf("GET\n/mqtt\n%s\n%s\nhost\n%s", mqttURL.RawQuery, canonicalHeaders, emptyStringSHA256)

	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s", awsAlgorithm, datetime, credentialScope, hex.EncodeToString(hashSHA256([]byte(canonicalString))))

	creds := deriveSigningKey(region, awsService, secretAccessKey, now)
	signature := hex.EncodeToString(hmacSHA256(creds, []byte(stringToSign)))

	mqttURL.RawQuery += "&" + signatureQueryKey + "=" + signature + "&X-Amz-Security-Token=" + url.QueryEscape(sessionToken)

	return mqttURL.String(), nil
}
