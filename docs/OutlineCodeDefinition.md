# OutlineCodeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | **string** | The Guid of an outline code. | [optional] [default to null]
**FieldId** | **string** | Corresponds to the field number of an outline code. | [optional] [default to null]
**FieldName** | **string** | The name of a custom outline code. | [optional] [default to null]
**Alias** | **string** | The alias of a custom outline code. | [optional] [default to null]
**PhoneticAlias** | **string** | The phonetic pronunciation of the alias of the custom outline code. | [optional] [default to null]
**Values** | [**[]OutlineValue**](OutlineValue.md) | Returns List&amp;lt;OutlineValue&amp;gt; Values. The values of the table associated with this outline code. | [optional] [default to null]
**Enterprise** | **bool** | Determines whether a custom outline code is an enterprise custom outline code. | [default to null]
**EnterpriseOutlineCodeAlias** | **int32** | The reference to another custom field for which this outline code definition is an alias. | [default to null]
**ResourceSubstitutionEnabled** | **bool** | Determines whether the custom outline code can be used by the Resource Substitution Wizard in Microsoft Project. | [default to null]
**LeafOnly** | **bool** | Determines whether the values specified in this outline code field must be leaf values. | [default to null]
**AllLevelsRequired** | **bool** | Determines whether the new codes must have all levels. Not available for Enterprise Codes. | [default to null]
**OnlyTableValuesAllowed** | **bool** | Determines whether the values specified must come from values table. | [default to null]
**Masks** | [**[]OutlineMask**](OutlineMask.md) | Returns List&amp;lt;OutlineMask&amp;gt; Masks. The table of entries that define the outline code mask. | [optional] [default to null]
**ShowIndent** | **bool** | Determines whether the indents of this outline code must be shown. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


