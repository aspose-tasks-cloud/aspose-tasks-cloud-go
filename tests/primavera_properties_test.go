/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="primavera_properties_test.go">
 *   Copyright (c) 2022 Aspose.Tasks Cloud
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

// Example of how to work with primavera properties.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get task primavera properties.
func Test_PrimaveraProperties_GetPrimaveraTaskProperties(t *testing.T) {
	localFileName := "p6_multiproject.xml"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetPrimaveraTaskPropertiesOpts{
		TaskUid: 1,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetPrimaveraTaskProperties(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.PrimaveraProperties)
	assert.Equal(t, int32(0), result.PrimaveraProperties.SequenceNumber)
	assert.Equal(t, "A1040", result.PrimaveraProperties.ActivityId)
	assert.Equal(t, CreateTime(2000, 10, 12, 8, 0, 0), result.PrimaveraProperties.RemainingEarlyStart)
	assert.Equal(t, CreateTime(2000, 10, 12, 17, 0, 0), result.PrimaveraProperties.RemainingEarlyFinish)
	assert.Equal(t, CreateTime(2000, 10, 12, 8, 0, 0), result.PrimaveraProperties.RemainingLateStart)
	assert.Equal(t, CreateTime(2000, 10, 12, 17, 0, 0), result.PrimaveraProperties.RemainingLateFinish)
	assert.Equal(t, "Fixed Units", result.PrimaveraProperties.RawDurationType)
	assert.Equal(t, "Task Dependent", result.PrimaveraProperties.RawActivityType)
	assert.Equal(t, "Units", result.PrimaveraProperties.RawCompletePercentType)
	assert.Equal(t, "Not Started", result.PrimaveraProperties.RawStatus)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
