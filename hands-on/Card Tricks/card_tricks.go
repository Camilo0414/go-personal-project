package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	rt := 0
	if index < 0 || index >= len(slice) {
		rt = -1
	} else {
		rt = slice[index]
	}
	return rt
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	var slicert []int
	if index < 0 || index >= len(slice) {
		slicert = append(slice, value)
	} else {
		initHalfSlice := append(slice[:index], value)
		endingHalfSlice := slice[index+1:]

		slicert = append(initHalfSlice, endingHalfSlice...)
	}

	return slicert
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	slicert := append(values, slice...)

	return slicert
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	} else {
		slicertA := slice[:index]
		slicertB := slice[index+1:]
		return append(slicertA, slicertB...)
	}
}
