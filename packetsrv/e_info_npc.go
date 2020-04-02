package packetsrv

import (
	"github.com/gilgames000/go-noskit/parser"
)

// NPCInfo packet
type NPCInfo struct {
	VNum           int    `json:"vnum"            parser:"'e_info' '10' @String"`
	Level          int    `json:"level"           parser:" @String"`
	AttributeType  int    `json:"attribute_type"  parser:" @String"`
	DamageType     int    `json:"damage_type"     parser:" @String"`
	AttributeLevel int    `json:"attribute_level" parser:" @String"`
	AttackLevel    int    `json:"attack_level"    parser:" @String"`
	MinDamage      int    `json:"min_damage"      parser:" @String"`
	MaxDamage      int    `json:"max_damage"      parser:" @String"`
	HitRate        int    `json:"hit_rate"        parser:" @String"`
	CriticalDamage int    `json:"critical_damage" parser:" @String"`
	CriticalRate   int    `json:"critical_rate"   parser:" @String"`
	DefenseLevel   int    `json:"defense_level"   parser:" @String"`
	MeleeDef       int    `json:"melee_def"       parser:" @String"`
	MeleeDodge     int    `json:"melee_dodge"     parser:" @String"`
	RangedDef      int    `json:"ranged_def"      parser:" @String"`
	RangedDodge    int    `json:"ranged_dodge"    parser:" @String"`
	MagicDef       int    `json:"magic_def"       parser:" @String"`
	FireRes        int    `json:"fire_res"        parser:" @String"`
	WaterRes       int    `json:"water_res"       parser:" @String"`
	LightRes       int    `json:"light_res"       parser:" @String"`
	DarkRes        int    `json:"dark_res"        parser:" @String"`
	MaxHP          int    `json:"max_hp"          parser:" @String"`
	MaxMP          int    `json:"max_mp"          parser:" @String"`
	NPCName        string `json:"npc_name"        parser:" '-1' @String"`
}

// Name of the packet
func (p NPCInfo) Name() string {
	return "e_info 10"
}

// Type of the packet
func (p NPCInfo) Type() parser.PacketType {
	return parser.SERVER
}
