/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="file_test.go">
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

// Example of how to work with files.
package api_test

import (
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"testing"
)

// Test for uploading file.
func Test_File_UploadFile(t *testing.T) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)
	remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
	fileName := "NewProductDev.mpp"

	options := &requests.UploadFileOpts{
		File: OpenFile(t, fileName),
		Path: remoteDataFolder + "/" + fileName,
	}
	_, _, err := client.TasksApi.UploadFile(ctx, options)

	if err != nil {
		t.Error(err)
	}
}

// Test for copy file.
func Test_File_CopyFile(t *testing.T) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)
	remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
	fileName := "NewProductDev.mpp"

	client.TasksApi.UploadFile(ctx, &requests.UploadFileOpts{
		File: OpenFile(t, fileName),
		Path: remoteDataFolder + "/" + fileName,
	})

	options := &requests.CopyFileOpts{
		SrcPath:  remoteDataFolder + "/" + fileName,
		DestPath: remoteDataFolder + "/TestCopyFileDest.mpp",
	}
	_, err := client.TasksApi.CopyFile(ctx, options)

	if err != nil {
		t.Error(err)
	}
}

// Test for move file.
func Test_File_MoveFile(t *testing.T) {
	remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
	localFileName := "NewProductDev.mpp"
	remoteFilePath := remoteDataFolder + "/" + CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteFilePath)

	options := &requests.MoveFileOpts{
		SrcPath:  remoteFilePath,
		DestPath: remoteDataFolder + "/TestMoveFileDest.mpp",
	}
	_, err := client.TasksApi.MoveFile(ctx, options)

	if err != nil {
		t.Error(err)
	}
}

// Test for delete file.
func Test_File_DeleteFile(t *testing.T) {
	remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
	localFileName := "NewProductDev.mpp"
	remoteFilePath := remoteDataFolder + "/" + CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteFilePath)

	options := &requests.DeleteFileOpts{
		Path: remoteFilePath,
	}
	_, err := client.TasksApi.DeleteFile(ctx, options)

	if err != nil {
		t.Error(err)
	}
}

// Test for download file.
func Test_File_DownloadFile(t *testing.T) {
	remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
	localFileName := "NewProductDev.mpp"
	remoteFilePath := remoteDataFolder + "/" + CreateRandomGuid() + localFileName

	client, ctx := UploadTestFileToStorage(t, localFileName, remoteFilePath)

	options := &requests.DownloadFileOpts{
		Path: remoteFilePath,
	}
	_, _, err := client.TasksApi.DownloadFile(ctx, options)

	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() { DeleteTestFileFromStorage(t, ctx, client) })
}
