/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tasks_api.go">
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

package api

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/requests"
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models"
)

// Linger please
var (
	_ context.Context
)

type TasksApiService service
/* 
TasksApiService Copy file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.CopyFileOpts - Request parameters:
     * @param "SrcPath" (string) - Source file path e.g. &#39;/folder/file.ext&#39;
     * @param "DestPath" (string) - Destination file path
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name
     * @param "VersionId" (optional.String) -  File version ID to copy


*/
func (a *TasksApiService) CopyFile(ctx context.Context, localVarOptionals *requests.CopyFileOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/file/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("destPath", fmt.Sprintf("%v", localVarOptionals.DestPath))
	if localVarOptionals != nil && localVarOptionals.SrcStorageName.IsSet() {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DestStorageName.IsSet() {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VersionId.IsSet() {
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()

    return localVarHttpResponse, err
}

/* 
TasksApiService Delete file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteFileOpts - Request parameters:
     * @param "Path" (string) - File path e.g. &#39;/folder/file.ext&#39;
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID to delete


*/
func (a *TasksApiService) DeleteFile(ctx context.Context, localVarOptionals *requests.DeleteFileOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VersionId.IsSet() {
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()

    return localVarHttpResponse, err
}

/* 
TasksApiService Download file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DownloadFileOpts - Request parameters:
     * @param "Path" (string) - File path e.g. &#39;/folder/file.ext&#39;
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID to download

@return []byte
*/
func (a *TasksApiService) DownloadFile(ctx context.Context, localVarOptionals *requests.DownloadFileOpts) ([]byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VersionId.IsSet() {
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Move file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.MoveFileOpts - Request parameters:
     * @param "SrcPath" (string) - Source file path e.g. &#39;/src.ext&#39;
     * @param "DestPath" (string) - Destination file path e.g. &#39;/dest.ext&#39;
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name
     * @param "VersionId" (optional.String) -  File version ID to move


*/
func (a *TasksApiService) MoveFile(ctx context.Context, localVarOptionals *requests.MoveFileOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/file/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("destPath", fmt.Sprintf("%v", localVarOptionals.DestPath))
	if localVarOptionals != nil && localVarOptionals.SrcStorageName.IsSet() {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DestStorageName.IsSet() {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VersionId.IsSet() {
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()

    return localVarHttpResponse, err
}

/* 
TasksApiService Upload file
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.UploadFileOpts - Request parameters:
     * @param "Path" (string) - Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.             
     * @param "File" (*os.File) - File to upload
     * @param "StorageName" (optional.String) -  Storage name

@return models.FilesUploadResult
*/
func (a *TasksApiService) UploadFile(ctx context.Context, localVarOptionals *requests.UploadFileOpts) (models.FilesUploadResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.FilesUploadResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFile := localVarOptionals.File
	if localVarFile != nil {
        fbs, _ := ioutil.ReadAll(localVarFile)
        localVarFile.Close()
        localVarFormParams = append(localVarFormParams, NewFileFormParamContainer(localVarFile.Name(), fbs))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Copy folder
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.CopyFolderOpts - Request parameters:
     * @param "SrcPath" (string) - Source folder path e.g. &#39;/src&#39;
     * @param "DestPath" (string) - Destination folder path e.g. &#39;/dst&#39;
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name


*/
func (a *TasksApiService) CopyFolder(ctx context.Context, localVarOptionals *requests.CopyFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/folder/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("destPath", fmt.Sprintf("%v", localVarOptionals.DestPath))
	if localVarOptionals != nil && localVarOptionals.SrcStorageName.IsSet() {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DestStorageName.IsSet() {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()

    return localVarHttpResponse, err
}

/* 
TasksApiService Create the folder
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.CreateFolderOpts - Request parameters:
     * @param "Path" (string) - Folder path to create e.g. &#39;folder_1/folder_2/&#39;
     * @param "StorageName" (optional.String) -  Storage name


*/
func (a *TasksApiService) CreateFolder(ctx context.Context, localVarOptionals *requests.CreateFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()

    return localVarHttpResponse, err
}

/* 
TasksApiService Delete folder
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteFolderOpts - Request parameters:
     * @param "Path" (string) - Folder path e.g. &#39;/folder&#39;
     * @param "StorageName" (optional.String) -  Storage name
     * @param "Recursive" (optional.Bool) -  Enable to delete folders, subfolders and files


*/
func (a *TasksApiService) DeleteFolder(ctx context.Context, localVarOptionals *requests.DeleteFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Recursive.IsSet() {
		localVarQueryParams.Add("recursive", parameterToString(localVarOptionals.Recursive.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()

    return localVarHttpResponse, err
}

/* 
TasksApiService Get all files and folders within a folder
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetFilesListOpts - Request parameters:
     * @param "Path" (string) - Folder path e.g. &#39;/folder&#39;
     * @param "StorageName" (optional.String) -  Storage name

@return models.FilesList
*/
func (a *TasksApiService) GetFilesList(ctx context.Context, localVarOptionals *requests.GetFilesListOpts) (models.FilesList, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.FilesList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Move folder
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.MoveFolderOpts - Request parameters:
     * @param "SrcPath" (string) - Folder path to move e.g. &#39;/folder&#39;
     * @param "DestPath" (string) - Destination folder path to move to e.g &#39;/dst&#39;
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name


*/
func (a *TasksApiService) MoveFolder(ctx context.Context, localVarOptionals *requests.MoveFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/folder/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("destPath", fmt.Sprintf("%v", localVarOptionals.DestPath))
	if localVarOptionals != nil && localVarOptionals.SrcStorageName.IsSet() {
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DestStorageName.IsSet() {
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarHttpResponse, err
        }

        return localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()

    return localVarHttpResponse, err
}
/* 
TasksApiService Get disc usage
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetDiscUsageOpts - Request parameters:
     * @param "StorageName" (optional.String) -  Storage name

@return models.DiscUsage
*/
func (a *TasksApiService) GetDiscUsage(ctx context.Context, localVarOptionals *requests.GetDiscUsageOpts) (models.DiscUsage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.DiscUsage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/disc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get file versions
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetFileVersionsOpts - Request parameters:
     * @param "Path" (string) - File path e.g. &#39;/file.ext&#39;
     * @param "StorageName" (optional.String) -  Storage name

@return models.FileVersions
*/
func (a *TasksApiService) GetFileVersions(ctx context.Context, localVarOptionals *requests.GetFileVersionsOpts) (models.FileVersions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.FileVersions
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/version/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Check if file or folder exists
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.ObjectExistsOpts - Request parameters:
     * @param "Path" (string) - File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39;
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID

@return models.ObjectExist
*/
func (a *TasksApiService) ObjectExists(ctx context.Context, localVarOptionals *requests.ObjectExistsOpts) (models.ObjectExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.ObjectExist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/exist/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.StorageName.IsSet() {
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VersionId.IsSet() {
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Check if storage exists
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.StorageExistsOpts - Request parameters:
     * @param "StorageName" (string) - Storage name

@return models.StorageExist
*/
func (a *TasksApiService) StorageExists(ctx context.Context, localVarOptionals *requests.StorageExistsOpts) (models.StorageExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.StorageExist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/storage/{storageName}/exist"
	localVarPath = strings.Replace(localVarPath, "{"+"storageName"+"}", fmt.Sprintf("%v", localVarOptionals.StorageName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Deletes a project assignment with all references to it.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteAssignmentOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "AssignmentUid" (int32) - assignment Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteAssignment(ctx context.Context, localVarOptionals *requests.DeleteAssignmentOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/assignments/{assignmentUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"assignmentUid"+"}", fmt.Sprintf("%v", localVarOptionals.AssignmentUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Read project assignment with the specified Uid.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetAssignmentOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "AssignmentUid" (int32) - Assignment Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AssignmentResponse
*/
func (a *TasksApiService) GetAssignment(ctx context.Context, localVarOptionals *requests.GetAssignmentOpts) (models.AssignmentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.AssignmentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/assignments/{assignmentUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"assignmentUid"+"}", fmt.Sprintf("%v", localVarOptionals.AssignmentUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get timescaled data for an assignment with the specified Uid.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetAssignmentTimephasedDataOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "AssignmentUid" (int32) - Uid of assignment to get timephased data for.
     * @param "Type_" (optional.String) -  Type of timephased data to get.
     * @param "StartDate" (*custom.TimeWithoutTZ) -  Start date.
     * @param "EndDate" (*custom.TimeWithoutTZ) -  End date.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.

@return models.TimephasedDataResponse
*/
func (a *TasksApiService) GetAssignmentTimephasedData(ctx context.Context, localVarOptionals *requests.GetAssignmentTimephasedDataOpts) (models.TimephasedDataResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.TimephasedDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/assignments/{assignmentUid}/timeScaleData"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"assignmentUid"+"}", fmt.Sprintf("%v", localVarOptionals.AssignmentUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(localVarOptionals.StartDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(localVarOptionals.EndDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get project&#39;s assignment items.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetAssignmentsOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AssignmentItemsResponse
*/
func (a *TasksApiService) GetAssignments(ctx context.Context, localVarOptionals *requests.GetAssignmentsOpts) (models.AssignmentItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.AssignmentItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Adds a new assignment to a project and returns assignment item in a response.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostAssignmentOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - The unique id of the task to be assigned.
     * @param "ResourceUid" (int32) - The unique id of the resource to be assigned.
     * @param "Units" (optional.Float64) -  The units for the new assignment. If not specified, &#39;cost&#39; value is used.
     * @param "Cost" (optional.Float32) -  The cost for a new assignment. If not specified, default value is used.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AssignmentItemResponse
*/
func (a *TasksApiService) PostAssignment(ctx context.Context, localVarOptionals *requests.PostAssignmentOpts) (models.AssignmentItemResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.AssignmentItemResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("taskUid", fmt.Sprintf("%v", localVarOptionals.TaskUid))
	localVarQueryParams.Add("resourceUid", fmt.Sprintf("%v", localVarOptionals.ResourceUid))
	if localVarOptionals != nil && localVarOptionals.Units.IsSet() {
		localVarQueryParams.Add("units", parameterToString(localVarOptionals.Units.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cost.IsSet() {
		localVarQueryParams.Add("cost", parameterToString(localVarOptionals.Cost.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Updates resources assignment with the specified Uid.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutAssignmentOpts - Request parameters:
     * @param "Name" (string) - The file name
     * @param "AssignmentUid" (int32) - Assignment UID
     * @param "Assignment" (ResourceAssignment) - Assignment DTO
     * @param "Mode" (optional.String) -  Project&#39;s calculation mode
     * @param "Recalculate" (optional.Bool) -  boolean value
     * @param "Storage" (optional.String) -  The document storage
     * @param "Folder" (optional.String) -  The document storage
     * @param "FileName" (optional.String) -  The filename to save Changes

@return models.AssignmentResponse
*/
func (a *TasksApiService) PutAssignment(ctx context.Context, localVarOptionals *requests.PutAssignmentOpts) (models.AssignmentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AssignmentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/assignments/{assignmentUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"assignmentUid"+"}", fmt.Sprintf("%v", localVarOptionals.AssignmentUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Mode.IsSet() {
		localVarQueryParams.Add("mode", parameterToString(localVarOptionals.Mode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Recalculate.IsSet() {
		localVarQueryParams.Add("recalculate", parameterToString(localVarOptionals.Recalculate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Assignment
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Deletes a project calendar
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteCalendarOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - Calendar&#39;s Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteCalendar(ctx context.Context, localVarOptionals *requests.DeleteCalendarOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars/{calendarUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarUid"+"}", fmt.Sprintf("%v", localVarOptionals.CalendarUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Deletes calendar exception from calendar exceptions collection.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteCalendarExceptionOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - Calendar Uid
     * @param "Index" (int32) - Index of the calendar exception. See CalendarException.Index property.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteCalendarException(ctx context.Context, localVarOptionals *requests.DeleteCalendarExceptionOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars/{calendarUid}/calendarExceptions/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarUid"+"}", fmt.Sprintf("%v", localVarOptionals.CalendarUid), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project&#39;s calendar with the specified Uid.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetCalendarOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - Calendar&#39;s Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.CalendarResponse
*/
func (a *TasksApiService) GetCalendar(ctx context.Context, localVarOptionals *requests.GetCalendarOpts) (models.CalendarResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.CalendarResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars/{calendarUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarUid"+"}", fmt.Sprintf("%v", localVarOptionals.CalendarUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a list of calendar&#39;s exceptions.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetCalendarExceptionsOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - Calendar&#39;s Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.CalendarExceptionsResponse
*/
func (a *TasksApiService) GetCalendarExceptions(ctx context.Context, localVarOptionals *requests.GetCalendarExceptionsOpts) (models.CalendarExceptionsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.CalendarExceptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars/{calendarUid}/calendarExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarUid"+"}", fmt.Sprintf("%v", localVarOptionals.CalendarUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Gets the collection of work weeks of the specified calendar.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetCalendarWorkWeeksOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - Calendar Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.CalendarWorkWeeksResponse
*/
func (a *TasksApiService) GetCalendarWorkWeeks(ctx context.Context, localVarOptionals *requests.GetCalendarWorkWeeksOpts) (models.CalendarWorkWeeksResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.CalendarWorkWeeksResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars/{calendarUid}/workWeeks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarUid"+"}", fmt.Sprintf("%v", localVarOptionals.CalendarUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Read project calendar items.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetCalendarsOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.CalendarItemsResponse
*/
func (a *TasksApiService) GetCalendars(ctx context.Context, localVarOptionals *requests.GetCalendarsOpts) (models.CalendarItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.CalendarItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Adds a new calendar to project file.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostCalendarOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Calendar" (Calendar) - Calendar DTO
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.CalendarItemResponse
*/
func (a *TasksApiService) PostCalendar(ctx context.Context, localVarOptionals *requests.PostCalendarOpts) (models.CalendarItemResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.CalendarItemResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Calendar
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Adds a new calendar exception to a calendar.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostCalendarExceptionOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - Calendar&#39;s Uid
     * @param "CalendarException" (CalendarException) - CalendarException DTO
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) PostCalendarException(ctx context.Context, localVarOptionals *requests.PostCalendarExceptionOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars/{calendarUid}/calendarExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarUid"+"}", fmt.Sprintf("%v", localVarOptionals.CalendarUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.CalendarException
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Edits an existing project calendar.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutCalendarOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - The Uid of an existing calendar to edit.
     * @param "Calendar" (Calendar) - Modified calendar DTO
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutCalendar(ctx context.Context, localVarOptionals *requests.PutCalendarOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("calendarUid", fmt.Sprintf("%v", localVarOptionals.CalendarUid))
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Calendar
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Updates calendar exception.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutCalendarExceptionOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "CalendarUid" (int32) - Calendar Uid
     * @param "Index" (int32) - Index of the calendar exception. See CalendarException.Index property.
     * @param "CalendarException" (CalendarException) - CalendarException DTO
     * @param "FileName" (optional.String) -  The name of the project document to save changes to. If this parameter              is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutCalendarException(ctx context.Context, localVarOptionals *requests.PutCalendarExceptionOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/calendars/{calendarUid}/calendarExceptions/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarUid"+"}", fmt.Sprintf("%v", localVarOptionals.CalendarUid), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.CalendarException
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Returns the list of the tasks which must be completed on time for a project to finish on schedule. Each task of the project is represented as a task item here, which is light-weighted task representation of the project task..
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetCriticalPathOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.TaskItemsResponse
*/
func (a *TasksApiService) GetCriticalPath(ctx context.Context, localVarOptionals *requests.GetCriticalPathOpts) (models.TaskItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/criticalPath"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Returns page count for the project to be rendered using given Timescale and PresentationFormat  or specified PageSize.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetPageCountOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "PageSize" (optional.String) -  PageSize to get page count for.
     * @param "PresentationFormat" (optional.String) -  PresentationFormat to get page count for.
     * @param "Timescale" (optional.String) -  Timescale to get page count for.
     * @param "StartDate" (*custom.TimeWithoutTZ) -  Start date to get page count for.
     * @param "EndDate" (*custom.TimeWithoutTZ) -  End date to get page count for.
     * @param "Folder" (optional.String) -  The document folder
     * @param "Storage" (optional.String) -  The document storage.

@return models.PageCountResponse
*/
func (a *TasksApiService) GetPageCount(ctx context.Context, localVarOptionals *requests.GetPageCountOpts) (models.PageCountResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.PageCountResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/pagecount"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PresentationFormat.IsSet() {
		localVarQueryParams.Add("presentationFormat", parameterToString(localVarOptionals.PresentationFormat.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timescale.IsSet() {
		localVarQueryParams.Add("timescale", parameterToString(localVarOptionals.Timescale.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(localVarOptionals.StartDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(localVarOptionals.EndDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get Uids of projects contained in the file.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetProjectIdsOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.ProjectIdsResponse
*/
func (a *TasksApiService) GetProjectIds(ctx context.Context, localVarOptionals *requests.GetProjectIdsOpts) (models.ProjectIdsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.ProjectIdsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/projectids"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project document.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTaskDocumentOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return []byte
*/
func (a *TasksApiService) GetTaskDocument(ctx context.Context, localVarOptionals *requests.GetTaskDocumentOpts) ([]byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project document in the specified format
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTaskDocumentWithFormatOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Format" (string) - The format of the resulting file.
     * @param "ReturnAsZipArchive" (optional.Bool) -  If parameter is true, HTML resources are included as separate files and returned along with the resulting html file as a zip package.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return []byte
*/
func (a *TasksApiService) GetTaskDocumentWithFormat(ctx context.Context, localVarOptionals *requests.GetTaskDocumentWithFormatOpts) ([]byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/format"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("format", fmt.Sprintf("%v", localVarOptionals.Format))
	if localVarOptionals != nil && localVarOptionals.ReturnAsZipArchive.IsSet() {
		localVarQueryParams.Add("returnAsZipArchive", parameterToString(localVarOptionals.ReturnAsZipArchive.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Imports project from database with the specified connection string and saves it to specified file with the specified format.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutImportProjectFromDbOpts - Request parameters:
     * @param "DatabaseType" (string) - The type of source database. The supported values are (Msp, Mpd, Primavera).
     * @param "ConnectionString" (string) - The connection string to the source database.
     * @param "ProjectUid" (string) - Uid of the project to import.
     * @param "Filename" (string) - The name of the resulting file.
     * @param "Format" (optional.String) -  Format of the resulting file.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "DatabaseSchema" (optional.String) -  Schema of Microsoft project database (if applicable)

@return models.AsposeResponse
*/
func (a *TasksApiService) PutImportProjectFromDb(ctx context.Context, localVarOptionals *requests.PutImportProjectFromDbOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/importfromdb"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("databaseType", fmt.Sprintf("%v", localVarOptionals.DatabaseType))
	localVarQueryParams.Add("projectUid", fmt.Sprintf("%v", localVarOptionals.ProjectUid))
	localVarQueryParams.Add("filename", fmt.Sprintf("%v", localVarOptionals.Filename))
	if localVarOptionals != nil && localVarOptionals.Format.IsSet() {
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DatabaseSchema.IsSet() {
		localVarQueryParams.Add("databaseSchema", parameterToString(localVarOptionals.DatabaseSchema.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.ConnectionString
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Imports project from primavera db formats (Primavera SQLite .db or Primavera xml) and saves it to specified file with the specified format.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutImportProjectFromFileOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "ProjectUid" (string) - Uid of the project to import.
     * @param "Filename" (string) - The name of the resulting file.
     * @param "FileType" (optional.String) -  The type of file to import. The supported values are (PrimaveraSqliteDb, PrimaveraXml).
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "OutputFileFormat" (optional.String) -  The format of the resulting file.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutImportProjectFromFile(ctx context.Context, localVarOptionals *requests.PutImportProjectFromFileOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/import"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("projectUid", fmt.Sprintf("%v", localVarOptionals.ProjectUid))
	localVarQueryParams.Add("filename", fmt.Sprintf("%v", localVarOptionals.Filename))
	if localVarOptionals != nil && localVarOptionals.FileType.IsSet() {
		localVarQueryParams.Add("fileType", parameterToString(localVarOptionals.FileType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OutputFileFormat.IsSet() {
		localVarQueryParams.Add("outputFileFormat", parameterToString(localVarOptionals.OutputFileFormat.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Imports project from Project Online and saves it to specified file.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutImportProjectFromProjectOnlineOpts - Request parameters:
     * @param "Name" (string) - The name of the resulting file.
     * @param "Guid" (string) - Guid of the project to import.
     * @param "SiteUrl" (string) - The URL of PWA (Project Web Access) API of Project Online
     * @param "UserName" (optional.String) -  The user name for the sharepoint site.
     * @param "Format" (optional.String) -  Format of the resulting file.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "XProjectOnlineToken" (optional.String) -  Authorization token for the SharePoint. For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly
     * @param "XSharepointPassword" (optional.String) -  The password for the SharePoint site.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutImportProjectFromProjectOnline(ctx context.Context, localVarOptionals *requests.PutImportProjectFromProjectOnlineOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/importfromprojectonline"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("siteUrl", fmt.Sprintf("%v", localVarOptionals.SiteUrl))
	if localVarOptionals != nil && localVarOptionals.UserName.IsSet() {
		localVarQueryParams.Add("userName", parameterToString(localVarOptionals.UserName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Format.IsSet() {
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XProjectOnlineToken.IsSet() {
		localVarHeaderParams["x-project-online-token"] = parameterToString(localVarOptionals.XProjectOnlineToken.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XSharepointPassword.IsSet() {
		localVarHeaderParams["x-sharepoint-password"] = parameterToString(localVarOptionals.XSharepointPassword.Value(), "")
	}
	// body params
	localVarPostBody = &localVarOptionals.Guid
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Get a collection of a project&#39;s document properties.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetDocumentPropertiesOpts - Request parameters:
     * @param "Name" (string) - The document name.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.DocumentPropertiesResponse
*/
func (a *TasksApiService) GetDocumentProperties(ctx context.Context, localVarOptionals *requests.GetDocumentPropertiesOpts) (models.DocumentPropertiesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.DocumentPropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/documentproperties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a document property by name.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetDocumentPropertyOpts - Request parameters:
     * @param "Name" (string) - The document name.
     * @param "PropertyName" (string) - The property name.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.DocumentPropertyResponse
*/
func (a *TasksApiService) GetDocumentProperty(ctx context.Context, localVarOptionals *requests.GetDocumentPropertyOpts) (models.DocumentPropertyResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", localVarOptionals.PropertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Set/create document property.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostDocumentPropertyOpts - Request parameters:
     * @param "Name" (string) - The document name.
     * @param "PropertyName" (string) - The property name.
     * @param "Property" (DocumentProperty) - DocumentProperty with new property value.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Filename" (optional.String) -  Name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document.

@return models.DocumentPropertyResponse
*/
func (a *TasksApiService) PostDocumentProperty(ctx context.Context, localVarOptionals *requests.PostDocumentPropertyOpts) (models.DocumentPropertyResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", localVarOptionals.PropertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filename.IsSet() {
		localVarQueryParams.Add("filename", parameterToString(localVarOptionals.Filename.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Property
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Set/create document property.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutDocumentPropertyOpts - Request parameters:
     * @param "Name" (string) - The document name.
     * @param "PropertyName" (string) - The property name.
     * @param "Property" (DocumentProperty) - DocumentProperty with new property value.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Filename" (optional.String) -  Name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document.

@return models.DocumentPropertyResponse
*/
func (a *TasksApiService) PutDocumentProperty(ctx context.Context, localVarOptionals *requests.PutDocumentPropertyOpts) (models.DocumentPropertyResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", localVarOptionals.PropertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filename.IsSet() {
		localVarQueryParams.Add("filename", parameterToString(localVarOptionals.Filename.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Property
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Delete a project extended attribute.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteExtendedAttributeByIndexOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Index" (int32) - Index (See ExtendedAttributeItem.Index property) or FieldId of the extended attribute.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteExtendedAttributeByIndex(ctx context.Context, localVarOptionals *requests.DeleteExtendedAttributeByIndexOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/extendedAttributes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project extended attribute definition with the specified index.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetExtendedAttributeByIndexOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Index" (int32) - Index (See ExtendedAttributeItem.Index property) or FieldId of the extended attribute.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.ExtendedAttributeResponse
*/
func (a *TasksApiService) GetExtendedAttributeByIndex(ctx context.Context, localVarOptionals *requests.GetExtendedAttributeByIndexOpts) (models.ExtendedAttributeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.ExtendedAttributeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/extendedAttributes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project&#39;s extended attributes.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetExtendedAttributesOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.ExtendedAttributeItemsResponse
*/
func (a *TasksApiService) GetExtendedAttributes(ctx context.Context, localVarOptionals *requests.GetExtendedAttributesOpts) (models.ExtendedAttributeItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.ExtendedAttributeItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/extendedAttributes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Add a new extended attribute definition to a project or  updates existing attribute definition (with the same FieldId).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutExtendedAttributeOpts - Request parameters:
     * @param "ExtendedAttribute" (ExtendedAttributeDefinition) - Definition of the extended attribute to add.
     * @param "Name" (string) - The name of the file.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.ExtendedAttributeItemResponse
*/
func (a *TasksApiService) PutExtendedAttribute(ctx context.Context, localVarOptionals *requests.PutExtendedAttributeOpts) (models.ExtendedAttributeItemResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.ExtendedAttributeItemResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/extendedAttributes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.ExtendedAttribute
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Deletes a project outline code
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteOutlineCodeByIndexOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Index" (int32) - Index of the outline code. See OutlineCodeItem.Index property.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteOutlineCodeByIndex(ctx context.Context, localVarOptionals *requests.DeleteOutlineCodeByIndexOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/outlineCodes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get outline code by index.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetOutlineCodeByIndexOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Index" (int32) - Index of the outline code. See OutlineCodeItem.Index property.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.OutlineCodeResponse
*/
func (a *TasksApiService) GetOutlineCodeByIndex(ctx context.Context, localVarOptionals *requests.GetOutlineCodeByIndexOpts) (models.OutlineCodeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.OutlineCodeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/outlineCodes/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project&#39;s outline codes.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetOutlineCodesOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.OutlineCodeItemsResponse
*/
func (a *TasksApiService) GetOutlineCodes(ctx context.Context, localVarOptionals *requests.GetOutlineCodesOpts) (models.OutlineCodeItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.OutlineCodeItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/outlineCodes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Creates new project in Project Server\\Project Online instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.CreateNewProjectOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "SiteUrl" (string) - The URL of PWA (Project Web Access) API of Project Online.
     * @param "UserName" (optional.String) -  The user name for the sharepoint site.
     * @param "SaveOptions" (optional.Interface of ProjectServerSaveOptionsDto) -  Dispensable save options for Project Server\\Project Online.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "XProjectOnlineToken" (optional.String) -  Authorization token for the SharePoint. For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly
     * @param "XSharepointPassword" (optional.String) -  The password for the SharePoint site.

@return models.AsposeResponse
*/
func (a *TasksApiService) CreateNewProject(ctx context.Context, localVarOptionals *requests.CreateNewProjectOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/projectOnline/{siteUrl}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteUrl"+"}", fmt.Sprintf("%v", localVarOptionals.SiteUrl), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.UserName.IsSet() {
		localVarQueryParams.Add("userName", parameterToString(localVarOptionals.UserName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XProjectOnlineToken.IsSet() {
		localVarHeaderParams["x-project-online-token"] = parameterToString(localVarOptionals.XProjectOnlineToken.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XSharepointPassword.IsSet() {
		localVarHeaderParams["x-sharepoint-password"] = parameterToString(localVarOptionals.XSharepointPassword.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.SaveOptions.IsSet() {
		
		localVarOptionalSaveOptions, localVarOptionalSaveOptionsok := localVarOptionals.SaveOptions.Value().(models.ProjectServerSaveOptionsDto)
		if !localVarOptionalSaveOptionsok {
				return localVarReturnValue, nil, reportError("saveOptions should be ProjectServerSaveOptionsDto")
		}
		localVarPostBody = &localVarOptionalSaveOptions
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Gets the list of published projects in the current Project Online account.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetProjectListOpts - Request parameters:
     * @param "SiteUrl" (string) - The URL of PWA (Project Web Access) API of Project Online.
     * @param "UserName" (optional.String) -  The user name for the sharepoint site.
     * @param "XProjectOnlineToken" (optional.String) -  Authorization token for the SharePoint. For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly
     * @param "XSharepointPassword" (optional.String) -  The password for the SharePoint site.

@return models.ProjectListResponse
*/
func (a *TasksApiService) GetProjectList(ctx context.Context, localVarOptionals *requests.GetProjectListOpts) (models.ProjectListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.ProjectListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/projectOnline/projectList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("siteUrl", fmt.Sprintf("%v", localVarOptionals.SiteUrl))
	if localVarOptionals != nil && localVarOptionals.UserName.IsSet() {
		localVarQueryParams.Add("userName", parameterToString(localVarOptionals.UserName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XProjectOnlineToken.IsSet() {
		localVarHeaderParams["x-project-online-token"] = parameterToString(localVarOptionals.XProjectOnlineToken.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XSharepointPassword.IsSet() {
		localVarHeaderParams["x-sharepoint-password"] = parameterToString(localVarOptionals.XSharepointPassword.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Updates existing project in Project Server\\Project Online instance. The existing project will be overwritten.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.UpdateProjectOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "SiteUrl" (string) - The URL of PWA (Project Web Access) API of Project Online.
     * @param "UserName" (optional.String) -  The user name for the sharepoint site.
     * @param "SaveOptions" (optional.Interface of ProjectServerSaveOptionsDto) -  Dispensable save options for Project Server\\Project Online.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "XProjectOnlineToken" (optional.String) -  Authorization token for the SharePoint. For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly
     * @param "XSharepointPassword" (optional.String) -  The password for the SharePoint site.

@return models.AsposeResponse
*/
func (a *TasksApiService) UpdateProject(ctx context.Context, localVarOptionals *requests.UpdateProjectOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/projectOnline/{siteUrl}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteUrl"+"}", fmt.Sprintf("%v", localVarOptionals.SiteUrl), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.UserName.IsSet() {
		localVarQueryParams.Add("userName", parameterToString(localVarOptionals.UserName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XProjectOnlineToken.IsSet() {
		localVarHeaderParams["x-project-online-token"] = parameterToString(localVarOptionals.XProjectOnlineToken.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XSharepointPassword.IsSet() {
		localVarHeaderParams["x-sharepoint-password"] = parameterToString(localVarOptionals.XSharepointPassword.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.SaveOptions.IsSet() {
		
		localVarOptionalSaveOptions, localVarOptionalSaveOptionsok := localVarOptionals.SaveOptions.Value().(models.ProjectServerSaveOptionsDto)
		if !localVarOptionalSaveOptionsok {
				return localVarReturnValue, nil, reportError("saveOptions should be ProjectServerSaveOptionsDto")
		}
		localVarPostBody = &localVarOptionalSaveOptions
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Reschedules all project tasks ids, outline levels, start/finish dates, sets early/late dates, calculates slacks, work and cost fields.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutRecalculateProjectOpts - Request parameters:
     * @param "Name" (string) - The filename
     * @param "Mode" (optional.String) -  Project&#39;s calculation mode
     * @param "Validate" (optional.Bool) -  If true the validation of recalculation will be performed. What data is validated:     At the moment only basic validation of task and task link date ranges is implemented.     Task&#39;s date ranges (e.g. ActualStart - ActualFinish, EarlyStart - EarlyFinish, etc.) as well as Task Links dates will be checked against the date criteria that start date is less or equal than finish date.
     * @param "FileName" (optional.String) -  The filename to save the changes
     * @param "Storage" (optional.String) -  The document storage
     * @param "Folder" (optional.String) -  The document folder

@return models.ProjectRecalculateResponse
*/
func (a *TasksApiService) PutRecalculateProject(ctx context.Context, localVarOptionals *requests.PutRecalculateProjectOpts) (models.ProjectRecalculateResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.ProjectRecalculateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/recalculate/project"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Mode.IsSet() {
		localVarQueryParams.Add("mode", parameterToString(localVarOptionals.Mode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Validate.IsSet() {
		localVarQueryParams.Add("validate", parameterToString(localVarOptionals.Validate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Recalculate project resource fields
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutRecalculateProjectResourceFieldsOpts - Request parameters:
     * @param "Name" (string) - The name of the file
     * @param "Storage" (optional.String) -  The document storage
     * @param "Folder" (optional.String) -  The document folder
     * @param "FileName" (optional.String) -  The document fileName

@return models.AsposeResponse
*/
func (a *TasksApiService) PutRecalculateProjectResourceFields(ctx context.Context, localVarOptionals *requests.PutRecalculateProjectResourceFieldsOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/recalculate/resourceFields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Recalculate project uncomplete work
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutRecalculateProjectUncompleteWorkToStartAfterOpts - Request parameters:
     * @param "Name" (string) - The file name
     * @param "After" (time.Time) - DateTime (from System lib)
     * @param "Storage" (optional.String) -  The document storage 
     * @param "Folder" (optional.String) -  The document folder
     * @param "FileName" (optional.String) -  The filename to save the changes

@return models.AsposeResponse
*/
func (a *TasksApiService) PutRecalculateProjectUncompleteWorkToStartAfter(ctx context.Context, localVarOptionals *requests.PutRecalculateProjectUncompleteWorkToStartAfterOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/recalculate/uncompleteWorkToStartAfter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.After
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Recalculate project work as complete 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutRecalculateProjectWorkAsCompleteOpts - Request parameters:
     * @param "Name" (string) - The filename
     * @param "CompleteThrough" (time.Time) - DateTime (type from System lib)
     * @param "SetZeroOrHundredPercentCompleteOnly" (optional.Bool) -  boolean value
     * @param "Storage" (optional.String) -  The document storage
     * @param "Folder" (optional.String) -  The document folder
     * @param "FileName" (optional.String) -  The filename to save the changes

@return models.AsposeResponse
*/
func (a *TasksApiService) PutRecalculateProjectWorkAsComplete(ctx context.Context, localVarOptionals *requests.PutRecalculateProjectWorkAsCompleteOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/recalculate/projectWorkAsComplete"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.SetZeroOrHundredPercentCompleteOnly.IsSet() {
		localVarQueryParams.Add("setZeroOrHundredPercentCompleteOnly", parameterToString(localVarOptionals.SetZeroOrHundredPercentCompleteOnly.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.CompleteThrough
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Returns created report in PDF format.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetReportPdfOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Type_" (string) - A type of the project&#39;s graphical report.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return []byte
*/
func (a *TasksApiService) GetReportPdf(ctx context.Context, localVarOptionals *requests.GetReportPdfOpts) ([]byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/report"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("type", fmt.Sprintf("%v", localVarOptionals.Type_))
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Deletes a project resource with all references to it
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteResourceOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "ResourceUid" (int32) - Resource Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteResource(ctx context.Context, localVarOptionals *requests.DeleteResourceOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/resources/{resourceUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUid"+"}", fmt.Sprintf("%v", localVarOptionals.ResourceUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get project resource.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetResourceOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "ResourceUid" (int32) - Resource Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.ResourceResponse
*/
func (a *TasksApiService) GetResource(ctx context.Context, localVarOptionals *requests.GetResourceOpts) (models.ResourceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.ResourceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/resources/{resourceUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUid"+"}", fmt.Sprintf("%v", localVarOptionals.ResourceUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get resource&#39;s assignments.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetResourceAssignmentsOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "ResourceUid" (int32) - Resource Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AssignmentsResponse
*/
func (a *TasksApiService) GetResourceAssignments(ctx context.Context, localVarOptionals *requests.GetResourceAssignmentsOpts) (models.AssignmentsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.AssignmentsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/resources/{resourceUid}/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUid"+"}", fmt.Sprintf("%v", localVarOptionals.ResourceUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get timescaled data for a resource with the specified Uid.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetResourceTimephasedDataOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "ResourceUid" (int32) - Uid of resource to get timephased data for.
     * @param "Type_" (optional.String) -  Type of timephased data to get.
     * @param "StartDate" (*custom.TimeWithoutTZ) -  Start date.
     * @param "EndDate" (*custom.TimeWithoutTZ) -  End date.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.

@return models.TimephasedDataResponse
*/
func (a *TasksApiService) GetResourceTimephasedData(ctx context.Context, localVarOptionals *requests.GetResourceTimephasedDataOpts) (models.TimephasedDataResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.TimephasedDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/resources/{resourceUid}/timeScaleData"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUid"+"}", fmt.Sprintf("%v", localVarOptionals.ResourceUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(localVarOptionals.StartDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(localVarOptionals.EndDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project&#39;s resources.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetResourcesOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.ResourceItemsResponse
*/
func (a *TasksApiService) GetResources(ctx context.Context, localVarOptionals *requests.GetResourcesOpts) (models.ResourceItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.ResourceItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/resources"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Add a new resource to a project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostResourceOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "ResourceName" (optional.String) -  The name of the new resource. The default value is an empty string
     * @param "BeforeResourceId" (optional.Int32) -  The id of the resource to insert the new resource before. The default value is the id of the last resource in a project file.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.ResourceItemResponse
*/
func (a *TasksApiService) PostResource(ctx context.Context, localVarOptionals *requests.PostResourceOpts) (models.ResourceItemResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.ResourceItemResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/resources"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.ResourceName.IsSet() {
		localVarQueryParams.Add("resourceName", parameterToString(localVarOptionals.ResourceName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeforeResourceId.IsSet() {
		localVarQueryParams.Add("beforeResourceId", parameterToString(localVarOptionals.BeforeResourceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Updates resource with the specified Uid
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutResourceOpts - Request parameters:
     * @param "Name" (string) - The file name
     * @param "ResourceUid" (int32) - The Uid of a resource
     * @param "Resource" (Resource) - The representation of the modified resource
     * @param "Mode" (optional.String) -  The calculation mode of a project
     * @param "Recalculate" (optional.Bool) -  Specifies whether the project&#39;s recalculation should be performed
     * @param "Storage" (optional.String) -  The document storage
     * @param "Folder" (optional.String) -  The document storage
     * @param "FileName" (optional.String) -  The filename to save Changes

@return models.ResourceResponse
*/
func (a *TasksApiService) PutResource(ctx context.Context, localVarOptionals *requests.PutResourceOpts) (models.ResourceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.ResourceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/resources/{resourceUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUid"+"}", fmt.Sprintf("%v", localVarOptionals.ResourceUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Mode.IsSet() {
		localVarQueryParams.Add("mode", parameterToString(localVarOptionals.Mode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Recalculate.IsSet() {
		localVarQueryParams.Add("recalculate", parameterToString(localVarOptionals.Recalculate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Resource
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Performs a risk analysys using Monte Carlo simulation and creates a risk analysis report.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetRiskAnalysisReportOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - The uid of the task for which this risk will be applied in Monte Carlo simulation.
     * @param "DistributionType" (optional.String) -  Gets or sets the probability distribution used in Monte Carlo simulation. The default value is ProbabilityDistributionType.Normal.
     * @param "Optimistic" (optional.Int32) -  The percentage of the most likely task duration which can happen in the best possible project scenario. The default value is 75, which means that if the estimated specified task duration is 4 days then the optimistic duration will be 3 days.
     * @param "Pessimistic" (optional.Int32) -  The percentage of the most likely task duration which can happen in the worst possible project scenario. The default value is 125, which means that if the estimated specified task duration is 4 days then the pessimistic duration will be 5 days.
     * @param "ConfidenceLevel" (optional.String) -  Gets or sets the confidence level that correspond to the percentage of the time the actual generated values will be within optimistic and pessimistic estimates. The default value is CL99.
     * @param "Iterations" (optional.Int32) -  The number of iterations to use in Monte Carlo simulation. The default value is 100.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The folder storage.
     * @param "FileName" (optional.String) -  The resulting report fileName.

@return []byte
*/
func (a *TasksApiService) GetRiskAnalysisReport(ctx context.Context, localVarOptionals *requests.GetRiskAnalysisReportOpts) ([]byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/riskAnalysisReport"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("taskUid", fmt.Sprintf("%v", localVarOptionals.TaskUid))
	if localVarOptionals != nil && localVarOptionals.DistributionType.IsSet() {
		localVarQueryParams.Add("distributionType", parameterToString(localVarOptionals.DistributionType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Optimistic.IsSet() {
		localVarQueryParams.Add("optimistic", parameterToString(localVarOptionals.Optimistic.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Pessimistic.IsSet() {
		localVarQueryParams.Add("pessimistic", parameterToString(localVarOptionals.Pessimistic.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConfidenceLevel.IsSet() {
		localVarQueryParams.Add("confidenceLevel", parameterToString(localVarOptionals.ConfidenceLevel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Iterations.IsSet() {
		localVarQueryParams.Add("iterations", parameterToString(localVarOptionals.Iterations.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Deletes a project task with all references to it and rebuilds tasks tree.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteTaskOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Task Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteTask(ctx context.Context, localVarOptionals *requests.DeleteTaskOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Read project task.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTaskOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Task Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.TaskResponse
*/
func (a *TasksApiService) GetTask(ctx context.Context, localVarOptionals *requests.GetTaskOpts) (models.TaskResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get task assignments.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTaskAssignmentsOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Task Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AssignmentsResponse
*/
func (a *TasksApiService) GetTaskAssignments(ctx context.Context, localVarOptionals *requests.GetTaskAssignmentsOpts) (models.AssignmentsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.AssignmentsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get recurring info for a task with the specified Uid
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTaskRecurringInfoOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Task Uid
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.RecurringInfoResponse
*/
func (a *TasksApiService) GetTaskRecurringInfo(ctx context.Context, localVarOptionals *requests.GetTaskRecurringInfoOpts) (models.RecurringInfoResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.RecurringInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}/recurringInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get timescaled data for a task with the specified Uid.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTaskTimephasedDataOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Uid of task to get timephased data for.
     * @param "Type_" (optional.String) -  Type of timephased data to get.
     * @param "StartDate" (*custom.TimeWithoutTZ) -  Start date.
     * @param "EndDate" (*custom.TimeWithoutTZ) -  End date.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  The document storage.

@return models.TimephasedDataResponse
*/
func (a *TasksApiService) GetTaskTimephasedData(ctx context.Context, localVarOptionals *requests.GetTaskTimephasedDataOpts) (models.TimephasedDataResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.TimephasedDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}/timeScaleData"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(localVarOptionals.StartDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(localVarOptionals.EndDate, ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get a project&#39;s tasks.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTasksOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.TaskItemsResponse
*/
func (a *TasksApiService) GetTasks(ctx context.Context, localVarOptionals *requests.GetTasksOpts) (models.TaskItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Add a new task to a project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostTaskOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskName" (optional.String) -  The name of the new task. The default value is an empty string
     * @param "BeforeTaskId" (optional.Int32) -  The id of the task to insert the new task before.              The default value is the id of the last task in a project file.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.TaskItemResponse
*/
func (a *TasksApiService) PostTask(ctx context.Context, localVarOptionals *requests.PostTaskOpts) (models.TaskItemResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskItemResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.TaskName.IsSet() {
		localVarQueryParams.Add("taskName", parameterToString(localVarOptionals.TaskName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeforeTaskId.IsSet() {
		localVarQueryParams.Add("beforeTaskId", parameterToString(localVarOptionals.BeforeTaskId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Adds a new recurring task.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostTaskRecurringInfoOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "ParentTaskUid" (int32) - The Uid of parent task for the created recurring task
     * @param "TaskName" (string) - Name of the task to create.
     * @param "RecurringInfo" (RecurringInfo) - DTO which defines task&#39;s recurring info.
     * @param "CalendarName" (string) - Name of the project&#39;s calendar to use for recurring task&#39;s scheduling.
     * @param "FileName" (optional.String) -  File name to save changes to.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.TaskItemResponse
*/
func (a *TasksApiService) PostTaskRecurringInfo(ctx context.Context, localVarOptionals *requests.PostTaskRecurringInfoOpts) (models.TaskItemResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskItemResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/recurringInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("parentTaskUid", fmt.Sprintf("%v", localVarOptionals.ParentTaskUid))
	localVarQueryParams.Add("taskName", fmt.Sprintf("%v", localVarOptionals.TaskName))
	localVarQueryParams.Add("calendarName", fmt.Sprintf("%v", localVarOptionals.CalendarName))
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.RecurringInfo
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Create multiple tasks by single request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostTasksOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Requests" ([]TaskCreationRequest) - Requests to create tasks.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.TaskItemsResponse
*/
func (a *TasksApiService) PostTasks(ctx context.Context, localVarOptionals *requests.PostTasksOpts) (models.TaskItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/batch"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Requests
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Move one task to another parent task.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutMoveTaskOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Unique id of the task to be moved.
     * @param "ParentTaskUid" (int32) - Unique id of the new parent task.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. 
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutMoveTask(ctx context.Context, localVarOptionals *requests.PutMoveTaskOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}/move"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("parentTaskUid", fmt.Sprintf("%v", localVarOptionals.ParentTaskUid))
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Move a task to another position under the same parent and the same outline level
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutMoveTaskToSiblingOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Task Unique Id.
     * @param "BeforeTaskUid" (int32) - Unique Id of the task after which the current task will be placed.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.             If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutMoveTaskToSibling(ctx context.Context, localVarOptionals *requests.PutMoveTaskToSiblingOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}/moveToSibling"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("beforeTaskUid", fmt.Sprintf("%v", localVarOptionals.BeforeTaskUid))
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Updates special task getting by task UID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutTaskOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Task UID
     * @param "Task" (Task) - TaskDTO
     * @param "Mode" (optional.String) -  CalculationModeDTO
     * @param "Recalculate" (optional.Bool) -  boolean value 
     * @param "Storage" (optional.String) -  The document strorage
     * @param "Folder" (optional.String) -  The document folder
     * @param "FileName" (optional.String) -  The name of the file to save changes

@return models.TaskResponse
*/
func (a *TasksApiService) PutTask(ctx context.Context, localVarOptionals *requests.PutTaskOpts) (models.TaskResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Mode.IsSet() {
		localVarQueryParams.Add("mode", parameterToString(localVarOptionals.Mode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Recalculate.IsSet() {
		localVarQueryParams.Add("recalculate", parameterToString(localVarOptionals.Recalculate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.Task
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Updates existing task&#39;s recurring info. Note that task should be already recurring.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutTaskRecurringInfoOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUid" (int32) - Task Uid.
     * @param "RecurringInfo" (RecurringInfo) - A modified DTO of task&#39;s recurring info.
     * @param "FileName" (optional.String) -  File name to save changes to.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutTaskRecurringInfo(ctx context.Context, localVarOptionals *requests.PutTaskRecurringInfoOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/tasks/{taskUid}/recurringInfo"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskUid"+"}", fmt.Sprintf("%v", localVarOptionals.TaskUid), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.RecurringInfo
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Delete task link.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.DeleteTaskLinkOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Index" (int32) - Index of the task link object. See TaskLink.Index property.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document.

@return models.AsposeResponse
*/
func (a *TasksApiService) DeleteTaskLink(ctx context.Context, localVarOptionals *requests.DeleteTaskLinkOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/taskLinks/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Get tasks&#39; links of a project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetTaskLinksOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.TaskLinksResponse
*/
func (a *TasksApiService) GetTaskLinks(ctx context.Context, localVarOptionals *requests.GetTaskLinksOpts) (models.TaskLinksResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskLinksResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/taskLinks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Adds a new task link to a project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostTaskLinkOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskLink" (TaskLink) - The TaskLink object representation that should be added.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document.

@return models.AsposeResponse
*/
func (a *TasksApiService) PostTaskLink(ctx context.Context, localVarOptionals *requests.PostTaskLinkOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/taskLinks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.TaskLink
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Updates existing task link.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutTaskLinkOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Index" (int32) - Index of the task link object. See TaskLink.Index property.
     * @param "TaskLink" (TaskLink) - The representation of the modified TaskLink object.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document.

@return models.TaskLinkResponse
*/
func (a *TasksApiService) PutTaskLink(ctx context.Context, localVarOptionals *requests.PutTaskLinkOpts) (models.TaskLinkResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.TaskLinkResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/taskLinks/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.TaskLink
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Returns VBA project.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetVbaProjectOpts - Request parameters:
     * @param "Name" (string) - The name of the file
     * @param "Folder" (optional.String) -  The folder storage
     * @param "Storage" (optional.String) -  The document storage.

@return models.VbaProjectResponse
*/
func (a *TasksApiService) GetVbaProject(ctx context.Context, localVarOptionals *requests.GetVbaProjectOpts) (models.VbaProjectResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.VbaProjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/vbaproject"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/* 
TasksApiService Get a project&#39;s WBS Definition.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.GetWbsDefinitionOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return models.WbsDefinitionResponse
*/
func (a *TasksApiService) GetWbsDefinition(ctx context.Context, localVarOptionals *requests.GetWbsDefinitionOpts) (models.WbsDefinitionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarReturnValue models.WbsDefinitionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/wbsDefinition"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}

/* 
TasksApiService Renumber WBS code of passed tasks (if specified) or all project&#39;s tasks.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PutRenumberWbsCodeOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "TaskUids" ([]int32) - Uids of task for WBS codes renumbering. If not specified, WBS codes for all tasks will be renumbered.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "FileName" (optional.String) -  The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.
     * @param "Folder" (optional.String) -  The document folder.

@return models.AsposeResponse
*/
func (a *TasksApiService) PutRenumberWbsCode(ctx context.Context, localVarOptionals *requests.PutRenumberWbsCodeOpts) (models.AsposeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarReturnValue models.AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/renumberWbsCode"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileName.IsSet() {
		localVarQueryParams.Add("fileName", parameterToString(localVarOptionals.FileName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.TaskUids
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
    if err != nil || localVarHttpResponse == nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
    if localVarHttpResponse.StatusCode == 401 {
        defer localVarHttpResponse.Body.Close()
        return localVarReturnValue, nil, errors.New("Access is denied")
    }
    if localVarHttpResponse.StatusCode >= 300 {
        defer localVarHttpResponse.Body.Close()

        var apiError TasksApiErrorResponse

        if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
            return localVarReturnValue, localVarHttpResponse, err
        }

        return localVarReturnValue, localVarHttpResponse, &apiError
    }
    defer localVarHttpResponse.Body.Close()
    if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

    return localVarReturnValue, localVarHttpResponse, err
}
/*
TasksApiService Get a project document in the specified format and with the specified save options
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param *requests.PostTaskDocumentWithFormatOpts - Request parameters:
     * @param "Name" (string) - The name of the file.
     * @param "Format" (string) - The format of the resulting file.
     * @param "SaveOptions" (interface{}) - Instance of inheritor of SaveOptions class which contains format-specific settings for saving a project.
     * @param "ReturnAsZipArchive" (optional.Bool) -  If parameter is true, HTML resources are included as separate files and returned along with the resulting html file as a zip package.
     * @param "Storage" (optional.String) -  The document storage.
     * @param "Folder" (optional.String) -  The document folder.

@return []byte
*/
func (a *TasksApiService) PostTaskDocumentWithFormat(ctx context.Context, localVarOptionals *requests.PostTaskDocumentWithFormatOpts) ([]byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarReturnValue []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/tasks/{name}/format"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make([]FormParamContainer, 0)

	localVarQueryParams.Add("format", fmt.Sprintf("%v", localVarOptionals.Format))
	if localVarOptionals != nil && localVarOptionals.ReturnAsZipArchive.IsSet() {
		localVarQueryParams.Add("returnAsZipArchive", parameterToString(localVarOptionals.ReturnAsZipArchive.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Storage.IsSet() {
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Folder.IsSet() {
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &localVarOptionals.SaveOptions
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode == 401 {
		defer localVarHttpResponse.Body.Close()
		return localVarReturnValue, nil, errors.New("Access is denied")
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()

		var apiError TasksApiErrorResponse

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}

		return localVarReturnValue, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}
