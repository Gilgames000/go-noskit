package testdoubles

import "github.com/gilgames000/go-noskit/entities"

type FakeNPCGateway struct {
	BazaarNPC entities.NPC
}

func (n FakeNPCGateway) All() []entities.NPC {
	return []entities.NPC{n.BazaarNPC}
}

func (n FakeNPCGateway) SearchByID(id int) (entities.NPC, bool) {
	if id != n.BazaarNPC.ID {
		return entities.NPC{}, false
	}

	return n.BazaarNPC, true
}

