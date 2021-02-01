/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="report_test.go">
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

// Example of how to work with reports.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get report pdf.
func Test_Report_GetReportPdf(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetReportPdfOpts{
		Type_:  string(models.MILESTONES_ReportType),
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	file, response, err := client.TasksApi.GetReportPdf(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.Less(t, 1, len(file))
	fileAsStrings := ConvertBytesToStrings(file)
	assert.Equal(t, "%PDF-1.5", fileAsStrings[0])
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get critical path.
func Test_Report_GetCriticalPath(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetCriticalPathOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetCriticalPath(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.Equal(t, 76, len(result.Tasks.TaskItem))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get risk analysis report.
func Test_Report_GetRiskAnalysisReport(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetRiskAnalysisReportOpts{
		Name:             remoteFileName,
		TaskUid:          1,
		DistributionType: optional.NewString(string(models.NORMAL_ProbabilityDistributionType)),
		ConfidenceLevel:  optional.NewString(string(models.CL85_ConfidenceLevel)),
		Optimistic:       optional.NewInt32(80),
		Pessimistic:      optional.NewInt32(130),
		Iterations:       optional.NewInt32(200),
		Folder:           optional.NewString(remoteBaseTestDataFolder),
	}
	file, response, err := client.TasksApi.GetRiskAnalysisReport(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.Less(t, 1, len(file))
	fileAsStrings := ConvertBytesToStrings(file)
	assert.Equal(t, "%PDF-1.5", fileAsStrings[0])
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
