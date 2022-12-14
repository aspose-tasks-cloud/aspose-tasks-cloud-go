# PrimaveraTaskProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceNumber** | **int32** | The sequence number of the WBS item (summary tasks). It is used to sort summary tasks in Primavera. | [default to null]
**ActivityId** | **string** | Activity id field - a task&#39;s unique identifier used by Primavera. | [optional] [default to null]
**RemainingEarlyFinish** | [**time.Time**](time.Time.md) | Remaining early finish date - the date when the remaining work for the activity is scheduled to be finished. | [default to null]
**RemainingEarlyStart** | [**time.Time**](time.Time.md) | Remaining early start date - the date when the remaining work for the activity is scheduled to begin. | [default to null]
**RemainingLateStart** | [**time.Time**](time.Time.md) | Remaining late start date. | [default to null]
**RemainingLateFinish** | [**time.Time**](time.Time.md) | Remaining late finish date. | [default to null]
**RawDurationType** | **string** | Raw text representation (as in source file) of &#39;Duration Type&#39; field of the activity. | [optional] [default to null]
**RawActivityType** | **string** | Raw text representation (as in source file) of &#39;Activity Type&#39; field of the activity. | [optional] [default to null]
**RawCompletePercentType** | **string** | Raw text representation (as in source file) of &#39;% Complete Type&#39; field of the activity. | [optional] [default to null]
**RawStatus** | **string** | Raw text representation (as in source file) of &#39;Status&#39; field of the activity. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


