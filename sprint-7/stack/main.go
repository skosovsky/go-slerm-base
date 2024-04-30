package main

type BigStackEntry struct {
	valueInt       int
	valueFloat     float64
	payload        string
	otherPayload   string
	embeddedParams struct {
		A int
		B int
		C int
	}
}

type BigStackByValue []BigStackEntry

func (s *BigStackByValue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *BigStackByValue) Push(value BigStackEntry) {
	*s = append(*s, value)
}

func (s *BigStackByValue) Pop() (BigStackEntry, bool) {
	if s.IsEmpty() {
		return BigStackEntry{}, false //nolint:exhaustruct // it's learning code
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}

type BigStackByPointer []*BigStackEntry

func (s *BigStackByPointer) IsEmpty() bool {
	return len(*s) == 0
}

func (s *BigStackByPointer) Push(value *BigStackEntry) {
	*s = append(*s, value)
}

func (s *BigStackByPointer) Pop() (*BigStackEntry, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}

type SmallStackEntry struct {
	valueInt   int
	valueFloat float64
}

type SmallStackByValue []SmallStackEntry

func (s *SmallStackByValue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *SmallStackByValue) Push(value SmallStackEntry) {
	*s = append(*s, value)
}

func (s *SmallStackByValue) Pop() (SmallStackEntry, bool) {
	if s.IsEmpty() {
		return SmallStackEntry{}, false //nolint:exhaustruct // it's learning code
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}

type SmallStackByPointer []*SmallStackEntry

func (s *SmallStackByPointer) IsEmpty() bool {
	return len(*s) == 0
}

func (s *SmallStackByPointer) Push(value *SmallStackEntry) {
	*s = append(*s, value)
}

func (s *SmallStackByPointer) Pop() (*SmallStackEntry, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}

func main() {

}
