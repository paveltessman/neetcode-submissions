func makeSet() []map[byte]struct{} {
	m := make([]map[byte]struct{}, 9)
	for i := range 9 {
		m[i] = make(map[byte]struct{})
	}
	return m
}

func isValidSudoku(board [][]byte) bool {

	rows := makeSet()
	cols := makeSet()
	squares := makeSet()

	for row, values := range board {
		for col, value := range values {
			if value == '.' {
				continue
			}

			if _, exists := rows[row][value]; exists {
				return false
			}

			if _, exists := cols[col][value]; exists {
				return false
			}

			sRow := row / 3
			sCol := col / 3
			squareIndex := sRow*3 + sCol

			if _, exists := squares[squareIndex][value]; exists {
				return false
			}

			rows[row][value] = struct{}{}
			cols[col][value] = struct{}{}
			squares[squareIndex][value] = struct{}{}
		}

	}
	return true
}
