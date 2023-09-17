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
	"encoding/base32"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testOtpSecret = "SEYAQXI5QGY7VSU4CQP2I4BSBU"

func TestGenerateNewSecret(t *testing.T) {
	s, e := GenerateNewSecret()
	assert.Nil(t, e)
	assert.Equal(t, 26, len(s))

	b, e := base32.StdEncoding.WithPadding(base32.NoPadding).
		DecodeString(s)
	assert.Nil(t, e)
	assert.Equal(t, 16, len(b))
}

func TestBuildTotpUri(t *testing.T) {
	expected := "otpauth://totp/Issuer:test@example.com" +
		"?issuer=Issuer&secret=SECRET"
	assert.Equal(t, expected, BuildTotpUri("test@example.com", "Issuer",
		"SECRET"))
}

func TestGenerateToken(t *testing.T) {
	_, err := GenerateToken("not base32 encoded", 1)
	assert.NotNil(t, err)
	assert.Equal(t, "illegal base32 data at input byte 0", err.Error())

	token, err := GenerateToken(testOtpSecret, 1)
	assert.Nil(t, err)
	assert.Equal(t, "400163", token)
}

func TestVerifyToken(t *testing.T) {
	_, err := VerifyToken("000000", "not base32 encoded")
	assert.NotNil(t, err)
	assert.Equal(t, "illegal base32 data at input byte 0", err.Error())

	ok, err := VerifyToken("000000", testOtpSecret)
	assert.Nil(t, err)
	assert.False(t, ok)

	interval := int(time.Now().UTC().Unix() / 30)
	token, err := GenerateToken(testOtpSecret, interval)
	assert.Nil(t, err)

	ok, err = VerifyToken(token, testOtpSecret)
	assert.Nil(t, err)
	assert.True(t, ok)
}
