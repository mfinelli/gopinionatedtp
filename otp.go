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
// tokens.
package otp

import (
	// "crypto/hmac"
	"crypto/rand"
	"encoding/base32"
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
