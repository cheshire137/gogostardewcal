package calendar

type Event interface {
	Day() int
	Season() string
	String() string
}
