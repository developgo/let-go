/*
 * Copyright (c) 2021 Marcin Gasperowicz <xnooga@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
 * documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
 * persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
 * WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 * OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package vm

type theIntType struct {
	zero Int
}

func (lt *theIntType) Name() string { return "Int" }

func (lt *theIntType) Box(bare interface{}) (Value, error) {
	raw, ok := bare.(int)
	if !ok {
		return IntType.zero, NewTypeError(bare, "can't be boxed as", lt)
	}
	return Int(raw), nil
}

// IntType is the type of IntValues
var IntType *theIntType

func init() {
	IntType = &theIntType{zero: 0}
}

// Int is boxed int
type Int int

// Type implements Value
func (l Int) Type() ValueType { return IntType }

// Unbox implements Unbox
func (l Int) Unbox() interface{} {
	return int(l)
}
