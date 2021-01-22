/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="import_project_test.go">
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

// Example of how to work with import project.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for put import project from file.
func Test_ImportProject_PutImportProjectFromFile(t *testing.T) {
	localFileName := "p6_multiproject.xml"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	putOptions := &requests.PutImportProjectFromFileOpts{
		Name:             remoteFileName,
		ProjectUid:       "111",
		Filename:         "imported_from_primavera.xml",
		Folder:           optional.NewString(remoteBaseTestDataFolder),
		FileType:         optional.NewString(string(models.PRIMAVERA_XML_ImportedProjectType)),
		OutputFileFormat: optional.NewString(string(models.P6XML_ProjectFileFormat)),
	}
	putProjectResult, _, err := client.TasksApi.PutImportProjectFromFile(ctx, putOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putProjectResult.Code)
	uploadedTestFiles.mx.Lock()
	uploadedTestFiles.items = append(uploadedTestFiles.items, "imported_from_primavera.xml")
	uploadedTestFiles.mx.Unlock()

	getTasksResult, _, err := client.TasksApi.GetTasks(ctx, &requests.GetTasksOpts{
		Name:   putOptions.Filename,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTasksResult.Code)
	assert.Equal(t, 12, len(getTasksResult.Tasks.TaskItem))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put import project from database.
func Test_ImportProject_PutImportProjectFromDb(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	putOptions := &requests.PutImportProjectFromDbOpts{
		ConnectionString: "Data Source=db.contoso.com;Initial Catalog=ProjectServer;Persist Security Info=True;User ID=sa;Password=pwd;",
		DatabaseType:     string(models.MSP_ProjectDatabaseType),
		Filename:         "imported_from_db.xml",
		Folder:           optional.NewString(remoteBaseTestDataFolder),
		ProjectUid:       "E6426C44-D6CB-4B9C-AF16-48910ACE0F54",
		DatabaseSchema:   optional.NewString("dbo"),
		Format:           optional.NewString(string(models.P6XML_ProjectFileFormat)),
	}
	putProjectResult, _, err := client.TasksApi.PutImportProjectFromDb(ctx, putOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putProjectResult.Code)
	uploadedTestFiles.mx.Lock()
	uploadedTestFiles.items = append(uploadedTestFiles.items, putOptions.Filename)
	uploadedTestFiles.mx.Unlock()

	getProjectIdsResult, _, err := client.TasksApi.GetProjectIds(ctx, &requests.GetProjectIdsOpts{
		Name:   putOptions.Filename,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getProjectIdsResult.Code)
	assert.Equal(t, []string{"1"}, getProjectIdsResult.ProjectIds)

	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put import project from project online / project server using token credentials.
func Test_ImportProject_PutImportProjectFromProjectOnlineByTokenCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	putOptions := &requests.PutImportProjectFromProjectOnlineOpts{
		Name:                "NewProductDev.mpp",
		Guid:                "E6426C44-D6CB-4B9C-AF16-48910ACE0F54",
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		XProjectOnlineToken: optional.NewString("SOMESECRETTOKEN"),
		Format:              optional.NewString(string(models.P6XML_ProjectFileFormat)),
		Folder:              optional.NewString(remoteBaseTestDataFolder),
	}
	putProjectResult, _, err := client.TasksApi.PutImportProjectFromProjectOnline(ctx, putOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putProjectResult.Code)
	uploadedTestFiles.mx.Lock()
	uploadedTestFiles.items = append(uploadedTestFiles.items, putOptions.Name)
	uploadedTestFiles.mx.Unlock()

	file, response, err := client.TasksApi.DownloadFile(ctx, &requests.DownloadFileOpts{
		Path: putOptions.Folder.Value() + "/" + putOptions.Name,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.Less(t, 1, len(file))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put import project from project online / project server using login and password credentials.
func Test_ImportProject_PutImportProjectFromProjectOnlineByLoginPasswordCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	putOptions := &requests.PutImportProjectFromProjectOnlineOpts{
		Name:                "NewProductDev.mpp",
		Guid:                "E6426C44-D6CB-4B9C-AF16-48910ACE0F54",
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		UserName:            optional.NewString("SomeLogin"),
		XSharepointPassword: optional.NewString("SomePassword"),
		Format:              optional.NewString(string(models.P6XML_ProjectFileFormat)),
		Folder:              optional.NewString(remoteBaseTestDataFolder),
	}
	putProjectResult, _, err := client.TasksApi.PutImportProjectFromProjectOnline(ctx, putOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putProjectResult.Code)
	uploadedTestFiles.mx.Lock()
	uploadedTestFiles.items = append(uploadedTestFiles.items, putOptions.Name)
	uploadedTestFiles.mx.Unlock()

	file, response, err := client.TasksApi.DownloadFile(ctx, &requests.DownloadFileOpts{
		Path: putOptions.Folder.Value() + "/" + putOptions.Name,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.Less(t, 1, len(file))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
