/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="recurring_info_test.go">
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

// Example of how to work with recurring info.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get task recurring info.
func Test_RecurringInfo_GetTaskRecurringInfo(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	result, _, err := client.TasksApi.GetTaskRecurringInfo(ctx, &requests.GetTaskRecurringInfoOpts{
		TaskUid: 6,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.Equal(t, int32(2), result.RecurringInfo.Occurrences)
	assert.Equal(t, models.MONTHLY_RecurrencePattern, *result.RecurringInfo.RecurrencePattern)
	assert.True(t, result.RecurringInfo.UseEndDate)
	assert.False(t, result.RecurringInfo.MonthlyUseOrdinalDay)
	assert.Equal(t, int32(1), result.RecurringInfo.MonthlyDay)
	assert.Equal(t, models.NONE_WeekDayType, *result.RecurringInfo.WeeklyDays)
	assert.Equal(t, models.SECOND_OrdinalNumber, *result.RecurringInfo.YearlyOrdinalNumber)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put task recurring info.
func Test_RecurringInfo_PutTaskRecurringInfo(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getOpts := &requests.GetTaskRecurringInfoOpts{
		TaskUid: 6,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetTaskRecurringInfo(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.RecurringInfo)

	getResult.RecurringInfo.Occurrences = 10
	putResult, _, err := client.TasksApi.PutTaskRecurringInfo(ctx, &requests.PutTaskRecurringInfoOpts{
		TaskUid:       6,
		RecurringInfo: *getResult.RecurringInfo,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getResult, _, err = client.TasksApi.GetTaskRecurringInfo(ctx, getOpts)
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, int32(10), getResult.RecurringInfo.Occurrences)
	assert.Equal(t, models.MONTHLY_RecurrencePattern, *getResult.RecurringInfo.RecurrencePattern)
	assert.True(t, getResult.RecurringInfo.UseEndDate)
	assert.False(t, getResult.RecurringInfo.MonthlyUseOrdinalDay)
	assert.Equal(t, int32(1), getResult.RecurringInfo.MonthlyDay)
	assert.Equal(t, models.NONE_WeekDayType, *getResult.RecurringInfo.WeeklyDays)
	assert.Equal(t, models.SECOND_OrdinalNumber, *getResult.RecurringInfo.YearlyOrdinalNumber)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post task recurring info.
func Test_RecurringInfo_PostTaskRecurringInfo(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newRecurrencePattern := models.WEEKLY_RecurrencePattern
	newWeeklyDays := models.THURSDAY_WeekDayType
	newMonthlyOrdinalNumber := models.FIRST_OrdinalNumber
	newMonthlyOrdinalDay := models.MONDAY_DayOfWeek
	newYearlyOrdinalNumber := models.FIRST_OrdinalNumber
	newYearlyOrdinalDay := models.MONDAY_DayOfWeek
	newYearlyOrdinalMonth := models.DECEMBER_Month
	recurringInfo := models.RecurringInfo{
		RecurrencePattern:    &newRecurrencePattern,
		WeeklyDays:           &newWeeklyDays,
		MonthlyOrdinalNumber: &newMonthlyOrdinalNumber,
		MonthlyOrdinalDay:    &newMonthlyOrdinalDay,
		YearlyOrdinalNumber:  &newYearlyOrdinalNumber,
		YearlyOrdinalDay:     &newYearlyOrdinalDay,
		YearlyOrdinalMonth:   &newYearlyOrdinalMonth,
		Occurrences:          4,
		WeeklyRepetitions:    3,
		StartDate:            CreateTime(2018, 1, 1, 8, 0, 0),
		EndDate:              CreateTime(2018, 12, 31, 0, 0, 0),
		UseEndDate:           true,
	}
	postResult, _, err := client.TasksApi.PostTaskRecurringInfo(ctx, &requests.PostTaskRecurringInfoOpts{
		ParentTaskUid: 0,
		TaskName:      "New recurring task",
		CalendarName:  "Standard",
		RecurringInfo: recurringInfo,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), postResult.Code)
	assert.NotNil(t, postResult.TaskItem)

	getRequestOpts := &requests.GetTaskOpts{
		TaskUid: postResult.TaskItem.Uid,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetTask(ctx, getRequestOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, 18, len(getResult.Task.SubtasksUids))

	var lastAddedSubtaskUid int32
	for _, el := range getResult.Task.SubtasksUids {
		if el > lastAddedSubtaskUid {
			lastAddedSubtaskUid = el
		}
	}
	getRequestOpts.TaskUid = lastAddedSubtaskUid
	getResult, _, err = client.TasksApi.GetTask(ctx, getRequestOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, CreateTime(2018, 12, 27, 8, 0, 0), getResult.Task.Start)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
