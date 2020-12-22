/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="calendar.go">
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

package models

// Represents a calendar used in a project.
type Calendar struct {
	// The name of the calendar.
	Name string `json:"name,omitempty"`
	// The unique identifier of the calendar.
	Uid int32 `json:"uid"`
	// The collection of weekdays that defines the calendar.
	Days []WeekDay `json:"days,omitempty"`
	// Determines whether the calendar is a base calendar.
	IsBaseCalendar bool `json:"isBaseCalendar"`
	// The base calendar on which this calendar depends. Only applicable if the calendar is not a base calendar.
	BaseCalendar *Calendar `json:"baseCalendar,omitempty"`
	// Specifies whether the calendar is a baseline calendar.
	IsBaselineCalendar bool `json:"isBaselineCalendar"`
}
