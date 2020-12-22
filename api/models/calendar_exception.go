/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="calendar_exception.go">
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

import (
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/custom"
)

// Represent exceptional time periods in a calendar.
type CalendarException struct {
	// Index of the current item in the collection of calendar's exceptions.
	Index int32 `json:"index"`
	// Determines whether the range of recurrence is defined by entering a number of occurrences. False specifies that the range of recurrence is defined by entering a finish date.
	EnteredByOccurrences bool `json:"enteredByOccurrences"`
	// The beginning of the exception time.
	FromDate custom.TimeWithoutTZ `json:"fromDate"`
	// The end of the exception time.
	ToDate custom.TimeWithoutTZ `json:"toDate"`
	// The number of occurrences for which the calendar exception is valid.
	Occurrences int32 `json:"occurrences"`
	// The name of the exception.
	Name string `json:"name,omitempty"`
	// The exception type.
	Type_ *CalendarExceptionType `json:"type"`
	// The period of recurrence for the exception.
	Period int32 `json:"period"`
	// The days of the week on which the exception is valid.
	DaysOfWeek []DayType `json:"daysOfWeek,omitempty"`
	// The month item for which an exception recurrence is scheduled.
	MonthItem *MonthItemType `json:"monthItem"`
	// The position of a month item within a month.
	MonthPosition *MonthPosition `json:"monthPosition"`
	// The month for which an exception recurrence is scheduled.
	Month *Month `json:"month"`
	// The day of a month on which an exception recurrence is scheduled.
	MonthDay int32 `json:"monthDay"`
	// Determines whether the specified date or day type is working.
	DayWorking bool `json:"dayWorking"`
	// The collection of working times that defines the time worked on the weekday.  At least one working time must present, and there can't be more than five.
	WorkingTimes []WorkingTime `json:"workingTimes,omitempty"`
}
