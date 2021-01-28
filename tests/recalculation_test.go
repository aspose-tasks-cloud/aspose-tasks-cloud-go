/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="recalculation_test.go">
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

// Example of how to recalculate project.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for put recalculate project.
func Test_Recalculation_PutRecalculateProject(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	postTaskResult, _, err := client.TasksApi.PostTask(ctx, &requests.PostTaskOpts{
		TaskName: optional.NewString("NewTaskName"),
		Name:     remoteFileName,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), postTaskResult.Code)

	taskUid := postTaskResult.TaskItem.Uid
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, &requests.GetTaskOpts{
		TaskUid: taskUid,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, int32(200), getTaskResult.Task)

	task := getTaskResult.Task
	task.Name = "New task Name"
	task.ActualStart = CreateTime(2000, 10, 20, 0, 0, 0)
	task.ActualFinish = CreateTime(2000, 10, 9, 0, 0, 0)
	task.Cost = 100
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &requests.PutTaskOpts{
		Task:        *task,
		TaskUid:     taskUid,
		Name:        remoteFileName,
		Recalculate: optional.NewBool(false),
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	putRecalculateOpts := &requests.PutRecalculateProjectOpts{
		Name:     remoteFileName,
		Mode:     optional.NewString(string(models.NONE_CalculationMode)),
		Validate: optional.NewBool(true),
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	putRecalculateResult, _, err := client.TasksApi.PutRecalculateProject(ctx, putRecalculateOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putRecalculateResult.Code)
	assert.Equal(t, models.VALID_ProjectValidationState, *putRecalculateResult.Result.ValidationState)

	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put recalculate project resource fields.
func Test_Recalculation_PutRecalculateProjectResourceFields(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.PutRecalculateProjectResourceFieldsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.PutRecalculateProjectResourceFields(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put recalculate project uncomplete work to start after.
func Test_Recalculation_PutRecalculateProjectUncompleteWorkToStartAfter(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.PutRecalculateProjectUncompleteWorkToStartAfterOpts{
		After:  CreateTime(2010, 10, 10, 0, 0, 0),
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.PutRecalculateProjectUncompleteWorkToStartAfter(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put recalculate project work as complete.
func Test_Recalculation_PutRecalculateProjectWorkAsComplete(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.PutRecalculateProjectWorkAsCompleteOpts{
		CompleteThrough:  CreateTime(2010, 10, 10, 0, 0, 0),
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.PutRecalculateProjectWorkAsComplete(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
