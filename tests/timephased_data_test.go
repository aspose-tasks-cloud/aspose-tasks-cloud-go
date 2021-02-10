/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="timephased_data_test.go">
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

// Example of how to work with timephased data.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get task timephased data.
func Test_TimephasedData_GetTaskTimephasedData(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTaskTimephasedDataOpts{
		TaskUid: 27,
		Type_:   optional.NewString(string(models.TASK_WORK_TimephasedDataType)),
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetTaskTimephasedData(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Items)
	assert.LessOrEqual(t, 1, len(result.Items))
	for _, timephasedData := range result.Items {
		assert.Equal(t, models.TASK_WORK_TimephasedDataType, *timephasedData.TimephasedDataType)
		assert.Equal(t, int32(27), timephasedData.Uid)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get resource timephased data.
func Test_TimephasedData_GetResourceTimephasedData(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetResourceTimephasedDataOpts{
		ResourceUid: 1,
		Type_:       optional.NewString(string(models.RESOURCE_WORK_TimephasedDataType)),
		Name:        remoteFileName,
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetResourceTimephasedData(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Items)
	assert.LessOrEqual(t, 1, len(result.Items))
	for _, timephasedData := range result.Items {
		assert.Equal(t, models.RESOURCE_WORK_TimephasedDataType, *timephasedData.TimephasedDataType)
		assert.Equal(t, int32(1), timephasedData.Uid)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get assignment timephased data.
func Test_TimephasedData_GetAssignmentTimephasedData(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetAssignmentTimephasedDataOpts{
		AssignmentUid: 66,
		Type_:         optional.NewString(string(models.ASSIGNMENT_WORK_TimephasedDataType)),
		Name:          remoteFileName,
		Folder:        optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetAssignmentTimephasedData(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Items)
	assert.LessOrEqual(t, 1, len(result.Items))
	for _, timephasedData := range result.Items {
		assert.Equal(t, models.ASSIGNMENT_WORK_TimephasedDataType, *timephasedData.TimephasedDataType)
		assert.Equal(t, int32(66), timephasedData.Uid)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
