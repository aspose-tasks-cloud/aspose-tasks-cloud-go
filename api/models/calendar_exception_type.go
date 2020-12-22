/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="calendar_exception_type.go">
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
// CalendarExceptionType : Specifies the calendar exception type.
type CalendarExceptionType string

// List of CalendarExceptionType
const (
	DAILY_CalendarExceptionType CalendarExceptionType = "Daily"
	YEARLY_BY_DAY_CalendarExceptionType CalendarExceptionType = "YearlyByDay"
	YEARLY_BY_POSITION_CalendarExceptionType CalendarExceptionType = "YearlyByPosition"
	MONTHLY_BY_DAY_CalendarExceptionType CalendarExceptionType = "MonthlyByDay"
	MONTHLY_BY_POSITION_CalendarExceptionType CalendarExceptionType = "MonthlyByPosition"
	WEEKLY_CalendarExceptionType CalendarExceptionType = "Weekly"
	BY_DAY_COUNT_CalendarExceptionType CalendarExceptionType = "ByDayCount"
	BY_WEEK_DAY_COUNT_CalendarExceptionType CalendarExceptionType = "ByWeekDayCount"
	NO_EXCEPTION_TYPE_CalendarExceptionType CalendarExceptionType = "NoExceptionType"
)
