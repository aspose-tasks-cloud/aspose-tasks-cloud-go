/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="primavera_constraint_type.go">
 *   Copyright (c) 2025 Aspose.Tasks Cloud
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

// PrimaveraConstraintType : Specifies a type of activity's constraint.
type PrimaveraConstraintType string

// List of PrimaveraConstraintType
const (
	NONE_PrimaveraConstraintType                PrimaveraConstraintType = "None"
	FINISH_ON_PrimaveraConstraintType           PrimaveraConstraintType = "FinishOn"
	FINISH_ON_OR_AFTER_PrimaveraConstraintType  PrimaveraConstraintType = "FinishOnOrAfter"
	FINISH_ON_OR_BEFORE_PrimaveraConstraintType PrimaveraConstraintType = "FinishOnOrBefore"
	MANDATORY_FINISH_PrimaveraConstraintType    PrimaveraConstraintType = "MandatoryFinish"
	MANDATORY_START_PrimaveraConstraintType     PrimaveraConstraintType = "MandatoryStart"
	START_ON_PrimaveraConstraintType            PrimaveraConstraintType = "StartOn"
	START_ON_OR_AFTER_PrimaveraConstraintType   PrimaveraConstraintType = "StartOnOrAfter"
	START_ON_OR_BEFORE_PrimaveraConstraintType  PrimaveraConstraintType = "StartOnOrBefore"
	AS_LATE_AS_POSSIBLE_PrimaveraConstraintType PrimaveraConstraintType = "AsLateAsPossible"
)
