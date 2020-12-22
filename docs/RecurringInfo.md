# RecurringInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecurrencePattern** | [***RecurrencePattern**](RecurrencePattern.md) | Represents a recurrence pattern of the recurring task. Can be one of the values of  enum. | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Specifies the date for the occurrences to begin. | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | Specifies the date for the occurrences to end. | [default to null]
**Duration** | **string** | Specifies the duration for one occurrence of the recurring task. the instance of  class. | [default to null]
**Occurrences** | **int32** | Specifies a number of occurrences of the recurring task. | [default to null]
**UseEndDate** | **bool** | Determines whether to use the end date or a number of occurrences for the recurring task. | [default to null]
**DailyRepetitions** | **int32** | Specifies an interval between repetitions for the daily recurrence pattern. | [default to null]
**DailyUseWorkdays** | **bool** | Determines whether to use workdays for the daily recurrence pattern. | [default to null]
**WeeklyRepetitions** | **int32** | Specifies an interval between repetitions for the weekly recurrence pattern. | [default to null]
**WeeklyDays** | [***WeekDayType**](WeekDayType.md) | Specifies a collection of days used in the weekly recurrence pattern. | [default to null]
**MonthlyUseOrdinalDay** | **bool** | Determines whether to use ordinal day for the monthly recurrence pattern. | [default to null]
**MonthlyOrdinalNumber** | [***OrdinalNumber**](OrdinalNumber.md) | Specifies an ordinal number of the monthly recurrence pattern. Can be one of the values of  enum. | [default to null]
**MonthlyOrdinalDay** | [***DayOfWeek**](DayOfWeek.md) | Specifies a day of the monthly recurrence pattern when using ordinal day. Can be one of the values of  enum. | [default to null]
**MonthlyOrdinalRepetitions** | **int32** | Specifies a number of repetitions for the monthly recurrence pattern when using ordinal day. | [default to null]
**MonthlyDay** | **int32** | Specifies a number of day of the monthly recurrence pattern. | [default to null]
**MonthlyRepetitions** | **int32** | Specifies a number of repetitions for the monthly recurrence pattern. | [default to null]
**YearlyUseOrdinalDay** | **bool** | Determines whether to use ordinal day for the yearly recurrence pattern. | [default to null]
**YearlyDate** | [**time.Time**](time.Time.md) | Specifies a date for the yearly recurrence pattern. | [default to null]
**YearlyOrdinalNumber** | [***OrdinalNumber**](OrdinalNumber.md) | Specifies an ordinal number of the yearly recurrence pattern. Can be one of the values of  enum. | [default to null]
**YearlyOrdinalDay** | [***DayOfWeek**](DayOfWeek.md) | Specifies a weekday of the yearly recurrence pattern when using ordinal day. Can be one of the values of  enum. | [default to null]
**YearlyOrdinalMonth** | [***Month**](Month.md) | Specifies a month of the yearly recurrence pattern when using ordinal day. Can be one of the values of  enum. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


