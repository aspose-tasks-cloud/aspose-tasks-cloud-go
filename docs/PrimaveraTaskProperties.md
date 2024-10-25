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
**DurationPercentComplete** | **float64** | Gets the value of duration percent complete. | [default to null]
**PhysicalPercentComplete** | **float64** | Gets the value of Physical Percent Complete. | [default to null]
**ActualNonLaborUnits** | **float64** | Gets the value of actual non labor units. | [default to null]
**ActualLaborUnits** | **float64** | Gets the value of actual labor units. | [default to null]
**UnitsPercentComplete** | **float64** | Gets the value of units percent complete. | [default to null]
**RemainingLaborUnits** | **float64** | Gets the value of remaining labor units. | [default to null]
**RemainingNonLaborUnits** | **float64** | Gets the value of remaining non labor units. | [default to null]
**DurationType** | [***PrimaveraDurationType**](PrimaveraDurationType.md) | Gets the value of &#39;Duration Type&#39; field of the activity. | [default to null]
**ActivityType** | [***PrimaveraActivityType**](PrimaveraActivityType.md) | Gets the value of &#39;Activity Type&#39; field. | [default to null]
**PercentCompleteType** | [***PrimaveraPercentCompleteType**](PrimaveraPercentCompleteType.md) | Gets the value of &#39;% Complete Type&#39; field of the activity. | [default to null]
**ActualLaborCost** | **float32** | Gets the value of actual labor cost. | [default to null]
**ActualNonlaborCost** | **float32** | Gets the value of actual non labor cost. | [default to null]
**ActualMaterialCost** | **float32** | Gets the value of actual material cost.              | [default to null]
**ActualExpenseCost** | **float32** | Gets the value of actual expense cost. | [default to null]
**RemainingExpenseCost** | **float32** | Gets the value of remaining expense cost. | [default to null]
**ActualTotalCost** | **float32** | Gets the total value of actual costs. | [default to null]
**BudgetedTotalCost** | **float32** | Gets the total value of budgeted (or planned) costs. | [default to null]
**BudgetedLaborCost** | **float32** | Gets the value of budgeted (or planned) labor cost. | [default to null]
**BudgetedNonlaborCost** | **float32** | Gets the value of budgeted (or planned) non labor cost. | [default to null]
**BudgetedMaterialCost** | **float32** | Gets the value of of budgeted (or planned) material cost. | [default to null]
**BudgetedExpenseCost** | **float32** | Gets the value of budgeted (or planned) expense cost. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


