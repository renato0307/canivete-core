/*
Copyright © 2021 Renato Torres <renato.torres@pm.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package programming

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/renato0307/canivete-core/interface/programming"
)

func (p *Service) DebugJwt(tokenString string) (programming.JwtDebuggerOutput, error) {
	output := programming.JwtDebuggerOutput{}

	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return output, fmt.Errorf("invalid token - it must contain 3 parts")
	}

	// handles header
	headerDecoded, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return output, fmt.Errorf("invalid token - header is not valid")
	}

	err = json.Unmarshal(headerDecoded, &output.Header)
	if err != nil {
		return output, fmt.Errorf("invalid token - header does not contain valid JSON")
	}

	// handles payload
	payloadDecoded, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return output, fmt.Errorf("invalid token - payload is not valid")
	}

	err = json.Unmarshal(payloadDecoded, &output.Payload)
	if err != nil {
		return output, fmt.Errorf("invalid token - payload does not contain valid JSON")
	}

	return output, nil
}
