/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="view_screen.go">
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
// ViewScreen : Specifies the screen type for a view.
type ViewScreen string

// List of ViewScreen
const (
	GANTT_ViewScreen ViewScreen = "Gantt"
	NETWORK_DIAGRAM_ViewScreen ViewScreen = "NetworkDiagram"
	RELATIONSHIP_DIAGRAM_ViewScreen ViewScreen = "RelationshipDiagram"
	TASK_FORM_ViewScreen ViewScreen = "TaskForm"
	TASK_SHEET_ViewScreen ViewScreen = "TaskSheet"
	RESOURCE_FORM_ViewScreen ViewScreen = "ResourceForm"
	RESOURCE_SHEET_ViewScreen ViewScreen = "ResourceSheet"
	RESOURCE_GRAPH_ViewScreen ViewScreen = "ResourceGraph"
	TASK_DETAILS_FORM_ViewScreen ViewScreen = "TaskDetailsForm"
	TASK_NAME_FORM_ViewScreen ViewScreen = "TaskNameForm"
	RESOURCE_NAME_FORM_ViewScreen ViewScreen = "ResourceNameForm"
	CALENDAR_ViewScreen ViewScreen = "Calendar"
	TASK_USAGE_ViewScreen ViewScreen = "TaskUsage"
	RESOURCE_USAGE_ViewScreen ViewScreen = "ResourceUsage"
)
