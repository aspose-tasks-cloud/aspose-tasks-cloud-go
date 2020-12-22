/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="file_version.go">
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

// File Version
type FileVersion struct {
	// File or folder name.
	Name string `json:"name,omitempty"`
	// True if it is a folder.
	IsFolder bool `json:"isFolder"`
	// File or folder last modified DateTime.
	ModifiedDate *custom.TimeWithoutTZ `json:"modifiedDate,omitempty""`
	// File or folder size.
	Size int64 `json:"size"`
	// File or folder path.
	Path string `json:"path,omitempty"`
	// File Version ID.
	VersionId string `json:"versionId,omitempty"`
	// Specifies whether the file is (true) or is not (false) the latest version of an file.
	IsLatest bool `json:"isLatest"`
}
