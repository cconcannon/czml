package czml

// CartographicDegreesValue is a geodetic, WGS84 position specified as [Longitude, Latitude, Height]
// where Longitude and Latitude are in degrees and Height is in meters. If the array has three
// elements, the value is constant. If it has four or more elements, they are time-tagged samples
// arranged as [Time, Longitude, Latitude, Height, Time, Longitude, Latitude, Height, ...], where
// Time is an ISO 8601 date and time string or seconds since epoch
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicDegreesValue
type CartographicDegreesValue struct {
	Lat    float64
	Lon    float64
	Height float64
	Time   string
}

// CartographicRadiansValue is a geodetic, WGS84 position specified as [Longitude, Latitude, Height]
// where Longitude and Latitude are in radians and Height is in meters. If the array has three
// elements, the value is constant. If it has four or more elements, they are time-tagged samples
// arranged as [Time, Longitude, Latitude, Height, Time, Longitude, Latitude, Height, ...], where
// Time is an ISO 8601 date and time string or seconds since epoch
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRadiansValue
type CartographicRadiansValue []float64

// CartographicRadiansListOfListsValue is a list of lists of geodetic, WGS84 positions specified as
// [Longitude, Latitude, Height, Longitude, Latitude, Height, ...], where Longitude and Latitude are
// in radians and Height is in meters.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRadiansListOfListsValue
type CartographicRadiansListOfListsValue []CartographicRadiansValue

// CartographicDegreesListOfListsValue is a list of lists of geodetic, WGS84 positions specified as
// [Longitude, Latitude, Height, Longitude, Latitude, Height, ...], where Longitude and Latitude are
// in degrees and Height is in meters.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicDegreesListOfListsValue
type CartographicDegreesListOfListsValue []CartographicDegreesValue

// CartographicRectangleRadiansValue is a two-dimensional region specified as [WestLongitude,
// SouthLatitude, EastLongitude, NorthLatitude], with values in radians. If the array has four
// elements, the value is constant. If it has five or more elements, they are time-tagged samples
// arranged as [Time, WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, Time,
// WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, ...], where Time is an ISO 8601 date
// and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRectangleRadiansValue
type CartographicRectangleRadiansValue []interface{}

// CartographicRectangleDegreesValue is a two-dimensional region specified as [WestLongitude,
// SouthLatitude, EastLongitude, NorthLatitude], with values in degrees. If the array has four
// elements, the value is constant. If it has five or more elements, they are time-tagged samples
// arranged as [Time, WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, Time,
// WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, ...], where Time is an ISO 8601 date
// and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRectangleDegreesValue
type CartographicRectangleDegreesValue []interface{}

// CartographicRadiansListValue is a list of geodetic, WGS84 positions specified as
// [Longitude, Latitude, Height, Longitude, Latitude, Height, ...], where Longitude and Latitude
// are in radians and Height is in meters.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRadiansListValue
type CartographicRadiansListValue []float64
