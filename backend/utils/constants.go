package utils

type ListingAmenities int

const (
	Amenity_Undefined ListingAmenities = iota
	Amenity_Dishwasher
	Amenity_WasherDryerInUnit
	Amenity_PrivateOutdoorSpace
	Amenity_GuarantorsAccepted
	Amenity_BoardApprovalRequired
	Amenity_Loft
	Amenity_Furnished
	Amenity_Fireplace
	Amenity_Sublet
	Amenity_CentralAir
	Amenity_CityView
	Amenity_GardenView
	Amenity_ParkView
	Amenity_SkylineView
	Amenity_WaterView
	Amenity_Elevator
	Amenity_Doorman
	Amenity_LaundryInBuilding
	Amenity_Gym
	Amenity_GarageParking
	Amenity_PetsAllowedDogOnly
	Amenity_PetsAllowedCatOnly
	Amenity_PetsAllowed
	Amenity_SwimmingPool
	Amenity_PublicOutdoorSpace
	Amenity_PiedATerreAllowed
	Amenity_GreenBuilding
	Amenity_ChildrenPlayroom
	Amenity_SmokeFree
)

var ListingAmenitiesName = map[int32]string{
	0:  "Amenity_Undefined",
	1:  "Amenity_Dishwasher",
	2:  "Amenity_WasherDryerInUnit",
	3:  "Amenity_PrivateOutdoorSpace",
	4:  "Amenity_GuarantorsAccepted",
	5:  "Amenity_BoardApprovalRequired",
	6:  "Amenity_Loft",
	7:  "Amenity_Furnished",
	8:  "Amenity_Fireplace",
	9:  "Amenity_Sublet",
	10: "Amenity_CentralAir",
	11: "Amenity_CityView",
	12: "Amenity_GardenView",
	13: "Amenity_ParkView",
	14: "Amenity_SkylineView",
	15: "Amenity_WaterView",
	16: "Amenity_Elevator",
	17: "Amenity_Doorman",
	18: "Amenity_LaundryInBuilding",
	19: "Amenity_Gym",
	20: "Amenity_GarageParking",
	21: "Amenity_PetsAllowedDogOnly",
	22: "Amenity_PetsAllowedCatOnly",
	23: "Amenity_PetsAllowed",
	24: "Amenity_SwimmingPool",
	25: "Amenity_PublicOutdoorSpace",
	26: "Amenity_PiedATerreAllowed",
	27: "Amenity_GreenBuilding",
	28: "Amenity_ChildrenPlayroom",
	29: "Amenity_SmokeFree",
}

var ListingAmenitiesValue = map[string]int32{
	"Amenity_Undefined":             0,
	"Amenity_Dishwasher":            1,
	"Amenity_WasherDryerInUnit":     2,
	"Amenity_PrivateOutdoorSpace":   3,
	"Amenity_GuarantorsAccepted":    4,
	"Amenity_BoardApprovalRequired": 5,
	"Amenity_Loft":                  6,
	"Amenity_Furnished":             7,
	"Amenity_Fireplace":             8,
	"Amenity_Sublet":                9,
	"Amenity_CentralAir":            10,
	"Amenity_CityView":              11,
	"Amenity_GardenView":            12,
	"Amenity_ParkView":              13,
	"Amenity_SkylineView":           14,
	"Amenity_WaterView":             15,
	"Amenity_Elevator":              16,
	"Amenity_Doorman":               17,
	"Amenity_LaundryInBuilding":     18,
	"Amenity_Gym":                   19,
	"Amenity_GarageParking":         20,
	"Amenity_PetsAllowedDogOnly":    21,
	"Amenity_PetsAllowedCatOnly":    22,
	"Amenity_PetsAllowed":           23,
	"Amenity_SwimmingPool":          24,
	"Amenity_PublicOutdoorSpace":    25,
	"Amenity_PiedATerreAllowed":     26,
	"Amenity_GreenBuilding":         27,
	"Amenity_ChildrenPlayroom":      28,
	"Amenity_SmokeFree":             29,
}

func (x ListingAmenities) String() string {
	return EnumName(ListingAmenitiesName, int32(x))
}
