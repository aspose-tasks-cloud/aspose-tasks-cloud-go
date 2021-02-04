/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="task_document_format_test.go">
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

// Example of how to work with task document format.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get task document with format.
func Test_TaskDocumentFormat_GetTaskDocumentWithFormat(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTaskDocumentWithFormatOpts{
		Name:   remoteFileName,
		Format: string(models.CSV_ProjectFileFormat),
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	data, response, err := client.TasksApi.GetTaskDocumentWithFormat(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, data)
	assert.Less(t, 1, len(data))
	fileAsStrings := ConvertBytesToStrings(data)
	assert.Equal(t, "ID;Task_Name;Outline_Level;Duration;Start_Date;Finish_Date;Percent_Comp;Cost;Work", fileAsStrings[0])
	assert.Equal(t, "1;Five to Eight Weeks Before Moving;1;16 days;Thu 01.01.04 08:00;Thu 22.01.04 17:00;0%;$0;0 hrs", fileAsStrings[1])
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get task document with format as zipped html.
func Test_TaskDocumentFormat_GetTaskDocumentWithFormatAsZippedHtml(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTaskDocumentWithFormatOpts{
		Name:               remoteFileName,
		Format:             string(models.HTML_ProjectFileFormat),
		ReturnAsZipArchive: optional.NewBool(true),
		Folder:             optional.NewString(remoteBaseTestDataFolder),
	}
	data, response, err := client.TasksApi.GetTaskDocumentWithFormat(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, data)
	assert.Less(t, 1, len(data))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get task document with format.
func Test_TaskDocumentFormat_PostTaskDocumentWithFormat(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	// SaveOptions parameters is a representation of
	// Aspose.Tasks's SaveOptions class or its format-specific inheritors (Like CsvOptions, etc):
	// See Aspose.Tasks reference: https://apireference.aspose.com/net/tasks/aspose.tasks.saving/saveoptions
	type (
		Column = struct {
			Type     string
			Name     string
			Property string
			Width    int
		}
		View = struct {
			Columns []Column
		}
		saveOptions = struct {
			TextDelimiter           string
			IncludeHeaders          bool
			NonExistingTestProperty bool
			View                    View
		}
	)
	so := saveOptions{
		"Comma",
		false,
		false,
		View{
			[]Column{
				{"GanttChartColumn", "TestColumn1", "Name", 120},
				{"GanttChartColumn", "TestColumn2", "Duration", 120},
			},
		},
	}
	opts := &requests.PostTaskDocumentWithFormatOpts{
		Name:        remoteFileName,
		SaveOptions: so,
		Format:      string(models.CSV_ProjectFileFormat),
		Folder:      optional.NewString(remoteBaseTestDataFolder),
	}
	data, response, err := client.TasksApi.PostTaskDocumentWithFormat(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, data)
	assert.Less(t, 1, len(data))
	fileAsStrings := ConvertBytesToStrings(data)
	assert.Equal(t, "Five to Eight Weeks Before Moving,16 days", fileAsStrings[0])
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
