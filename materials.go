package czml

// Material is a definition of how a surface is colored or shaded
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Material
type Material struct {
	SolidColor   *SolidColorMaterial   `json:"solidColor,omitempty"`
	Image        *ImageMaterial        `json:"image,omitempty"`
	Grid         *GridMaterial         `json:"grid,omitempty"`
	Stripe       *StripeMaterial       `json:"stripe,omitempty"`
	Checkerboard *CheckerboardMaterial `json:"checkerboard,omitempty"`
}

type PolylineMaterial struct {
	SolidColor      *SolidColorMaterial      `json:"solidColor,omitempty"`
	PolylineOutline *PolylineOutlineMaterial `json:"polylineOutline,omitempty"`
	PolylineArrow   *PolylineArrowMaterial   `json:"polylineArrow,omitempty"`
	PolylineDash    *PolylineDashMaterial    `json:"polylineDash,omitempty"`
	PolylineGlow    *PolylineGlowMaterial    `json:"polylineGlow,omitempty"`
	Image           *ImageMaterial           `json:"image,omitempty"`
	Grid            *GridMaterial            `json:"grid,omitempty"`
	Stripe          *StripeMaterial          `json:"stripe,omitempty"`
	Checkerboard    *CheckerboardMaterial    `json:"checkerboard,omitempty"`
}

// SolidColorMaterial is a material that fills the surface with a solid color
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/SolidColorMaterial
type SolidColorMaterial struct {
	Color *Color `json:"color,omitempty"`
}

// PolylineOutlineMaterial is a material that fills the surface of a line with an outlined color.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineOutlineMaterial
type PolylineOutlineMaterial struct {
	Color        *Color   `json:"color,omitempty"`
	OutlineColor *Color   `json:"outlineColor,omitempty"`
	OutlineWidth *float64 `json:"outlineWidth,omitempty"`
}

// PolylineArrowMaterial is a material that fills the surface of a line with an arrow.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineArrowMaterial
type PolylineArrowMaterial struct {
	Color *Color `json:"color,omitempty"`
}

// PolylineDashMaterial is a material that fills the surface of a line with a pattern of dashes.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineDashMaterial
type PolylineDashMaterial struct {
	Color       *Color   `json:"color,omitempty"`
	GapColor    *Color   `json:"gapColor,omitempty"`
	DashLength  *float64 `json:"dashLength,omitempty"`
	DashPattern *int     `json:"dashPattern,omitempty"`
}

// PolylineGlowMaterial is a material that fills the surface of a line with a glowing color.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineGlowMaterial
type PolylineGlowMaterial struct {
	Color      *Color   `json:"color,omitempty"`
	GlowPower  *float64 `json:"glowPower,omitempty"`
	TaperPower *float64 `json:"taperPower,omitempty"`
}

// ImageMaterial is a material that fills the surface with an image
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ImageMaterial
type ImageMaterial struct {
	Image       *Uri    `json:"image,omitempty"`
	Repeat      *Repeat `json:"repeat,omitempty"`
	Color       *Color  `json:"color,omitempty"`
	Transparent *bool   `json:"transparent,omitempty"`
}

// Repeat is the number of times an image repeats along each axis.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Repeat
type Repeat struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// GridMaterial is a material that fills the surface with a two-dimensional grid.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/GridMaterial
type GridMaterial struct {
	Color         *Color         `json:"color,omitempty"`
	CellAlpha     *float64       `json:"cellAlpha,omitempty"`
	LineCount     *LineCount     `json:"lineCount,omitempty"`
	LineThickness *LineThickness `json:"lineThickness,omitempty"`
	LineOffset    *LineOffset    `json:"lineOffset,omitempty"`
}

// LineCount is the number of grid lines along each axis
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LineCount
type LineCount struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// LineThickness is the thickness of grid lines along each axis, in pixels
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LineThickness
type LineThickness struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// LineOffset is the offset of grid lines along each axis, as a percentage from 0 to 1
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LineOffset
type LineOffset struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// StripeMaterial is a material that fills the surface with alternating colors
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/StripeMaterial
type StripeMaterial struct {
	Orientation StripeOrientation `json:"orientation,omitempty"`
	EvenColor   *Color            `json:"evenColor,omitempty"`
	OddColor    *Color            `json:"oddColor,omitempty"`
	Offset      *float64          `json:"offset,omitempty"`
	Repeat      *float64          `json:"repeat,omitempty"`
}

// StripeOrientation describes the orientation of stripes in a stripe material
// Valid values are `HORIZONTAL` and `VERTICAL`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/StripeOrientation
type StripeOrientation string

// CheckerboardMaterial is a material that fills the surface with a checkerboard pattern.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CheckerboardMaterial
type CheckerboardMaterial struct {
	EvenColor *Color  `json:"evenColor,omitempty"`
	OddColor  *Color  `json:"oddColor,omitempty"`
	Repeat    *Repeat `json:"repeat,omitempty"`
}
