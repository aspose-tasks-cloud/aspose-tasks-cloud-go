/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="folder_test.go">
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

// Example of how to work with folders.
package api_test

import (
    "github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
    "testing"
)

// Test for delete folder.
func Test_Folder_DeleteFolder(t *testing.T) {
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFileName := "NewProductDev.mpp"
    testDeleteFolder := remoteDataFolder + "/TestDeleteFolder"

    client, ctx := UploadTestFileToStorage(t, localFileName, testDeleteFolder + "/TestDeleteFolder.mpp")

    options := &requests.DeleteFolderOpts{
        Path : testDeleteFolder,
    }
    _, err := client.TasksApi.DeleteFolder(ctx, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for create folder.
func Test_Folder_CreateFolder(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage/TestCreateFolder"


    options := &requests.CreateFolderOpts{
        Path : remoteDataFolder,
    }
    _, err := client.TasksApi.CreateFolder(ctx, options)

    if err != nil {
        t.Error(err)
    }

    client.TasksApi.DeleteFolder(ctx, &requests.DeleteFolderOpts{
        Path : remoteDataFolder,
    })
}

// Test for get file list of folder.
func Test_Folder_GetFilesList(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"

    options := &requests.GetFilesListOpts{
        Path : remoteDataFolder,
    }
    _, _, err := client.TasksApi.GetFilesList(ctx, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for copy folder.
func Test_Folder_CopyFolder(t *testing.T) {

    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFileName := "NewProductDev.mpp"
    folderToCopy := remoteDataFolder + "/TestCopyFolder"

    client, ctx := UploadTestFileToStorage(t, localFileName, folderToCopy + "Src/TestCopyFolder.mpp")

    options := &requests.CopyFolderOpts{
        SrcPath : folderToCopy + "Src",
        DestPath: folderToCopy + "Dest",
    }
    _, err := client.TasksApi.CopyFolder(ctx, options)

    if err != nil {
        t.Error(err)
    }

    client.TasksApi.DeleteFolder(ctx, &requests.DeleteFolderOpts{
        Path : folderToCopy + "Dest",
    })
}

// Test for move folder.
func Test_Folder_MoveFolder(t *testing.T) {
    remoteDataFolder := remoteBaseTestDataFolder + "/Storage"
    localFileName := "NewProductDev.mpp"
    folderToMove := remoteDataFolder + "/TestMoveFolder"

    client, ctx := UploadTestFileToStorage(t, localFileName, folderToMove + "Src/TestMoveFolder.mpp")

    options := &requests.MoveFolderOpts{
        SrcPath : folderToMove + "Src",
        DestPath: folderToMove + "Dest",
    }
    _, err := client.TasksApi.MoveFolder(ctx, options)

    if err != nil {
        t.Error(err)
    }

    client.TasksApi.DeleteFolder(ctx, &requests.DeleteFolderOpts{
        Path : folderToMove + "Dest",
    })
}
