/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="project_online_test.go">
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

// Example of how to work with project online / project server.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get project list from project online / project server using token credentials.
func Test_ProjectOnline_GetProjectListByTokenCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	getOptions := &requests.GetProjectListOpts{
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		XProjectOnlineToken: optional.NewString("SOMESECRETTOKEN"),
	}
	result, _, err := client.TasksApi.GetProjectList(ctx, getOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Projects)
	assert.LessOrEqual(t, 1, len(result.Projects.ProjectInfo))
}

// Test for get project list from project online / project server using login and password credentials.
func Test_ProjectOnline_GetProjectListByLoginPasswordCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	getOptions := &requests.GetProjectListOpts{
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		UserName:            optional.NewString("SomeLogin"),
		XSharepointPassword: optional.NewString("SomePassword"),
	}
	result, _, err := client.TasksApi.GetProjectList(ctx, getOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Projects)
	assert.LessOrEqual(t, 1, len(result.Projects.ProjectInfo))
}

// Test for create new project in project online / project server using token credentials.
func Test_ProjectOnline_CreateNewProjectByTokenCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.CreateNewProjectOpts{
		Name:                remoteFileName,
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		XProjectOnlineToken: optional.NewString("SOMESECRETTOKEN"),
		Folder:              optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.CreateNewProject(ctx, options)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for create new project in project online / project server using login and password credentials.
func Test_ProjectOnline_CreateNewProjectByLoginPasswordCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.CreateNewProjectOpts{
		Name:                remoteFileName,
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		UserName:            optional.NewString("SomeLogin"),
		XSharepointPassword: optional.NewString("SomePassword"),
		Folder:              optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.CreateNewProject(ctx, options)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for update project in project online / project server using token credentials.
func Test_ProjectOnline_UpdateProjectByTokenCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.UpdateProjectOpts{
		Name:                remoteFileName,
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		XProjectOnlineToken: optional.NewString("SOMESECRETTOKEN"),
		Folder:              optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.UpdateProject(ctx, options)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for update project in project online / project server using login and password credentials.
func Test_ProjectOnline_UpdateProjectByLoginPasswordCredentials(t *testing.T) {
	t.Skip("Ignored because real credentials is required")
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.UpdateProjectOpts{
		Name:                remoteFileName,
		SiteUrl:             "http://project_server_instance.local/sites/pwa",
		UserName:            optional.NewString("SomeLogin"),
		XSharepointPassword: optional.NewString("SomePassword"),
		Folder:              optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.UpdateProject(ctx, options)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
