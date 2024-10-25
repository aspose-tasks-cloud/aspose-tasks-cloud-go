/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="view.go">
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

// Represents a view in Project
type View struct {
	// Gets or sets a value indicating whether Microsoft Project shows the single view  name in the View or the Other Views drop-down lists in the Ribbon
	ShowInMenu bool `json:"showInMenu"`
	// Gets the type of item in the single view, such as tasks or resources. Read-only.
	Type_ *ItemType `json:"type"`
	// Gets the screen type for the single view. Read-only.
	Screen *ViewScreen `json:"screen"`
	// Gets or sets the name of a view object.
	Name string `json:"name,omitempty"`
	// Gets the unique identifier of a view.
	Uid int32 `json:"uid"`
}
