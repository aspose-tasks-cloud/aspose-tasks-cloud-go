/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_properties_test.go">
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

// Example of how to work with document properties.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get document properties.
func Test_DocumentProperties_GetDocumentProperties(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetDocumentPropertiesOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetDocumentProperties(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Properties)
	assert.Equal(t, 52, len(result.Properties.List))
	assert.Equal(t, "Title", result.Properties.List[0].Name)
	assert.Equal(t, "Home Move", result.Properties.List[0].Value)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get document property.
func Test_DocumentProperties_GetDocumentProperty(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "Title",
	}
	result, _, err := client.TasksApi.GetDocumentProperty(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.Property)
	assert.Equal(t, "Title", result.Property.Name)
	assert.Equal(t, "Home Move", result.Property.Value)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put document property.
func Test_DocumentProperties_PutDocumentProperty(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	putOptions := &requests.PutDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "Title",
		Property: models.DocumentProperty{
			Name: "Title",
			Value: "New title value",
		},
	}
	putResult, _, err := client.TasksApi.PutDocumentProperty(ctx, putOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)
	assert.NotNil(t, putResult.Property)
	assert.Equal(t, "Title", putResult.Property.Name)
	assert.Equal(t, "New title value", putResult.Property.Value)

	getOptions := &requests.GetDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "Title",
	}
	getResult, _, err := client.TasksApi.GetDocumentProperty(ctx, getOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Property)
	assert.Equal(t, "Title", getResult.Property.Name)
	assert.Equal(t, "New title value", getResult.Property.Value)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put document property should ignore changes of nonexistent property.
func Test_DocumentProperties_PutDocumentPropertyShouldIgnoreNonexistentProperty(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	putOptions := &requests.PutDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "new property",
		Property: models.DocumentProperty{
			Name: "new property",
			Value: "new property value",
		},
	}
	putResult, _, err := client.TasksApi.PutDocumentProperty(ctx, putOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)
	assert.Nil(t, putResult.Property)

	getOptions := &requests.GetDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "new property",
	}
	getResult, _, err := client.TasksApi.GetDocumentProperty(ctx, getOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Nil(t, getResult.Property)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post document property.
func Test_DocumentProperties_PostDocumentProperty(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	postOptions := &requests.PostDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "Title",
		Property: models.DocumentProperty{
			Name: "Title",
			Value: "New title value",
		},
	}
	postResult, _, err := client.TasksApi.PostDocumentProperty(ctx, postOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), postResult.Code)
	assert.NotNil(t, postResult.Property)
	assert.Equal(t, "Title", postResult.Property.Name)
	assert.Equal(t, "New title value", postResult.Property.Value)

	getOptions := &requests.GetDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "Title",
	}
	getResult, _, err := client.TasksApi.GetDocumentProperty(ctx, getOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.Property)
	assert.Equal(t, "Title", getResult.Property.Name)
	assert.Equal(t, "New title value", getResult.Property.Value)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for post document property should ignore changes of nonexistent property.
func Test_DocumentProperties_PostDocumentPropertyShouldIgnoreNonexistentProperty(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	postOptions := &requests.PostDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "new property",
		Property: models.DocumentProperty{
			Name: "new property",
			Value: "new property value",
		},
	}
	postResult, _, err := client.TasksApi.PostDocumentProperty(ctx, postOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), postResult.Code)
	assert.Nil(t, postResult.Property)

	getOptions := &requests.GetDocumentPropertyOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		PropertyName: "new property",
	}
	getResult, _, err := client.TasksApi.GetDocumentProperty(ctx, getOptions)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Nil(t, getResult.Property)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
