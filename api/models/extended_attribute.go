/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="extended_attribute.go">
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

import (
	"github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/custom"
)

// Represents extended attribute.
type ExtendedAttribute struct {
	// Gets or sets the id of a field.
	FieldId string `json:"fieldId,omitempty"`
	// Gets the type of a custom field.
	AttributeType *CustomFieldType `json:"attributeType"`
	// Gets or sets the guid of a value.
	ValueGuid string `json:"valueGuid,omitempty"`
	// Gets or sets Id of the lookup value (if value is lookup value)
	LookupValueId int32 `json:"lookupValueId,omitempty"`
	// Gets or sets value for attributes with 'Duration' type.
	DurationValue *Duration `json:"durationValue,omitempty"`
	// Gets or sets a value for attributes with numeric types (Cost, Number).
	NumericValue float32 `json:"numericValue"`
	// Gets or sets a value for attributes with date types (Date, Start, Finish).
	DateValue custom.TimeWithoutTZ `json:"dateValue"`
	// Gets or sets a value indicating whether a flag is set for an attribute with 'Flag' type.
	FlagValue bool `json:"flagValue"`
	// Gets or sets a value for attributes with 'Text' type.
	TextValue string `json:"textValue,omitempty"`
	// Gets whether calculation of extended attribute's value resulted in an error.             
	IsErrorValue bool `json:"isErrorValue"`
}
