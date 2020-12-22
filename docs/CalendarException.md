# CalendarException

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Index of the current item in the collection of calendar&#39;s exceptions. | [default to null]
**EnteredByOccurrences** | **bool** | Determines whether the range of recurrence is defined by entering a number of occurrences. False specifies that the range of recurrence is defined by entering a finish date. | [default to null]
**FromDate** | [**time.Time**](time.Time.md) | The beginning of the exception time. | [default to null]
**ToDate** | [**time.Time**](time.Time.md) | The end of the exception time. | [default to null]
**Occurrences** | **int32** | The number of occurrences for which the calendar exception is valid. | [default to null]
**Name** | **string** | The name of the exception. | [optional] [default to null]
**Type_** | [***CalendarExceptionType**](CalendarExceptionType.md) | The exception type. | [default to null]
**Period** | **int32** | The period of recurrence for the exception. | [default to null]
**DaysOfWeek** | [**[]DayType**](DayType.md) | The days of the week on which the exception is valid. | [optional] [default to null]
**MonthItem** | [***MonthItemType**](MonthItemType.md) | The month item for which an exception recurrence is scheduled. | [default to null]
**MonthPosition** | [***MonthPosition**](MonthPosition.md) | The position of a month item within a month. | [default to null]
**Month** | [***Month**](Month.md) | The month for which an exception recurrence is scheduled. | [default to null]
**MonthDay** | **int32** | The day of a month on which an exception recurrence is scheduled. | [default to null]
**DayWorking** | **bool** | Determines whether the specified date or day type is working. | [default to null]
**WorkingTimes** | [**[]WorkingTime**](WorkingTime.md) | The collection of working times that defines the time worked on the weekday.  At least one working time must present, and there can&#39;t be more than five. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


