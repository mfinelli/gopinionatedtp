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
	"fmt"

	"github.com/pterm/pterm"
	"github.com/yeqown/go-qrcode/v2"
)

// QrCodeToTerminal constructs a TOTP URI and then outputs a terminal-based
// (using pterm) QR code that can be scanned by an authentication app.
func QrCodeToTerminal(username, issuer, secret string) error {
	uri := BuildTotpUri(username, issuer, secret)
	qrc, _ := qrcode.New(uri)

	w := newTerminalWriter()

	if err := qrc.Save(w); err != nil {
		return err
	}

	return nil
}

// implements a terminal writer using pterm for qrcode
// https://github.com/yeqown/go-qrcode/tree/main/writer
var _ qrcode.Writer = (*terminalWriter)(nil)

type terminalWriter struct{}

func newTerminalWriter() *terminalWriter {
	w := &terminalWriter{}
	return w
}

func (w terminalWriter) Write(mat qrcode.Matrix) error {
	chr := "██"

	ww, _ := mat.Width(), mat.Height()
	wwWithBorder := ww + 4

	mat.Iterate(qrcode.IterDirection_COLUMN,
		func(x int, y int, state qrcode.QRValue) {
			if x == 0 && y == 0 {
				// add the top row border
				for range wwWithBorder {
					pterm.FgWhite.Print(chr)
				}
				fmt.Println()
				for range wwWithBorder {
					pterm.FgWhite.Print(chr)
				}
				fmt.Println()
			}

			if y == 0 {
				pterm.FgWhite.Print(chr)
				pterm.FgWhite.Print(chr)
			}

			if state.IsSet() {
				pterm.FgBlack.Print(chr)
			} else {
				pterm.FgWhite.Print(chr)
			}

			if y == ww-1 {
				pterm.FgWhite.Print(chr)
				pterm.FgWhite.Print(chr)
				fmt.Println()
			}
		})

	// add the bottom row border
	for range wwWithBorder {
		pterm.FgWhite.Print(chr)
	}
	fmt.Println()
	for range wwWithBorder {
		pterm.FgWhite.Print(chr)
	}
	fmt.Println()

	return nil
}

func (w terminalWriter) Close() error {
	return nil
}
