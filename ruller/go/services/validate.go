package services

var validDirections = []string{
	"->",
	"<>",
}

func validateDirection(direction string) bool {
	for _, valid := range validDirections {
		if direction == valid {
			return true
		}
	}
	return false
}