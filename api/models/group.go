/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="group.go">
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

// Represents a group definition. A Group object is a member of the ResourceGroups collection or the TaskGroups collection.
type Group struct {
	// Gets or sets a unique identifier of a group.
	Uid int32 `json:"uid"`
	// Gets or sets a name of a Group object.
	Name string `json:"name,omitempty"`
	// Gets or sets a value indicating whether Project shows the group name in the Group drop-down list in the Ribbon.
	ShowInMenu bool `json:"showInMenu"`
	// Gets or sets a value indicating whether summary rows are displayed for the group.
	ShowSummary bool `json:"showSummary"`
	// Gets or sets a value indicating whether to show all the levels of summary tasks for subtasks within group.
	MaintainHierarchy bool `json:"maintainHierarchy"`
	// Gets or sets a value indicating whether assignments should be grouped instead of tasks.
	GroupAssignments bool `json:"groupAssignments"`
	// Gets or sets a collection of criteria representing the fields in a group definition.
	GroupCriteria []GroupCriterion `json:"groupCriteria,omitempty"`
}
