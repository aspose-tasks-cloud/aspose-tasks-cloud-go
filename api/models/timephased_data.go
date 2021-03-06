/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="timephased_data.go">
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

// Represents a time phased data.
type TimephasedData struct {
	// The unique identifier of a timephased data
	Uid int32 `json:"uid"`
	// The start date of a timephased data period.
	Start custom.TimeWithoutTZ `json:"start"`
	// The finish date of a timephased data period.
	Finish custom.TimeWithoutTZ `json:"finish"`
	// The time unit of a timephased data period.
	Unit *TimeUnitType `json:"unit"`
	// The value per unit of time for a timephased data period.
	Value string `json:"value,omitempty"`
	// The type of a timephased data.
	TimephasedDataType *TimephasedDataType `json:"timephasedDataType"`
}
