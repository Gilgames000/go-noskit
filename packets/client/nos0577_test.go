package packetclt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNoS0577String(t *testing.T) {
	packet := &NoS0577{
		LoginCode:        "65333833663961382D273134302D346235632D623563652D643937383332343836303966",
		InstallationUUID: "46a8719e-b25f-482b-bc85-d9424e519e18",
		RandomHex:        "0073D2BA",
		CountryID:        0,
		ClientVersion:    "0.9.3.3126",
		ClientHash:       "E9BBDAF4510A63EB13240EB6960670ED",
	}

	expected := "NoS0577 65333833663961382D273134302D346235632D623563652D643937383332343836303966 46a8719e-b25f-482b-bc85-d9424e519e18 0073D2BA 0\x0B0.9.3.3126 0 E9BBDAF4510A63EB13240EB6960670ED"

	out := packet.String()

	if !cmp.Equal(out, expected) {
		t.Errorf("\nNoS0577 packet string representation failed\npacket: %+v\nexpected: %+v\nparsed: %+v", packet, expected, out)
	}
}
