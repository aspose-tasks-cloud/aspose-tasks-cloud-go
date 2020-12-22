# AssignmentBaseline

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaselineNumber** | [***BaselineType**](BaselineType.md) | The unique number of a baseline data record. | [default to null]
**Work** | **string** | The work assigned to a resource when the baseline is saved. | [default to null]
**Cost** | **float32** | The projected cost of a resource when the baseline is saved. | [default to null]
**Bcws** | **float64** | The budget cost of a work scheduled for a resource. | [default to null]
**Bcwp** | **float64** | The budgeted cost of a work performed by a resource for a project to-date. | [default to null]
**Start** | [**time.Time**](time.Time.md) | Gets or sets the scheduled start date of the resource assignment when the baseline was saved. The start date of the resource assignment when this baseline was saved. | [optional] [default to null]
**Finish** | [**time.Time**](time.Time.md) | Gets or sets the scheduled finish date of the resource assignment when the baseline was saved. The finish date of the resource assignment when this baseline was saved. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


