/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="calendars_test.go">
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

// Example of how to work with calendars.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get calendars.
func Test_Calendars_GetCalendars(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetCalendarsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetCalendars(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Calendars.List)
	assert.Equal(t, 1, len(result.Calendars.List))
	assert.Equal(t, "Standard", result.Calendars.List[0].Name)
	assert.Equal(t, int32(1), result.Calendars.List[0].Uid)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get calendar.
func Test_Calendars_GetCalendar(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetCalendarOpts{
		CalendarUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetCalendar(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Calendar)
	assert.Equal(t, "3F979F74-B9D3-4E5F-98DC-5E08060A0C30", result.Calendar.Guid)
	assert.Equal(t, "Standard", result.Calendar.Name)
	assert.Equal(t, int32(1), result.Calendar.Uid)
	assert.True(t, result.Calendar.IsBaseCalendar)
	assert.False(t, result.Calendar.IsBaselineCalendar)
	assert.Equal(t, 7, len(result.Calendar.Days))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post calendar.
func Test_Calendars_PostCalendar(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	firstDayType := models.SUNDAY_DayType
	secondDayType := models.MONDAY_DayType
	newFirstWorkingTime := models.WorkingTime{
		FromTime: CreateTime(1, 1, 1, 8, 0, 0),
		ToTime:   CreateTime(1, 1, 1, 13, 0, 0),
	}
	newSecondWorkingTime := models.WorkingTime{
		FromTime: CreateTime(1, 1, 1, 14, 0, 0),
		ToTime:   CreateTime(1, 1, 1, 17, 0, 0),
	}
	postCalendarOpts := &requests.PostCalendarOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Calendar: models.Calendar{
			Name:               "My new calendar",
			IsBaseCalendar:     false,
			IsBaselineCalendar: false,
			Days: []models.WeekDay{
				{
					DayType:    &firstDayType,
					DayWorking: false,
				},
				{
					DayType:    &secondDayType,
					DayWorking: true,
					FromDate:   CreateTime(1, 1, 1, 8, 0, 0),
					ToDate:     CreateTime(1, 1, 1, 17, 0, 0),
					WorkingTimes: []models.WorkingTime{
						newFirstWorkingTime,
						newSecondWorkingTime,
					},
				},
			},
		},
	}
	postResult, _, err := client.TasksApi.PostCalendar(ctx, postCalendarOpts)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), postResult.Code)

	createdCalendarUid := postResult.CalendarItem.Uid
	getResult, _, err := client.TasksApi.GetCalendar(ctx, &requests.GetCalendarOpts{
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
		CalendarUid: createdCalendarUid,
	})
	assert.NotNil(t, getResult.Calendar)
	assert.Equal(t, postCalendarOpts.Calendar.Name, getResult.Calendar.Name)
	assert.Equal(t, 7, len(getResult.Calendar.Days))

	var monday models.WeekDay
	for _, day := range getResult.Calendar.Days {
		if *day.DayType == models.MONDAY_DayType {
			monday = day
			break
		}
	}
	assert.NotNil(t, monday)
	assert.Equal(t, 2, len(monday.WorkingTimes))
	assert.Equal(t, newFirstWorkingTime.FromTime, monday.WorkingTimes[0].FromTime)
	assert.Equal(t, newFirstWorkingTime.ToTime, monday.WorkingTimes[0].ToTime)
	assert.Equal(t, newSecondWorkingTime.FromTime, monday.WorkingTimes[1].FromTime)
	assert.Equal(t, newSecondWorkingTime.ToTime, monday.WorkingTimes[1].ToTime)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post calendar should return response with code 400 if workingTimes is empty for any working day.
func Test_Calendars_PostCalendarShouldReturn400IfWorkingTimesIsEmpty(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	firstDayType := models.SUNDAY_DayType
	secondDayType := models.MONDAY_DayType
	postCalendarOpts := &requests.PostCalendarOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Calendar: models.Calendar{
			Name:               "My new calendar",
			IsBaseCalendar:     false,
			IsBaselineCalendar: false,
			Days: []models.WeekDay{
				{
					DayType:    &firstDayType,
					DayWorking: false,
				},
				{
					DayType:    &secondDayType,
					DayWorking: true,
					FromDate:   CreateTime(1, 1, 1, 8, 0, 0),
					ToDate:     CreateTime(1, 1, 1, 17, 0, 0),
				},
			},
		},
	}
	_, response, err := client.TasksApi.PostCalendar(ctx, postCalendarOpts)

	if err == nil {
		t.Error("The error must be produced")
	} else {
		assert.Equal(t, "Calendar should have at least one working time defined", err.Error())
	}
	assert.Equal(t, 400, response.StatusCode)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put calendar.
func Test_Calendars_PutCalendar(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	firstDayType := models.SUNDAY_DayType
	secondDayType := models.MONDAY_DayType
	newFirstWorkingTime := models.WorkingTime{
		FromTime: CreateTime(1, 1, 1, 8, 0, 0),
		ToTime:   CreateTime(1, 1, 1, 13, 0, 0),
	}
	newSecondWorkingTime := models.WorkingTime{
		FromTime: CreateTime(1, 1, 1, 14, 0, 0),
		ToTime:   CreateTime(1, 1, 1, 17, 0, 0),
	}
	putCalendarOpts := &requests.PutCalendarOpts{
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
		CalendarUid: 1,
		Calendar: models.Calendar{
			Uid:                1,
			Name:               "Modified calendar",
			IsBaseCalendar:     false,
			IsBaselineCalendar: false,
			Days: []models.WeekDay{
				{
					DayType:    &firstDayType,
					DayWorking: false,
				},
				{
					DayType:    &secondDayType,
					DayWorking: true,
					FromDate:   CreateTime(1, 1, 1, 8, 0, 0),
					ToDate:     CreateTime(1, 1, 1, 17, 0, 0),
					WorkingTimes: []models.WorkingTime{
						newFirstWorkingTime,
						newSecondWorkingTime,
					},
				},
			},
		},
	}
	putResult, _, err := client.TasksApi.PutCalendar(ctx, putCalendarOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getResult, _, err := client.TasksApi.GetCalendar(ctx, &requests.GetCalendarOpts{
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
		CalendarUid: putCalendarOpts.CalendarUid,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Calendar)
	assert.Equal(t, putCalendarOpts.Calendar.Name, getResult.Calendar.Name)
	assert.Equal(t, 7, len(getResult.Calendar.Days))

	var monday models.WeekDay
	for _, day := range getResult.Calendar.Days {
		if *day.DayType == models.MONDAY_DayType {
			monday = day
			break
		}
	}
	assert.NotNil(t, monday)
	assert.Equal(t, 2, len(monday.WorkingTimes))
	assert.Equal(t, newFirstWorkingTime.FromTime, monday.WorkingTimes[0].FromTime)
	assert.Equal(t, newFirstWorkingTime.ToTime, monday.WorkingTimes[0].ToTime)
	assert.Equal(t, newSecondWorkingTime.FromTime, monday.WorkingTimes[1].FromTime)
	assert.Equal(t, newSecondWorkingTime.ToTime, monday.WorkingTimes[1].ToTime)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete calendar.
func Test_Calendars_DeleteCalendar(t *testing.T) {
	localFileName := "CalendarWorkWeeks.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	deleteCalendarOpts := &requests.DeleteCalendarOpts{
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
		CalendarUid: 3,
	}
	deleteResult, _, err := client.TasksApi.DeleteCalendar(ctx, deleteCalendarOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getResult, _, err := client.TasksApi.GetCalendars(ctx, &requests.GetCalendarsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Calendars.List)
	assert.Equal(t, 1, len(getResult.Calendars.List))
	assert.NotEqual(t, deleteCalendarOpts.CalendarUid, getResult.Calendars.List[0].Uid)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
