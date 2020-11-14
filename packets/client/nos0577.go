package packetclt

import (
	"fmt"

	"github.com/gilgames000/go-noskit/packets"
)

// NoS0577 packet
type NoS0577 struct {
	LoginCode        string `json:"login_code"`
	InstallationUUID string `json:"installation_uuid"`
	RandomHex        string `json:"random_hex"`
	CountryID        int    `json:"country_id"`
	ClientVersion    string `json:"client_version"`
	ClientHash       string `json:"client_hash"`
}

// Name of the packet
func (p NoS0577) Name() string {
	return "NoS0577"
}

// Type of the packet
func (p NoS0577) Type() packets.PacketType {
	return packets.CLIENT
}

// String representation of the packet
func (p NoS0577) String() string {
	return fmt.Sprintf("%s %s %s %s %d\x0B%s 0 %s",
		p.Name(),
		p.LoginCode,
		p.InstallationUUID,
		p.RandomHex,
		p.CountryID,
		p.ClientVersion,
		p.ClientHash)
}
