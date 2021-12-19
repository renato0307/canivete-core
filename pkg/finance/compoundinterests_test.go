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
package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCompoundInterest(t *testing.T) {
	// act
	p := Service{}
	output, err := p.CalculateCompoundInterests(1000, 1, 10, 0, 0, 5)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1628.9, output.Total.FinalAmount)
}

func TestCalculateCompoundInterestWithRegularContributions(t *testing.T) {
	// act
	p := Service{}
	output, err := p.CalculateCompoundInterests(5000, 12, 10, 100, 12, 5)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 23763.28, output.Total.FinalAmount)
}

func TestCompoundInterestsWithRegularContributionsInvalidValues(t *testing.T) {
	// act
	p := Service{}
	_, err := p.CalculateCompoundInterests(5000, 12, 10, 100, 0, 5)

	// assert
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "y must be bigger than zero")
}
