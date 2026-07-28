/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="group_on.go">
 *   Copyright (c) 2025 Aspose.Tasks Cloud
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

// GroupOn : Specifies the type of grouping for a field used as a criterion in a group definition.
type GroupOn string

// List of GroupOn
const (
	EACH_VALUE_GroupOn          GroupOn = "EachValue"
	INTERVAL_GroupOn            GroupOn = "Interval"
	DATE_EACH_VALUE_GroupOn     GroupOn = "DateEachValue"
	DATE_MINUTE_GroupOn         GroupOn = "DateMinute"
	DATE_HOUR_GroupOn           GroupOn = "DateHour"
	DATE_DAY_GroupOn            GroupOn = "DateDay"
	DATE_WEEK_GroupOn           GroupOn = "DateWeek"
	DATE_THIRD_OF_MONTH_GroupOn GroupOn = "DateThirdOfMonth"
	DATE_MONTH_GroupOn          GroupOn = "DateMonth"
	DATE_QTR_GroupOn            GroupOn = "DateQtr"
	DATE_YEAR_GroupOn           GroupOn = "DateYear"
	DURATION_EACH_VALUE_GroupOn GroupOn = "DurationEachValue"
	DURATION_MINUTES_GroupOn    GroupOn = "DurationMinutes"
	DURATION_HOURS_GroupOn      GroupOn = "DurationHours"
	DURATION_DAYS_GroupOn       GroupOn = "DurationDays"
	DURATION_WEEKS_GroupOn      GroupOn = "DurationWeeks"
	DURATION_MONTHS_GroupOn     GroupOn = "DurationMonths"
	OUTLINE_EACH_VALUE_GroupOn  GroupOn = "OutlineEachValue"
	OUTLINE_LEVEL_GroupOn       GroupOn = "OutlineLevel"
	PCT_EACH_VALUE_GroupOn      GroupOn = "PctEachValue"
	PCT_INTERVAL_GroupOn        GroupOn = "PctInterval"
	PCT199_GroupOn              GroupOn = "Pct199"
	PCT150_GroupOn              GroupOn = "Pct150"
	PCT125_GroupOn              GroupOn = "Pct125"
	PCT110_GroupOn              GroupOn = "Pct110"
	TEXT_EACH_VALUE_GroupOn     GroupOn = "TextEachValue"
	TEXT_PREFIX_GroupOn         GroupOn = "TextPrefix"
)
