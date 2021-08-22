# \TasksApi

All URIs are relative to *https://api.aspose.cloud/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFile**](TasksApi.md#copyfile) | **Put** /tasks/storage/file/copy/{srcPath} | Copy file
[**DeleteFile**](TasksApi.md#deletefile) | **Delete** /tasks/storage/file/{path} | Delete file
[**DownloadFile**](TasksApi.md#downloadfile) | **Get** /tasks/storage/file/{path} | Download file
[**MoveFile**](TasksApi.md#movefile) | **Put** /tasks/storage/file/move/{srcPath} | Move file
[**UploadFile**](TasksApi.md#uploadfile) | **Put** /tasks/storage/file/{path} | Upload file
[**CopyFolder**](TasksApi.md#copyfolder) | **Put** /tasks/storage/folder/copy/{srcPath} | Copy folder
[**CreateFolder**](TasksApi.md#createfolder) | **Put** /tasks/storage/folder/{path} | Create the folder
[**DeleteFolder**](TasksApi.md#deletefolder) | **Delete** /tasks/storage/folder/{path} | Delete folder
[**GetFilesList**](TasksApi.md#getfileslist) | **Get** /tasks/storage/folder/{path} | Get all files and folders within a folder
[**MoveFolder**](TasksApi.md#movefolder) | **Put** /tasks/storage/folder/move/{srcPath} | Move folder
[**GetDiscUsage**](TasksApi.md#getdiscusage) | **Get** /tasks/storage/disc | Get disc usage
[**GetFileVersions**](TasksApi.md#getfileversions) | **Get** /tasks/storage/version/{path} | Get file versions
[**ObjectExists**](TasksApi.md#objectexists) | **Get** /tasks/storage/exist/{path} | Check if file or folder exists
[**StorageExists**](TasksApi.md#storageexists) | **Get** /tasks/storage/{storageName}/exist | Check if storage exists
[**DeleteAssignment**](TasksApi.md#deleteassignment) | **Delete** /tasks/{name}/assignments/{assignmentUid} | Deletes a project assignment with all references to it.
[**GetAssignment**](TasksApi.md#getassignment) | **Get** /tasks/{name}/assignments/{assignmentUid} | Read project assignment with the specified Uid.
[**GetAssignmentTimephasedData**](TasksApi.md#getassignmenttimephaseddata) | **Get** /tasks/{name}/assignments/{assignmentUid}/timeScaleData | Get timescaled data for an assignment with the specified Uid.
[**GetAssignments**](TasksApi.md#getassignments) | **Get** /tasks/{name}/assignments | Get project&#39;s assignment items.
[**PostAssignment**](TasksApi.md#postassignment) | **Post** /tasks/{name}/assignments | Adds a new assignment to a project and returns assignment item in a response.
[**PutAssignment**](TasksApi.md#putassignment) | **Put** /tasks/{name}/assignments/{assignmentUid} | Updates resources assignment with the specified Uid.
[**DeleteCalendar**](TasksApi.md#deletecalendar) | **Delete** /tasks/{name}/calendars/{calendarUid} | Deletes a project calendar
[**DeleteCalendarException**](TasksApi.md#deletecalendarexception) | **Delete** /tasks/{name}/calendars/{calendarUid}/calendarExceptions/{index} | Deletes calendar exception from calendar exceptions collection.
[**GetCalendar**](TasksApi.md#getcalendar) | **Get** /tasks/{name}/calendars/{calendarUid} | Get a project&#39;s calendar with the specified Uid.
[**GetCalendarExceptions**](TasksApi.md#getcalendarexceptions) | **Get** /tasks/{name}/calendars/{calendarUid}/calendarExceptions | Get a list of calendar&#39;s exceptions.
[**GetCalendarWorkWeeks**](TasksApi.md#getcalendarworkweeks) | **Get** /tasks/{name}/calendars/{calendarUid}/workWeeks | Gets the collection of work weeks of the specified calendar.
[**GetCalendars**](TasksApi.md#getcalendars) | **Get** /tasks/{name}/calendars | Read project calendar items.
[**PostCalendar**](TasksApi.md#postcalendar) | **Post** /tasks/{name}/calendars | Adds a new calendar to project file.
[**PostCalendarException**](TasksApi.md#postcalendarexception) | **Post** /tasks/{name}/calendars/{calendarUid}/calendarExceptions | Adds a new calendar exception to a calendar.
[**PutCalendar**](TasksApi.md#putcalendar) | **Put** /tasks/{name}/calendars | Edits an existing project calendar.
[**PutCalendarException**](TasksApi.md#putcalendarexception) | **Put** /tasks/{name}/calendars/{calendarUid}/calendarExceptions/{index} | Updates calendar exception.
[**GetCriticalPath**](TasksApi.md#getcriticalpath) | **Get** /tasks/{name}/criticalPath | Returns the list of the tasks which must be completed on time for a project to finish on schedule. Each task of the project is represented as a task item here, which is light-weighted task representation of the project task..
[**GetPageCount**](TasksApi.md#getpagecount) | **Get** /tasks/{name}/pagecount | Returns page count for the project to be rendered using given Timescale and PresentationFormat  or specified PageSize.
[**GetProjectIds**](TasksApi.md#getprojectids) | **Get** /tasks/{name}/projectids | Get Uids of projects contained in the file.
[**GetTaskDocument**](TasksApi.md#gettaskdocument) | **Get** /tasks/{name} | Get a project document.
[**GetTaskDocumentWithFormat**](TasksApi.md#gettaskdocumentwithformat) | **Get** /tasks/{name}/format | Get a project document in the specified format
[**PutImportProjectFromDb**](TasksApi.md#putimportprojectfromdb) | **Put** /tasks/importfromdb | Imports project from database with the specified connection string and saves it to specified file with the specified format.
[**PutImportProjectFromFile**](TasksApi.md#putimportprojectfromfile) | **Put** /tasks/{name}/import | Imports project from primavera db formats (Primavera SQLite .db or Primavera xml) and saves it to specified file with the specified format.
[**PutImportProjectFromProjectOnline**](TasksApi.md#putimportprojectfromprojectonline) | **Put** /tasks/{name}/importfromprojectonline | Imports project from Project Online and saves it to specified file.
[**GetDocumentProperties**](TasksApi.md#getdocumentproperties) | **Get** /tasks/{name}/documentproperties | Get a collection of a project&#39;s document properties.
[**GetDocumentProperty**](TasksApi.md#getdocumentproperty) | **Get** /tasks/{name}/documentproperties/{propertyName} | Get a document property by name.
[**PostDocumentProperty**](TasksApi.md#postdocumentproperty) | **Post** /tasks/{name}/documentproperties/{propertyName} | Set/create document property.
[**PutDocumentProperty**](TasksApi.md#putdocumentproperty) | **Put** /tasks/{name}/documentproperties/{propertyName} | Set/create document property.
[**DeleteExtendedAttributeByIndex**](TasksApi.md#deleteextendedattributebyindex) | **Delete** /tasks/{name}/extendedAttributes/{index} | Delete a project extended attribute.
[**GetExtendedAttributeByIndex**](TasksApi.md#getextendedattributebyindex) | **Get** /tasks/{name}/extendedAttributes/{index} | Get a project extended attribute definition with the specified index.
[**GetExtendedAttributes**](TasksApi.md#getextendedattributes) | **Get** /tasks/{name}/extendedAttributes | Get a project&#39;s extended attributes.
[**PutExtendedAttribute**](TasksApi.md#putextendedattribute) | **Put** /tasks/{name}/extendedAttributes | Add a new extended attribute definition to a project or  updates existing attribute definition (with the same FieldId).
[**DeleteOutlineCodeByIndex**](TasksApi.md#deleteoutlinecodebyindex) | **Delete** /tasks/{name}/outlineCodes/{index} | Deletes a project outline code
[**GetOutlineCodeByIndex**](TasksApi.md#getoutlinecodebyindex) | **Get** /tasks/{name}/outlineCodes/{index} | Get outline code by index.
[**GetOutlineCodes**](TasksApi.md#getoutlinecodes) | **Get** /tasks/{name}/outlineCodes | Get a project&#39;s outline codes.
[**CreateNewProject**](TasksApi.md#createnewproject) | **Post** /tasks/projectOnline/{siteUrl}/{name} | Creates new project in Project Server\\Project Online instance.
[**GetProjectList**](TasksApi.md#getprojectlist) | **Get** /tasks/projectOnline/projectList | Gets the list of published projects in the current Project Online account.
[**UpdateProject**](TasksApi.md#updateproject) | **Put** /tasks/projectOnline/{siteUrl}/{name} | Updates existing project in Project Server\\Project Online instance. The existing project will be overwritten.
[**PutRecalculateProject**](TasksApi.md#putrecalculateproject) | **Put** /tasks/{name}/recalculate/project | Reschedules all project tasks ids, outline levels, start/finish dates, sets early/late dates, calculates slacks, work and cost fields.
[**PutRecalculateProjectResourceFields**](TasksApi.md#putrecalculateprojectresourcefields) | **Put** /tasks/{name}/recalculate/resourceFields | Recalculate project resource fields
[**PutRecalculateProjectUncompleteWorkToStartAfter**](TasksApi.md#putrecalculateprojectuncompleteworktostartafter) | **Put** /tasks/{name}/recalculate/uncompleteWorkToStartAfter | Recalculate project uncomplete work
[**PutRecalculateProjectWorkAsComplete**](TasksApi.md#putrecalculateprojectworkascomplete) | **Put** /tasks/{name}/recalculate/projectWorkAsComplete | Recalculate project work as complete 
[**GetReportPdf**](TasksApi.md#getreportpdf) | **Get** /tasks/{name}/report | Returns created report in PDF format.
[**DeleteResource**](TasksApi.md#deleteresource) | **Delete** /tasks/{name}/resources/{resourceUid} | Deletes a project resource with all references to it
[**GetResource**](TasksApi.md#getresource) | **Get** /tasks/{name}/resources/{resourceUid} | Get project resource.
[**GetResourceAssignments**](TasksApi.md#getresourceassignments) | **Get** /tasks/{name}/resources/{resourceUid}/assignments | Get resource&#39;s assignments.
[**GetResourceTimephasedData**](TasksApi.md#getresourcetimephaseddata) | **Get** /tasks/{name}/resources/{resourceUid}/timeScaleData | Get timescaled data for a resource with the specified Uid.
[**GetResources**](TasksApi.md#getresources) | **Get** /tasks/{name}/resources | Get a project&#39;s resources.
[**PostResource**](TasksApi.md#postresource) | **Post** /tasks/{name}/resources | Add a new resource to a project.
[**PutResource**](TasksApi.md#putresource) | **Put** /tasks/{name}/resources/{resourceUid} | Updates resource with the specified Uid
[**GetRiskAnalysisReport**](TasksApi.md#getriskanalysisreport) | **Get** /tasks/{name}/riskAnalysisReport | Performs a risk analysys using Monte Carlo simulation and creates a risk analysis report.
[**DeleteTask**](TasksApi.md#deletetask) | **Delete** /tasks/{name}/tasks/{taskUid} | Deletes a project task with all references to it and rebuilds tasks tree.
[**GetTask**](TasksApi.md#gettask) | **Get** /tasks/{name}/tasks/{taskUid} | Read project task.
[**GetTaskAssignments**](TasksApi.md#gettaskassignments) | **Get** /tasks/{name}/tasks/{taskUid}/assignments | Get task assignments.
[**GetTaskRecurringInfo**](TasksApi.md#gettaskrecurringinfo) | **Get** /tasks/{name}/tasks/{taskUid}/recurringInfo | Get recurring info for a task with the specified Uid
[**GetTaskTimephasedData**](TasksApi.md#gettasktimephaseddata) | **Get** /tasks/{name}/tasks/{taskUid}/timeScaleData | Get timescaled data for a task with the specified Uid.
[**GetTasks**](TasksApi.md#gettasks) | **Get** /tasks/{name}/tasks | Get a project&#39;s tasks.
[**PostTask**](TasksApi.md#posttask) | **Post** /tasks/{name}/tasks | Add a new task to a project.
[**PostTaskRecurringInfo**](TasksApi.md#posttaskrecurringinfo) | **Post** /tasks/{name}/tasks/recurringInfo | Adds a new recurring task.
[**PostTasks**](TasksApi.md#posttasks) | **Post** /tasks/{name}/tasks/batch | Create multiple tasks by single request.
[**PutMoveTask**](TasksApi.md#putmovetask) | **Put** /tasks/{name}/tasks/{taskUid}/move | Move one task to another parent task.
[**PutMoveTaskToSibling**](TasksApi.md#putmovetasktosibling) | **Put** /tasks/{name}/tasks/{taskUid}/moveToSibling | Move a task to another position under the same parent and the same outline level
[**PutTask**](TasksApi.md#puttask) | **Put** /tasks/{name}/tasks/{taskUid} | Updates special task getting by task UID
[**PutTaskRecurringInfo**](TasksApi.md#puttaskrecurringinfo) | **Put** /tasks/{name}/tasks/{taskUid}/recurringInfo | Updates existing task&#39;s recurring info. Note that task should be already recurring.
[**DeleteTaskLink**](TasksApi.md#deletetasklink) | **Delete** /tasks/{name}/taskLinks/{index} | Delete task link.
[**GetTaskLinks**](TasksApi.md#gettasklinks) | **Get** /tasks/{name}/taskLinks | Get tasks&#39; links of a project.
[**PostTaskLink**](TasksApi.md#posttasklink) | **Post** /tasks/{name}/taskLinks | Adds a new task link to a project.
[**PutTaskLink**](TasksApi.md#puttasklink) | **Put** /tasks/{name}/taskLinks/{index} | Updates existing task link.
[**GetVbaProject**](TasksApi.md#getvbaproject) | **Get** /tasks/{name}/vbaproject | Returns VBA project.
[**GetWbsDefinition**](TasksApi.md#getwbsdefinition) | **Get** /tasks/{name}/wbsDefinition | Get a project&#39;s WBS Definition.
[**PutRenumberWbsCode**](TasksApi.md#putrenumberwbscode) | **Put** /tasks/{name}/renumberWbsCode | Renumber WBS code of passed tasks (if specified) or all project&#39;s tasks.


# **CopyFile**
> CopyFile(ctx, srcPath, destPath, optional)
Copy file

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **srcPath** | **string**| Source file path e.g. &#39;/folder/file.ext&#39; | 
  **destPath** | **string**| Destination file path | 
 **optional** | ***CopyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **srcStorageName** | **optional.String**| Source storage name | 
 **destStorageName** | **optional.String**| Destination storage name | 
 **versionId** | **optional.String**| File version ID to copy | 

### Return type
 (empty response body)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> DeleteFile(ctx, path, optional)
Delete file

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| File path e.g. &#39;/folder/file.ext&#39; | 
 **optional** | ***DeleteFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 
 **versionId** | **optional.String**| File version ID to delete | 

### Return type
 (empty response body)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFile**
> *os.File DownloadFile(ctx, path, optional)
Download file

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| File path e.g. &#39;/folder/file.ext&#39; | 
 **optional** | ***DownloadFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 
 **versionId** | **optional.String**| File version ID to download | 

### Return type
[***os.File**](*os.File.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFile**
> MoveFile(ctx, srcPath, destPath, optional)
Move file

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **srcPath** | **string**| Source file path e.g. &#39;/src.ext&#39; | 
  **destPath** | **string**| Destination file path e.g. &#39;/dest.ext&#39; | 
 **optional** | ***MoveFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **srcStorageName** | **optional.String**| Source storage name | 
 **destStorageName** | **optional.String**| Destination storage name | 
 **versionId** | **optional.String**| File version ID to move | 

### Return type
 (empty response body)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFile**
> FilesUploadResult UploadFile(ctx, path, file, optional)
Upload file

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.              | 
  **file** | ***os.File**| File to upload | 
 **optional** | ***UploadFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storageName** | **optional.String**| Storage name | 

### Return type
[**FilesUploadResult**](FilesUploadResult.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyFolder**
> CopyFolder(ctx, srcPath, destPath, optional)
Copy folder

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **srcPath** | **string**| Source folder path e.g. &#39;/src&#39; | 
  **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; | 
 **optional** | ***CopyFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **srcStorageName** | **optional.String**| Source storage name | 
 **destStorageName** | **optional.String**| Destination storage name | 

### Return type
 (empty response body)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> CreateFolder(ctx, path, optional)
Create the folder

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Folder path to create e.g. &#39;folder_1/folder_2/&#39; | 
 **optional** | ***CreateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 

### Return type
 (empty response body)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> DeleteFolder(ctx, path, optional)
Delete folder

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Folder path e.g. &#39;/folder&#39; | 
 **optional** | ***DeleteFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 
 **recursive** | **optional.Bool**| Enable to delete folders, subfolders and files | [default to false]

### Return type
 (empty response body)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesList**
> FilesList GetFilesList(ctx, path, optional)
Get all files and folders within a folder

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Folder path e.g. &#39;/folder&#39; | 
 **optional** | ***GetFilesListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 

### Return type
[**FilesList**](FilesList.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFolder**
> MoveFolder(ctx, srcPath, destPath, optional)
Move folder

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **srcPath** | **string**| Folder path to move e.g. &#39;/folder&#39; | 
  **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; | 
 **optional** | ***MoveFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **srcStorageName** | **optional.String**| Source storage name | 
 **destStorageName** | **optional.String**| Destination storage name | 

### Return type
 (empty response body)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiscUsage**
> DiscUsage GetDiscUsage(ctx, optional)
Get disc usage

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDiscUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageName** | **optional.String**| Storage name | 

### Return type
[**DiscUsage**](DiscUsage.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileVersions**
> FileVersions GetFileVersions(ctx, path, optional)
Get file versions

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| File path e.g. &#39;/file.ext&#39; | 
 **optional** | ***GetFileVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 

### Return type
[**FileVersions**](FileVersions.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectExists**
> ObjectExist ObjectExists(ctx, path, optional)
Check if file or folder exists

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39; | 
 **optional** | ***ObjectExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageName** | **optional.String**| Storage name | 
 **versionId** | **optional.String**| File version ID | 

### Return type
[**ObjectExist**](ObjectExist.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StorageExists**
> StorageExist StorageExists(ctx, storageName)
Check if storage exists

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storageName** | **string**| Storage name | 

### Return type
[**StorageExist**](StorageExist.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAssignment**
> AsposeResponse DeleteAssignment(ctx, name, assignmentUid, optional)
Deletes a project assignment with all references to it.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **assignmentUid** | **int32**| assignment Uid | 
 **optional** | ***DeleteAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **fileName** | **optional.String**| The name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssignment**
> AssignmentResponse GetAssignment(ctx, name, assignmentUid, optional)
Read project assignment with the specified Uid.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **assignmentUid** | **int32**| Assignment Uid | 
 **optional** | ***GetAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AssignmentResponse**](AssignmentResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssignmentTimephasedData**
> TimephasedDataResponse GetAssignmentTimephasedData(ctx, name, assignmentUid, optional)
Get timescaled data for an assignment with the specified Uid.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **assignmentUid** | **int32**| Uid of assignment to get timephased data for. | 
 **optional** | ***GetAssignmentTimephasedDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| Type of timephased data to get. | 
 **startDate** | **optional.Time**| Start date. | 
 **endDate** | **optional.Time**| End date. | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 

### Return type
[**TimephasedDataResponse**](TimephasedDataResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssignments**
> AssignmentItemsResponse GetAssignments(ctx, name, optional)
Get project's assignment items.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AssignmentItemsResponse**](AssignmentItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAssignment**
> AssignmentItemResponse PostAssignment(ctx, name, taskUid, resourceUid, optional)
Adds a new assignment to a project and returns assignment item in a response.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| The unique id of the task to be assigned. | 
  **resourceUid** | **int32**| The unique id of the resource to be assigned. | 
 **optional** | ***PostAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **units** | **optional.Float64**| The units for the new assignment. If not specified, &#39;cost&#39; value is used. | 
 **cost** | **optional.Float32**| The cost for a new assignment. If not specified, default value is used. | 
 **fileName** | **optional.String**| The name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AssignmentItemResponse**](AssignmentItemResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAssignment**
> AssignmentResponse PutAssignment(ctx, name, assignmentUid, assignment, optional)
Updates resources assignment with the specified Uid.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The file name | 
  **assignmentUid** | **int32**| Assignment UID | 
  **assignment** | [**ResourceAssignment**](ResourceAssignment.md)| Assignment DTO | 
 **optional** | ***PutAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mode** | **optional.String**| Project&#39;s calculation mode | [default to 0]
 **recalculate** | **optional.Bool**| boolean value | [default to true]
 **storage** | **optional.String**| The document storage | 
 **folder** | **optional.String**| The document storage | 
 **fileName** | **optional.String**| The filename to save Changes | 

### Return type
[**AssignmentResponse**](AssignmentResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCalendar**
> AsposeResponse DeleteCalendar(ctx, name, calendarUid, optional)
Deletes a project calendar

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| Calendar&#39;s Uid | 
 **optional** | ***DeleteCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCalendarException**
> AsposeResponse DeleteCalendarException(ctx, name, calendarUid, index, optional)
Deletes calendar exception from calendar exceptions collection.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| Calendar Uid | 
  **index** | **int32**| Index of the calendar exception. See CalendarException.Index property. | 
 **optional** | ***DeleteCalendarExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCalendar**
> CalendarResponse GetCalendar(ctx, name, calendarUid, optional)
Get a project's calendar with the specified Uid.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| Calendar&#39;s Uid | 
 **optional** | ***GetCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**CalendarResponse**](CalendarResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCalendarExceptions**
> CalendarExceptionsResponse GetCalendarExceptions(ctx, name, calendarUid, optional)
Get a list of calendar's exceptions.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| Calendar&#39;s Uid | 
 **optional** | ***GetCalendarExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**CalendarExceptionsResponse**](CalendarExceptionsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCalendarWorkWeeks**
> CalendarWorkWeeksResponse GetCalendarWorkWeeks(ctx, name, calendarUid, optional)
Gets the collection of work weeks of the specified calendar.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| Calendar Uid | 
 **optional** | ***GetCalendarWorkWeeksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**CalendarWorkWeeksResponse**](CalendarWorkWeeksResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCalendars**
> CalendarItemsResponse GetCalendars(ctx, name, optional)
Read project calendar items.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetCalendarsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**CalendarItemsResponse**](CalendarItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCalendar**
> CalendarItemResponse PostCalendar(ctx, name, calendar, optional)
Adds a new calendar to project file.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendar** | [**Calendar**](Calendar.md)| Calendar DTO | 
 **optional** | ***PostCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**CalendarItemResponse**](CalendarItemResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCalendarException**
> AsposeResponse PostCalendarException(ctx, name, calendarUid, calendarException, optional)
Adds a new calendar exception to a calendar.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| Calendar&#39;s Uid | 
  **calendarException** | [**CalendarException**](CalendarException.md)| CalendarException DTO | 
 **optional** | ***PostCalendarExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCalendar**
> AsposeResponse PutCalendar(ctx, name, calendarUid, calendar, optional)
Edits an existing project calendar.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| The Uid of an existing calendar to edit. | 
  **calendar** | [**Calendar**](Calendar.md)| Modified calendar DTO | 
 **optional** | ***PutCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCalendarException**
> AsposeResponse PutCalendarException(ctx, name, calendarUid, index, calendarException, optional)
Updates calendar exception.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **calendarUid** | **int32**| Calendar Uid | 
  **index** | **int32**| Index of the calendar exception. See CalendarException.Index property. | 
  **calendarException** | [**CalendarException**](CalendarException.md)| CalendarException DTO | 
 **optional** | ***PutCalendarExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fileName** | **optional.String**| The name of the project document to save changes to. If this parameter              is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCriticalPath**
> TaskItemsResponse GetCriticalPath(ctx, name, optional)
Returns the list of the tasks which must be completed on time for a project to finish on schedule. Each task of the project is represented as a task item here, which is light-weighted task representation of the project task..

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetCriticalPathOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**TaskItemsResponse**](TaskItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageCount**
> PageCountResponse GetPageCount(ctx, name, optional)
Returns page count for the project to be rendered using given Timescale and PresentationFormat  or specified PageSize.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetPageCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.String**| PageSize to get page count for. | 
 **presentationFormat** | **optional.String**| PresentationFormat to get page count for. | 
 **timescale** | **optional.String**| Timescale to get page count for. | [default to 1]
 **startDate** | **optional.Time**| Start date to get page count for. | 
 **endDate** | **optional.Time**| End date to get page count for. | 
 **folder** | **optional.String**| The document folder | 
 **storage** | **optional.String**| The document storage. | 

### Return type
[**PageCountResponse**](PageCountResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectIds**
> ProjectIdsResponse GetProjectIds(ctx, name, optional)
Get Uids of projects contained in the file.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetProjectIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**ProjectIdsResponse**](ProjectIdsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskDocument**
> *os.File GetTaskDocument(ctx, name, optional)
Get a project document.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetTaskDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[***os.File**](*os.File.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskDocumentWithFormat**
> *os.File GetTaskDocumentWithFormat(ctx, name, format, optional)
Get a project document in the specified format

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **format** | **string**| The format of the resulting file. | 
 **optional** | ***GetTaskDocumentWithFormatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **returnAsZipArchive** | **optional.Bool**| If parameter is true, HTML resources are included as separate files and returned along with the resulting html file as a zip package. | [default to false]
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[***os.File**](*os.File.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImportProjectFromDb**
> AsposeResponse PutImportProjectFromDb(ctx, databaseType, connectionString, projectUid, filename, optional)
Imports project from database with the specified connection string and saves it to specified file with the specified format.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseType** | **string**| The type of source database. The supported values are (Msp, Mpd, Primavera). | 
  **connectionString** | **string**| The connection string to the source database. | 
  **projectUid** | **string**| Uid of the project to import. | 
  **filename** | **string**| The name of the resulting file. | 
 **optional** | ***PutImportProjectFromDbOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | **optional.String**| Format of the resulting file. | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 
 **databaseSchema** | **optional.String**| Schema of Microsoft project database (if applicable) | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImportProjectFromFile**
> AsposeResponse PutImportProjectFromFile(ctx, name, projectUid, filename, optional)
Imports project from primavera db formats (Primavera SQLite .db or Primavera xml) and saves it to specified file with the specified format.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **projectUid** | **string**| Uid of the project to import. | 
  **filename** | **string**| The name of the resulting file. | 
 **optional** | ***PutImportProjectFromFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileType** | **optional.String**| The type of file to import. The supported values are (PrimaveraSqliteDb, PrimaveraXml). | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 
 **outputFileFormat** | **optional.String**| The format of the resulting file. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImportProjectFromProjectOnline**
> AsposeResponse PutImportProjectFromProjectOnline(ctx, name, guid, siteUrl, optional)
Imports project from Project Online and saves it to specified file.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the resulting file. | 
  **guid** | **string**| Guid of the project to import. | 
  **siteUrl** | **string**| The URL of PWA (Project Web Access) API of Project Online | 
 **optional** | ***PutImportProjectFromProjectOnlineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userName** | **optional.String**| The user name for the sharepoint site. | 
 **format** | **optional.String**| Format of the resulting file. | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 
 **xProjectOnlineToken** | **optional.String**| Authorization token (SPOIDCRL) for SharePoint&#39;s PWA (Project Web Access). For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly | 
 **xSharepointPassword** | **optional.String**| The password for the SharePoint site. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperties**
> DocumentPropertiesResponse GetDocumentProperties(ctx, name, optional)
Get a collection of a project's document properties.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The document name. | 
 **optional** | ***GetDocumentPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**DocumentPropertiesResponse**](DocumentPropertiesResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperty**
> DocumentPropertyResponse GetDocumentProperty(ctx, name, propertyName, optional)
Get a document property by name.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The document name. | 
  **propertyName** | **string**| The property name. | 
 **optional** | ***GetDocumentPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentProperty**
> DocumentPropertyResponse PostDocumentProperty(ctx, name, propertyName, property, optional)
Set/create document property.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The document name. | 
  **propertyName** | **string**| The property name. | 
  **property** | [**DocumentProperty**](DocumentProperty.md)| DocumentProperty with new property value. | 
 **optional** | ***PostDocumentPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **filename** | **optional.String**| Name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutDocumentProperty**
> DocumentPropertyResponse PutDocumentProperty(ctx, name, propertyName, property, optional)
Set/create document property.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The document name. | 
  **propertyName** | **string**| The property name. | 
  **property** | [**DocumentProperty**](DocumentProperty.md)| DocumentProperty with new property value. | 
 **optional** | ***PutDocumentPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **filename** | **optional.String**| Name of the project document to save changes to. If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExtendedAttributeByIndex**
> AsposeResponse DeleteExtendedAttributeByIndex(ctx, name, index, optional)
Delete a project extended attribute.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **index** | **int32**| Index (See ExtendedAttributeItem.Index property) or FieldId of the extended attribute. | 
 **optional** | ***DeleteExtendedAttributeByIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtendedAttributeByIndex**
> ExtendedAttributeResponse GetExtendedAttributeByIndex(ctx, name, index, optional)
Get a project extended attribute definition with the specified index.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **index** | **int32**| Index (See ExtendedAttributeItem.Index property) or FieldId of the extended attribute. | 
 **optional** | ***GetExtendedAttributeByIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**ExtendedAttributeResponse**](ExtendedAttributeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtendedAttributes**
> ExtendedAttributeItemsResponse GetExtendedAttributes(ctx, name, optional)
Get a project's extended attributes.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetExtendedAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**ExtendedAttributeItemsResponse**](ExtendedAttributeItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutExtendedAttribute**
> ExtendedAttributeItemResponse PutExtendedAttribute(ctx, extendedAttribute, name, optional)
Add a new extended attribute definition to a project or  updates existing attribute definition (with the same FieldId).

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extendedAttribute** | [**ExtendedAttributeDefinition**](ExtendedAttributeDefinition.md)| Definition of the extended attribute to add. | 
  **name** | **string**| The name of the file. | 
 **optional** | ***PutExtendedAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**ExtendedAttributeItemResponse**](ExtendedAttributeItemResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOutlineCodeByIndex**
> AsposeResponse DeleteOutlineCodeByIndex(ctx, name, index, optional)
Deletes a project outline code

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **index** | **int32**| Index of the outline code. See OutlineCodeItem.Index property. | 
 **optional** | ***DeleteOutlineCodeByIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutlineCodeByIndex**
> OutlineCodeResponse GetOutlineCodeByIndex(ctx, name, index, optional)
Get outline code by index.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **index** | **int32**| Index of the outline code. See OutlineCodeItem.Index property. | 
 **optional** | ***GetOutlineCodeByIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**OutlineCodeResponse**](OutlineCodeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutlineCodes**
> OutlineCodeItemsResponse GetOutlineCodes(ctx, name, optional)
Get a project's outline codes.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetOutlineCodesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**OutlineCodeItemsResponse**](OutlineCodeItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNewProject**
> AsposeResponse CreateNewProject(ctx, name, siteUrl, optional)
Creates new project in Project Server\\Project Online instance.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **siteUrl** | **string**| The URL of PWA (Project Web Access) API of Project Online. | 
 **optional** | ***CreateNewProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userName** | **optional.String**| The user name for the sharepoint site. | 
 **saveOptions** | [**optional.Interface of ProjectServerSaveOptionsDto**](ProjectServerSaveOptionsDto.md)| Dispensable save options for Project Server\\Project Online. | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 
 **xProjectOnlineToken** | **optional.String**| Authorization token (SPOIDCRL) for SharePoint&#39;s PWA (Project Web Access). For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly | 
 **xSharepointPassword** | **optional.String**| The password for the SharePoint site. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectList**
> ProjectListResponse GetProjectList(ctx, siteUrl, optional)
Gets the list of published projects in the current Project Online account.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**| The URL of PWA (Project Web Access) API of Project Online. | 
 **optional** | ***GetProjectListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userName** | **optional.String**| The user name for the sharepoint site. | 
 **xProjectOnlineToken** | **optional.String**| Authorization token (SPOIDCRL) for SharePoint&#39;s PWA (Project Web Access). For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly | 
 **xSharepointPassword** | **optional.String**| The password for the SharePoint site. | 

### Return type
[**ProjectListResponse**](ProjectListResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> AsposeResponse UpdateProject(ctx, name, siteUrl, optional)
Updates existing project in Project Server\\Project Online instance. The existing project will be overwritten.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **siteUrl** | **string**| The URL of PWA (Project Web Access) API of Project Online. | 
 **optional** | ***UpdateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userName** | **optional.String**| The user name for the sharepoint site. | 
 **saveOptions** | [**optional.Interface of ProjectServerSaveOptionsDto**](ProjectServerSaveOptionsDto.md)| Dispensable save options for Project Server\\Project Online. | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 
 **xProjectOnlineToken** | **optional.String**| Authorization token (SPOIDCRL) for SharePoint&#39;s PWA (Project Web Access). For example, in c# it can be retrieved using SharePointOnlineCredentials class from Microsoft.SharePoint.Client.Runtime assembly | 
 **xSharepointPassword** | **optional.String**| The password for the SharePoint site. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRecalculateProject**
> ProjectRecalculateResponse PutRecalculateProject(ctx, name, optional)
Reschedules all project tasks ids, outline levels, start/finish dates, sets early/late dates, calculates slacks, work and cost fields.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The filename | 
 **optional** | ***PutRecalculateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **optional.String**| Project&#39;s calculation mode | [default to 0]
 **validate** | **optional.Bool**| If true the validation of recalculation will be performed. What data is validated:     At the moment only basic validation of task and task link date ranges is implemented.     Task&#39;s date ranges (e.g. ActualStart - ActualFinish, EarlyStart - EarlyFinish, etc.) as well as Task Links dates will be checked against the date criteria that start date is less or equal than finish date. | [default to false]
 **fileName** | **optional.String**| The filename to save the changes | 
 **storage** | **optional.String**| The document storage | 
 **folder** | **optional.String**| The document folder | 

### Return type
[**ProjectRecalculateResponse**](ProjectRecalculateResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRecalculateProjectResourceFields**
> AsposeResponse PutRecalculateProjectResourceFields(ctx, name, optional)
Recalculate project resource fields

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file | 
 **optional** | ***PutRecalculateProjectResourceFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage | 
 **folder** | **optional.String**| The document folder | 
 **fileName** | **optional.String**| The document fileName | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRecalculateProjectUncompleteWorkToStartAfter**
> AsposeResponse PutRecalculateProjectUncompleteWorkToStartAfter(ctx, name, after, optional)
Recalculate project uncomplete work

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The file name | 
  **after** | **time.Time**| DateTime (from System lib) | 
 **optional** | ***PutRecalculateProjectUncompleteWorkToStartAfterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage  | 
 **folder** | **optional.String**| The document folder | 
 **fileName** | **optional.String**| The filename to save the changes | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRecalculateProjectWorkAsComplete**
> AsposeResponse PutRecalculateProjectWorkAsComplete(ctx, name, completeThrough, optional)
Recalculate project work as complete 

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The filename | 
  **completeThrough** | **time.Time**| DateTime (type from System lib) | 
 **optional** | ***PutRecalculateProjectWorkAsCompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setZeroOrHundredPercentCompleteOnly** | **optional.Bool**| boolean value | [default to true]
 **storage** | **optional.String**| The document storage | 
 **folder** | **optional.String**| The document folder | 
 **fileName** | **optional.String**| The filename to save the changes | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportPdf**
> *os.File GetReportPdf(ctx, name, type_, optional)
Returns created report in PDF format.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **type_** | **string**| A type of the project&#39;s graphical report. | 
 **optional** | ***GetReportPdfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[***os.File**](*os.File.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResource**
> AsposeResponse DeleteResource(ctx, name, resourceUid, optional)
Deletes a project resource with all references to it

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **resourceUid** | **int32**| Resource Uid | 
 **optional** | ***DeleteResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResource**
> ResourceResponse GetResource(ctx, name, resourceUid, optional)
Get project resource.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **resourceUid** | **int32**| Resource Uid | 
 **optional** | ***GetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**ResourceResponse**](ResourceResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceAssignments**
> AssignmentsResponse GetResourceAssignments(ctx, name, resourceUid, optional)
Get resource's assignments.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **resourceUid** | **int32**| Resource Uid | 
 **optional** | ***GetResourceAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AssignmentsResponse**](AssignmentsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceTimephasedData**
> TimephasedDataResponse GetResourceTimephasedData(ctx, name, resourceUid, optional)
Get timescaled data for a resource with the specified Uid.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **resourceUid** | **int32**| Uid of resource to get timephased data for. | 
 **optional** | ***GetResourceTimephasedDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| Type of timephased data to get. | 
 **startDate** | **optional.Time**| Start date. | 
 **endDate** | **optional.Time**| End date. | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 

### Return type
[**TimephasedDataResponse**](TimephasedDataResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResources**
> ResourceItemsResponse GetResources(ctx, name, optional)
Get a project's resources.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**ResourceItemsResponse**](ResourceItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostResource**
> ResourceItemResponse PostResource(ctx, name, optional)
Add a new resource to a project.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***PostResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceName** | **optional.String**| The name of the new resource. The default value is an empty string | 
 **beforeResourceId** | **optional.Int32**| The id of the resource to insert the new resource before. The default value is the id of the last resource in a project file. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**ResourceItemResponse**](ResourceItemResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutResource**
> ResourceResponse PutResource(ctx, name, resourceUid, resource, optional)
Updates resource with the specified Uid

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The file name | 
  **resourceUid** | **int32**| The Uid of a resource | 
  **resource** | [**Resource**](Resource.md)| The representation of the modified resource | 
 **optional** | ***PutResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mode** | **optional.String**| The calculation mode of a project | [default to 0]
 **recalculate** | **optional.Bool**| Specifies whether the project&#39;s recalculation should be performed | [default to true]
 **storage** | **optional.String**| The document storage | 
 **folder** | **optional.String**| The document storage | 
 **fileName** | **optional.String**| The filename to save Changes | 

### Return type
[**ResourceResponse**](ResourceResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRiskAnalysisReport**
> *os.File GetRiskAnalysisReport(ctx, name, taskUid, optional)
Performs a risk analysys using Monte Carlo simulation and creates a risk analysis report.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| The uid of the task for which this risk will be applied in Monte Carlo simulation. | 
 **optional** | ***GetRiskAnalysisReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **distributionType** | **optional.String**| Gets or sets the probability distribution used in Monte Carlo simulation. The default value is ProbabilityDistributionType.Normal. | [default to 1]
 **optimistic** | **optional.Int32**| The percentage of the most likely task duration which can happen in the best possible project scenario. The default value is 75, which means that if the estimated specified task duration is 4 days then the optimistic duration will be 3 days. | [default to 70]
 **pessimistic** | **optional.Int32**| The percentage of the most likely task duration which can happen in the worst possible project scenario. The default value is 125, which means that if the estimated specified task duration is 4 days then the pessimistic duration will be 5 days. | [default to 130]
 **confidenceLevel** | **optional.String**| Gets or sets the confidence level that correspond to the percentage of the time the actual generated values will be within optimistic and pessimistic estimates. The default value is CL99. | [default to 75]
 **iterations** | **optional.Int32**| The number of iterations to use in Monte Carlo simulation. The default value is 100. | [default to 100]
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The folder storage. | 
 **fileName** | **optional.String**| The resulting report fileName. | 

### Return type
[***os.File**](*os.File.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTask**
> AsposeResponse DeleteTask(ctx, name, taskUid, optional)
Deletes a project task with all references to it and rebuilds tasks tree.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Task Uid | 
 **optional** | ***DeleteTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTask**
> TaskResponse GetTask(ctx, name, taskUid, optional)
Read project task.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Task Uid | 
 **optional** | ***GetTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**TaskResponse**](TaskResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskAssignments**
> AssignmentsResponse GetTaskAssignments(ctx, name, taskUid, optional)
Get task assignments.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Task Uid | 
 **optional** | ***GetTaskAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AssignmentsResponse**](AssignmentsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskRecurringInfo**
> RecurringInfoResponse GetTaskRecurringInfo(ctx, name, taskUid, optional)
Get recurring info for a task with the specified Uid

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Task Uid | 
 **optional** | ***GetTaskRecurringInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**RecurringInfoResponse**](RecurringInfoResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskTimephasedData**
> TimephasedDataResponse GetTaskTimephasedData(ctx, name, taskUid, optional)
Get timescaled data for a task with the specified Uid.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Uid of task to get timephased data for. | 
 **optional** | ***GetTaskTimephasedDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **optional.String**| Type of timephased data to get. | 
 **startDate** | **optional.Time**| Start date. | 
 **endDate** | **optional.Time**| End date. | 
 **folder** | **optional.String**| The document folder. | 
 **storage** | **optional.String**| The document storage. | 

### Return type
[**TimephasedDataResponse**](TimephasedDataResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTasks**
> TaskItemsResponse GetTasks(ctx, name, optional)
Get a project's tasks.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**TaskItemsResponse**](TaskItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTask**
> TaskItemResponse PostTask(ctx, name, optional)
Add a new task to a project.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***PostTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskName** | **optional.String**| The name of the new task. The default value is an empty string | 
 **beforeTaskId** | **optional.Int32**| The id of the task to insert the new task before.              The default value is the id of the last task in a project file. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**TaskItemResponse**](TaskItemResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTaskRecurringInfo**
> TaskItemResponse PostTaskRecurringInfo(ctx, name, parentTaskUid, taskName, recurringInfo, calendarName, optional)
Adds a new recurring task.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **parentTaskUid** | **int32**| The Uid of parent task for the created recurring task | 
  **taskName** | **string**| Name of the task to create. | 
  **recurringInfo** | [**RecurringInfo**](RecurringInfo.md)| DTO which defines task&#39;s recurring info. | 
  **calendarName** | **string**| Name of the project&#39;s calendar to use for recurring task&#39;s scheduling. | 
 **optional** | ***PostTaskRecurringInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **fileName** | **optional.String**| File name to save changes to. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**TaskItemResponse**](TaskItemResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTasks**
> TaskItemsResponse PostTasks(ctx, name, requests, optional)
Create multiple tasks by single request.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **requests** | [**[]TaskCreationRequest**](TaskCreationRequest.md)| Requests to create tasks. | 
 **optional** | ***PostTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileName** | **optional.String**| The name of the project document to save changes to. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**TaskItemsResponse**](TaskItemsResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMoveTask**
> AsposeResponse PutMoveTask(ctx, name, taskUid, parentTaskUid, optional)
Move one task to another parent task.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Unique id of the task to be moved. | 
  **parentTaskUid** | **int32**| Unique id of the new parent task. | 
 **optional** | ***PutMoveTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document.  | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMoveTaskToSibling**
> AsposeResponse PutMoveTaskToSibling(ctx, name, taskUid, beforeTaskUid, optional)
Move a task to another position under the same parent and the same outline level

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Task Unique Id. | 
  **beforeTaskUid** | **int32**| Unique Id of the task after which the current task will be placed. | 
 **optional** | ***PutMoveTaskToSiblingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileName** | **optional.String**| The name of the project document to save changes to.             If this parameter is omitted then the changes will be saved to the source project document. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTask**
> TaskResponse PutTask(ctx, name, taskUid, task, optional)
Updates special task getting by task UID

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Task UID | 
  **task** | [**Task**](Task.md)| TaskDTO | 
 **optional** | ***PutTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mode** | **optional.String**| CalculationModeDTO | [default to 0]
 **recalculate** | **optional.Bool**| boolean value  | [default to true]
 **storage** | **optional.String**| The document strorage | 
 **folder** | **optional.String**| The document folder | 
 **fileName** | **optional.String**| The name of the file to save changes | 

### Return type
[**TaskResponse**](TaskResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTaskRecurringInfo**
> AsposeResponse PutTaskRecurringInfo(ctx, name, taskUid, recurringInfo, optional)
Updates existing task's recurring info. Note that task should be already recurring.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUid** | **int32**| Task Uid. | 
  **recurringInfo** | [**RecurringInfo**](RecurringInfo.md)| A modified DTO of task&#39;s recurring info. | 
 **optional** | ***PutTaskRecurringInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileName** | **optional.String**| File name to save changes to. | 
 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaskLink**
> AsposeResponse DeleteTaskLink(ctx, name, index, optional)
Delete task link.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **index** | **int32**| Index of the task link object. See TaskLink.Index property. | 
 **optional** | ***DeleteTaskLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskLinks**
> TaskLinksResponse GetTaskLinks(ctx, name, optional)
Get tasks' links of a project.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetTaskLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**TaskLinksResponse**](TaskLinksResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTaskLink**
> AsposeResponse PostTaskLink(ctx, name, taskLink, optional)
Adds a new task link to a project.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskLink** | [**TaskLink**](TaskLink.md)| The TaskLink object representation that should be added. | 
 **optional** | ***PostTaskLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTaskLink**
> TaskLinkResponse PutTaskLink(ctx, name, index, taskLink, optional)
Updates existing task link.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **index** | **int32**| Index of the task link object. See TaskLink.Index property. | 
  **taskLink** | [**TaskLink**](TaskLink.md)| The representation of the modified TaskLink object. | 
 **optional** | ***PutTaskLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.  If this parameter is omitted then the changes will be saved to the source project document. | 

### Return type
[**TaskLinkResponse**](TaskLinkResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVbaProject**
> VbaProjectResponse GetVbaProject(ctx, name, optional)
Returns VBA project.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file | 
 **optional** | ***GetVbaProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folder** | **optional.String**| The folder storage | 
 **storage** | **optional.String**| The document storage. | 

### Return type
[**VbaProjectResponse**](VbaProjectResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWbsDefinition**
> WbsDefinitionResponse GetWbsDefinition(ctx, name, optional)
Get a project's WBS Definition.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
 **optional** | ***GetWbsDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storage** | **optional.String**| The document storage. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**WbsDefinitionResponse**](WBSDefinitionResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRenumberWbsCode**
> AsposeResponse PutRenumberWbsCode(ctx, name, taskUids, optional)
Renumber WBS code of passed tasks (if specified) or all project's tasks.

### Required Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the file. | 
  **taskUids** | **[]int32**| Uids of task for WBS codes renumbering. If not specified, WBS codes for all tasks will be renumbered. | 
 **optional** | ***PutRenumberWbsCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storage** | **optional.String**| The document storage. | 
 **fileName** | **optional.String**| The name of the project document to save changes to.              If this parameter is omitted then the changes will be saved to the source project document. | 
 **folder** | **optional.String**| The document folder. | 

### Return type
[**AsposeResponse**](AsposeResponse.md)

### Authorization
[JWT](../README.md#JWT)

### HTTP request headers
 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

