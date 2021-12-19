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
package internet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertMediumToMd(t *testing.T) {
	// arrange
	postId := "f744fbff033e"

	// act
	s := Service{}
	output, err := s.ConvertMediumToMd(postId)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, postId, output.PostId)
	assert.NotEmpty(t, output.Markdown)
}
