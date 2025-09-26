package space

type Planet string
var earthYearSecond float64 = 31557600

func Age(second float64, planet Planet) float64 {
	var space_year float64
	switch planet {
	case "Earth":
		space_year=  1
	case "Mercury":
		space_year = 0.2408467
	case "Venus":
		space_year = 0.61519726
	case "Mars":
		space_year = 1.8808158
	case "Jupiter":
		space_year = 11.862615
	case "Saturn":
		space_year = 29.447498
	case "Uranus":
		space_year = 84.016846
	case "Neptune":
		space_year = 164.79132
	}
	second = second / earthYearSecond
	return   second / space_year
}