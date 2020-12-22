/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="resource.go">
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

// Represents a project resource.
type Resource struct {
	// The name of a resource.
	Name string `json:"name,omitempty"`
	// The unique identifier of a resource.
	Uid int32 `json:"uid"`
	// The position identifier of a resource within the list of resources.
	Id int32 `json:"id"`
	// Contains the generated unique identification code for the resource.             
	Guid string `json:"guid,omitempty"`
	// The type of a resource.
	Type_ *ResourceType `json:"type"`
	// Determines whether a resource is null.
	IsNull bool `json:"isNull"`
	// The initials of a resource.
	Initials string `json:"initials,omitempty"`
	// The phonetic spelling of the resource name. For use with Japanese only.
	Phonetics string `json:"phonetics,omitempty"`
	// The NT account associated with a resource.
	NtAccount string `json:"ntAccount,omitempty"`
	// The NT account associated with a resource.
	WindowsUserAccount string `json:"windowsUserAccount,omitempty"`
	// The type of a workgroup to which a resource belongs.  type.
	Workgroup *WorkGroupType `json:"workgroup"`
	// The unit of measure for the material resource. Read/write String.
	MaterialLabel string `json:"materialLabel,omitempty"`
	// The code or other information about a resource.
	Code string `json:"code,omitempty"`
	// The group to which a resource belongs.
	Group string `json:"group,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	// The title of a hyperlink associated with a resource.
	Hyperlink string `json:"hyperlink,omitempty"`
	// The hyperlink associated with a resource.
	HyperlinkAddress string `json:"hyperlinkAddress,omitempty"`
	// The document bookmark of a hyperlink associated with a resource. Read/write String.
	HyperlinkSubAddress string `json:"hyperlinkSubAddress,omitempty"`
	// The maximum number of units of a resource is available.
	MaxUnits float64 `json:"maxUnits"`
	// The largest number of units assigned to a resource at any time.
	PeakUnits float64 `json:"peakUnits"`
	OverAllocated bool `json:"overAllocated"`
	// The first date when a resource is available.
	AvailableFrom custom.TimeWithoutTZ `json:"availableFrom"`
	// The last date when a resource is available.
	AvailableTo custom.TimeWithoutTZ `json:"availableTo"`
	// The scheduled start date of a resource.
	Start custom.TimeWithoutTZ `json:"start"`
	// The scheduled finish date of a resource.
	Finish custom.TimeWithoutTZ `json:"finish"`
	// Determines whether a resource can be leveled.
	CanLevel bool `json:"canLevel"`
	// Determines how cost is accrued against the resource.
	AccrueAt *CostAccrualType `json:"accrueAt"`
	// The total work assigned to a resource across all assigned tasks.
	Work *string `json:"work,omitempty"`
	// The amount of non-overtime work assigned to a resource.
	RegularWork *string `json:"regularWork,omitempty"`
	// The amount of overtime work assigned to a resource.
	OvertimeWork *string `json:"overtimeWork,omitempty"`
	// The amount of actual work performed by a resource.
	ActualWork *string `json:"actualWork,omitempty"`
	// The amount of remaining work required to complete all assigned tasks.
	RemainingWork *string `json:"remainingWork,omitempty"`
	// The amount of actual overtime work performed by a resource.
	ActualOvertimeWork *string `json:"actualOvertimeWork,omitempty"`
	// The amount of remaining overtime work required to complete all tasks.
	RemainingOvertimeWork *string `json:"remainingOvertimeWork,omitempty"`
	// The percentage of work completed across all tasks.
	PercentWorkComplete int32 `json:"percentWorkComplete"`
	// The standard rate of a resource. This value retrieved from the current date if a rate table exists for a resource.
	StandardRate float32 `json:"standardRate"`
	// The units used by Microsoft Project to display the standard rate.
	StandardRateFormat *RateFormatType `json:"standardRateFormat"`
	// The total project cost for a resource across all assigned tasks.
	Cost float32 `json:"cost"`
	// The units used by Microsoft Project to display the overtime rate.
	OvertimeRateFormat *RateFormatType `json:"overtimeRateFormat"`
	// The total overtime cost of a resource including actual and remaining overtime costs.
	OvertimeCost float32 `json:"overtimeCost"`
	// The cost per use of a resource. This value retrieved from the current date if a rate table exists for the resource.
	CostPerUse float32 `json:"costPerUse"`
	// The actual cost incurred by the resource across all assigned tasks.
	ActualCost float32 `json:"actualCost"`
	// The actual overtime cost incurred by the resource across all assigned tasks.
	ActualOvertimeCost float32 `json:"actualOvertimeCost"`
	// The remaining projected cost of a resource to complete all assigned tasks.
	RemainingCost float32 `json:"remainingCost"`
	// The remaining projected overtime cost of a resource to complete all assigned tasks.
	RemainingOvertimeCost float32 `json:"remainingOvertimeCost"`
	// The difference between a baseline work and a work
	WorkVariance float64 `json:"workVariance"`
	// The difference between a baseline cost and a cost.
	CostVariance float64 `json:"costVariance"`
	// The earned value schedule variance, through the project status date.
	Sv float64 `json:"sv"`
	// The earned value cost variance, through the project status date.
	Cv float64 `json:"cv"`
	// The actual cost of a work performed by a resource for the project to-date.
	Acwp float64 `json:"acwp"`
	// The calendar of a resource.
	CalendarUid int32 `json:"calendarUid"`
	// Notes' plain text extracted from RTF data.
	NotesText string `json:"notesText,omitempty"`
	// The text notes associated with a resource.
	Notes string `json:"notes,omitempty"`
	// The text notes in RTF format. Supported for MPP formats only.
	NotesRTF string `json:"notesRTF,omitempty"`
	// The budget cost of a work scheduled for a resource.
	Bcws float64 `json:"bcws"`
	// The budgeted cost of a work performed by a resource for the project to-date.
	Bcwp float64 `json:"bcwp"`
	// Determines whether a resource is generic.
	IsGeneric bool `json:"isGeneric"`
	// Determines whether a resource is inactive.
	IsInactive bool `json:"isInactive"`
	// Determines whether a resource is an Enterprise resource.
	IsEnterprise bool `json:"isEnterprise"`
	// The booking type of a resource.
	BookingType *BookingType `json:"bookingType"`
	// The duration through which actual work is protected.
	ActualWorkProtected *string `json:"actualWorkProtected,omitempty"`
	// The duration through which actual overtime work is protected.
	ActualOvertimeWorkProtected *string `json:"actualOvertimeWorkProtected,omitempty"`
	// The Active Directory Guid for a resource.
	ActiveDirectoryGuid string `json:"activeDirectoryGuid,omitempty"`
	// The date when a resource was created.
	CreationDate custom.TimeWithoutTZ `json:"creationDate"`
	// Indicates which cost center the costs accrued by the resource should be charged to.
	CostCenter string `json:"costCenter,omitempty"`
	// Determines whether a resource is a cost resource.
	IsCostResource bool `json:"isCostResource"`
	// Determines whether the current resource is a team resource.             
	TeamAssignmentPool bool `json:"teamAssignmentPool"`
	// The name of an assignment owner.
	AssignmentOwner string `json:"assignmentOwner,omitempty"`
	// The GUID of an assignment owner.
	AssignmentOwnerGuid string `json:"assignmentOwnerGuid,omitempty"`
	// Determines whether a resource is a budget resource.
	IsBudget bool `json:"isBudget"`
	// The budget work for a budget work or material resource.
	BudgetWork *string `json:"budgetWork,omitempty"`
	// The budget cost for a budget cost resource.
	BudgetCost float32 `json:"budgetCost"`
	// The overtime rate of a resource. This value retrieved from the current date if a rate table exists for a resource.
	OvertimeRate float32 `json:"overtimeRate"`
	// Gets or sets the collection of baseline values of the resource.
	Baselines []Baseline `json:"baselines,omitempty"`
	// Resource extended attributes.
	ExtendedAttributes []ExtendedAttribute `json:"extendedAttributes,omitempty"`
	// Resource outline codes.
	OutlineCodes []OutlineCode `json:"outlineCodes,omitempty"`
}
