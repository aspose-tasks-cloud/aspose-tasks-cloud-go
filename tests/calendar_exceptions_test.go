/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="calendar_exceptions_test.go">
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

// Example of how to work with calendar exceptions.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get calendar exceptions.
func Test_CalendarExceptions_GetCalendarExceptions(t *testing.T) {
	localFileName := "Calenar_with_exception.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetCalendarExceptionsOpts{
		CalendarUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetCalendarExceptions(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.CalendarExceptions)
	assert.Equal(t, 1, len(result.CalendarExceptions))
	calendarException := result.CalendarExceptions[0]
	assert.True(t, calendarException.DayWorking)
	assert.Equal(t, []models.DayType{models.MONDAY_DayType}, calendarException.DaysOfWeek)
	assert.Equal(t, CreateTime(2018, 2, 13, 0, 0, 0), calendarException.FromDate)
	assert.Equal(t, CreateTime(2018, 4, 9, 23, 59, 0), calendarException.ToDate)
	assert.Equal(t, models.UNDEFINED_Month, *calendarException.Month)
	assert.Equal(t, models.UNDEFINED_MonthItemType, *calendarException.MonthItem)
	assert.Equal(t, models.UNDEFINED_MonthPosition, *calendarException.MonthPosition)
	assert.Equal(t, models.WEEKLY_CalendarExceptionType, *calendarException.Type_)
	assert.Equal(t, int32(8), calendarException.Occurrences)
	assert.Equal(t, int32(1), calendarException.Period)
	assert.Equal(t, 2, len(calendarException.WorkingTimes))
	assert.Equal(t, "09:00:00", ExtractTimeOnly(calendarException.WorkingTimes[0].FromTime))
	assert.Equal(t, "12:34:00", ExtractTimeOnly(calendarException.WorkingTimes[0].ToTime))
	assert.Equal(t, "15:11:00", ExtractTimeOnly(calendarException.WorkingTimes[1].FromTime))
	assert.Equal(t, "17:30:00", ExtractTimeOnly(calendarException.WorkingTimes[1].ToTime))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post calendar exception.
func Test_CalendarExceptions_PostCalendarException(t *testing.T) {
	localFileName := "New_project_2013.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newExceptionType := models.MONTHLY_BY_DAY_CalendarExceptionType
	newMonthItemType := models.UNDEFINED_MonthItemType
	newMonthPosition := models.UNDEFINED_MonthPosition
	newMonth := models.UNDEFINED_Month
	newException := models.CalendarException{
		Name:                 "Non-working day exception",
		DayWorking:           false,
		FromDate:             CreateTime(2014, 10, 27, 0, 0, 0),
		ToDate:               CreateTime(2015, 8, 5, 23, 59, 0),
		Occurrences:          10,
		Type_:                &newExceptionType,
		MonthItem:            &newMonthItemType,
		MonthPosition:        &newMonthPosition,
		Month:                &newMonth,
		EnteredByOccurrences: true,
		MonthDay:             5,
		Period:               1,
	}

	postResult, _, err := client.TasksApi.PostCalendarException(ctx, &requests.PostCalendarExceptionOpts{
		CalendarUid:       1,
		CalendarException: newException,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), postResult.Code)

	getResult, _, err := client.TasksApi.GetCalendarExceptions(ctx, &requests.GetCalendarExceptionsOpts{
		CalendarUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.CalendarExceptions)
	assert.Equal(t, 1, len(getResult.CalendarExceptions))
	assertCalendarExceptionsAreEqual(t, &newException, &getResult.CalendarExceptions[0])
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put calendar exception.
func Test_CalendarExceptions_PutCalendarException(t *testing.T) {
	localFileName := "Calenar_with_exception.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getResult, _, err := client.TasksApi.GetCalendarExceptions(ctx, &requests.GetCalendarExceptionsOpts{
		CalendarUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)

	exception := &getResult.CalendarExceptions[0]
	exception.WorkingTimes = []models.WorkingTime{
		{
			FromTime: CreateTime(1, 1, 1, 9, 0, 0),
			ToTime:   CreateTime(1, 1, 1, 17, 0, 0),
		},
	}
	exception.DaysOfWeek = []models.DayType{models.THURSDAY_DayType, models.FRIDAY_DayType}
	exception.Occurrences = 10
	exception.EnteredByOccurrences = true
	exception.Period = 1
	exception.Name = "Non-working day exception"
	exception.DayWorking = true
	exception.FromDate = CreateTime(2014, 10, 27, 0, 0, 0)
	exception.ToDate = CreateTime(2015, 8, 5, 23, 59, 0)

	putResult, _, err := client.TasksApi.PutCalendarException(ctx, &requests.PutCalendarExceptionOpts{
		CalendarUid: 1,
		Index: exception.Index,
		CalendarException: *exception,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getResult, _, err = client.TasksApi.GetCalendarExceptions(ctx, &requests.GetCalendarExceptionsOpts{
		CalendarUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)

	assert.NotNil(t, getResult.CalendarExceptions)
	assert.Equal(t, 1, len(getResult.CalendarExceptions))
	assertCalendarExceptionsAreEqual(t, exception, &getResult.CalendarExceptions[0])
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete calendar exception.
func Test_CalendarExceptions_DeleteCalendarException(t *testing.T) {
	localFileName := "Calenar_with_exception.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	deleteResult, _, err := client.TasksApi.DeleteCalendarException(ctx, &requests.DeleteCalendarExceptionOpts{
		CalendarUid: 1,
		Index: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getResult, _, err := client.TasksApi.GetCalendarExceptions(ctx, &requests.GetCalendarExceptionsOpts{
		CalendarUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)

	assert.NotNil(t, getResult.CalendarExceptions)
	assert.Equal(t, 0, len(getResult.CalendarExceptions))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func assertCalendarExceptionsAreEqual(t *testing.T, actual *models.CalendarException, expected *models.CalendarException) {
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.DayWorking, expected.DayWorking)
	assert.Equal(t, actual.EnteredByOccurrences, expected.EnteredByOccurrences)
	assert.Equal(t, actual.FromDate, expected.FromDate)
	assert.Equal(t, actual.ToDate, expected.ToDate)
	assert.Equal(t, actual.MonthDay, expected.MonthDay)
	assert.Equal(t, actual.Occurrences, expected.Occurrences)
	assert.Equal(t, actual.Period, expected.Period)
	assert.Equal(t, actual.Type_, expected.Type_)

	assert.Equal(t, len(actual.WorkingTimes), len(expected.WorkingTimes))
	for i := range actual.WorkingTimes {
		assert.Equal(t, actual.WorkingTimes[i].FromTime, expected.WorkingTimes[i].FromTime)
		assert.Equal(t, actual.WorkingTimes[i].ToTime, expected.WorkingTimes[i].ToTime)
	}

	assert.Equal(t, len(actual.DaysOfWeek), len(expected.DaysOfWeek))
	for i := range actual.DaysOfWeek {
		assert.Equal(t, actual.DaysOfWeek[i], expected.DaysOfWeek[i])
	}
}
