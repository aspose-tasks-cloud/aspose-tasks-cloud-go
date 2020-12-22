/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="api_client_test.go">
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

package api_test

import (
    "context"
    "os"
    "testing"
    "github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api"
    "github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
    "github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
    "github.com/stretchr/testify/assert"
)

func TestAppSidAndAppKey(t *testing.T) {
    p := []struct {
        appKey string
        appSid string
    }{
        {"", "x"},
        {"x", ""},
    }
    for _, cp := range p {
        config := models.Configuration{
            AppKey: cp.appKey,
            AppSid: cp.appSid,
        }
        _, err := api.NewAPIClient(&config)

        assert.Error(t, err)
    }
}

func TestBaseUrl(t *testing.T) {
    config := models.Configuration{
        AppKey:  "x",
        AppSid:  "x",
        BaseUrl: "x",
    }
    _, err := api.NewAPIClient(&config)

    assert.Error(t, err)
}

func TestWrongAppKey(t *testing.T) {
    config := ReadConfiguration(t)
    config.AppKey = "x"
    _, _, err := api.CreateTasksApi(config)

    assert.Error(t, err)
}

func TestWrongAppSid(t *testing.T) {
    config := ReadConfiguration(t)
    config.AppSid = "x"
    _, _, err := api.CreateTasksApi(config)

    assert.Error(t, err)
}

func TestUnathorizedAccess(t *testing.T) {
    config := ReadConfiguration(t)
    client, err := api.NewAPIClient(config)

    if err != nil {
        t.Error(err)
    }
    fileName := "NewProductDev.mpp"
    remoteFilePath := remoteBaseTestDataFolder + "/Storage/" + fileName
    localFilePath := GetLocalPath(fileName)
    file, fileErr := os.Open(localFilePath)
    if fileErr != nil {
        t.Error(fileErr)
    }

    request := requests.UploadFileOpts{
        Path: remoteFilePath,
        File:        file,
    }

    _, _, err = client.TasksApi.UploadFile(context.Background(), &request)

    assert.Error(t, err)
}

