package ui

// twoColumnsTableRow represents a row in a two columns table
type twoColumnsTableRow [2]string

func (t *twoColumnsTableRow) toArray() []string {
	return t[:]
}

func (t *twoColumnsTableRow) firstColumnValue() string {
	return t[0]
}

func (t *twoColumnsTableRow) secondColumnValue() string {
	return t[1]
}
