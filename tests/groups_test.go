/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="groups_test.go">
 *   Copyright (c) 2026 Aspose.Tasks Cloud
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

// Example of how to work with groups.
package api_test

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
)

// Test for get resource groups.
func Test_Groups_GetResourceGroups(t *testing.T) {
	localFileName := "ProjectWithGroups.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetResourceGroupsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetResourceGroups(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Groups)
	assert.NotNil(t, result.Groups.List)
	assert.True(t, len(result.Groups.List) > 0)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get resource group.
func Test_Groups_GetResourceGroup(t *testing.T) {
	localFileName := "ProjectWithGroups.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	listOpts := &requests.GetResourceGroupsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	listResult, _, err := client.TasksApi.GetResourceGroups(ctx, listOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), listResult.Code)
	assert.True(t, len(listResult.Groups.List) > 0)

	groupUid := listResult.Groups.List[0].Uid

	opts := &requests.GetResourceGroupOpts{
		Name:     remoteFileName,
		GroupUid: groupUid,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetResourceGroup(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Group)
	assert.Equal(t, groupUid, result.Group.Uid)
	assert.NotEmpty(t, result.Group.Name)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get grouped resources.
func Test_Groups_GetGroupedResources(t *testing.T) {
	localFileName := "ProjectWithGroups.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	listOpts := &requests.GetResourceGroupsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	listResult, _, err := client.TasksApi.GetResourceGroups(ctx, listOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), listResult.Code)
	assert.True(t, len(listResult.Groups.List) > 1)

	groupUid := listResult.Groups.List[1].Uid

	opts := &requests.GetGroupedResourcesOpts{
		Name:     remoteFileName,
		GroupUid: groupUid,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetGroupedResources(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Buckets)
	assert.True(t, len(result.Buckets) > 0)

	for _, bucket := range result.Buckets {
		assert.NotEmpty(t, bucket.Key)
		assert.NotNil(t, bucket.Resources)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get task groups.
func Test_Groups_GetTaskGroups(t *testing.T) {
	localFileName := "ProjectWithGroups.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTaskGroupsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetTaskGroups(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Groups)
	assert.NotNil(t, result.Groups.List)
	assert.True(t, len(result.Groups.List) > 1)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get task group.
func Test_Groups_GetTaskGroup(t *testing.T) {
	localFileName := "ProjectWithGroups.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	listOpts := &requests.GetTaskGroupsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	listResult, _, err := client.TasksApi.GetTaskGroups(ctx, listOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), listResult.Code)
	assert.True(t, len(listResult.Groups.List) > 1)

	groupUid := listResult.Groups.List[0].Uid

	opts := &requests.GetTaskGroupOpts{
		Name:     remoteFileName,
		GroupUid: groupUid,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetTaskGroup(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Group)
	assert.Equal(t, groupUid, result.Group.Uid)
	assert.Equal(t, "Percent complete", result.Group.Name)
	assert.NotNil(t, result.Group.GroupCriteria)
	assert.True(t, len(result.Group.GroupCriteria) > 0)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get grouped tasks.
func Test_Groups_GetGroupedTasks(t *testing.T) {
	localFileName := "ProjectWithGroups.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	listOpts := &requests.GetTaskGroupsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	listResult, _, err := client.TasksApi.GetTaskGroups(ctx, listOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), listResult.Code)
	assert.True(t, len(listResult.Groups.List) > 1)

	groupUid := listResult.Groups.List[0].Uid

	opts := &requests.GetGroupedTasksOpts{
		Name:     remoteFileName,
		GroupUid: groupUid,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetGroupedTasks(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Buckets)
	assert.True(t, len(result.Buckets) > 0)

	for _, bucket := range result.Buckets {
		assert.NotEmpty(t, bucket.Key)
		assert.NotNil(t, bucket.Tasks)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
