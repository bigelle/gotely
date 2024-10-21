package types

type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareId    *string  `json:"foursquare_id,omitempty"`
	FourSquareType  *string  `json:"four_square_type,omitempty"`
	GooglePlaceId   *string  `json:"google_place_id,omitempty"`
	GooglePlaceType *string  `json:"google_place_type,omitempty"`
}
