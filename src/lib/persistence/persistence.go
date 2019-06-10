package persistence

type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
	AddUser(User) ([]byte, error)
	AddBookingForUser([]byte, Booking) error
	AddLocation(Location) (Location, error)
	FindUser(string, string) (User, error)
	FindBookingForUser([]byte) ([]Booking, error)
	FindLocation(string) (Location, error)
	FindAllLocations() ([]Location, error)
	
}