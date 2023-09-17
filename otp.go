/*!
 * Copyright 2023 Mario Finelli
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// The otp package provides an opinionated way to generate and verify TOTP
// tokens as well as providing an easy way to encrypt the secrets for
// persistent storage (e.g., in the database alongside user records). It also
// provides a facility to dump QR codes to the terminal so that they can be
// scanned by an authenticator application.
package otp

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"time"
)

// GenerateNewSecret does what it says on the tin: it generates a new secret
// that is suitable for use as a TOTP secret.
func GenerateNewSecret() (string, error) {
	// the google-authenticator-libpam project uses 16 bytes
	key := make([]byte, 16)

	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	return base32.StdEncoding.WithPadding(base32.NoPadding).
		EncodeToString(key), nil
}

// BuildTotpUri returns a otpauth:// scheme URI suitable for input into a QR
// code generator. It requires the username of the user, an issuer and the
// user's OTP secret (generated with GenerateNewSecret) and reurns the URI:
//
// otpauth://totp/[issuer]:[username]?issuer=[issuer]&secret=[secret]
func BuildTotpUri(username, issuer, secret string) string {
	u := url.URL{
		Scheme: "otpauth",
		Host:   "totp",
	}

	u.Path = fmt.Sprintf("%s:%s", issuer, username)

	v := url.Values{}
	v.Set("secret", secret)
	v.Set("issuer", issuer)
	u.RawQuery = v.Encode()

	return u.String()
}

// GenerateToken returns the current time-based token for the given secret
// and interval.
func GenerateToken(secret string, interval int) (string, error) {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).
		DecodeString(secret)
	if err != nil {
		return "", err
	}

	length := 6 // 6 character codes

	// totp: https://datatracker.ietf.org/doc/html/rfc6238
	// hotp: https://datatracker.ietf.org/doc/html/rfc4226
	hash := hmac.New(sha1.New, key)
	err = binary.Write(hash, binary.BigEndian, int64(interval))
	if err != nil {
		return "", err
	}
	sign := hash.Sum(nil)

	offset := sign[19] & 15
	truncated := binary.BigEndian.Uint32(sign[offset : offset+4])
	return fmt.Sprintf("%0*d", length, (truncated&0x7fffffff)%1000000), nil
}

// VerifyToken returns true if the provided token matches the current token
// (+/- one interval) and false otherwise.
func VerifyToken(providedToken, secret string) (bool, error) {
	interval := int(time.Now().UTC().Unix() / 30) // 30s period

	for i := interval - 1; i <= interval+1; i++ {
		token, err := GenerateToken(secret, i)
		if err != nil {
			return false, err
		}

		if token == providedToken {
			return true, nil
		}
	}

	return false, nil
}
