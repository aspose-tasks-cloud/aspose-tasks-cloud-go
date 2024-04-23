/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="views_test.go">
 *   Copyright (c) 2024 Aspose.Tasks Cloud
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

// Example of how to work with views.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get views.
func Test_Views_getViews(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetViewsOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetViews(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	firstView := result.Views[0]
	assert.NotNil(t, firstView)
	assert.Equal(t, true, firstView.ShowInMenu)
	assert.Equal(t, models.TASK_ITEM_ItemType, *firstView.Type_)
	assert.Equal(t, models.GANTT_ViewScreen, *firstView.Screen)
	assert.Equal(t, "&Gantt Chart", firstView.Name)
	assert.Equal(t, int32(1), firstView.Uid)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get all table text styles.
func Test_Views_getAllTableTextStyles(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetAllTableTextStylesOpts{
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
		ViewUid: 2,
	}
	result, _, err := client.TasksApi.GetAllTableTextStyles(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.Equal(t, 8, len(result.Items))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get table text style.
func Test_Views_getTableTextStyle(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetTableTextStyleOpts{
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
		ViewUid: 2,
		RowUid:  29,
	}
	result, _, err := client.TasksApi.GetTableTextStyle(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.TableTextStyle)
	assert.Equal(t, int32(29), result.TableTextStyle.RowUid)
	assert.Equal(t, models.UNDEFINED_Field, *result.TableTextStyle.Field)
	assert.Equal(t, models.ALLOCATED_TextItemType, *result.TableTextStyle.ItemType)
	assert.Equal(t, models.TRANSPARENT_Colors, *result.TableTextStyle.Color)
	assert.Equal(t, models.HOLLOW_BackgroundPattern, *result.TableTextStyle.BackgroundPattern)
	assert.Equal(t, models.TRANSPARENT_Colors, *result.TableTextStyle.BackgroundColor)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for create table text style.
func Test_Views_createTableTextStyle(t *testing.T) {
	localFileName := "Home_move_plan.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newField := models.TASK_NAME_Field
	newItemType := models.ALLOCATED_TextItemType
	newColor := models.RED_Colors
	newBackgroundPattern := models.HOLLOW_BackgroundPattern
	newBackgroundColor := models.GREEN_YELLOW_Colors
	newTableTextStyle := models.TableTextStyle{
		RowUid:            3,
		Field:             &newField,
		Color:             &newColor,
		ItemType:          &newItemType,
		BackgroundPattern: &newBackgroundPattern,
		BackgroundColor:   &newBackgroundColor,
	}

	createOpts := &requests.CreateTableTextStyleOpts{
		Name:           remoteFileName,
		Folder:         optional.NewString(remoteBaseTestDataFolder),
		ViewUid:        1,
		TableTextStyle: newTableTextStyle,
	}
	createResult, _, err := client.TasksApi.CreateTableTextStyle(ctx, createOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(201), createResult.Code)

	getOpts := &requests.GetTableTextStyleOpts{
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
		ViewUid: createOpts.ViewUid,
		RowUid:  createOpts.TableTextStyle.RowUid,
	}
	getResult, _, err := client.TasksApi.GetTableTextStyle(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, getResult.TableTextStyle)
	assert.Equal(t, newTableTextStyle.RowUid, getResult.TableTextStyle.RowUid)
	assert.Equal(t, newTableTextStyle.Field, getResult.TableTextStyle.Field)
	assert.Equal(t, newTableTextStyle.ItemType, getResult.TableTextStyle.ItemType)
	assert.Equal(t, newTableTextStyle.Color, getResult.TableTextStyle.Color)
	assert.Equal(t, newTableTextStyle.BackgroundPattern, getResult.TableTextStyle.BackgroundPattern)
	assert.Equal(t, newTableTextStyle.BackgroundColor, getResult.TableTextStyle.BackgroundColor)

	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for update table text style.
func Test_Views_updateTableTextStyle(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getOpts := &requests.GetTableTextStyleOpts{
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
		ViewUid: 2,
		RowUid:  29,
	}
	getResult, _, err := client.TasksApi.GetTableTextStyle(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, models.TRANSPARENT_Colors, *getResult.TableTextStyle.Color)
	assert.Equal(t, models.TRANSPARENT_Colors, *getResult.TableTextStyle.BackgroundColor)

	newColor := models.ORANGE_RED_Colors
	newBackgroundColor := models.DODGER_BLUE_Colors
	getResult.TableTextStyle.Color = &newColor
	getResult.TableTextStyle.BackgroundColor = &newBackgroundColor

	updateOpts := &requests.UpdateTableTextStyleOpts{
		Name:           remoteFileName,
		Folder:         optional.NewString(remoteBaseTestDataFolder),
		ViewUid:        getOpts.ViewUid,
		TableTextStyle: *getResult.TableTextStyle,
	}
	updateResult, _, err := client.TasksApi.UpdateTableTextStyle(ctx, updateOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), updateResult.Code)

	getResult, _, err = client.TasksApi.GetTableTextStyle(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, newColor, *getResult.TableTextStyle.Color)
	assert.Equal(t, newBackgroundColor, *getResult.TableTextStyle.BackgroundColor)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete table text style.
func Test_Views_deleteTableTextStyle(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getOpts := &requests.GetAllTableTextStylesOpts{
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
		ViewUid: 2,
	}
	getResult, _, err := client.TasksApi.GetAllTableTextStyles(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, 8, len(getResult.Items))

	deleteOpts := &requests.DeleteTableTextStyleOpts{
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
		ViewUid: getOpts.ViewUid,
		RowUid:  29,
	}
	deleteResult, _, err := client.TasksApi.DeleteTableTextStyle(ctx, deleteOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getResult, _, err = client.TasksApi.GetAllTableTextStyles(ctx, getOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.Equal(t, 7, len(getResult.Items))

	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
