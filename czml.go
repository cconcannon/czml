package czml

import (
	"encoding/json"
	"errors"
)

// Czml is a .czml file, which is a valid JSON document that contains an array of packets
type Czml struct {
	Packets []Packet
}

// Unmarshal accepts raw byte data (i.e. a .czml file) and returns
// the data Unmarshaled into the CZML structured interface
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Marshal accepts a name for the document and a Packet array,
// and returns the Marshaled czml data
func Marshal(c Czml) ([]byte, error) {
	return json.Marshal(c.Packets)
}

// MarshalIndent accepts a name for the document, a Packet array,
// a prefix string, and an indentation string, and returns the
// Marshaled and Indented czml data
func MarshalIndent(c Czml, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(&c.Packets, prefix, indent)
}

// InitializeDocument initializes .czml file data with the "document" packet as the first packet
func (c *Czml) InitializeDocument(name string) {
	var packet Packet
	packet.Id = "document"
	packet.Name = name
	packet.Version = "1.0"
	c.AddPacket(packet)
}

// AddPacket adds a Packet to the .czml document data
func (c *Czml) AddPacket(p Packet) {
	c.Packets = append(c.Packets, p)
}

func (c *Czml) AddClock(interval, currentTime string, multiplier float64) error {
	if len(c.Packets) == 0 || c.Packets[0].Id != "document" {
		return errors.New("initialize document before adding properties")
	}

	c.Packets[0].Clock = &Clock{
		Interval:    interval,
		CurrentTime: currentTime,
		Multiplier:  &multiplier,
	}

	return nil
}
