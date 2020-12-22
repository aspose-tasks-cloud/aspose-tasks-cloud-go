/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="report_type.go">
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
// ReportType : Specifies a type of the project's graphical report.
type ReportType string

// List of ReportType
const (
	PROJECT_OVERVIEW_ReportType ReportType = "ProjectOverview"
	COST_OVERVIEW_ReportType ReportType = "CostOverview"
	WORK_OVERVIEW_ReportType ReportType = "WorkOverview"
	RESOURCE_OVERVIEW_ReportType ReportType = "ResourceOverview"
	RESOURCE_COST_OVERVIEW_ReportType ReportType = "ResourceCostOverview"
	CRITICAL_TASKS_ReportType ReportType = "CriticalTasks"
	LATE_TASKS_ReportType ReportType = "LateTasks"
	MILESTONES_ReportType ReportType = "Milestones"
	UPCOMING_TASK_ReportType ReportType = "UpcomingTask"
	COST_OVERRUNS_ReportType ReportType = "CostOverruns"
	TASK_COST_OVERVIEW_ReportType ReportType = "TaskCostOverview"
	OVERALLOCATED_RESOURCES_ReportType ReportType = "OverallocatedResources"
	SLIPPING_TASKS_ReportType ReportType = "SlippingTasks"
	BEST_PRACTICE_ANALYZER_ReportType ReportType = "BestPracticeAnalyzer"
	BURNDOWN_ReportType ReportType = "Burndown"
	CASH_FLOW_ReportType ReportType = "CashFlow"
)
