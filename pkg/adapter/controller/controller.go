package controller

// Controller holds the controllers for the entire across application
type Controller struct {
	Employee interface{ Employee }
	User interface{User}
	Link interface{Link}
	Group interface{Group}
	Todo interface{Todo}
}