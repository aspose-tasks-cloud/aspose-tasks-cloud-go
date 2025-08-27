/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="resource_level_test.go">
 *   Copyright (c) 2025 Aspose.Tasks Cloud
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

// Test for level tasks without body.
func Test_ResourceLevel_LevelTasks_WithoutBody(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.LevelTasksOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.LevelTasks(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.AffectedTaskUids)
	assert.Equal(t, 2, len(result.AffectedTaskUids))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for level tasks with body.
func Test_ResourceLevel_LevelTasks_WithBody(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	levelingOrder := models.STANDARD_LevelingOrder
	options := models.LevelingOptions{
		LevelingOrder: &levelingOrder,
		ResourceUids:  []int32{2, 999},
	}
	opts := &requests.LevelTasksOpts{
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
		Options: optional.NewInterface(options),
	}
	result, _, err := client.TasksApi.LevelTasks(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.AffectedTaskUids)
	assert.Equal(t, 1, len(result.AffectedTaskUids))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for clear leveling without body.
func Test_ResourceLevel_ClearLeveling_WithoutBody(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.ClearLevelingOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.ClearLeveling(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for clear leveling with body.
func Test_ResourceLevel_ClearLeveling_WithBody(t *testing.T) {
	localFileName := "sample.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.ClearLevelingOpts{
		Name:     remoteFileName,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
		TaskUids: optional.NewInterface([]int32{24, 9999}),
	}
	result, _, err := client.TasksApi.ClearLeveling(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
