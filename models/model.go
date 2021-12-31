package models

type Shop struct {
	Name      string
	ID        string
	Items     []Item
	IsHoliday bool
	Holidays  map[string]string
	Holiday   int
}

type Item struct {
	ID    string
	Owner string
	Name  string
	Price int64
	No    int
}

type Person struct {
	Company    string
	Base       Human
	IsSleeping bool
}

type Human struct {
	ID      string
	Name    string
	Age     int16
	Planet  Planet
	Interpt chan string
}

type Planet struct {
	Name       string
	Position   int32
	Galaxy     Galaxy
	Revolution int16
	ID         string
}

type Galaxy struct {
	Name     string
	Position int32
	ID       string
}
