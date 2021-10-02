package czml

// Label is a string of text
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Label
type Label struct {
	Show                       *bool                     `json:"show,omitempty"`
	Text                       string                    `json:"text,omitempty"`
	Font                       *Font                     `json:"font,omitempty"`
	Style                      *LabelStyle               `json:"style,omitempty"`
	Scale                      *float64                  `json:"scale,omitempty"`
	ShowBackground             *bool                     `json:"showBackground,omitempty"`
	BackgroundColor            *Color                    `json:"backgroundColor,omitempty"`
	BackgroundPadding          *BackgroundPadding        `json:"backgroundPadding,omitempty"`
	PixelOffset                *PixelOffset              `json:"pixelOffset,omitempty"`
	EyeOffset                  *EyeOffset                `json:"eyeOffset,omitempty"`
	HorizontalOrigin           *HorizontalOrigin         `json:"horizontalOrigin,omitempty"`
	VerticalOrigin             *VerticalOrigin           `json:"verticalOrigin,omitempty"`
	HeightReference            *HeightReference          `json:"heightReference,omitempty"`
	FillColor                  *Color                    `json:"fillColor,omitempty"`
	OutlineColor               *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth               *float64                  `json:"outlineWidth,omitempty"`
	TranslucencyByDistance     *NearFarScaler            `json:"translucencyByDistance,omitempty"`
	PixelOffsetScaleByDistance *NearFarScaler            `json:"pixelOffsetScaleByDistance,omitempty"`
	ScaleByDistance            *NearFarScaler            `json:"scaleByDistance,omitempty"`
	DistanceDisplayCondition   *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	DisableDepthTestDistance   *float64                  `json:"disableDepthTestDistance,omitempty"`
}

// BackgroundPadding describes the amount of horizontal and vertical padding, in pixels, between a
// label's text and its background.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/BackgroundPadding
type BackgroundPadding struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// Font describes a font used to draw text. Fonts are specified using the same syntax as the CSS
// "font" property.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Font
type Font struct {
	Font      FontValue      `json:"font,omitempty"`
	Reference ReferenceValue `json:"reference,omitempty"`
}

// FontValue is a font, specified using the same syntax as the CSS "font" property.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/FontValue
type FontValue string

// LabelStyle describes the style of a label
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LabelStyle
type LabelStyle struct {
	LabelStyle LabelStyleValue `json:"labelStyle,omitempty"`
	Reference  ReferenceValue  `json:"reference,omitempty"`
}

// LabelStyleValue is the style of a label
// valid values are `FILL`, `OUTLINE`, and `FILL_AND_OUTLINE`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LabelStyleValue
type LabelStyleValue string
