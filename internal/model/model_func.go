package model

func (v Vehicle) IsEmpty() bool {
	var empty Vehicle
	return v == empty
}
