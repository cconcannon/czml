package czml

// Color describes a color. The color can optionally vary over time
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Color
type Color struct {
	Rgba      RgbaValue      `json:"rgba,omitempty"`
	Rgbaf     RgbafValue     `json:"rgbaf,omitempty"`
	Reference ReferenceValue `json:"reference,omitempty"`
}

// RgbaValue is a color specified as an array of color components [Red, Green, Blue, Alpha] where
// each component is in the range 0-255. If the array has four elements, the color is constant. If
// it has five or more elements, they are time-tagged samples arranged as [Time, Red, Green, Blue,
// Alpha, Time, Red, Green, Blue, Alpha, ...], where Time is an ISO 8601 date and time string or
// seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/RgbaValue
type RgbaValue []int

// RgbafValue is a color specified as an array of color components [Red, Green, Blue, Alpha] where
// each component is in the range 0.0-1.0. If the array has four elements, the color is constant.
// If it has five or more elements, they are time-tagged samples arranged as [Time, Red, Green,
// Blue, Alpha, Time, Red, Green, Blue, Alpha, ...], where Time is an ISO 8601 date and time string
// or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/RgbafValue
type RgbafValue []float32

func translateColor(color string) (rgba RgbaValue) {
	switch color {
	case "red":
		rgba = []int{255, 0, 0, 255}
	case "green":
		rgba = []int{0, 255, 0, 255}
	case "blue":
		rgba = []int{0, 0, 255, 255}
	case "purple":
		rgba = []int{128, 0, 128, 255}
	case "yellow":
		rgba = []int{255, 255, 0, 255}
	case "white":
		rgba = []int{255, 255, 255, 255}
	case "black":
		rgba = []int{0, 0, 0, 255}
	default:
		rgba = []int{128, 128, 128, 255}
	}

	return rgba
}
