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

type CompoundInterestsDetailOutput struct {
	FinalAmount        float64
	TotalContributions float64
	Interests          float64
}

type CompoundInterestsHistoryEntryOutput struct {
	Period string
	Totals CompoundInterestsDetailOutput
}

type CompoundInterestsOutput struct {
	Total   CompoundInterestsDetailOutput
	History []CompoundInterestsHistoryEntryOutput
}

type Interface interface {
	CalculateCompoundInterests(p, n, t, m, y, rInt float64) (CompoundInterestsOutput, error)
}
