/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="resource_assignment.go">
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

// Represents a resource assignment in a project.
type ResourceAssignment struct {
	// Returns or sets a task unique id to which a resource is assigned.
	TaskUid int32 `json:"taskUid"`
	// Returns or sets a resource unique id assigned to a task.
	ResourceUid int32 `json:"resourceUid"`
	// Returns or sets the global unique identifier of an assignment.
	Guid string `json:"guid,omitempty"`
	// Returns or sets the unique identifier of an assignment.
	Uid int32 `json:"uid"`
	// Returns or sets the amount of a work completed on an assignment.
	PercentWorkComplete int32 `json:"percentWorkComplete"`
	// Returns or sets the actual cost incurred on an assignment.
	ActualCost float32 `json:"actualCost"`
	// Returns or sets the actual finish date of an assignment.
	ActualFinish custom.TimeWithoutTZ `json:"actualFinish"`
	// Returns or sets the actual overtime cost incurred on an assignment.
	ActualOvertimeCost float32 `json:"actualOvertimeCost"`
	// Returns or sets the actual amount of an overtime work incurred on an assignment.
	ActualOvertimeWork *string `json:"actualOvertimeWork,omitempty"`
	// Returns or sets the actual start date of an assignment.
	ActualStart custom.TimeWithoutTZ `json:"actualStart"`
	// Returns or sets the actual amount of a work incurred on an assignment.
	ActualWork *string `json:"actualWork,omitempty"`
	// Returns or sets the actual cost of a work performed on an assignment to-date.
	Acwp float64 `json:"acwp"`
	// Determines whether a resource has accepted all of its assignments.
	Confirmed bool `json:"confirmed"`
	// Returns or sets the projected or scheduled cost of an assignment.
	Cost float32 `json:"cost"`
	// Returns or sets the cost rate table used for this assignment.
	CostRateTableType *RateType `json:"costRateTableType"`
	// Returns or sets the difference between the cost and baseline cost of a resource.
	CostVariance float64 `json:"costVariance"`
	// Returns or sets the earned value cost variance.
	Cv float64 `json:"cv"`
	// Returns or sets the delay of an assignment.
	Delay int32 `json:"delay"`
	// Returns or sets the scheduled finish date of an assignment.
	Finish custom.TimeWithoutTZ `json:"finish"`
	// Returns or sets the variance of an assignment finish date from a baseline finish date.
	FinishVariance int32 `json:"finishVariance"`
	// Returns or sets the title of the hyperlink associated with an assignment.
	Hyperlink string `json:"hyperlink,omitempty"`
	// Returns or sets the hyperlink associated with an assignment.
	HyperlinkAddress string `json:"hyperlinkAddress,omitempty"`
	// Returns or sets the document bookmark of the hyperlink associated with an assignment.
	HyperlinkSubAddress string `json:"hyperlinkSubAddress,omitempty"`
	// Returns or sets the variance of an assignment work from the baseline work as minutes.
	WorkVariance float64 `json:"workVariance"`
	// Determines whether the Units have Fixed Rate.
	HasFixedRateUnits bool `json:"hasFixedRateUnits"`
	// Determines whether the consumption of an assigned material resource occurs in a single, fixed amount.
	FixedMaterial bool `json:"fixedMaterial"`
	// Returns or sets the delay caused by leveling.
	LevelingDelay int32 `json:"levelingDelay"`
	// Returns or sets the duration format of a delay.
	LevelingDelayFormat *TimeUnitType `json:"levelingDelayFormat"`
	// Determines whether the Project is linked to another OLE object.
	LinkedFields bool `json:"linkedFields"`
	// Determines whether the assignment is a milestone.
	Milestone bool `json:"milestone"`
	// Returns or sets the text notes associated with an assignment.
	Notes string `json:"notes,omitempty"`
	// Determines whether the assignment is overallocated.
	Overallocated bool `json:"overallocated"`
	// Returns or sets the sum of the actual and remaining overtime cost of an assignment.
	OvertimeCost float32 `json:"overtimeCost"`
	// Returns or sets the scheduled overtime work of an assignment.
	OvertimeWork *string `json:"overtimeWork,omitempty"`
	// Returns or sets the largest number of a resource's units assigned to a task.
	PeakUnits float64 `json:"peakUnits"`
	// Returns or sets the amount of a non-overtime work scheduled for an assignment.
	RegularWork *string `json:"regularWork,omitempty"`
	// Returns or sets the remaining projected cost of completing an assignment.
	RemainingCost float32 `json:"remainingCost"`
	// Returns or sets the remaining projected overtime cost of completing an assignment.
	RemainingOvertimeCost float32 `json:"remainingOvertimeCost"`
	// Returns or sets the remaining overtime work scheduled to complete an assignment.
	RemainingOvertimeWork *string `json:"remainingOvertimeWork,omitempty"`
	// Returns or sets the remaining work scheduled to complete an assignment.
	RemainingWork *string `json:"remainingWork,omitempty"`
	// Determines whether the response has been received for a TeamAssign message.
	ResponsePending bool `json:"responsePending"`
	// Returns or sets the scheduled start date of an assignment.
	Start custom.TimeWithoutTZ `json:"start"`
	// Returns or sets the date when assignment is stopped.
	Stop custom.TimeWithoutTZ `json:"stop"`
	// Returns or sets the date when assignment is resumed.
	Resume custom.TimeWithoutTZ `json:"resume"`
	// Returns or sets the variance of an assignment start date from a baseline start date.
	StartVariance int32 `json:"startVariance"`
	// Determines whether the task is a summary task.
	Summary bool `json:"summary"`
	// Returns or sets the earned value schedule variance, through the project status date.
	Sv float64 `json:"sv"`
	// Returns or sets the number of units for an assignment.
	Units float64 `json:"units"`
	// Determines whether the resource assigned to a task needs to be updated as to the status of the task.
	UpdateNeeded bool `json:"updateNeeded"`
	// Returns or sets the difference between basline cost and total cost. Read/write Double.
	Vac float64 `json:"vac"`
	// Returns or sets the amount of scheduled work for an assignment. Read/write TimeSpan.
	Work *string `json:"work,omitempty"`
	// Returns or sets the work contour of an assignment.
	WorkContour *WorkContourType `json:"workContour"`
	// Returns or sets the budgeted cost of a work on assignment.
	Bcws float64 `json:"bcws"`
	// Returns or sets the budgeted cost of a work performed on assignment to-date.
	Bcwp float64 `json:"bcwp"`
	// Returns or sets the booking type of an assignment.
	BookingType *BookingType `json:"bookingType"`
	// Returns or sets the duration through which actual work is protected.
	ActualWorkProtected *string `json:"actualWorkProtected,omitempty"`
	// Returns or sets the duration through which actual overtime work is protected.
	ActualOvertimeWorkProtected *string `json:"actualOvertimeWorkProtected,omitempty"`
	// Returns or sets the date that the assignment was created.
	CreationDate custom.TimeWithoutTZ `json:"creationDate"`
	// Returns or sets the name of an assignment owner.
	AssnOwner string `json:"assnOwner,omitempty"`
	// Returns or sets the Guid of an assignment owner.
	AssnOwnerGuid string `json:"assnOwnerGuid,omitempty"`
	// Returns or sets the budgeted cost of resources on an assignment.
	BudgetCost float32 `json:"budgetCost"`
	// Returns or sets the budgeted work amount for a work or material resources on an assignment.
	BudgetWork *string `json:"budgetWork,omitempty"`
	// Returns the time unit for the usage rate of the material resource assignment.
	RateScale *RateScaleType `json:"rateScale"`
	// List of ResourceAssignment's Baseline values.
	Baselines []AssignmentBaseline `json:"baselines,omitempty"`
	// ResourceAssignment extended attributes.
	ExtendedAttributes []ExtendedAttribute `json:"extendedAttributes,omitempty"`
	// Represents a collection of TimephasedData objects.
	TimephasedData []TimephasedData `json:"timephasedData,omitempty"`
}
