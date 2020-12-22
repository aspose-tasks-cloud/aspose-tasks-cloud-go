/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="confidence_level.go">
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
// ConfidenceLevel : Specifies supported confidence levels used in risk analysis that correspond to  the percentage of the time the actual values will be within optimistic and pessimistic estimates.
type ConfidenceLevel string

// List of ConfidenceLevel
const (
	CL75_ConfidenceLevel ConfidenceLevel = "CL75"
	CL85_ConfidenceLevel ConfidenceLevel = "CL85"
	CL90_ConfidenceLevel ConfidenceLevel = "CL90"
	CL95_ConfidenceLevel ConfidenceLevel = "CL95"
	CL99_ConfidenceLevel ConfidenceLevel = "CL99"
)
