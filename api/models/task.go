/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="task.go">
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

// Represents project task.
type Task struct {
	// The unique id of a task.
	Uid int32 `json:"uid"`
	// The position of a task in collection.
	Id int32 `json:"id"`
	// The name of a task.
	Name string `json:"name,omitempty"`
	// The duration of a task entered by the user as a text.
	DurationText string `json:"durationText,omitempty"`
	// The duration of a task.
	Duration *string `json:"duration,omitempty"`
	// The start date of a task.
	Start custom.TimeWithoutTZ `json:"start"`
	// The finish date of a task.
	Finish custom.TimeWithoutTZ `json:"finish"`
	// Returns the task's start text.
	StartText string `json:"startText,omitempty"`
	// Returns the task's finish text.
	FinishText string `json:"finishText,omitempty"`
	// The percent complete of a task.
	PercentComplete int32 `json:"percentComplete"`
	// The percent work complete of a task.
	PercentWorkComplete int32 `json:"percentWorkComplete"`
	// Determines if a task is active.
	IsActive bool `json:"isActive"`
	// The actual cost of a task.
	ActualCost float32 `json:"actualCost"`
	// The actual duration of a task.
	ActualDuration *string `json:"actualDuration,omitempty"`
	// The actual finish date of a task.
	ActualFinish custom.TimeWithoutTZ `json:"actualFinish"`
	// The actual overtime cost of a task.
	ActualOvertimeCost float32 `json:"actualOvertimeCost"`
	// The actual overtime work of a task.
	ActualOvertimeWork *string `json:"actualOvertimeWork,omitempty"`
	// The duration through which actual work is protected. Reading supported for XML format only.
	ActualWorkProtected *string `json:"actualWorkProtected,omitempty"`
	// The duration through which actual overtime work is protected. Reading supported for XML format only.
	ActualOvertimeWorkProtected *string `json:"actualOvertimeWorkProtected,omitempty"`
	// The actual start date of a task.
	ActualStart custom.TimeWithoutTZ `json:"actualStart"`
	// The amount of budgeted work for a project root task.
	BudgetWork *string `json:"budgetWork,omitempty"`
	// The amount of budgeted cost for a project root task.
	BudgetCost float32 `json:"budgetCost"`
	// Shows the specific date associated with certain constraint types,  such as Must Start On, Must Finish On, Start No Earlier Than, Start No Later Than, Finish No Earlier Than, and Finish No Later Than.
	ConstraintDate custom.TimeWithoutTZ `json:"constraintDate"`
	// Provides choices for the type of constraint that can be applied for scheduling a task.
	ConstraintType *ConstraintType `json:"constraintType"`
	// The contact person for a task.
	Contact string `json:"contact,omitempty"`
	// The projected or scheduled cost of a task.
	Cost float32 `json:"cost"`
	// The difference between the baseline cost and total cost for a task.
	Cv float64 `json:"cv"`
	// The deadline for a task to be completed.
	Deadline custom.TimeWithoutTZ `json:"deadline"`
	// Contains the difference between the total duration of a task and the baseline duration of a task.
	DurationVariance *string `json:"durationVariance,omitempty"`
	// The early finish date of a task.
	EarlyFinish custom.TimeWithoutTZ `json:"earlyFinish"`
	// The early start date of a task.
	EarlyStart custom.TimeWithoutTZ `json:"earlyStart"`
	// Determines whether a task is effort-driven.
	IsEffortDriven bool `json:"isEffortDriven"`
	// Determines whether a task is external.
	IsExternalTask bool `json:"isExternalTask"`
	// The source location and task identifier of an external task.
	ExternalTaskProject string `json:"externalTaskProject,omitempty"`
	// If a task is an external task the property contains the task's external Id.  type.
	ExternalId int32 `json:"externalId"`
	// Contains the duration between the Early Finish and Late Finish dates.
	FinishSlack int32 `json:"finishSlack"`
	// The variance of the task finish date from the baseline finish date as minutes.
	FinishVariance int32 `json:"finishVariance"`
	// The fixed cost of a task.
	FixedCost float64 `json:"fixedCost"`
	// Determines how the fixed cost is accrued against a task.
	FixedCostAccrual *CostAccrualType `json:"fixedCostAccrual"`
	// The amount of a free slack.
	FreeSlack int32 `json:"freeSlack"`
	Guid string `json:"guid,omitempty"`
	// Determines whether the GANTT bar of a task is hidden when displayed in Microsoft Project.
	HideBar bool `json:"hideBar"`
	// Determines whether a task ignores the resource calendar.
	IgnoreResourceCalendar bool `json:"ignoreResourceCalendar"`
	// The late finish date of a task.
	LateFinish custom.TimeWithoutTZ `json:"lateFinish"`
	// The late start date of a task.
	LateStart custom.TimeWithoutTZ `json:"lateStart"`
	IsLevelAssignments bool `json:"isLevelAssignments"`
	CanLevelingSplit bool `json:"canLevelingSplit"`
	// The delay caused by leveling a task.
	LevelingDelay int32 `json:"levelingDelay"`
	// Shows whether a task is marked for further action or identification of some kind.             
	IsMarked bool `json:"isMarked"`
	// Determines whether a task is a milestone.
	IsMilestone bool `json:"isMilestone"`
	// Determines whether a task is in the critical chain.
	IsCritical bool `json:"isCritical"`
	// Determines whether a task is an inserted project.
	IsSubproject bool `json:"isSubproject"`
	// Determines whether a subproject is read-only.
	IsSubprojectReadOnly bool `json:"isSubprojectReadOnly"`
	// The source location of a subproject. Read/write String.
	SubprojectName string `json:"subprojectName,omitempty"`
	// Determines whether a task is a summary task.
	IsSummary bool `json:"isSummary"`
	// Unique ids of all subtasks.
	SubtasksUids []int32 `json:"subtasksUids,omitempty"`
	// The outline level of a task.
	OutlineLevel int32 `json:"outlineLevel"`
	IsOverAllocated bool `json:"isOverAllocated"`
	// Determines whether a task is estimated.
	IsEstimated bool `json:"isEstimated"`
	// The sum of an actual and remaining overtime cost of a task.
	OvertimeCost float32 `json:"overtimeCost"`
	// The amount of an overtime work scheduled for a task.
	OvertimeWork *string `json:"overtimeWork,omitempty"`
	// The percentage complete value entered by the Project Manager.
	PhysicalPercentComplete int32 `json:"physicalPercentComplete"`
	PreLeveledFinish custom.TimeWithoutTZ `json:"preLeveledFinish"`
	PreLeveledStart custom.TimeWithoutTZ `json:"preLeveledStart"`
	// Determines whether a task is a recurring task.
	IsRecurring bool `json:"isRecurring"`
	// The amount of non-overtime work scheduled for a task.
	RegularWork *string `json:"regularWork,omitempty"`
	// The remaining projected cost of completing a task.
	RemainingCost float32 `json:"remainingCost"`
	// The amount of time required to complete the unfinished portion of a task.
	RemainingDuration *string `json:"remainingDuration,omitempty"`
	// The remaining overtime cost projected to finish a task.
	RemainingOvertimeCost float32 `json:"remainingOvertimeCost"`
	// The remaining overtime work scheduled to finish a task.
	RemainingOvertimeWork *string `json:"remainingOvertimeWork,omitempty"`
	// The remaining work scheduled to complete a task.
	RemainingWork *string `json:"remainingWork,omitempty"`
	// The date when a task resumed.
	Resume custom.TimeWithoutTZ `json:"resume"`
	// Determines whether a task can be resumed.
	IsResumeValid bool `json:"isResumeValid,omitempty"`
	// The date that represents the end of the actual portion of a task.
	Stop custom.TimeWithoutTZ `json:"stop"`
	// Determines whether a task is rolled up.
	IsRollup bool `json:"isRollup"`
	// Returns the task's start slack.
	StartSlack int32 `json:"startSlack"`
	// The variance of the task start date from the baseline start date as minutes. 
	StartVariance int32 `json:"startVariance"`
	// The unique id of task calendar.
	CalendarUid int32 `json:"calendarUid"`
	// Determines whether a task is manually scheduled.
	IsManual bool `json:"isManual"`
	// Defines manually scheduled start of a task.
	ManualStart custom.TimeWithoutTZ `json:"manualStart"`
	// Defines manually scheduled finish of a task.
	ManualFinish custom.TimeWithoutTZ `json:"manualFinish"`
	// Defines manually scheduled duration of a task.
	ManualDuration *string `json:"manualDuration,omitempty"`
	// The amount of a total slack.
	TotalSlack int32 `json:"totalSlack"`
	// The type of a task.
	Type_ *TaskType `json:"type"`
	// The work breakdown structure code of a task.
	Wbs string `json:"wbs,omitempty"`
	// The priority of a task from 0 to 1000.
	Priority int32 `json:"priority"`
	// The amount of the scheduled work for a task.
	Work *string `json:"work,omitempty"`
	// The variance of the task work from the baseline task work as minutes.
	WorkVariance float64 `json:"workVariance"`
	// Notes' plain text extracted from RTF data.
	NotesText string `json:"notesText,omitempty"`
	// The text notes in RTF format.
	NotesRTF string `json:"notesRTF,omitempty"`
	Acwp float64 `json:"acwp"`
	Bcws float64 `json:"bcws"`
	Bcwp float64 `json:"bcwp"`
	// LevelingDelayFormat
	LevelingDelayFormat *TimeUnitType `json:"levelingDelayFormat"`
	// The task Uid numbers for the predecessor tasks on which the task depends before it can be started or finished.
	Predecessors string `json:"predecessors,omitempty"`
	// The task Uid numbers for the successor tasks to a task.
	Successors string `json:"successors,omitempty"`
	// Indicates whether to hide the schedule conflict warning indicator in Microsoft Project.             
	IgnoreWarnings bool `json:"ignoreWarnings"`
	// Determines whether a summary task is expanded or not in GanttChart view.
	IsExpanded bool `json:"isExpanded"`
	// Specifies whether a task should be displayed on a timeline view.
	DisplayOnTimeline bool `json:"displayOnTimeline"`
	// Determines whether the task should be displayed as a summary task. Reading supported for XML format only.
	DisplayAsSummary bool `json:"displayAsSummary"`
	// The title or explanatory text for a hyperlink associated with a task.
	Hyperlink string `json:"hyperlink,omitempty"`
	// The address for a hyperlink associated with a task.
	HyperlinkAddress string `json:"hyperlinkAddress,omitempty"`
	// The specific location in a document in a hyperlink associated with a task.  type.
	HyperlinkSubAddress string `json:"hyperlinkSubAddress,omitempty"`
	// Determines whether the % Complete or Physical % Complete field should be used to calculate budgeted cost of work performed (BCWP).
	EarnedValueMethod *EarnedValueMethodType `json:"earnedValueMethod"`
	// Determines whether the current task should be published to Project Server with the rest of the project.
	IsPublished bool `json:"isPublished"`
	// The name of the enterprise resource who is to receive status updates for the current task from resources.
	StatusManager string `json:"statusManager,omitempty"`
	// The start date of a delivery. Reading supported for XML format only.
	CommitmentStart custom.TimeWithoutTZ `json:"commitmentStart"`
	// The finish date of a delivery. Reading supported for XML format only.
	CommitmentFinish custom.TimeWithoutTZ `json:"commitmentFinish"`
	// Determines whether a task has an associated delivery or a dependency on an associated delivery. Reading supported for XML format only.
	CommitmentType int32 `json:"commitmentType"`
	// Gets or sets the collection of baseline values of the task.
	Baselines []TaskBaseline `json:"baselines,omitempty"`
	// Task extended attributes.
	ExtendedAttributes []ExtendedAttribute `json:"extendedAttributes,omitempty"`
	// Task outline codes.
	OutlineCodes []OutlineCode `json:"outlineCodes,omitempty"`
	// Represents the flag which indicates that task has schedule discrepancies.
	Warning bool `json:"warning"`
}
