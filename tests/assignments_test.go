/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="assignments_test.go">
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

// Example of how to work with assignments.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get assignments.
func Test_Assignments_GetAssignments(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetAssignmentsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetAssignments(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Assignments)
	assert.Equal(t, 6, len(result.Assignments.AssignmentItem))
	assert.Equal(t, int32(34), result.Assignments.AssignmentItem[0].TaskUid)
	assert.Equal(t, int32(63), result.Assignments.AssignmentItem[0].Uid)
	assert.Equal(t, "/63", result.Assignments.AssignmentItem[0].Link.Href)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get assignment.
func Test_Assignments_GetAssignment(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetAssignmentOpts{
		AssignmentUid: 63,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetAssignment(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Assignment)
	assert.Equal(t, "08:00:00", *result.Assignment.RegularWork)
	assert.Equal(t, "08:00:00", *result.Assignment.RemainingWork)
	assert.Equal(t, CreateTime(2012, 7, 9, 8, 0, 0), result.Assignment.Start)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get resource assignments.
func Test_Assignments_GetResourceAssignments(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetResourceAssignmentsOpts{
		ResourceUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetResourceAssignments(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Assignments)
	assert.Equal(t, 6, len(result.Assignments.List))
	for _, resourceAssignment := range result.Assignments.List {
		assert.Equal(t, int32(1), resourceAssignment.ResourceUid)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post assignment.
func Test_Assignments_PostAssignment(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	postAssignmentOpts := &requests.PostAssignmentOpts{
		ResourceUid: 1,
		Units:       optional.NewFloat64(0.5),
		TaskUid:     0,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.PostAssignment(ctx, postAssignmentOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.AssignmentItem)

	assignmentUid := result.AssignmentItem.Uid
	assignmentResult, _, err := client.TasksApi.GetAssignment(ctx, &requests.GetAssignmentOpts{
		AssignmentUid: assignmentUid,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), assignmentResult.Code)
	assert.NotNil(t, assignmentResult.Assignment)
	assert.Equal(t, int32(0), assignmentResult.Assignment.TaskUid)
	assert.Equal(t, int32(1), assignmentResult.Assignment.ResourceUid)

	taskResult, _, err := client.TasksApi.GetTask(ctx, &requests.GetTaskOpts{
		TaskUid: postAssignmentOpts.TaskUid,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), taskResult.Code)
	assert.Equal(t, assignmentResult.Assignment.Start, taskResult.Task.Start)
	assert.Equal(t, assignmentResult.Assignment.Finish, taskResult.Task.Finish)
	assert.Equal(t, assignmentResult.Assignment.Work, taskResult.Task.Work)
	assert.Equal(t, assignmentResult.Assignment.Cost, taskResult.Task.Cost)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post assignment with assignment cost usage.
func Test_Assignments_PostAssignment_UsingCostInsteadOfUnits(t *testing.T) {
	localFileName := "Cost_Res.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	postAssignmentOpts := &requests.PostAssignmentOpts{
		ResourceUid: 1,
		Cost:        optional.NewFloat32(2),
		TaskUid:     0,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.PostAssignment(ctx, postAssignmentOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.AssignmentItem)

	assignmentUid := result.AssignmentItem.Uid
	assignmentResult, _, err := client.TasksApi.GetAssignment(ctx, &requests.GetAssignmentOpts{
		AssignmentUid: assignmentUid,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), assignmentResult.Code)
	assert.NotNil(t, assignmentResult.Assignment)
	assert.Equal(t, postAssignmentOpts.Cost.Value(), assignmentResult.Assignment.Cost)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put assignment.
func Test_Assignments_PutAssignment(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getResult, _, err := client.TasksApi.GetAssignment(ctx, &requests.GetAssignmentOpts{
		AssignmentUid: 63,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Assignment)

	getResult.Assignment.TaskUid = 0

	putResponse, _, err := client.TasksApi.PutAssignment(ctx, &requests.PutAssignmentOpts{
		Mode:          optional.NewString(string(models.AUTOMATIC_CalculationMode)),
		Recalculate:   optional.NewBool(true),
		AssignmentUid: 63,
		Assignment:    *getResult.Assignment,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	assert.Equal(t, int32(200), putResponse.Code)

	assignmentsResult, _, err := client.TasksApi.GetAssignments(ctx, &requests.GetAssignmentsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), assignmentsResult.Code)
	assert.NotNil(t, assignmentsResult.Assignments)
	for _, assignment := range assignmentsResult.Assignments.AssignmentItem {
		if assignment.TaskUid == 34 {
			assert.NotEqual(t, int32(1), assignment.ResourceUid)
		}
		if assignment.TaskUid == 0 {
			assert.Equal(t, int32(1), assignment.ResourceUid)
			assert.Equal(t, int32(63), assignment.Uid)
		}
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put assignment with changing TimephasedData and Baselines.
func Test_Assignments_PutAssignmentChangingTimephasedDataAndBaselines(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getResult, _, err := client.TasksApi.GetAssignment(ctx, &requests.GetAssignmentOpts{
		AssignmentUid: 63,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Assignment)

	getResult.Assignment.Cost = 100
	getResult.Assignment.Start = CreateTime(2001, 9, 10, 0, 0, 0)
	getResult.Assignment.Finish = CreateTime(2002, 9, 10, 0, 0, 0)
	newBaselineType := models.BASELINE1_BaselineType
	newStart := CreateTime(2002, 9, 10, 0, 0, 0)
	getResult.Assignment.Baselines = []models.AssignmentBaseline{
		{
			BaselineNumber: &newBaselineType,
			Start:          &newStart,
		},
	}
	allocatedActualWork := "10:10:10"
	getResult.Assignment.ActualWork = &allocatedActualWork
	getResult.Assignment.ActualCost = 100
	getResult.Assignment.ActualStart = CreateTime(2001, 9, 10, 0, 0, 0)
	getResult.Assignment.ActualFinish = CreateTime(2002, 9, 10, 0, 0, 0)
	allocatedOvertimeActualWork := "100:10:10"
	getResult.Assignment.ActualOvertimeWork = &allocatedOvertimeActualWork
	allocatedWork := "80:10:10"
	getResult.Assignment.Work = &allocatedWork
	getResult.Assignment.Uid = 63
	getResult.Assignment.Vac = 10
	newWorkContourType := models.CONTOURED_WorkContourType
	getResult.Assignment.WorkContour = &newWorkContourType
	newTimeUnitType := models.HOUR_TimeUnitType
	newTimephasedDataType := models.ASSIGNMENT_REMAINING_WORK_TimephasedDataType
	getResult.Assignment.TimephasedData = []models.TimephasedData{
		{
			Uid:                getResult.Assignment.Uid,
			Start:              CreateTime(2001, 9, 10, 9, 0, 0),
			Finish:             CreateTime(2001, 9, 10, 14, 0, 0),
			Unit:               &newTimeUnitType,
			Value:              "4:0:0",
			TimephasedDataType: &newTimephasedDataType,
		},
	}

	putResponse, _, err := client.TasksApi.PutAssignment(ctx, &requests.PutAssignmentOpts{
		Mode:          optional.NewString(string(models.NONE_CalculationMode)),
		Recalculate:   optional.NewBool(false),
		AssignmentUid: getResult.Assignment.Uid,
		Assignment:    *getResult.Assignment,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResponse.Code)
	assert.NotNil(t, putResponse.Assignment)
	assert.Equal(t, getResult.Assignment.Uid, putResponse.Assignment.Uid)
	assert.Equal(t, getResult.Assignment.Vac, putResponse.Assignment.Vac)
	assert.NotEqual(t, getResult.Assignment.Cost, putResponse.Assignment.Cost, "Calculated fields must be overwritten")
	assert.Equal(t, getResult.Assignment.Start, putResponse.Assignment.Start)
	assert.Equal(t, getResult.Assignment.Finish, putResponse.Assignment.Finish)
	assert.Equal(t, "80.10:10:00", *putResponse.Assignment.Work)
	assert.Equal(t, *getResult.Assignment.ActualWork, *putResponse.Assignment.ActualWork)
	assert.Equal(t, getResult.Assignment.ActualStart, putResponse.Assignment.ActualStart)
	assert.Equal(t, getResult.Assignment.ActualFinish, putResponse.Assignment.ActualFinish)
	assert.Equal(t, "100.10:10:00", *putResponse.Assignment.ActualOvertimeWork)
	assert.Equal(t, 1, len(putResponse.Assignment.Baselines))
	assert.Equal(t, models.BASELINE1_BaselineType, *putResponse.Assignment.Baselines[0].BaselineNumber)
	assert.Equal(t, getResult.Assignment.Baselines[0].Start, putResponse.Assignment.Baselines[0].Start)
	assert.Equal(t, 1, len(putResponse.Assignment.TimephasedData))
	assert.Equal(t, getResult.Assignment.TimephasedData[0].Uid, putResponse.Assignment.TimephasedData[0].Uid)
	assert.Equal(t, "PT4H0M0S", putResponse.Assignment.TimephasedData[0].Value)
	assert.Equal(t, getResult.Assignment.TimephasedData[0].Start, putResponse.Assignment.TimephasedData[0].Start)
	assert.Equal(t, getResult.Assignment.TimephasedData[0].Finish, putResponse.Assignment.TimephasedData[0].Finish)
	assert.Equal(t, getResult.Assignment.TimephasedData[0].TimephasedDataType, putResponse.Assignment.TimephasedData[0].TimephasedDataType)

	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete assignment.
func Test_Assignments_DeleteAssignment(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	deleteResult, _, err := client.TasksApi.DeleteAssignment(ctx, &requests.DeleteAssignmentOpts{
		AssignmentUid: 63,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getResponse, _, err := client.TasksApi.GetAssignments(ctx, &requests.GetAssignmentsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	assert.Equal(t, int32(200), getResponse.Code)

	assignmentsResult, _, err := client.TasksApi.GetAssignments(ctx, &requests.GetAssignmentsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), assignmentsResult.Code)
	assert.NotNil(t, assignmentsResult.Assignments)
	for _, assignment := range assignmentsResult.Assignments.AssignmentItem {
		assert.False(t, assignment.Uid == 63)
		assert.False(t, assignment.TaskUid == 34 && assignment.ResourceUid == 1)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
