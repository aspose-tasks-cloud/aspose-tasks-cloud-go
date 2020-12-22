# ProjectServerSaveOptionsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | Gets or sets name of a project which is displayed in Project Server \\ Project     Online projects list. Should be unique within Project Server \\ Project Online     instance. Is the value is omitted, the value of Prj.Name property will be used     instead. | [optional] [default to null]
**ProjectGuid** | **string** | Gets or sets unique identifier of a project. Should be unique within Project     Server \\ Project Online instance. | [optional] [default to null]
**Timeout** | **string** | Gets or sets timeout used when waiting for processing of save project request     by a Project Server&#39;s queue processing service. The default value for this property     is 1 minute. The processing time may be longer for large projects or in case when Project     Server instance is too busy responding to other requests. | [default to null]
**PollingInterval** | **string** | Gets or sets interval between queue job status requests. The default value is     2 seconds. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


