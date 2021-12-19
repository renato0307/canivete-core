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
	"errors"
	"fmt"
	"math"

	"github.com/renato0307/canivete-core/interface/finance"
)

// CalculateCompoundInterests calculates compound interests.
//
// The formula for compound interests is a = p*((1+r/n)^(n * t))
// With different periodic payments an extra is needed:
// 	a_series = m * (y/n) {[(1 + r/n)^(n * t) - 1] / (r/n)}
// 	total = a + a_series
//
// Where:
// 	a = the future value of the investment/loan, including interest
// 	p = the principal investment amount (the initial deposit or loan amount)
// 	r = the annual interest rate (decimal)
// 	n = the number of times that interest is compounded per unit t
// 	t = the time the money is invested or borrowed for
// 	m = the regular contribution
// 	y = regular contributions in the compounded period
func (s *Service) CalculateCompoundInterests(p, n, t, m, y, rInt float64) (finance.CompoundInterestsOutput, error) {
	output := finance.CompoundInterestsOutput{
		Total:   finance.CompoundInterestsDetailOutput{},
		History: []finance.CompoundInterestsHistoryEntryOutput{},
	}

	if m > 0 && y <= 0 {
		return output, errors.New("y must be bigger than zero")
	}

	r := rInt / 100

	output.Total = calculateValues(p, n, t, m, y, r)

	// history
	for i := 1; i <= int(t); i++ {
		historyEntry := finance.CompoundInterestsHistoryEntryOutput{}
		historyEntry.Totals = calculateValues(p, n, float64(i), m, y, r)
		historyEntry.Period = fmt.Sprint(i)
		output.History = append(output.History, historyEntry)
	}

	return output, nil
}

func calculateValues(p, n, t, m, y, r float64) finance.CompoundInterestsDetailOutput {
	// base calculation
	a := p * math.Pow(1+r/n, n*t)

	// calculation for regular contributions
	aseries := 0.0
	if m > 0 {
		aseries = m * (y / n) * ((math.Pow(1+r/n, n*t) - 1) / (r / n))
	}

	// set output
	output := finance.CompoundInterestsDetailOutput{}
	output.FinalAmount = roundTwoDecimalPlaces(a + aseries)
	output.TotalContributions = roundTwoDecimalPlaces(p + (m * y * t))
	output.Interests = roundTwoDecimalPlaces(output.FinalAmount - output.TotalContributions)

	return output
}

func roundTwoDecimalPlaces(value float64) float64 {
	return math.Ceil(value*100) / 100
}
