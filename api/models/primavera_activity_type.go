/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="primavera_activity_type.go">
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
// PrimaveraActivityType : Specifies type of Primavera activity.
type PrimaveraActivityType string

// List of PrimaveraActivityType
const (
	NONE_PrimaveraActivityType PrimaveraActivityType = "None"
	START_MILESTONE_PrimaveraActivityType PrimaveraActivityType = "StartMilestone"
	FINISH_MILESTONE_PrimaveraActivityType PrimaveraActivityType = "FinishMilestone"
	TASK_DEPENDENT_PrimaveraActivityType PrimaveraActivityType = "TaskDependent"
	RESOURCE_DEPENDENT_PrimaveraActivityType PrimaveraActivityType = "ResourceDependent"
	LEVEL_OF_EFFORT_PrimaveraActivityType PrimaveraActivityType = "LevelOfEffort"
	WBS_SUMMARY_PrimaveraActivityType PrimaveraActivityType = "WbsSummary"
)
