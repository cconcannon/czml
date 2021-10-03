package czml

// PolylineMaterial is a definition of how a polyline is colored or shaded
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineMaterial
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

func (m *PolylineMaterial) UpdateColor(c SolidColorMaterial) {
	m.SolidColor = &c
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
