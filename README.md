# (gO)pinionatedTP

[![CI](https://github.com/mfinelli/gopinionatedtp/actions/workflows/default.yml/badge.svg)](https://github.com/mfinelli/gopinionatedtp/actions/workflows/default.yml)
[![Go Reference](https://pkg.go.dev/badge/go.finelli.dev/gopinionatedtp.svg)](https://pkg.go.dev/go.finelli.dev/gopinionatedtp)

An opinionated TOTP library for Golang.

## background

The goal of this library is basically to work with the standard OTP provider
applications (2FAS, Authy, Google Authenticator, etc) without needing to
figure out what the correct options to use are. There are a few other Golang
libraries that provide HOTP/TOTP functionality but they provide too many other
configuration options for the use-case outlined above. Therefore this library
is extremely opinionated to only provide what's necessary to work with the
standard OTP applications.

That being said, this library takes inspiration from those libraries as well
as another blog post on implementing OTP in Go:

- https://github.com/RijulGulati/otpgen
- https://github.com/jltorresm/otpgo
- http://www.inanzzz.com/index.php/post/y5nu/creating-a-one-time-password-otp-library-for-two-factor-authentication-2fa-with-golang

## usage

Below is an example how to generate a new secret for some user, encrypt it
(saving it to the database is left as an exercise for the reader), create a QR
code and then generate a token and verify it.

**N.B.** the secrets need to be stored in the database _encrypted_ not hashed
(like for passwords) as it's necessary to access the secret in order to verify
a provided token.

```go
package main

import "crypto/rand"
import "encoding/base64"
import "fmt"

import gotp "go.finelli.dev/gopinionatedtp"

func main() {
	// generate an encryption key (you should persist this somewhere
	// secure!)
	r := make([]byte, 32)
	rand.Read(r)
	key := base64.StdEncoding.EncodeToString(r)

	// generate a new secret for a given user
	secret, err := gotp.GenerateNewSecret()
	if err != nil {
		panic(err)
	}

	// encrypt the secret to store it in the database
	crypt, err := gotp.EncryptOtpSecret(secret, key)
	if err != nil {
		panic(err)
	}

	// ...save the result to the user record

	// dump a QR code to the terminal to scan with an authenticator app
	err = gotp.QrCodeToTerminal("user@example.com", "yourapp", secret)
	if err != nil {
		panic(err)
	}

	// ...retrieve the encrypted secret from the database

	// decrypt the secret for use
	secret, err := gotp.DecryptOtpSecret(crypt, key)
	if err != nil {
		panic(err)
	}

	// calculate the current token (for fun!)
	token, err := gotp.GenerateToken(secret, gotp.DefaultInterval())
	if err != nil {
		panic(err)
	}

	// validate a given token
	valid, err := gotp.VerifyToken("123456", secret)
	if err != nil {
		panic(err)
	}

	if valid {
		fmt.Println("Valid token!")
	} else {
		fmt.Println("Invalid token!")
	}
}
```

## license

```
Copyright 2023-2024 Mario Finelli

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
