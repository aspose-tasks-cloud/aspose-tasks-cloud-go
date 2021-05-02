/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tasks_extended_attributes_test.go">
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

// Example of how to work with tasks extended attributes.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for put extended attribute with lookup value.
func Test_TasksExtendedAttributes_PutExtendedAttribute_WithLookupValue(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newCalculationType := models.LOOKUP_CalculationType
	newCfType := models.TEXT_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttributeDef := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Text3",
		Alias:           "New Field",
		ValueList: []models.Value{
			{
				Description: "descr1",
				Val:         "Internal",
				Id:          111,
			},
			{
				Description: "descr2",
				Val:         "External",
				Id:          112,
			},
		},
	}
	putOpts := &requests.PutExtendedAttributeOpts{
		ExtendedAttribute: newExtendedAttributeDef,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 27,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)

	newAttributeType := models.NULL_CustomFieldType
	newExtendedAttribute := models.ExtendedAttribute{
		AttributeType: &newAttributeType,
		LookupValueId: newExtendedAttributeDef.ValueList[1].Id,
		FieldId:       putResult.ExtendedAttribute.FieldId,
	}
	getTaskResult.Task.ExtendedAttributes = append(getTaskResult.Task.ExtendedAttributes, newExtendedAttribute)
	putTaskRequest := requests.PutTaskOpts{
		TaskUid: getTaskOpts.TaskUid,
		Task:    *getTaskResult.Task,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &putTaskRequest)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	getTaskResult, _, err = client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)
	assert.Equal(t, 1, len(getTaskResult.Task.ExtendedAttributes))
	assert.Equal(t, newExtendedAttribute.FieldId, getTaskResult.Task.ExtendedAttributes[0].FieldId)
	assert.Equal(t, newExtendedAttribute.LookupValueId, getTaskResult.Task.ExtendedAttributes[0].LookupValueId)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute with date value.
func Test_TasksExtendedAttributes_PutExtendedAttribute_WithDateValue(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newCalculationType := models.NONE_CalculationType
	newCfType := models.FINISH_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttributeDef := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Finish4",
		Alias:           "Custom Finish Field",
	}
	putOpts := &requests.PutExtendedAttributeOpts{
		ExtendedAttribute: newExtendedAttributeDef,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 27,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)

	newAttributeType := models.NULL_CustomFieldType
	newExtendedAttribute := models.ExtendedAttribute{
		AttributeType: &newAttributeType,
		DateValue:     CreateTime(2018, 3, 4, 12, 30, 0),
		FieldId:       putResult.ExtendedAttribute.FieldId,
	}
	getTaskResult.Task.ExtendedAttributes = append(getTaskResult.Task.ExtendedAttributes, newExtendedAttribute)
	putTaskRequest := requests.PutTaskOpts{
		TaskUid: getTaskOpts.TaskUid,
		Task:    *getTaskResult.Task,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &putTaskRequest)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	getTaskResult, _, err = client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)
	assert.Equal(t, 1, len(getTaskResult.Task.ExtendedAttributes))
	assert.Equal(t, newExtendedAttribute.FieldId, getTaskResult.Task.ExtendedAttributes[0].FieldId)
	assert.Equal(t, newExtendedAttribute.DateValue, getTaskResult.Task.ExtendedAttributes[0].DateValue)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute with duration value.
func Test_TasksExtendedAttributes_PutExtendedAttribute_WithDurationValue(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newCalculationType := models.NONE_CalculationType
	newCfType := models.DURATION_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttributeDef := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Duration3",
		Alias:           "Custom Duration Field",
	}
	putOpts := &requests.PutExtendedAttributeOpts{
		ExtendedAttribute: newExtendedAttributeDef,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 27,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)

	newAttributeType := models.NULL_CustomFieldType
	newTimeSpan := "04:00:00"
	newTimeUnitType := models.MINUTE_TimeUnitType
	newExtendedAttribute := models.ExtendedAttribute{
		AttributeType: &newAttributeType,
		DurationValue: &models.Duration{
			TimeSpan: &newTimeSpan,
			TimeUnit: &newTimeUnitType,
		},
		FieldId: putResult.ExtendedAttribute.FieldId,
	}
	getTaskResult.Task.ExtendedAttributes = append(getTaskResult.Task.ExtendedAttributes, newExtendedAttribute)
	putTaskRequest := requests.PutTaskOpts{
		TaskUid: getTaskOpts.TaskUid,
		Task:    *getTaskResult.Task,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &putTaskRequest)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	getTaskResult, _, err = client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)
	assert.Equal(t, 1, len(getTaskResult.Task.ExtendedAttributes))
	assert.Equal(t, newExtendedAttribute.FieldId, getTaskResult.Task.ExtendedAttributes[0].FieldId)
	assert.NotNil(t, getTaskResult.Task.ExtendedAttributes[0].DurationValue)
	assert.Equal(t, &newExtendedAttribute.DurationValue.TimeSpan, &getTaskResult.Task.ExtendedAttributes[0].DurationValue.TimeSpan)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute with flag value.
func Test_TasksExtendedAttributes_PutExtendedAttribute_WithFlagValue(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newCalculationType := models.NONE_CalculationType
	newCfType := models.FLAG_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttributeDef := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Flag12",
		Alias:           "Custom Flag Field",
	}
	putOpts := &requests.PutExtendedAttributeOpts{
		ExtendedAttribute: newExtendedAttributeDef,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 27,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)

	newAttributeType := models.NULL_CustomFieldType
	newExtendedAttribute := models.ExtendedAttribute{
		AttributeType: &newAttributeType,
		FlagValue:     true,
		FieldId:       putResult.ExtendedAttribute.FieldId,
	}
	getTaskResult.Task.ExtendedAttributes = append(getTaskResult.Task.ExtendedAttributes, newExtendedAttribute)
	putTaskRequest := requests.PutTaskOpts{
		TaskUid: getTaskOpts.TaskUid,
		Task:    *getTaskResult.Task,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &putTaskRequest)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	getTaskResult, _, err = client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)
	assert.Equal(t, 1, len(getTaskResult.Task.ExtendedAttributes))
	assert.Equal(t, newExtendedAttribute.FieldId, getTaskResult.Task.ExtendedAttributes[0].FieldId)
	assert.True(t, getTaskResult.Task.ExtendedAttributes[0].FlagValue)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute with cost value.
func Test_TasksExtendedAttributes_PutExtendedAttribute_WithCostValue(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newCalculationType := models.NONE_CalculationType
	newCfType := models.COST_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttributeDef := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Cost10",
		Alias:           "Custom Cost Field",
	}
	putOpts := &requests.PutExtendedAttributeOpts{
		ExtendedAttribute: newExtendedAttributeDef,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 27,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)

	newAttributeType := models.NULL_CustomFieldType
	newExtendedAttribute := models.ExtendedAttribute{
		AttributeType: &newAttributeType,
		NumericValue:  115.99,
		FieldId:       putResult.ExtendedAttribute.FieldId,
	}
	getTaskResult.Task.ExtendedAttributes = append(getTaskResult.Task.ExtendedAttributes, newExtendedAttribute)
	putTaskRequest := requests.PutTaskOpts{
		TaskUid: getTaskOpts.TaskUid,
		Task:    *getTaskResult.Task,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &putTaskRequest)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	getTaskResult, _, err = client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)
	assert.Equal(t, 1, len(getTaskResult.Task.ExtendedAttributes))
	assert.Equal(t, newExtendedAttribute.FieldId, getTaskResult.Task.ExtendedAttributes[0].FieldId)
	assert.Equal(t, newExtendedAttribute.NumericValue, getTaskResult.Task.ExtendedAttributes[0].NumericValue)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute with number value.
func Test_TasksExtendedAttributes_PutExtendedAttribute_WithNumberValue(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newCalculationType := models.NONE_CalculationType
	newCfType := models.NUMBER_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttributeDef := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Number9",
		Alias:           "Custom Number Field",
	}
	putOpts := &requests.PutExtendedAttributeOpts{
		ExtendedAttribute: newExtendedAttributeDef,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 27,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)

	newAttributeType := models.NULL_CustomFieldType
	newExtendedAttribute := models.ExtendedAttribute{
		AttributeType: &newAttributeType,
		NumericValue:  99.99,
		FieldId:       putResult.ExtendedAttribute.FieldId,
	}
	getTaskResult.Task.ExtendedAttributes = append(getTaskResult.Task.ExtendedAttributes, newExtendedAttribute)
	putTaskRequest := requests.PutTaskOpts{
		TaskUid: getTaskOpts.TaskUid,
		Task:    *getTaskResult.Task,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &putTaskRequest)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	getTaskResult, _, err = client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)
	assert.Equal(t, 1, len(getTaskResult.Task.ExtendedAttributes))
	assert.Equal(t, newExtendedAttribute.FieldId, getTaskResult.Task.ExtendedAttributes[0].FieldId)
	assert.Equal(t, newExtendedAttribute.NumericValue, getTaskResult.Task.ExtendedAttributes[0].NumericValue)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put extended attribute with text value.
func Test_TasksExtendedAttributes_PutExtendedAttribute_WithTextValue(t *testing.T) {
	localFileName := "NewProductDev.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	newCalculationType := models.NONE_CalculationType
	newCfType := models.TEXT_CustomFieldType
	newElementType := models.TASK_ElementType
	newRollupType := models.NULL_RollupType
	newExtendedAttributeDef := models.ExtendedAttributeDefinition{
		CalculationType: &newCalculationType,
		CfType:          &newCfType,
		ElementType:     &newElementType,
		RollupType:      &newRollupType,
		FieldName:       "Text1",
		Alias:           "Custom Text Field",
	}
	putOpts := &requests.PutExtendedAttributeOpts{
		ExtendedAttribute: newExtendedAttributeDef,
		Name:              remoteFileName,
		Folder:            optional.NewString(remoteBaseTestDataFolder),
	}
	putResult, _, err := client.TasksApi.PutExtendedAttribute(ctx, putOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putResult.Code)

	getTaskOpts := &requests.GetTaskOpts{
		TaskUid: 27,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	getTaskResult, _, err := client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)

	newAttributeType := models.NULL_CustomFieldType
	newExtendedAttribute := models.ExtendedAttribute{
		AttributeType: &newAttributeType,
		TextValue:     "Test value",
		FieldId:       putResult.ExtendedAttribute.FieldId,
	}
	getTaskResult.Task.ExtendedAttributes = append(getTaskResult.Task.ExtendedAttributes, newExtendedAttribute)
	putTaskRequest := requests.PutTaskOpts{
		TaskUid: getTaskOpts.TaskUid,
		Task:    *getTaskResult.Task,
		Name:    remoteFileName,
		Folder:  optional.NewString(remoteBaseTestDataFolder),
	}
	putTaskResult, _, err := client.TasksApi.PutTask(ctx, &putTaskRequest)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), putTaskResult.Code)

	getTaskResult, _, err = client.TasksApi.GetTask(ctx, getTaskOpts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), getTaskResult.Code)
	assert.NotNil(t, getTaskResult.Task)
	assert.Equal(t, 1, len(getTaskResult.Task.ExtendedAttributes))
	assert.Equal(t, newExtendedAttribute.FieldId, getTaskResult.Task.ExtendedAttributes[0].FieldId)
	assert.Equal(t, newExtendedAttribute.TextValue, getTaskResult.Task.ExtendedAttributes[0].TextValue)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
