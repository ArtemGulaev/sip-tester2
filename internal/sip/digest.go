package sip

import (
	"crypto/md5"
	"fmt"
)

func md5Hash(s string) string {
	sum := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", sum)
}

func BuildDigest(
	username,
	password,
	realm,
	nonce,
	method,
	uri string,
) string {

	ha1 := md5Hash(
		fmt.Sprintf("%s:%s:%s",
			username,
			realm,
			password,
		),
	)

	ha2 := md5Hash(
		fmt.Sprintf("%s:%s",
			method,
			uri,
		),
	)

	return md5Hash(
		fmt.Sprintf("%s:%s:%s",
			ha1,
			nonce,
			ha2,
		),
	)
}
