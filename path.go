package czml

// Path is a polyline defined by the motion of an object over time. The possible vertices of the
// path are specified by the position property. Note that because clients cannot render a truly
// infinite path, the path must be limited, either by defining availability for this object, or by
// using the leadTime and trailTime properties.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Path
type Path struct {
	Show                     *bool                     `json:"show,omitempty"`
	LeadTime                 *float64                  `json:"leadTime,omitempty"`
	TrailTime                *float64                  `json:"trailTime,omitempty"`
	Width                    *float64                  `json:"width,omitempty"`
	Resolution               *float64                  `json:"resolution,omitempty"`
	Material                 *PolylineMaterial         `json:"material,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// UpdateColor adds or updates a solid-colored line specified by rgba value
func (p *Path) UpdateColor(rgba []int) {
	color := Color{Rgba: rgba}
	show := true
	width := float64(5)
	leadTime := float64(0)
	trailTime := float64(1000000000)
	resolution := float64(10)
	colorMaterial := SolidColorMaterial{Color: &color}
	material := PolylineMaterial{
		SolidColor: &colorMaterial,
	}
	p.Show = &show
	p.Width = &width
	p.LeadTime = &leadTime
	p.TrailTime = &trailTime
	p.Resolution = &resolution
	p.Material = &material
}
