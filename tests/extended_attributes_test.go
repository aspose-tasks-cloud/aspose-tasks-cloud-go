/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="extended_attributes_test.go">
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

// Example of how to work with extended attributes.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get extended attributes.
func Test_ExtendedAttributes_GetExtendedAttributes(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetExtendedAttributesOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetExtendedAttributes(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.ExtendedAttributes)
	assert.Equal(t, 2, len(result.ExtendedAttributes.List))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for get extended attribute by index.
func Test_ExtendedAttributes_GetExtendedAttributeByIndex(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	options := &requests.GetExtendedAttributeByIndexOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Index:  1,
	}
	result, _, err := client.TasksApi.GetExtendedAttributeByIndex(ctx, options)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.ExtendedAttribute)
	assert.Equal(t, "Text1", result.ExtendedAttribute.FieldName)
	assert.Equal(t, models.LOOKUP_CalculationType, *result.ExtendedAttribute.CalculationType)
	assert.Equal(t, 1, len(result.ExtendedAttribute.ValueList))
	assert.Equal(t, "descr", result.ExtendedAttribute.ValueList[0].Description)
	assert.Equal(t, int32(1), result.ExtendedAttribute.ValueList[0].Id)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute, it should add new ExtendedAttribute.
func Test_ExtendedAttributes_PutExtendedAttribute(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)
	firstValue := models.Value{
		Description: "descr1",
		Val:         "Internal",
		Id:          111,
	}
	secondValue := models.Value{
		Description: "descr2",
		Val:         "External",
		Id:          112,
	}
	newCalculationType := models.LOOKUP_CalculationType
	newCfType := models.TEXT_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttribute := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Text3",
		Alias:           "New Field",
		ValueList:       []models.Value{firstValue, secondValue},
	}

	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, &requests.PutExtendedAttributeOpts{
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
		ExtendedAttribute: newExtendedAttribute,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)
	assert.NotNil(t, putResult.ExtendedAttribute)
	assert.Equal(t, newExtendedAttribute.FieldName, putResult.ExtendedAttribute.FieldName)
	assert.Equal(t, newExtendedAttribute.Alias, putResult.ExtendedAttribute.Alias)
	assert.Equal(t, "188743737", putResult.ExtendedAttribute.FieldId)

	addedAttributeIndex := &putResult.ExtendedAttribute.Index

	getResult, _, err := client.TasksApi.GetExtendedAttributeByIndex(ctx, &requests.GetExtendedAttributeByIndexOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Index:  *addedAttributeIndex,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.ExtendedAttribute)
	assert.Equal(t, newExtendedAttribute.CfType, getResult.ExtendedAttribute.CfType)
	assert.Equal(t, newExtendedAttribute.FieldName, getResult.ExtendedAttribute.FieldName)
	assert.Equal(t, newExtendedAttribute.Alias, putResult.ExtendedAttribute.Alias)
	assert.Equal(t, putResult.ExtendedAttribute.FieldId, getResult.ExtendedAttribute.FieldId)
	assert.Equal(t, 2, len(getResult.ExtendedAttribute.ValueList))
	assert.Equal(t, firstValue.Id, getResult.ExtendedAttribute.ValueList[0].Id)
	assert.Equal(t, firstValue.Val, getResult.ExtendedAttribute.ValueList[0].Val)
	assert.Equal(t, firstValue.Description, getResult.ExtendedAttribute.ValueList[0].Description)
	assert.Equal(t, secondValue.Id, getResult.ExtendedAttribute.ValueList[1].Id)
	assert.Equal(t, secondValue.Val, getResult.ExtendedAttribute.ValueList[1].Val)
	assert.Equal(t, secondValue.Description, getResult.ExtendedAttribute.ValueList[1].Description)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute, it should edit ExtendedAttribute.
func Test_ExtendedAttributes_PutExtendedAttributeShouldEditExtendedAttribute(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	getResult, _, err := client.TasksApi.GetExtendedAttributeByIndex(ctx, &requests.GetExtendedAttributeByIndexOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Index:  1,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.ExtendedAttribute)

	newCalculationType := models.NONE_CalculationType
	newElementType := models.TASK_ElementType
	newCfType := models.TEXT_CustomFieldType
	extendedAttributeToEdit := getResult.ExtendedAttribute
	extendedAttributeToEdit.CalculationType = &newCalculationType
	extendedAttributeToEdit.CfType = &newCfType
	extendedAttributeToEdit.ElementType = &newElementType
	extendedAttributeToEdit.FieldName = "Text1"
	extendedAttributeToEdit.ValueList = []models.Value{}
	extendedAttributeToEdit.Alias = "Edited field"

	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, &requests.PutExtendedAttributeOpts{
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
		ExtendedAttribute: *extendedAttributeToEdit,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)
	assert.NotNil(t, putResult.ExtendedAttribute)
	assert.Equal(t, extendedAttributeToEdit.FieldName, putResult.ExtendedAttribute.FieldName)
	assert.Equal(t, extendedAttributeToEdit.Alias, putResult.ExtendedAttribute.Alias)
	assert.Equal(t, "188743731", putResult.ExtendedAttribute.FieldId)

	addedAttributeIndex := &putResult.ExtendedAttribute.Index

	getResult, _, err = client.TasksApi.GetExtendedAttributeByIndex(ctx, &requests.GetExtendedAttributeByIndexOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Index:  *addedAttributeIndex,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.ExtendedAttribute)
	assert.Equal(t, putResult.ExtendedAttribute.FieldId, getResult.ExtendedAttribute.FieldId)
	assert.Equal(t, extendedAttributeToEdit.CfType, getResult.ExtendedAttribute.CfType)
	assert.Equal(t, extendedAttributeToEdit.FieldName, getResult.ExtendedAttribute.FieldName)
	assert.Equal(t, extendedAttributeToEdit.Alias, getResult.ExtendedAttribute.Alias)
	assert.Equal(t, putResult.ExtendedAttribute.FieldId, getResult.ExtendedAttribute.FieldId)
	assert.Equal(t, 0, len(getResult.ExtendedAttribute.ValueList))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for delete extended attribute by index.
func Test_ExtendedAttributes_DeleteExtendedAttributeByIndex(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	deleteResult, _, err := client.TasksApi.DeleteExtendedAttributeByIndex(ctx, &requests.DeleteExtendedAttributeByIndexOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
		Index:  1,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), deleteResult.Code)

	getResult, _, err := client.TasksApi.GetExtendedAttributes(ctx, &requests.GetExtendedAttributesOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getResult.Code)
	assert.NotNil(t, getResult.ExtendedAttributes)
	assert.Equal(t, 1, len(getResult.ExtendedAttributes.List))
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
