/*
 * C-style assertions for Go
 * Copyright © 2016 Gunnar Þór Magnússon
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package assert provides C-style assertions for Go.
package assert

// Asserter is an on/off switch for assertions.
// Recall that bool variables initialize to false.
type Asserter struct {
	on bool
}

// New returns a new Asserted with assertions turned on.
func New() Asserter {
	return Asserter{
		on: true,
	}
}

// On turns assertions on.
func (a *Asserter) On() {
	a.on = true
}

// Off turns assertions off.
func (a *Asserter) Off() {
	a.on = false
}

// Assert panics and prints message if condition is not fulfilled.
func (a Asserter) Assert(condition bool, message string) {
	if a.on {
		if !condition {
			panic(message)
		}
	}
}
