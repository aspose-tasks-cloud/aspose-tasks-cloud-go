/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="project_info.go">
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

// Brief info about the published project available on Project Online.
type ProjectInfo struct {
	// The unique identifier of the project.
	Id string `json:"id"`
	// The name of the project.
	Name string `json:"name,omitempty"`
	// The date and time when the project was created.
	CreatedDate custom.TimeWithoutTZ `json:"createdDate"`
	// Value indicating whether the project is checked out.
	IsCheckedOut bool `json:"isCheckedOut"`
	// The most recent date when the project was published.
	LastPublishedDate custom.TimeWithoutTZ `json:"lastPublishedDate"`
	// The most recent date when the project was saved.
	LastSavedDate custom.TimeWithoutTZ `json:"lastSavedDate"`
	// The description of the project.
	Description string `json:"description,omitempty"`
}
