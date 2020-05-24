package gamestate

import (
	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
	packetsrv "github.com/gilgames000/go-noskit/packets/server"
)

var _ actions.NPCGateway = &NPCGateway{}

type NPCGateway struct {
	gameSocket GameSocket
	npcs       map[int]entities.NPC
}

func NewNPCGateway(gameSocket GameSocket) *NPCGateway {
	npcGateway := &NPCGateway{
		gameSocket: gameSocket,
		npcs:       make(map[int]entities.NPC),
	}

	go npcGateway.updater()

	return npcGateway
}

func (ng *NPCGateway) updater() {
	l := ng.gameSocket.NewListener([]string{
		packetsrv.SpawnNPC{}.Name(),
	}...)
	defer ng.gameSocket.CloseListener(l)

	for {
		packet := <-l
		switch p := packet.(type) {
		case packetsrv.SpawnNPC:
			ng.npcs[p.ID] = newNPC(p)
			// TODO: add DespawnNPC
		default:
			continue
		}
	}
}

func (ng *NPCGateway) All() []entities.NPC {
	var npcs []entities.NPC

	for _, npc := range ng.npcs {
		npcs = append(npcs, npc)
	}

	return npcs
}

func (ng *NPCGateway) SearchByID(id int) (entities.NPC, bool) {
	npc, ok := ng.npcs[id]

	return npc, ok
}

func newNPC(npc packetsrv.SpawnNPC) entities.NPC {
	return entities.NPC{
		Actor: entities.Actor{
			Thing: entities.Thing{
				ID: npc.ID,
				Position: entities.Point{
					X: npc.PositionX,
					Y: npc.PositionY,
				},
			},
			CombatLevel:  0,
			HPPercentage: npc.CurrentHP,
			MPPercentage: npc.CurrentMP,
		},
		VNum: npc.VNum,
	}
}
