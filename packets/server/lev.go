package packetsrv

import (
	"github.com/gilgames000/go-noskit/packets"
)

// CharacterLevel packet
type CharacterLevel struct {
	CombatLevel     int `json:"combat_level"      parser:"'lev' @String"`
	CurrentCombatXP int `json:"current_combat_xp" parser:"@String"`
	JobLevel        int `json:"job_level"         parser:"@String"`
	CurrentJobXP    int `json:"current_job_xp"    parser:"@String"`
	MaxCombatXP     int `json:"max_combat_xp"     parser:"@String"`
	MaxJobXP        int `json:"max_job_xp"        parser:"@String"`
	Reputation      int `json:"reputation"        parser:"@String"`
	SkillCP         int `json:"skill_cp"          parser:"@String"`
	HeroXP          int `json:"hero_xp"           parser:"@String"`
	HeroLevel       int `json:"hero_level"        parser:"@String"`
	HeroXPLoad      int `json:"hero_xp_load"      parser:"@String"`
	Unknown         int `json:"unknown"           parser:"@String"` // maybe R Mode? need to test it
}

// Name of the packet
func (p CharacterLevel) Name() string {
	return "lev"
}

// Type of the packet
func (p CharacterLevel) Type() packets.PacketType {
	return packets.SERVER
}
