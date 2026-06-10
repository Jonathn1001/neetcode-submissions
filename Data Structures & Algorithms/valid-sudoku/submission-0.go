// isValidSudoku checks if a 9x9 Sudoku board is valid according to Sudoku rules.
func isValidSudoku(board [][]byte) bool {
	// Use boolean arrays to track seen digits for each row, column, and 3x3 sub-box.
	// The size is 9x9 because there are 9 rows/cols/boxes, and digits 1-9 (mapped to 0-8).
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{} // boxes[box_idx][digit_idx]

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			char := board[r][c]
			if char == '.' {
				continue // Empty cell, skip validation
			}

			// Convert character digit to 0-indexed integer (e.g., '1' -> 0, '9' -> 8)
			digitIndex := int(char - '1')

			// Check row for duplicates
			if rows[r][digitIndex] {
				return false // Duplicate found in this row
			}
			rows[r][digitIndex] = true

			// Check column for duplicates
			if cols[c][digitIndex] {
				return false // Duplicate found in this column
			}
			cols[c][digitIndex] = true

			// Check 3x3 sub-box for duplicates
			// Calculate box index: (row / 3) * 3 + (col / 3)
			boxIndex := (r/3)*3 + (c/3)
			if boxes[boxIndex][digitIndex] {
				return false // Duplicate found in this 3x3 sub-box
			}
			boxes[boxIndex][digitIndex] = true
		}
	}

	return true // All checks passed, the board is valid
}