/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_text_style.go">
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

// Represents a text style in a view table.
type TableTextStyle struct {
	// Gets a row unique id. Return -1 if the style is to be applied to all rows of a view.
	RowUid int32 `json:"rowUid"`
	// Gets or sets a field the style is to be applied to.
	Field *Field `json:"field"`
	// Returns a value of the TextItemType enum.
	ItemType *TextItemType `json:"itemType"`
	// Gets or sets color of the text.
	Color *Colors `json:"color"`
	// Gets or sets background pattern of the text style.
	BackgroundPattern *BackgroundPattern `json:"backgroundPattern"`
	// Gets or sets background color of the text style.
	BackgroundColor *Colors `json:"backgroundColor"`
}
