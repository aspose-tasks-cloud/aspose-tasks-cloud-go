/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="timephased_data_type.go">
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
// TimephasedDataType : 
type TimephasedDataType string

// List of TimephasedDataType
const (
	ASSIGNMENT_REMAINING_WORK_TimephasedDataType TimephasedDataType = "AssignmentRemainingWork"
	ASSIGNMENT_ACTUAL_WORK_TimephasedDataType TimephasedDataType = "AssignmentActualWork"
	ASSIGNMENT_ACTUAL_OVERTIME_WORK_TimephasedDataType TimephasedDataType = "AssignmentActualOvertimeWork"
	ASSIGNMENT_BASELINE_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaselineWork"
	ASSIGNMENT_BASELINE_COST_TimephasedDataType TimephasedDataType = "AssignmentBaselineCost"
	ASSIGNMENT_ACTUAL_COST_TimephasedDataType TimephasedDataType = "AssignmentActualCost"
	RESOURCE_BASELINE_WORK_TimephasedDataType TimephasedDataType = "ResourceBaselineWork"
	RESOURCE_BASELINE_COST_TimephasedDataType TimephasedDataType = "ResourceBaselineCost"
	TASK_BASELINE_WORK_TimephasedDataType TimephasedDataType = "TaskBaselineWork"
	TASK_BASELINE_COST_TimephasedDataType TimephasedDataType = "TaskBaselineCost"
	TASK_PERCENT_COMPLETE_TimephasedDataType TimephasedDataType = "TaskPercentComplete"
	ASSIGNMENT_BASELINE1_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline1Work"
	ASSIGNMENT_BASELINE1_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline1Cost"
	TASK_BASELINE1_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline1Work"
	TASK_BASELINE1_COST_TimephasedDataType TimephasedDataType = "TaskBaseline1Cost"
	RESOURCE_BASELINE1_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline1Work"
	RESOURCE_BASELINE1_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline1Cost"
	ASSIGNMENT_BASELINE2_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline2Work"
	ASSIGNMENT_BASELINE2_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline2Cost"
	TASK_BASELINE2_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline2Work"
	TASK_BASELINE2_COST_TimephasedDataType TimephasedDataType = "TaskBaseline2Cost"
	RESOURCE_BASELINE2_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline2Work"
	RESOURCE_BASELINE2_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline2Cost"
	ASSIGNMENT_BASELINE3_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline3Work"
	ASSIGNMENT_BASELINE3_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline3Cost"
	TASK_BASELINE3_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline3Work"
	TASK_BASELINE3_COST_TimephasedDataType TimephasedDataType = "TaskBaseline3Cost"
	RESOURCE_BASELINE3_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline3Work"
	RESOURCE_BASELINE3_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline3Cost"
	ASSIGNMENT_BASELINE4_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline4Work"
	ASSIGNMENT_BASELINE4_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline4Cost"
	TASK_BASELINE4_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline4Work"
	TASK_BASELINE4_COST_TimephasedDataType TimephasedDataType = "TaskBaseline4Cost"
	RESOURCE_BASELINE4_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline4Work"
	RESOURCE_BASELINE4_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline4Cost"
	ASSIGNMENT_BASELINE5_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline5Work"
	ASSIGNMENT_BASELINE5_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline5Cost"
	TASK_BASELINE5_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline5Work"
	TASK_BASELINE5_COST_TimephasedDataType TimephasedDataType = "TaskBaseline5Cost"
	RESOURCE_BASELINE5_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline5Work"
	RESOURCE_BASELINE5_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline5Cost"
	ASSIGNMENT_BASELINE6_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline6Work"
	ASSIGNMENT_BASELINE6_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline6Cost"
	TASK_BASELINE6_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline6Work"
	TASK_BASELINE6_COST_TimephasedDataType TimephasedDataType = "TaskBaseline6Cost"
	RESOURCE_BASELINE6_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline6Work"
	RESOURCE_BASELINE6_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline6Cost"
	ASSIGNMENT_BASELINE7_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline7Work"
	ASSIGNMENT_BASELINE7_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline7Cost"
	TASK_BASELINE7_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline7Work"
	TASK_BASELINE7_COST_TimephasedDataType TimephasedDataType = "TaskBaseline7Cost"
	RESOURCE_BASELINE7_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline7Work"
	RESOURCE_BASELINE7_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline7Cost"
	ASSIGNMENT_BASELINE8_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline8Work"
	ASSIGNMENT_BASELINE8_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline8Cost"
	TASK_BASELINE8_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline8Work"
	TASK_BASELINE8_COST_TimephasedDataType TimephasedDataType = "TaskBaseline8Cost"
	RESOURCE_BASELINE8_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline8Work"
	RESOURCE_BASELINE8_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline8Cost"
	ASSIGNMENT_BASELINE9_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline9Work"
	ASSIGNMENT_BASELINE9_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline9Cost"
	TASK_BASELINE9_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline9Work"
	TASK_BASELINE9_COST_TimephasedDataType TimephasedDataType = "TaskBaseline9Cost"
	RESOURCE_BASELINE9_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline9Work"
	RESOURCE_BASELINE9_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline9Cost"
	ASSIGNMENT_BASELINE10_WORK_TimephasedDataType TimephasedDataType = "AssignmentBaseline10Work"
	ASSIGNMENT_BASELINE10_COST_TimephasedDataType TimephasedDataType = "AssignmentBaseline10Cost"
	TASK_BASELINE10_WORK_TimephasedDataType TimephasedDataType = "TaskBaseline10Work"
	TASK_BASELINE10_COST_TimephasedDataType TimephasedDataType = "TaskBaseline10Cost"
	RESOURCE_BASELINE10_WORK_TimephasedDataType TimephasedDataType = "ResourceBaseline10Work"
	RESOURCE_BASELINE10_COST_TimephasedDataType TimephasedDataType = "ResourceBaseline10Cost"
	PHYSICAL_PERCENT_COMPLETE_TimephasedDataType TimephasedDataType = "PhysicalPercentComplete"
	TASK_WORK_TimephasedDataType TimephasedDataType = "TaskWork"
	TASK_COST_TimephasedDataType TimephasedDataType = "TaskCost"
	RESOURCE_WORK_TimephasedDataType TimephasedDataType = "ResourceWork"
	RESOURCE_COST_TimephasedDataType TimephasedDataType = "ResourceCost"
	ASSIGNMENT_WORK_TimephasedDataType TimephasedDataType = "AssignmentWork"
	ASSIGNMENT_COST_TimephasedDataType TimephasedDataType = "AssignmentCost"
	UNDEFINED_TimephasedDataType TimephasedDataType = "Undefined"
)
