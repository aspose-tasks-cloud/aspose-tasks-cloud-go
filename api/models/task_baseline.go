/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="task_baseline.go">
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

// Represents baseline values of a task.
type TaskBaseline struct {
	// The unique number of a baseline data record.
	BaselineNumber *BaselineType `json:"baselineNumber"`
	// The work assigned to a resource when the baseline is saved.
	Work *string `json:"work,omitempty"`
	// The projected cost of a resource when the baseline is saved.
	Cost float32 `json:"cost"`
	// The budget cost of a work scheduled for a resource.
	Bcws float64 `json:"bcws"`
	// The budgeted cost of a work performed by a resource for a project to-date.
	Bcwp float64 `json:"bcwp"`
	// The scheduled start date of the task when the baseline was saved.
	Start custom.TimeWithoutTZ `json:"start"`
	// The scheduled finish date of the task when the baseline was saved.
	Finish custom.TimeWithoutTZ `json:"finish"`
	// The scheduled duration of the task when the baseline was saved.
	Duration *string `json:"duration,omitempty"`
	// The fixed cost of the task when the baseline was saved.
	FixedCost float64 `json:"fixedCost"`
	// The format for expressing the duration of the task baseline.
	DurationFormat *TimeUnitType `json:"durationFormat"`
	// Indicates whether the baseline duration of the task was estimated.
	EstimatedDuration bool `json:"estimatedDuration"`
}
