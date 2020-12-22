/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="extended_attribute_definition.go">
 *   Copyright (c) 2021 Aspose.Tasks Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------
 */

package models

// Extended attribute definition's brief into.
type ExtendedAttributeDefinition struct {
	// Corresponds to the Pid of a custom field.
	FieldId string `json:"fieldId,omitempty"`
	// The name of a custom field.
	FieldName string `json:"fieldName,omitempty"`
	// The type of a custom field.
	CfType *CustomFieldType `json:"cfType"`
	// The Guid of a custom field.
	Guid string `json:"guid,omitempty"`
	// Determines whether the extended attribute is associated with a task or a resource
	ElementType *ElementType `json:"elementType"`
	// The maximum number of values you can set in a picklist.
	MaxMultiValues int32 `json:"maxMultiValues"`
	// Determines whether a custom field is user defined.
	UserDef bool `json:"userDef"`
	// The alias of a custom field.
	Alias string `json:"alias,omitempty"`
	// The secondary Pid of a custom field.
	SecondaryPid string `json:"secondaryPid,omitempty"`
	// Determines whether an automatic rolldown to assignments is enabled.
	AutoRollDown bool `json:"autoRollDown"`
	// The Guid of the default lookup table entry.
	DefaultGuid string `json:"defaultGuid,omitempty"`
	// The Guid of the lookup table associated with a custom field.
	LookupUid string `json:"lookupUid,omitempty"`
	// The phonetic pronunciation of the alias of a custom field.
	PhoneticsAlias string `json:"phoneticsAlias,omitempty"`
	// The way rollups are calculated.
	RollupType *RollupType `json:"rollupType"`
	// Determines whether rollups are calculated for a task and group summary rows.
	CalculationType *CalculationType `json:"calculationType"`
	// The formula that Microsoft Project uses to populate a custom task field.
	Formula string `json:"formula,omitempty"`
	// Determines whether only values in the list are allowed in a file.
	RestrictValues bool `json:"restrictValues"`
	// The way value lists are sorted. Values are: 0=Descending, 1=Ascending.
	ValuelistSortOrder int32 `json:"valuelistSortOrder"`
	// Determines whether new values added to a project are automatically added to the list.
	AppendNewValues bool `json:"appendNewValues"`
	// The default value in the list.
	Default_ string `json:"default,omitempty"`
	// Returns list of Extended Attribute Values.
	ValueList []Value `json:"valueList,omitempty"`
	// Secondary guid of extended attribute.
	SecondaryGuid string `json:"secondaryGuid,omitempty"`
}
