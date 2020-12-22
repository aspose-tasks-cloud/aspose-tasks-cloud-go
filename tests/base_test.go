/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="base_test.go">
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
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/custom"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

var remoteBaseTestDataFolder = "Temp/SdkTests/Go/TestData"

var uploadedTestFiles = &struct {
	mx    sync.Mutex
	items []string
}{}

func ConvertBytesToStrings(bytes []byte) []string {
	asString := string(bytes)
	if strings.Index(asString, "\r\n") >= 0 {
		return strings.Split(asString, "\r\n")
	} else {
		return strings.Split(asString, "\n")
	}
}

func CreateRandomGuid() string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func GetLocalPath(filename string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentFolder := filepath.Dir(currentFilePath)
	return filepath.Join(currentFolder, "testdata", filename)
}

func GetLocalFile(filePath string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentFolder := filepath.Dir(currentFilePath)
	return filepath.Join(currentFolder, "testdata", filePath)
}

func ReadFile(t *testing.T, path string) (result string) {
	data, err := ioutil.ReadFile(GetLocalFile(path))
	if err != nil {
		t.Error(err)
	}

	return string(data)
}

func OpenFile(t *testing.T, path string) (result *os.File) {
	data, err := os.Open(GetLocalFile(path))
	if err != nil {
		t.Error(err)
	}

	return data
}

func CreateTime(year int, month int, day int, hour int, min int, sec int) custom.TimeWithoutTZ {
	return custom.TimeWithoutTZ(time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC))
}

func ExtractTimeOnly(datetime custom.TimeWithoutTZ) string {
	return time.Time(datetime).Format("15:04:05")
}

func ExtractDateOnly(datetime custom.TimeWithoutTZ) string {
	return time.Time(datetime).Format("2006-01-02")
}

func ReadConfiguration(t *testing.T) (cfg *models.Configuration) {
	configuration, err := models.NewConfiguration(GetConfigFilePath())

	if err != nil {
		t.Error(err)
	}

	return configuration
}

func GetConfigFilePath() (path string) {
	_, filename, _, _ := runtime.Caller(0)
	configFilePath := filepath.Join(filepath.Dir(filename), "../config.json")
	_, fileErr := os.Stat(configFilePath)
	if os.IsNotExist(fileErr) {
		configFilePath = filepath.Join(filepath.Dir(filename), "../../config.json")
	}
	return configFilePath
}

func PrepareTest(t *testing.T, config *models.Configuration) (apiClient *api.APIClient, ctx context.Context) {
	client, err := api.NewAPIClient(config)

	if err != nil {
		t.Error(err)
	}

	context, err := client.NewContextWithToken(nil)

	if err != nil {
		t.Error(err)
	}

	return client, context
}

func UploadTestFileToStorage(t *testing.T, fileName string, path string) (*api.APIClient, context.Context) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	uploadedTestFiles.mx.Lock()
	defer uploadedTestFiles.mx.Unlock()
	UploadNextFileToStorage(t, ctx, client, fileName, path)
	uploadedTestFiles.items = append(uploadedTestFiles.items, path)
	return client, ctx
}

func DeleteTestFileFromStorage(t *testing.T, ctx context.Context, client *api.APIClient) {
	uploadedTestFiles.mx.Lock()
	defer uploadedTestFiles.mx.Unlock()
	for _, uploadedFile := range uploadedTestFiles.items {
		_, err := client.TasksApi.DeleteFile(ctx, &requests.DeleteFileOpts{
			Path: uploadedFile,
		})
		if err != nil {
			t.Error(err)
		}
	}
	uploadedTestFiles.items = []string{}
}

func UploadNextFileToStorage(t *testing.T, ctx context.Context, client *api.APIClient, fileName string, path string) {
	file := OpenFile(t, fileName)

	request := requests.UploadFileOpts{
		File: file,
		Path: path,
	}

	_, _, err := client.TasksApi.UploadFile(ctx, &request)
	if err != nil {
		t.Error(err)
	}
}

func TestDebugMode(t *testing.T) {
	config := ReadConfiguration(t)
	config.DebugMode = true
	client, ctx := PrepareTest(t, config)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	fileName := "NewProductDev.mpp"
	remoteFilePath := remoteBaseTestDataFolder + "/Storage/" + fileName
	localFilePath := GetLocalPath(fileName)
	file, fileErr := os.Open(localFilePath)
	if fileErr != nil {
		t.Error(fileErr)
	}

	request := requests.UploadFileOpts{
		Path: remoteFilePath,
		File: file,
	}

	_, _, err := client.TasksApi.UploadFile(ctx, &request)

	if err != nil {
		t.Error(err)
	}

	logData := buf.String()

	url := "PUT /v3.0/tasks/storage/file/Temp/SdkTests/Go/TestData/Storage/NewProductDev.mpp"
	success := " 200 OK"

	if !strings.Contains(logData, url) {
		t.Errorf("%s doesn't contain the string '%s'", logData, url)
	}

	if !strings.Contains(logData, success) {
		t.Errorf("%s doesn't contain the string '%s'", logData, success)
	}
}

func TestCoverage(t *testing.T) {
	_, currentFileName, _, _ := runtime.Caller(0)
	currentFolder := filepath.Dir(currentFileName)
	apiFileName := filepath.Join(currentFolder, "../api/tasks_api.go")
	api, err := ioutil.ReadFile(apiFileName)

	if err != nil {
		t.Error(err)
	}

	declareFuncRegex := regexp.MustCompile(`func \(a \*TasksApiService\) (\w+)`)
	allApiFuncNames := declareFuncRegex.FindAllSubmatch(api, -1)

	var exists = struct{}{}

	notTestedApiFuncNames := make(map[string]struct{})

	for _, name := range allApiFuncNames {
		notTestedApiFuncNames[string(name[1])] = exists
	}

	testFileNames, _ := filepath.Glob(filepath.Join(currentFolder, "*.go"))

	for _, testFileName := range testFileNames {
		tests, err := ioutil.ReadFile(testFileName)
		if err != nil {
			t.Error(err)
		}

		callFuncRegex := regexp.MustCompile(`\.TasksApi\.(\w+)`)
		testedlApiFuncNames := callFuncRegex.FindAllSubmatch(tests, -1)

		for _, name := range testedlApiFuncNames {
			delete(notTestedApiFuncNames, string(name[1]))
		}
	}

	if len(notTestedApiFuncNames) > 0 {
		report := "Not covered methods:\n"

		for name := range notTestedApiFuncNames {
			report = report + fmt.Sprintf("\t%s\n", name)
		}

		t.Error(report)
	}
}

func TestCreateTasksApi(t *testing.T) {
	config := ReadConfiguration(t)
	_, _, err := api.CreateTasksApi(config)

	if err != nil {
		t.Error(err)
	}
}
