package czml

// Polyline is a line in the scene composed of multiple segments.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Polyline
type Polyline struct {
	Show                     *bool                     `json:"show,omitempty"`
	Positions                *PositionList             `json:"positions"`
	ArcType                  *ArcType                  `json:"arcType,omitempty"`
	Width                    *float64                  `json:"width,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Material                 *PolylineMaterial         `json:"material,omitempty"`
	FollowSurface            *bool                     `json:"followSurface,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DepthFailMaterial        *PolylineMaterial         `json:"depthFailMaterial,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	ClampToGround            *bool                     `json:"clampToGround,omitempty"`
	ClassificationType       ClassificationType        `json:"classificationType,omitempty"`
	ZIndex                   *int                      `json:"zIndex,omitempty"`
}

// PolylineVolume is a polyline with a volume, defined as a 2D shape extruded along a polyline
// that conforms to the curvature of the globe.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineVolume
type PolylineVolume struct {
	Show                     *bool                     `json:"show,omitempty"`
	Positions                *PositionList             `json:"positions"`
	Shape                    *Shape                    `json:"shape"`
	CornerType               *CornerType               `json:"cornerType,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *PolylineMaterial         `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// AddColor adds a solid-colored line of color specified by rgba value
func (p *Polyline) UpdateColor(rgba []int) {
	c := Color{Rgba: rgba}
	s := SolidColorMaterial{Color: &c}
	m := PolylineMaterial{SolidColor: &s}
	p.Material = &m
}

func (p *Polyline) AddPoint(lat, lon, ele float64) {
	if p.Positions == nil {
		p.Positions = &PositionList{}
	} else if p.Positions.CartographicDegrees == nil {
		p.Positions.CartographicDegrees = CartographicDegreesValue{}
	}

	p.Positions.CartographicDegrees = append(p.Positions.CartographicDegrees, lon, lat, ele)
}
