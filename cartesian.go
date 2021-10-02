package czml

// Cartesian2 is two-dimensional Cartesian value specified as [X, Y]. If the array has two elements,
// the value is constant. If it has three or more elements, they are time-tagged samples arranged as
// [Time, X, Y, Time, X, Y, ...], where Time is an ISO 8601 date and time string or seconds since
// epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian2Value
type Cartesian2Value []float64

// Cartesian2ListValue is a list of two-dimensional Cartesian values specified as
// [X, Y, X, Y, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian2ListValue
type Cartesian2ListValue *[]float64

// Cartesian3Value is the position specified as a three-dimensional
// Cartesian value [X, Y, Z] in meters relative to the referenceFrame
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3Value
type Cartesian3Value []float64

// Cartesian3VelocityValue is a three-dimensional Cartesian value and its derivative specified as
// [X, Y, Z, dX, dY, dZ]. If the array has six elements, the value is constant. If it has seven or
// more elements, they are time-tagged samples arranged as [Time, X, Y, Z, dX, dY, dZ, Time, X,...],
// where Time is an ISO 8601 date and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3VelocityValue
type Cartesian3VelocityValue []float64

// UnitCartesian3Value is a three-dimensional unit magnitude Cartesian value specified as [X, Y, Z].
// If the array has three elements, the value is constant. If it has four or more elements, they are
// time-tagged samples arranged as [Time, X, Y, Z, Time, X, Y, Z, ...], where Time is an ISO 8601
// date and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitCartesian3Value
type UnitCartesian3Value []interface{}

// UnitCartesian3ListValue is a list of three-dimensional unit magnitude Cartesian values,
// specified as [X, Y, Z, X, Y, Z, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitCartesian3ListValue
type UnitCartesian3ListValue []interface{}

// Cartesian3ListValue is a list of three-dimensional Cartesian values specified as
// [X, Y, Z, X, Y, Z, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3ListValue
type Cartesian3ListValue []float64

// Cartesian3ListOfListsValue is a list of lists of three-dimensional Cartesian values specified
// as [X, Y, Z, X, Y, Z, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3ListOfListsValue
type Cartesian3ListOfListsValue []Cartesian3Value
