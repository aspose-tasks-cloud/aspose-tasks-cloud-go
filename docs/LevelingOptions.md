# LevelingOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | [**time.Time**](time.Time.md) | Leveling period start date. The default value is the project&#x60;s start date. | [optional] [default to null]
**FinishDate** | [**time.Time**](time.Time.md) | Leveling period end date. The default value is the project&#x60;s finish date. | [optional] [default to null]
**ResourceUids** | **[]int32** | The list of the resource uids which will be leveled. If null is set,  all project resources will be leveled. | [optional] [default to null]
**LevelingOrder** | [***LevelingOrder**](LevelingOrder.md) | The order in which the leveling algorithm delays tasks that have overallocations. After determination of tasks causing the overallocation and which tasks can be delayed, the specified order is used which task should be delayed first. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


