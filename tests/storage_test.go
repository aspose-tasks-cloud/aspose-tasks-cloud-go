/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="storage_test.go">
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

// Example of how to work with storage.
package api_test

import (
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get disc usage.
func Test_Storage_GetDiscUsage(t *testing.T) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	opts := &requests.GetDiscUsageOpts{}
	result, response, err := client.TasksApi.GetDiscUsage(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.Less(t, int64(0), result.TotalSize)
}

// Test for get file versions.
func Test_Storage_GetFileVersions(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetFileVersionsOpts{
		Path: remoteBaseTestDataFolder + "/" + remoteFileName,
	}
	result, response, err := client.TasksApi.GetFileVersions(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	isUploadedFileFound := false
	for _, fv := range result.Value {
		if fv.Name == remoteFileName {
			isUploadedFileFound = true
			break
		}
	}
	assert.True(t, isUploadedFileFound)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for object exists check.
func Test_Storage_ObjectExists(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.ObjectExistsOpts{
		Path: remoteBaseTestDataFolder + "/" + remoteFileName,
	}
	result, response, err := client.TasksApi.ObjectExists(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.True(t, result.Exists)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for storage exists check.
func Test_Storage_StorageExists(t *testing.T) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	opts := &requests.StorageExistsOpts{}

	result, response, err := client.TasksApi.StorageExists(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.True(t, result.Exists)
}
