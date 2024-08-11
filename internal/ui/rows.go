package ui

// TwoColumnsTableRow represents a row in a two columns table
type TwoColumnsTableRow [2]string

func (t *TwoColumnsTableRow) toArray() []string {
	return t[:]
}

func (t *TwoColumnsTableRow) firstColumnValue() string {
	return t[0]
}

func (t *TwoColumnsTableRow) secondColumnValue() string {
	return t[1]
}
