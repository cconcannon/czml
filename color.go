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
