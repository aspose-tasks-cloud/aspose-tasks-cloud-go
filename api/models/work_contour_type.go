/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="work_contour_type.go">
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
// WorkContourType : Specifies the contour of a work.
type WorkContourType string

// List of WorkContourType
const (
	FLAT_WorkContourType WorkContourType = "Flat"
	BACK_LOADED_WorkContourType WorkContourType = "BackLoaded"
	FRONT_LOADED_WorkContourType WorkContourType = "FrontLoaded"
	DOUBLE_PEAK_WorkContourType WorkContourType = "DoublePeak"
	EARLY_PEAK_WorkContourType WorkContourType = "EarlyPeak"
	LATE_PEAK_WorkContourType WorkContourType = "LatePeak"
	BELL_WorkContourType WorkContourType = "Bell"
	TURTLE_WorkContourType WorkContourType = "Turtle"
	CONTOURED_WorkContourType WorkContourType = "Contoured"
	UNDEFINED_WorkContourType WorkContourType = "Undefined"
)
