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

package otp

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

// https://pkg.go.dev/golang.org/x/crypto/nacl/secretbox#example-package

func EncryptOtpSecret(secret string, key *[32]byte) (string, error) {
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return "", err
	}

	encrypted := secretbox.Seal(nonce[:], []byte(secret), &nonce, key)

	r := base64.StdEncoding.EncodeToString(encrypted)
	return r, nil
}

func DecryptOtpSecret(b64Secret string, key *[32]byte) (string, error) {
	secret, err := base64.StdEncoding.DecodeString(b64Secret)
	if err != nil {
		return "", err
	}

	// during the seal, the secret is appended to the nonce so we can
	// extract the nonce from the first 24 bytes of the secret
	var nonce [24]byte
	copy(nonce[:], secret[:24])

	decrypted, ok := secretbox.Open(nil, secret[24:], &nonce, key)
	if !ok {
		return "", fmt.Errorf("decryption error")
	}

	return string(decrypted), nil
}

func RawKeyToBytes(rawKey string) (*[32]byte, error) {
	if rawKey == "" {
		return nil, fmt.Errorf("encryption key is unset")
	}

	key, err := base64.StdEncoding.DecodeString(rawKey)
	if err != nil {
		return nil, err
	}

	var keyBytes [32]byte
	copy(keyBytes[:], key)

	return &keyBytes, nil
}
