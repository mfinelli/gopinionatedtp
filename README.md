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

## license

```
Copyright 2023 Mario Finelli

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
