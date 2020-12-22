/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="outline_code_definition.go">
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

// Represents an outline code definition.
type OutlineCodeDefinition struct {
	// The Guid of an outline code.
	Guid string `json:"guid,omitempty"`
	// Corresponds to the field number of an outline code.
	FieldId string `json:"fieldId,omitempty"`
	// The name of a custom outline code.
	FieldName string `json:"fieldName,omitempty"`
	// The alias of a custom outline code.
	Alias string `json:"alias,omitempty"`
	// The phonetic pronunciation of the alias of the custom outline code.
	PhoneticAlias string `json:"phoneticAlias,omitempty"`
	// Returns List&lt;OutlineValue&gt; Values. The values of the table associated with this outline code.
	Values []OutlineValue `json:"values,omitempty"`
	// Determines whether a custom outline code is an enterprise custom outline code.
	Enterprise bool `json:"enterprise"`
	// The reference to another custom field for which this outline code definition is an alias.
	EnterpriseOutlineCodeAlias int32 `json:"enterpriseOutlineCodeAlias"`
	// Determines whether the custom outline code can be used by the Resource Substitution Wizard in Microsoft Project.
	ResourceSubstitutionEnabled bool `json:"resourceSubstitutionEnabled"`
	// Determines whether the values specified in this outline code field must be leaf values.
	LeafOnly bool `json:"leafOnly"`
	// Determines whether the new codes must have all levels. Not available for Enterprise Codes.
	AllLevelsRequired bool `json:"allLevelsRequired"`
	// Determines whether the values specified must come from values table.
	OnlyTableValuesAllowed bool `json:"onlyTableValuesAllowed"`
	// Returns List&lt;OutlineMask&gt; Masks. The table of entries that define the outline code mask.
	Masks []OutlineMask `json:"masks,omitempty"`
	// Determines whether the indents of this outline code must be shown.
	ShowIndent bool `json:"showIndent"`
}
