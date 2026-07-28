/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="group_criterion.go">
 *   Copyright (c) 2025 Aspose.Tasks Cloud
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

// Represents a criterion in a group definition.
type GroupCriterion struct {
	// Gets or sets the field being grouped by.
	Field *Field `json:"field"`
	// Gets or sets a value indicating whether a field used as a criterion in a group definition is sorted in ascending order. False if the field is sorted in descending order.
	Ascending bool `json:"ascending"`
	// Gets or sets the type of grouping for a field used as a criterion in a group definition.
	GroupOn *GroupOn `json:"groupOn"`
	// Gets or sets the interval for a field used as a criterion in a group definition.
	GroupInterval string `json:"groupInterval,omitempty"`
	// Gets or sets the start of the intervals for a field used as a criterion in a group definition.
	StartAt string `json:"startAt,omitempty"`
	// Gets or sets the color of the cell background for a field used as a criterion in a group definition.
	CellColor *Colors `json:"cellColor"`
	// Gets or sets the color of the font for a field used as a criterion in a group definition.
	FontColor *Colors `json:"fontColor"`
	// Gets or sets the pattern of the cell for a field used as a criterion in a group definition.
	Pattern *BackgroundPattern `json:"pattern"`
	// Gets or sets the font for a criterion in a group definition.
	Font *FontInfo `json:"font,omitempty"`
}
