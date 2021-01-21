/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="get_page_count_test.go">
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

// Example of how to work with getting page count.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get page count.
func Test_GetPageCount(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetPageCountOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PresentationFormat: optional.NewString(string(models.TASK_USAGE_PresentationFormat)),
		Timescale: optional.NewString(string(models.MONTHS_Timescale)),
	}
	result, _, err := client.TasksApi.GetPageCount(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.Equal(t, int32(4), result.PageCount)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get page count with specified presentation format and date interval.
func Test_GetPageCountWithSpecifiedPresentationFormatAndDateInterval(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newStartDate := CreateTime(2004, 1, 1,0,0,0)
	newEndDate := CreateTime(2004, 1, 28,0,0,0)
	options := &requests.GetPageCountOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PresentationFormat: optional.NewString(string(models.TASK_USAGE_PresentationFormat)),
		Timescale: optional.NewString(string(models.MONTHS_Timescale)),
		StartDate: &newStartDate,
		EndDate: &newEndDate,
	}
	result, _, err := client.TasksApi.GetPageCount(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.Equal(t, int32(4), result.PageCount)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}