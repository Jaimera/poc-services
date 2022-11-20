package entity

// Port domain representation any business logic strict to port should be here
type Port struct {
	ID          uint32
	Slug        string
	Name        string
	Coordinates LatLng
	City        string
	Country     string
	Alias       *string
	Province    string
	Timezone    string
	Unlocs      string
	Regions     *string
	Code        *string
}
