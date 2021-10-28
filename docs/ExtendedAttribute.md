# ExtendedAttribute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** | Gets or sets the id of a field. | [optional] [default to null]
**AttributeType** | [***CustomFieldType**](CustomFieldType.md) | Gets the type of a custom field. | [default to null]
**ValueGuid** | **string** | Gets or sets the guid of a value. | [optional] [default to null]
**LookupValueId** | **int32** | Gets or sets Id of the lookup value (if value is lookup value) | [optional] [default to null]
**DurationValue** | [***Duration**](Duration.md) | Gets or sets value for attributes with &#39;Duration&#39; type. | [optional] [default to null]
**NumericValue** | **float32** | Gets or sets a value for attributes with numeric types (Cost, Number). | [default to null]
**DateValue** | [**time.Time**](time.Time.md) | Gets or sets a value for attributes with date types (Date, Start, Finish). | [default to null]
**FlagValue** | **bool** | Gets or sets a value indicating whether a flag is set for an attribute with &#39;Flag&#39; type. | [default to null]
**TextValue** | **string** | Gets or sets a value for attributes with &#39;Text&#39; type. | [optional] [default to null]
**IsErrorValue** | **bool** | Gets whether calculation of extended attribute&#39;s value resulted in an error.              | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


