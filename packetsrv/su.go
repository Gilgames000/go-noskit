package packetsrv

import (
	nospktparser "github.com/gilgames000/go-noskit/parser"
)

// SkillUsage packet
type SkillUsage struct {
	AttackerType      int `json:"attacker_type"        parser:"'su' @String"`
	AttackerID        int `json:"attacker_id"          parser:" @String"`
	TargetType        int `json:"target_type"          parser:" @String"`
	TargetID          int `json:"target_id"            parser:" @String"`
	SkillID           int `json:"skill_id"             parser:" @String"`
	SkillCooldown     int `json:"skill_cooldown"       parser:" @String"`
	AttackAnimation   int `json:"attack_animation"     parser:" @String"`
	SkillEffect       int `json:"skill_effect"         parser:" @String"`
	PositionX         int `json:"position_x"           parser:" @String"`
	PositionY         int `json:"position_y"           parser:" @String"`
	TargetIsAlive     int `json:"target_is_alive"      parser:" @String"`
	ResultingHP       int `json:"resulting_hp"         parser:" @String"`
	Damage            int `json:"damage"               parser:" @String"`
	HitMode           int `json:"hit_mode"             parser:" @String"`
	SkillTypeMinusOne int `json:"skill_type_minus_one" parser:" @String"`
}

// Name of the packet
func (p SkillUsage) Name() string {
	return "su"
}

// Type of the packet
func (p SkillUsage) Type() nospktparser.PacketType {
	return nospktparser.SERVER
}
