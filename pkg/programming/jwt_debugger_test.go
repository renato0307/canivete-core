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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugJwt(t *testing.T) {
	// arrange
	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTYzOTgyODY0NiwiZXhwIjoxNjM5ODMyMjQ2fQ.ujQ7wTsos4hYgipdnxSjLICDdfSLq9pYbpwS0WvUKc4"

	// act
	p := Service{}
	output, err := p.DebugJwt(tokenString)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, "JWT", output.Header["typ"])
	assert.Equal(t, "HS256", output.Header["alg"])
	assert.Equal(t, "John Doe", output.Payload["name"])
	assert.Equal(t, "1639828646", strconv.FormatFloat(output.Payload["iat"].(float64), 'f', 0, 64))
}

func TestDebugJwtEmptyToken(t *testing.T) {
	// arrange
	tokenString := ""

	// act
	p := Service{}
	output, err := p.DebugJwt(tokenString)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, output.Header)
	assert.Nil(t, output.Payload)
}

func TestDebugJwtInvalidHeader(t *testing.T) {
	// arrange
	tokenString := "THIS_IS_INVALID.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTYzOTgyODY0NiwiZXhwIjoxNjM5ODMyMjQ2fQ.ujQ7wTsos4hYgipdnxSjLICDdfSLq9pYbpwS0WvUKc4"

	// act
	p := Service{}
	output, err := p.DebugJwt(tokenString)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, output.Header)
	assert.Nil(t, output.Payload)
}

func TestDebugJwtHeaderNotJson(t *testing.T) {
	// arrange
	tokenString := "c3h4eGNhZGE.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTYzOTgyODY0NiwiZXhwIjoxNjM5ODMyMjQ2fQ.ujQ7wTsos4hYgipdnxSjLICDdfSLq9pYbpwS0WvUKc4"

	// act
	p := Service{}
	output, err := p.DebugJwt(tokenString)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, output.Header)
	assert.Nil(t, output.Payload)
}

func TestDebugJwtInvalidPayload(t *testing.T) {
	// arrange
	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.THIS_IS_INVALID.ujQ7wTsos4hYgipdnxSjLICDdfSLq9pYbpwS0WvUKc4"

	// act
	p := Service{}
	output, err := p.DebugJwt(tokenString)

	// assert
	assert.NotNil(t, err)
	assert.NotNil(t, output.Header)
	assert.Nil(t, output.Payload)
}

func TestDebugJwtPayloadNotJson(t *testing.T) {
	// arrange
	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.c3h4eGNhZGE.ujQ7wTsos4hYgipdnxSjLICDdfSLq9pYbpwS0WvUKc4"

	// act
	p := Service{}
	output, err := p.DebugJwt(tokenString)

	// assert
	assert.NotNil(t, err)
	assert.NotNil(t, output.Header)
	assert.Nil(t, output.Payload)
}
