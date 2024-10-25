# Calendar

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | **string** | Gets calendar&#39;s Guid. | [optional] [default to null]
**Name** | **string** | The name of the calendar. | [optional] [default to null]
**Uid** | **int32** | The unique identifier of the calendar. | [default to null]
**Days** | [**[]WeekDay**](WeekDay.md) | The collection of weekdays that defines the calendar. | [optional] [default to null]
**IsBaseCalendar** | **bool** | Determines whether the calendar is a base calendar. | [default to null]
**BaseCalendar** | [***Calendar**](Calendar.md) | The base calendar on which this calendar depends. Only applicable if the calendar is not a base calendar. | [optional] [default to null]
**IsBaselineCalendar** | **bool** | Specifies whether the calendar is a baseline calendar. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


