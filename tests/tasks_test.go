/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tasks_test.go">
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

// Example of how to work with tasks.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get tasks.
func Test_Tasks_GetTasks(t *testing.T) {
	localFileName := "Project2016.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTasksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetTasks(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Tasks)
	assert.Equal(t, 6, len(result.Tasks.TaskItem))
	var firstTask models.TaskItem
	for _, t := range result.Tasks.TaskItem {
		if t.Uid == 5 {
			firstTask = t
			break
		}
	}
	assert.NotNil(t, firstTask)
	assert.Equal(t, "Summary Task 1", firstTask.Name)
	assert.Equal(t, CreateTime(2015, 8, 3, 8, 0, 0), firstTask.Start)
	assert.Equal(t, CreateTime(2015, 8, 6, 17, 0, 0), firstTask.Finish)
	assert.Equal(t, "/5", firstTask.Link.Href)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get task.
func Test_Tasks_GetTask(t *testing.T) {
	localFileName := "Project2016.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTaskOpts{
		TaskUid: 5,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetTask(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Task)
	assert.Equal(t, int32(5), result.Task.Uid)
	assert.Equal(t, []int32{1, 2, 3, 4}, result.Task.SubtasksUids)
	assert.Equal(t, "Summary Task 1", result.Task.Name)
	assert.Equal(t, CreateTime(2015, 8, 3, 8, 0, 0), result.Task.Start)
	assert.Equal(t, CreateTime(2015, 8, 6, 17, 0, 0), result.Task.Finish)
	assert.Equal(t, "1.08:00:00", *result.Task.RemainingWork)
	assert.Equal(t, float64(1920), result.Task.WorkVariance)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete task.
func Test_Tasks_DeleteTask(t *testing.T) {
	localFileName := "Project2016.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	deleteOpts := &requests.DeleteTaskOpts{
		TaskUid: 4,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	deleteResult, _, err := client.TasksApi.DeleteTask(ctx, deleteOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	result, _, err := client.TasksApi.GetTasks(ctx, &requests.GetTasksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Tasks)
	assert.Equal(t, 5, len(result.Tasks.TaskItem))
	isDeletedTaskFound := false
	for _, t := range result.Tasks.TaskItem {
		if t.Uid == deleteOpts.TaskUid {
			isDeletedTaskFound = true
			break
		}
	}
	assert.False(t, isDeletedTaskFound)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post tasks.
func Test_Tasks_PostTasks(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	postTasksOpts := &requests.PostTasksOpts{
		Requests: []models.TaskCreationRequest{
			{
				TaskName: "SomeFirstTaskName",
			},
			{
				TaskName:      "SomeSecondTaskNameWithParent",
				ParentTaskUid: 2,
			},
		},
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	postResult, _, err := client.TasksApi.PostTasks(ctx, postTasksOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), postResult.Code)
	assert.NotNil(t, postResult.Tasks)
	assert.Equal(t, len(postTasksOpts.Requests), len(postResult.Tasks.TaskItem))

	var newSubtaskUid int32
	for _, t := range postResult.Tasks.TaskItem {
		if t.Name == postTasksOpts.Requests[1].TaskName {
			newSubtaskUid = t.Uid
			break
		}
	}
	getResult, _, err := client.TasksApi.GetTask(ctx, &requests.GetTaskOpts{
		TaskUid: postTasksOpts.Requests[1].ParentTaskUid,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Task)
	isTaskContainsNewSubtaskUid := false
	for _, uid := range getResult.Task.SubtasksUids {
		if uid == newSubtaskUid {
			isTaskContainsNewSubtaskUid = true
			break
		}
	}
	assert.True(t, isTaskContainsNewSubtaskUid)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post task.
func Test_Tasks_PostTask(t *testing.T) {
	localFileName := "Project2016.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	postTaskOpts := &requests.PostTaskOpts{
		Name:         remoteFileName,
		TaskName:     optional.NewString("New task name"),
		BeforeTaskId: optional.NewInt32(4),
		Folder:       optional.NewString(remoteBaseTestDataFolder),
	}
	postResult, _, err := client.TasksApi.PostTask(ctx, postTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), postResult.Code)
	assert.NotNil(t, postResult.TaskItem)

	getResult, _, err := client.TasksApi.GetTask(ctx, &requests.GetTaskOpts{
		TaskUid: postResult.TaskItem.Uid,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Task)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put task.
func Test_Tasks_PutTask(t *testing.T) {
	localFileName := "Project2016.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 4,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Task)

	taskToModify := getResult.Task
	taskToModify.Name = "Modified task name"
	taskToModify.IsManual = true
	taskToModify.ManualStart = CreateTime(2015, 10, 1, 9, 15, 0)
	taskToModify.ManualFinish = CreateTime(2015, 10, 1, 17, 15, 0)
	putTaskOpts := &requests.PutTaskOpts{
		TaskUid:     getTaskOpts.TaskUid,
		Recalculate: optional.NewBool(false),
		Mode:        optional.NewString(string(models.NONE_CalculationMode)),
		Task:        *taskToModify,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutTask(ctx, putTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getResult, _, err = client.TasksApi.GetTask(ctx, &requests.GetTaskOpts{
		TaskUid: putTaskOpts.TaskUid,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Task)
	assert.Equal(t, taskToModify.Name, getResult.Task.Name)
	assert.Equal(t, taskToModify.IsManual, getResult.Task.IsManual)
	assert.Equal(t, taskToModify.ManualStart, getResult.Task.ManualStart)
	assert.Equal(t, taskToModify.ManualFinish, getResult.Task.ManualFinish)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get task assignments.
func Test_Tasks_GetTaskAssignments(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTaskAssignmentsOpts{
		TaskUid: 1,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetTaskAssignments(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Assignments)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put move task.
func Test_Tasks_PutMoveTask(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getOpts := &requests.GetTaskOpts{
		TaskUid: 6,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetTask(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Task)
	for _, uid := range getResult.Task.SubtasksUids {
		assert.NotEqual(t, int32(10), uid)
	}

	putMoveOpts := &requests.PutMoveTaskOpts{
		TaskUid:       10,
		ParentTaskUid: getOpts.TaskUid,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	}
	putMoveResult, _, err := client.TasksApi.PutMoveTask(ctx, putMoveOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putMoveResult.Code)
	getResult, _, err = client.TasksApi.GetTask(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Task)
	isTaskContainsNewSubtaskUid := false
	for _, uid := range getResult.Task.SubtasksUids {
		if uid == putMoveOpts.TaskUid {
			isTaskContainsNewSubtaskUid = true
			break
		}
	}
	assert.True(t, isTaskContainsNewSubtaskUid)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put move task to sibling.
func Test_Tasks_PutMoveTaskToSibling(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	putMoveToSiblingOpts := &requests.PutMoveTaskToSiblingOpts{
		TaskUid:       40,
		BeforeTaskUid: 41,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	}
	putMoveToSiblingResult, _, err := client.TasksApi.PutMoveTaskToSibling(ctx, putMoveToSiblingOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putMoveToSiblingResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 38,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Task)
	assert.Equal(t, []int32{39, 40, 41}, getResult.Task.SubtasksUids)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put move task to sibling, should return response with code 404 if input uid is not found.
func Test_Tasks_PutMoveTaskToSibling_Return404IfUidIsNotFound(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	putMoveToSiblingOpts := &requests.PutMoveTaskToSiblingOpts{
		TaskUid:       99999,
		BeforeTaskUid: -1,
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	}
	_, _, err := client.TasksApi.PutMoveTaskToSibling(ctx, putMoveToSiblingOpts)
	if err != nil {
		assert.Equal(t, "Task with 99999 Uid doesn't exist", err.Error())
	} else {
		assert.Fail(t, "error must be produced")
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
