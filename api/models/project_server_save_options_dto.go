/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="project_server_save_options_dto.go">
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

// Allows to specify additional options when project is saved to Project Server or Project Online.
type ProjectServerSaveOptionsDto struct {
	// Gets or sets name of a project which is displayed in Project Server \\ Project     Online projects list. Should be unique within Project Server \\ Project Online     instance. Is the value is omitted, the value of Prj.Name property will be used     instead.
	ProjectName string `json:"projectName,omitempty"`
	// Gets or sets unique identifier of a project. Should be unique within Project     Server \\ Project Online instance.
	ProjectGuid string `json:"projectGuid,omitempty"`
	// Gets or sets timeout used when waiting for processing of save project request     by a Project Server's queue processing service. The default value for this property     is 1 minute. The processing time may be longer for large projects or in case when Project     Server instance is too busy responding to other requests.
	Timeout *string `json:"timeout,omitempty"`
	// Gets or sets interval between queue job status requests. The default value is     2 seconds.
	PollingInterval *string `json:"pollingInterval,omitempty"`
}
