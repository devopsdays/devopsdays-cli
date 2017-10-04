package helpers

// GetTalks takes in the city and year and returns a list of the talks
func GetTalks(city, year string) []string {

	s := make([]string, 3)
	s = append(s, "This is a talk", "So is this", "Rainbows are fun", "Never mind")
	return s
}
