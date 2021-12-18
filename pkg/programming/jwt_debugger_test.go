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
	assert.Equal(t, output.Header["typ"], "JWT")
	assert.Equal(t, output.Header["alg"], "HS256")
	assert.Equal(t, output.Payload["name"], "John Doe")
	assert.Equal(t, output.Payload["iat"], "1639828646")
}
