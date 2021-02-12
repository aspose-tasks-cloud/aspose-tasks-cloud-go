/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="wbs_test.go">
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

// Example of how to work with wbs.
package api_test

import (
	"github.com/antihax/optional"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test for get wbs definition.
func Test_Wbs_GetWbsDefinition(t *testing.T) {
	localFileName := "WBSDefinition.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.GetWbsDefinitionOpts{
		Name:   remoteFileName,
		Folder: optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.GetWbsDefinition(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	assert.NotNil(t, result.WbsDefinition)
	assert.True(t, result.WbsDefinition.GenerateWBSCode)
	assert.True(t, result.WbsDefinition.VerifyUniqueness)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}

// Test for put  renumber wbs code.
func Test_Wbs_PutRenumberWbsCode(t *testing.T) {
	localFileName := "WBSDefinition.mpp"
	remoteFileName := CreateRandomGuid() + localFileName
	client, ctx := UploadTestFileToStorage(t, localFileName, remoteBaseTestDataFolder+"/"+remoteFileName)

	opts := &requests.PutRenumberWbsCodeOpts{
		TaskUids: []int32{},
		Name:     remoteFileName,
		Folder:   optional.NewString(remoteBaseTestDataFolder),
	}
	result, _, err := client.TasksApi.PutRenumberWbsCode(ctx, opts)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int32(200), result.Code)
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
