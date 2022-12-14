/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="primavera_task_properties.go">
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

// Represents Primavera-specific properties for a task read from Primavera format (XER of P6XML).
type PrimaveraTaskProperties struct {
	// The sequence number of the WBS item (summary tasks). It is used to sort summary tasks in Primavera.
	SequenceNumber int32 `json:"sequenceNumber"`
	// Activity id field - a task's unique identifier used by Primavera.
	ActivityId string `json:"activityId,omitempty"`
	// Remaining early finish date - the date when the remaining work for the activity is scheduled to be finished.
	RemainingEarlyFinish custom.TimeWithoutTZ `json:"remainingEarlyFinish"`
	// Remaining early start date - the date when the remaining work for the activity is scheduled to begin.
	RemainingEarlyStart custom.TimeWithoutTZ `json:"remainingEarlyStart"`
	// Remaining late start date.
	RemainingLateStart custom.TimeWithoutTZ `json:"remainingLateStart"`
	// Remaining late finish date.
	RemainingLateFinish custom.TimeWithoutTZ `json:"remainingLateFinish"`
	// Raw text representation (as in source file) of 'Duration Type' field of the activity.
	RawDurationType string `json:"rawDurationType,omitempty"`
	// Raw text representation (as in source file) of 'Activity Type' field of the activity.
	RawActivityType string `json:"rawActivityType,omitempty"`
	// Raw text representation (as in source file) of '% Complete Type' field of the activity.
	RawCompletePercentType string `json:"rawCompletePercentType,omitempty"`
	// Raw text representation (as in source file) of 'Status' field of the activity.
	RawStatus string `json:"rawStatus,omitempty"`
}
