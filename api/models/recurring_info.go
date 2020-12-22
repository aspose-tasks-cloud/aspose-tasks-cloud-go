/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="recurring_info.go">
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

// Represents the details of a recurring task in a project.
type RecurringInfo struct {
	// Represents a recurrence pattern of the recurring task. Can be one of the values of  enum.
	RecurrencePattern *RecurrencePattern `json:"recurrencePattern"`
	// Specifies the date for the occurrences to begin.
	StartDate custom.TimeWithoutTZ `json:"startDate"`
	// Specifies the date for the occurrences to end.
	EndDate custom.TimeWithoutTZ `json:"endDate"`
	// Specifies the duration for one occurrence of the recurring task. the instance of  class.
	Duration *string `json:"duration,omitempty"`
	// Specifies a number of occurrences of the recurring task.
	Occurrences int32 `json:"occurrences"`
	// Determines whether to use the end date or a number of occurrences for the recurring task.
	UseEndDate bool `json:"useEndDate"`
	// Specifies an interval between repetitions for the daily recurrence pattern.
	DailyRepetitions int32 `json:"dailyRepetitions"`
	// Determines whether to use workdays for the daily recurrence pattern.
	DailyUseWorkdays bool `json:"dailyUseWorkdays"`
	// Specifies an interval between repetitions for the weekly recurrence pattern.
	WeeklyRepetitions int32 `json:"weeklyRepetitions"`
	// Specifies a collection of days used in the weekly recurrence pattern.
	WeeklyDays *WeekDayType `json:"weeklyDays"`
	// Determines whether to use ordinal day for the monthly recurrence pattern.
	MonthlyUseOrdinalDay bool `json:"monthlyUseOrdinalDay"`
	// Specifies an ordinal number of the monthly recurrence pattern. Can be one of the values of  enum.
	MonthlyOrdinalNumber *OrdinalNumber `json:"monthlyOrdinalNumber"`
	// Specifies a day of the monthly recurrence pattern when using ordinal day. Can be one of the values of  enum.
	MonthlyOrdinalDay *DayOfWeek `json:"monthlyOrdinalDay"`
	// Specifies a number of repetitions for the monthly recurrence pattern when using ordinal day.
	MonthlyOrdinalRepetitions int32 `json:"monthlyOrdinalRepetitions"`
	// Specifies a number of day of the monthly recurrence pattern.
	MonthlyDay int32 `json:"monthlyDay"`
	// Specifies a number of repetitions for the monthly recurrence pattern.
	MonthlyRepetitions int32 `json:"monthlyRepetitions"`
	// Determines whether to use ordinal day for the yearly recurrence pattern.
	YearlyUseOrdinalDay bool `json:"yearlyUseOrdinalDay"`
	// Specifies a date for the yearly recurrence pattern.
	YearlyDate custom.TimeWithoutTZ `json:"yearlyDate"`
	// Specifies an ordinal number of the yearly recurrence pattern. Can be one of the values of  enum.
	YearlyOrdinalNumber *OrdinalNumber `json:"yearlyOrdinalNumber"`
	// Specifies a weekday of the yearly recurrence pattern when using ordinal day. Can be one of the values of  enum.
	YearlyOrdinalDay *DayOfWeek `json:"yearlyOrdinalDay"`
	// Specifies a month of the yearly recurrence pattern when using ordinal day. Can be one of the values of  enum.
	YearlyOrdinalMonth *Month `json:"yearlyOrdinalMonth"`
}
