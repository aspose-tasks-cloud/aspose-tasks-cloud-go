# ExtendedAttributeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** | Corresponds to the Pid of a custom field. | [optional] [default to null]
**FieldName** | **string** | The name of a custom field. | [optional] [default to null]
**CfType** | [***CustomFieldType**](CustomFieldType.md) | The type of a custom field. | [default to null]
**Guid** | **string** | The Guid of a custom field. | [optional] [default to null]
**ElementType** | [***ElementType**](ElementType.md) | Determines whether the extended attribute is associated with a task or a resource | [default to null]
**MaxMultiValues** | **int32** | The maximum number of values you can set in a picklist. | [default to null]
**UserDef** | **bool** | Determines whether a custom field is user defined. | [default to null]
**Alias** | **string** | The alias of a custom field. | [optional] [default to null]
**SecondaryPid** | **string** | The secondary Pid of a custom field. | [optional] [default to null]
**AutoRollDown** | **bool** | Determines whether an automatic rolldown to assignments is enabled. | [default to null]
**DefaultGuid** | **string** | The Guid of the default lookup table entry. | [optional] [default to null]
**LookupUid** | **string** | The Guid of the lookup table associated with a custom field. | [optional] [default to null]
**PhoneticsAlias** | **string** | The phonetic pronunciation of the alias of a custom field. | [optional] [default to null]
**RollupType** | [***RollupType**](RollupType.md) | The way rollups are calculated. | [default to null]
**CalculationType** | [***CalculationType**](CalculationType.md) | Determines whether rollups are calculated for a task and group summary rows. | [default to null]
**Formula** | **string** | The formula that Microsoft Project uses to populate a custom task field. | [optional] [default to null]
**RestrictValues** | **bool** | Determines whether only values in the list are allowed in a file. | [default to null]
**ValuelistSortOrder** | **int32** | The way value lists are sorted. Values are: 0&#x3D;Descending, 1&#x3D;Ascending. | [default to null]
**AppendNewValues** | **bool** | Determines whether new values added to a project are automatically added to the list. | [default to null]
**Default_** | **string** | The default value in the list. | [optional] [default to null]
**ValueList** | [**[]Value**](Value.md) | Returns list of Extended Attribute Values. | [optional] [default to null]
**SecondaryGuid** | **string** | Secondary guid of extended attribute. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


