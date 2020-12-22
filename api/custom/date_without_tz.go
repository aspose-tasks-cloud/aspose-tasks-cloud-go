/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="date_without_tz.go">
 *   Copyright (c) 2021 Aspose.Tasks Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------
 */

package custom

import (
	"fmt"
	"regexp"
	"time"
)

const formatStringTimeWithoutTZ = "2006-01-02T15:04:05"
var regexFinder = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)

// TimeWithoutTZ provides a yyyy-MM-ddThh:mm:ss date.
type TimeWithoutTZ time.Time

// MarshalJSON outputs JSON.
func (d TimeWithoutTZ) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(d).Format(formatStringTimeWithoutTZ) + "\""), nil
}

// UnmarshalJSON handles incoming JSON.
func (d *TimeWithoutTZ) UnmarshalJSON(b []byte) (err error) {
	str := regexFinder.FindString(string(b))
	if str == ""{
		err = newErrInvalidDateFormatTimeWithoutTZ(string(b))
		return
	}
	t, err := time.ParseInLocation(formatStringTimeWithoutTZ, str, time.UTC)
	if err != nil {
		return
	}
	*d = TimeWithoutTZ(t)
	return
}

// String // String returns the time formatted using the format string
//	"2006-01-02T15:04:05"
//
// If the time has a monotonic clock reading, the returned string
// includes a final field "m=±<value>", where value is the monotonic
// clock reading formatted as a decimal number of seconds.
//
// The returned string is meant for transformation as query param or debugging;
// for a stable serialized representation, use t.MarshalText, t.MarshalBinary,
// or t.Format with an explicit format string.
func (d TimeWithoutTZ) String() string {
	return time.Time(d).Format(formatStringTimeWithoutTZ)
}

// ErrInvalidDateFormatTimeWithoutTZ is returned when the field cannot be parsed.
type ErrInvalidDateFormatTimeWithoutTZ error

func newErrInvalidDateFormatTimeWithoutTZ(input string) ErrInvalidDateFormatTimeWithoutTZ {
	return ErrInvalidDateFormatTimeWithoutTZ(fmt.Errorf(" string '%s' did not contain any date in yyyy-MM-ddThh:mm:ss format", input))
}

