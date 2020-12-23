/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="calendar_workWeeks_test.go">
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

// Example of how to work with calendar work weeks.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get calendar workWeeks.
func Test_CalendarWorkWeeks_GetCalendarWorkWeeks(t *testing.T) {
	localFileName := "CalendarWorkWeeks.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getCalendarsResult, _, err := client.TasksApi.GetCalendars(ctx, &requests.GetCalendarsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}

	var calendarUid int32
	for _, calendar := range getCalendarsResult.Calendars.List {
		if calendar.Name == "Test work weeks" {
			calendarUid = calendar.Uid
			break
		}
	}

	getCalendarWorkWeeksResult, _, err := client.TasksApi.GetCalendarWorkWeeks(ctx, &requests.GetCalendarWorkWeeksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		CalendarUid: calendarUid,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getCalendarWorkWeeksResult.Code)
	assert.Equal(t, 1, len(getCalendarWorkWeeksResult.CalendarWorkWeeks))
	workWeek := getCalendarWorkWeeksResult.CalendarWorkWeeks[0]
	assert.Equal(t, "Work week 1", workWeek.Name)
	assert.Equal(t, "2018-01-01", ExtractDateOnly(workWeek.FromDate))
	assert.Equal(t, "2018-01-07", ExtractDateOnly(workWeek.ToDate))
	assert.Equal(t, 4, len(workWeek.WeekDays))
	assert.True(t, workWeek.WeekDays[0].DayWorking)
	assert.Equal(t, models.MONDAY_DayType, *workWeek.WeekDays[0].DayType)
	assert.Equal(t, 1, len(workWeek.WeekDays[0].WorkingTimes))
	assert.Equal(t, "11:30:00", ExtractTimeOnly(workWeek.WeekDays[0].WorkingTimes[0].FromTime))
	assert.Equal(t, "12:30:00", ExtractTimeOnly(workWeek.WeekDays[0].WorkingTimes[0].ToTime))

	assert.False(t, workWeek.WeekDays[1].DayWorking)
	assert.Equal(t, models.TUESDAY_DayType, *workWeek.WeekDays[1].DayType)
	assert.Equal(t, 0, len(workWeek.WeekDays[1].WorkingTimes))

	assert.True(t, workWeek.WeekDays[2].DayWorking)
	assert.Equal(t, models.WEDNESDAY_DayType, *workWeek.WeekDays[2].DayType)
	assert.Equal(t, 2, len(workWeek.WeekDays[2].WorkingTimes))
	assert.Equal(t, "09:30:00", ExtractTimeOnly(workWeek.WeekDays[2].WorkingTimes[0].FromTime))
	assert.Equal(t, "13:23:00", ExtractTimeOnly(workWeek.WeekDays[2].WorkingTimes[0].ToTime))
	assert.Equal(t, "14:45:00", ExtractTimeOnly(workWeek.WeekDays[2].WorkingTimes[1].FromTime))
	assert.Equal(t, "18:45:00", ExtractTimeOnly(workWeek.WeekDays[2].WorkingTimes[1].ToTime))

	assert.True(t, workWeek.WeekDays[0].DayWorking)
	assert.Equal(t, models.SATURDAY_DayType, *workWeek.WeekDays[3].DayType)
	assert.Equal(t, 1, len(workWeek.WeekDays[3].WorkingTimes))
	assert.Equal(t, "09:00:00", ExtractTimeOnly(workWeek.WeekDays[3].WorkingTimes[0].FromTime))
	assert.Equal(t, "10:00:00", ExtractTimeOnly(workWeek.WeekDays[3].WorkingTimes[0].ToTime))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}