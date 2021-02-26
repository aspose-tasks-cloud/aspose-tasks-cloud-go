/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="resources_test.go">
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

// Example of how to work with resources.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get resources.
func Test_Resources_GetResources(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetResourcesOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetResources(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Resources)
	assert.Equal(t, 2, len(result.Resources.ResourceItem))
	assert.Equal(t, "Project manager", result.Resources.ResourceItem[1].Name)
	assert.Equal(t, int32(1), result.Resources.ResourceItem[1].Uid)
	assert.Equal(t, int32(1), result.Resources.ResourceItem[1].Id)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post resource.
func Test_Resources_PostResource(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getOpts := &requests.GetResourcesOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetResources(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Resources)

	startCount := len(getResult.Resources.ResourceItem)
	postOpts := &requests.PostResourceOpts{
		ResourceName: optional.NewString("new resource"),
		Name:         remoteFileName,
		Folder:       optional.NewString(remoteBaseTestDataFolder),
	}
	postResult, _, err := client.TasksApi.PostResource(ctx, postOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), postResult.Code)

	getResult, _, err = client.TasksApi.GetResources(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Resources)
	assert.Equal(t, startCount+1, len(getResult.Resources.ResourceItem))
	var addedResource models.ResourceItem
	for _, el := range getResult.Resources.ResourceItem {
		if el.Uid == postResult.ResourceItem.Uid {
			addedResource = el
		}
	}
	assert.Equal(t, postOpts.ResourceName.Value(), addedResource.Name)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put resource.
func Test_Resources_PutResource(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getOpts := &requests.GetResourceOpts{
		ResourceUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetResource(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Resource)

	newBaselineNumber := models.BASELINE1_BaselineType
	newBaseline := models.Baseline{
		BaselineNumber: &newBaselineNumber,
		Cost:           44,
	}
	getResult.Resource.Name = "Modified Resource1"
	getResult.Resource.Cost = 200
	getResult.Resource.Start = CreateTime(2000, 10, 10, 0, 0, 0)
	newWork := "4.04:10:00"
	getResult.Resource.Work = &newWork
	newOvertimeWork := "4.04:00:00"
	getResult.Resource.OvertimeWork = &newOvertimeWork
	newWorkgroup := models.NONE_WorkGroupType
	getResult.Resource.Workgroup = &newWorkgroup
	getResult.Resource.StandardRate = 101
	getResult.Resource.Finish = CreateTime(2000, 10, 10, 0, 0, 0)
	getResult.Resource.Baselines = []models.Baseline{newBaseline}
	putOpts := &requests.PutResourceOpts{
		ResourceUid: 1,
		Resource:    *getResult.Resource,
		Recalculate: optional.NewBool(false),
		Mode:        optional.NewString(string(models.NONE_CalculationMode)),
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutResource(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)
	assert.Equal(t, 1, len(putResult.Resource.Baselines))
	assert.Equal(t, *newBaseline.BaselineNumber, *putResult.Resource.Baselines[0].BaselineNumber)
	assert.Equal(t, newBaseline.Cost, putResult.Resource.Baselines[0].Cost)
	assert.Equal(t, getResult.Resource.StandardRate, putResult.Resource.StandardRate)
	assert.Equal(t, getResult.Resource.Start, putResult.Resource.Start)
	assert.Equal(t, getResult.Resource.Finish, putResult.Resource.Finish)
	assert.Equal(t, *getResult.Resource.Work, *putResult.Resource.Work)
	assert.Equal(t, *getResult.Resource.Workgroup, *putResult.Resource.Workgroup)
	assert.Equal(t, *getResult.Resource.OvertimeWork, *putResult.Resource.OvertimeWork)
	assert.Equal(t, getResult.Resource.Cost, putResult.Resource.Cost)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete resource.
func Test_Resources_DeleteResource(t *testing.T) {
	localFileName := "Plan_with_resource.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getOpts := &requests.GetResourcesOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	getResult, _, err := client.TasksApi.GetResources(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}

	resourceCountBeforeDelete := len(getResult.Resources.ResourceItem)

	deleteOpts := &requests.DeleteResourceOpts{
		ResourceUid: 1,
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	deleteResult, _, err := client.TasksApi.DeleteResource(ctx, deleteOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getResult, _, err = client.TasksApi.GetResources(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Resources)
	assert.Greater(t, resourceCountBeforeDelete, len(getResult.Resources.ResourceItem))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
