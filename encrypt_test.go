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
	"testing"

	"github.com/stretchr/testify/assert"
)

const testEncryptionKey = "7oCf0u+l1nCr/01cmLd2yrLjXjgiuH17VhmfjhXlreA="

func TestEncryptOtpSecret(t *testing.T) {
	key, err := RawKeyToBytes(testEncryptionKey)
	assert.Nil(t, err)
	assert.Equal(t, 32, len(*key))

	sec, err := GenerateNewSecret()
	assert.Nil(t, err)
	assert.Equal(t, 26, len(sec))

	enc, err := EncryptOtpSecret(sec, key)
	assert.Nil(t, err)
	assert.Equal(t, 88, len(enc))
}

func TestDecryptOtpSecret(t *testing.T) {
	key, err := RawKeyToBytes(testEncryptionKey)
	assert.Nil(t, err)
	assert.Equal(t, 32, len(*key))

	enc, err := EncryptOtpSecret("test secret", key)
	assert.Nil(t, err)
	assert.Equal(t, 68, len(enc))

	dec, err := DecryptOtpSecret(enc, key)
	assert.Nil(t, err)
	assert.Equal(t, "test secret", dec)
}
