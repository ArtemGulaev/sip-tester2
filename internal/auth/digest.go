package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func md5Hex(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

func BuildResponse(
	username string,
	password string,
	realm string,
	nonce string,
	method string,
	uri string,
) string {

	ha1 := md5Hex(
		fmt.Sprintf(
			"%s:%s:%s",
			username,
			realm,
			password,
		),
	)

	ha2 := md5Hex(
		fmt.Sprintf(
			"%s:%s",
			method,
			uri,
		),
	)

	return md5Hex(
		fmt.Sprintf(
			"%s:%s:%s",
			ha1,
			nonce,
			ha2,
		),
	)
}
