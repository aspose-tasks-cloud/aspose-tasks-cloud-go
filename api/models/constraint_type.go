/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="constraint_type.go">
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
// ConstraintType : Specifies a constraint type for a project task.
type ConstraintType string

// List of ConstraintType
const (
	AS_SOON_AS_POSSIBLE_ConstraintType ConstraintType = "AsSoonAsPossible"
	AS_LATE_AS_POSSIBLE_ConstraintType ConstraintType = "AsLateAsPossible"
	MUST_START_ON_ConstraintType ConstraintType = "MustStartOn"
	MUST_FINISH_ON_ConstraintType ConstraintType = "MustFinishOn"
	START_NO_EARLIER_THAN_ConstraintType ConstraintType = "StartNoEarlierThan"
	START_NO_LATER_THAN_ConstraintType ConstraintType = "StartNoLaterThan"
	FINISH_NO_EARLIER_THAN_ConstraintType ConstraintType = "FinishNoEarlierThan"
	FINISH_NO_LATER_THAN_ConstraintType ConstraintType = "FinishNoLaterThan"
	UNDEFINED_ConstraintType ConstraintType = "Undefined"
)
