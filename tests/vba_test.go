/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="vba_test.go">
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

// Example of how to work with vba.
package api_test

import (
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
)

// Test for get vba project.
func Test_Vba_GetVbaProject(t *testing.T) {
	localFileName := "VbaProject3.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetVbaProjectOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetVbaProject(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.VbaProject)
	assert.NotNil(t, result.VbaProject.Modules)
	assert.Equal(t, 8, len(result.VbaProject.Modules))
	assert.Equal(t, "Module1", result.VbaProject.Modules[1].Name)
	assert.True(t, strings.HasPrefix(result.VbaProject.Modules[1].SourceCode, "Type MEMORYSTATUS"))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func Test_Vba_GetVbaModule_ByIndex(t *testing.T) {
	localFileName := "VbaProject3.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: "1",
	}
	result, _, err := client.TasksApi.GetVbaModule(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Module)
	assert.Equal(t, "Module1", result.Module.Name)
	assert.NotEmpty(t, result.Module.SourceCode)
	assert.True(t, strings.HasPrefix(result.Module.SourceCode, "Type MEMORYSTATUS"))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func Test_Vba_GetVbaModule_ByName(t *testing.T) {
	localFileName := "VbaProject3.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: "Module1",
	}
	result, _, err := client.TasksApi.GetVbaModule(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Module)
	assert.Equal(t, "Module1", result.Module.Name)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func Test_Vba_PostVbaModule(t *testing.T) {
	localFileName := "VbaProject3.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	module := models.CreateVbaModuleRequest{
		Name:       "NewTestModule",
		Type_:      ptrVbaModuleType(models.PROCEDURAL_MODULE_VbaModuleType),
		SourceCode: "Sub TestProc()\n    MsgBox \"Hello\"\nEnd Sub",
	}

	opts := &requests.PostVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: module,
	}
	result, httpResponse, err := client.TasksApi.PostVbaModule(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, httpResponse.StatusCode == 200 || httpResponse.StatusCode == 201, "Expected status 200 or 201, got %d", httpResponse.StatusCode)
	assert.NotNil(t, result.Module)
	assert.Equal(t, "NewTestModule", result.Module.Name)
	assert.Equal(t, "Sub TestProc()\n    MsgBox \"Hello\"\nEnd Sub", result.Module.SourceCode)

	getOpts := &requests.GetVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: "NewTestModule",
	}
	getResult, _, err := client.TasksApi.GetVbaModule(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Module)
	assert.Equal(t, "NewTestModule", getResult.Module.Name)
	assert.Equal(t, "Sub TestProc()\n    MsgBox \"Hello\"\nEnd Sub", getResult.Module.SourceCode)
	assert.Equal(t, models.PROCEDURAL_MODULE_VbaModuleType, *getResult.Module.Type_)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func Test_Vba_PutVbaModule(t *testing.T) {
	localFileName := "VbaProject3.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newSourceCode := "Sub UpdatedProc()\n    MsgBox \"Updated\"\nEnd Sub"

	opts := &requests.PutVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: "Module1",
		UpdateRequest: models.UpdateVbaModuleRequest{
			SourceCode: newSourceCode,
		},
	}
	result, httpResponse, err := client.TasksApi.PutVbaModule(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, httpResponse.StatusCode == 200 || httpResponse.StatusCode == 201, "Expected status 200 or 201, got %d", httpResponse.StatusCode)
	assert.NotNil(t, result.Module)
	assert.Equal(t, "Module1", result.Module.Name)
	assert.Equal(t, "Sub UpdatedProc()\n    MsgBox \"Updated\"\nEnd Sub", result.Module.SourceCode)

	getOpts := &requests.GetVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: "Module1",
	}
	getResult, _, err := client.TasksApi.GetVbaModule(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Module)
	assert.Equal(t, "Module1", getResult.Module.Name)
	assert.Equal(t, "Sub UpdatedProc()\n    MsgBox \"Updated\"\nEnd Sub", getResult.Module.SourceCode)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func Test_Vba_DeleteVbaModule(t *testing.T) {
	localFileName := "VbaProject3.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getBeforeOpts := &requests.GetVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: "Module1",
	}
	getBeforeResult, _, err := client.TasksApi.GetVbaModule(ctx, getBeforeOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getBeforeResult.Code)
	assert.Equal(t, "Module1", getBeforeResult.Module.Name)

	deleteOpts := &requests.DeleteVbaModuleOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Module: "Module1",
	}
	deleteResult, _, err := client.TasksApi.DeleteVbaModule(ctx, deleteOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getProjectOpts := &requests.GetVbaProjectOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	getProjectResult, _, err := client.TasksApi.GetVbaProject(ctx, getProjectOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getProjectResult.Code)
	assert.NotNil(t, getProjectResult.VbaProject)
	moduleNames := make([]string, 0)
	for _, m := range getProjectResult.VbaProject.Modules {
		moduleNames = append(moduleNames, m.Name)
	}
	assert.NotContains(t, moduleNames, "Module1")
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func Test_Vba_ClearVba(t *testing.T) {
	localFileName := "VbaProject3.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getBeforeOpts := &requests.GetVbaProjectOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	getBeforeResult, _, err := client.TasksApi.GetVbaProject(ctx, getBeforeOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getBeforeResult.Code)
	assert.Greater(t, len(getBeforeResult.VbaProject.Modules), 0)

	clearOpts := &requests.ClearVbaOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	clearResult, _, err := client.TasksApi.ClearVba(ctx, clearOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), clearResult.Code)

	getAfterOpts := &requests.GetVbaProjectOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	getAfterResult, _, err := client.TasksApi.GetVbaProject(ctx, getAfterOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getAfterResult.Code)
	assert.Equal(t, 0, len(getAfterResult.VbaProject.Modules))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

func ptrVbaModuleType(value models.VbaModuleType) *models.VbaModuleType {
	return &value
}
