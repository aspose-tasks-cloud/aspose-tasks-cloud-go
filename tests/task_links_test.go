/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="task_links_test.go">
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

// Example of how to work with task links.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get task links.
func Test_TaskLinks_GetTaskLinks(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTaskLinksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetTaskLinks(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.TaskLinks)
	assert.Equal(t, 24, len(result.TaskLinks))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post task link.
func Test_TaskLinks_PostTaskLink(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newTaskLinkType := models.START_TO_START_TaskLinkType
	newTimeUnitType := models.DAY_TimeUnitType
	taskLink := models.TaskLink{
		PredecessorUid: 28,
		SuccessorUid:   30,
		Lag:            9600,
		LagFormat:      &newTimeUnitType,
		LinkType:       &newTaskLinkType,
	}
	postOpts := &requests.PostTaskLinkOpts{
		TaskLink: taskLink,
		Name:     remoteFileName,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	postResult, _, err := client.TasksApi.PostTaskLink(ctx, postOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), postResult.Code)

	getResult, _, err := client.TasksApi.GetTaskLinks(ctx, &requests.GetTaskLinksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, 25, len(getResult.TaskLinks))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put task link.
func Test_TaskLinks_PutTaskLink(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getResult, _, err := client.TasksApi.GetTaskLinks(ctx, &requests.GetTaskLinksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, 24, len(getResult.TaskLinks))
	taskLinkToEdit := &getResult.TaskLinks[0]
	newTaskLinkType := models.START_TO_FINISH_TaskLinkType
	newTimeUnitType := models.DAY_TimeUnitType
	// Modification of PredecessorUid and SuccessorUid fields is not supported.
	taskLinkToEdit.Lag = 9600
	taskLinkToEdit.LagFormat = &newTimeUnitType
	taskLinkToEdit.LinkType = &newTaskLinkType
	putOpts := &requests.PutTaskLinkOpts{
		Index: 1,
		TaskLink: *taskLinkToEdit,
		Name:     remoteFileName,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutTaskLink(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getResult, _, err = client.TasksApi.GetTaskLinks(ctx, &requests.GetTaskLinksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, 24, len(getResult.TaskLinks))
	assert.Equal(t, *taskLinkToEdit.LinkType, *getResult.TaskLinks[0].LinkType)
	assert.Equal(t, taskLinkToEdit.Lag, getResult.TaskLinks[0].Lag)
	assert.Equal(t, *taskLinkToEdit.LagFormat, *getResult.TaskLinks[0].LagFormat)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete task link.
func Test_TaskLinks_DeleteTaskLink(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	deleteOpts := &requests.DeleteTaskLinkOpts{
		Index: 1,
		Name:     remoteFileName,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	deleteResult, _, err := client.TasksApi.DeleteTaskLink(ctx, deleteOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getResult, _, err := client.TasksApi.GetTaskLinks(ctx, &requests.GetTaskLinksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, 23, len(getResult.TaskLinks))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
