package packetsrv

import (
	nospktparser "github.com/gilgames000/go-noskit/parser"
)

// CharacterEquipment packet
type CharacterEquipment struct {
	MainWeaponType       int `json:"main_weapon_type"        parser:"'sc' @String"`
	MainWeaponUp         int `json:"main_weapon_up"          parser:" @String"`
	MainWeaponMinDmg     int `json:"main_weapon_min_dmg"     parser:" @String"`
	MainWeaponMaxDmg     int `json:"main_weapon_max_dmg"     parser:" @String"`
	MainWeaponHitRate    int `json:"main_weapon_hit_rate"    parser:" @String"`
	MainWeaponCritRate   int `json:"main_weapon_crit_rate"   parser:" @String"`
	MainWeaponCritDmg    int `json:"main_weapon_crit_dmg"    parser:" @String"`
	SecondWeaponType     int `json:"second_weapon_type"      parser:" @String"`
	SecondWeaponUp       int `json:"second_weapon_up"        parser:" @String"`
	SecondWeaponMinDmg   int `json:"second_weapon_min_dmg"   parser:" @String"`
	SecondWeaponMaxDmg   int `json:"second_weapon_max_dmg"   parser:" @String"`
	SecondWeaponHitRate  int `json:"second_weapon_hit_rate"  parser:" @String"`
	SecondWeaponCritRate int `json:"second_weapon_crit_rate" parser:" @String"`
	SecondWeaponCritDmg  int `json:"second_weapon_crit_dmg"  parser:" @String"`
	ArmorUp              int `json:"armor_up"                parser:" @String"`
	MeleeDef             int `json:"melee_def"               parser:" @String"`
	MeleeDodge           int `json:"melee_dodge"             parser:" @String"`
	RangedDef            int `json:"ranged_def"              parser:" @String"`
	RangedDodge          int `json:"ranged_dodge"            parser:" @String"`
	MagicDef             int `json:"magic_def"               parser:" @String"`
	FireRes              int `json:"fire_res"                parser:" @String"`
	WaterRes             int `json:"water_res"               parser:" @String"`
	LightRes             int `json:"light_res"               parser:" @String"`
	DarkRes              int `json:"dark_resistance"         parser:" @String"`
}

// Name of the packet
func (p CharacterEquipment) Name() string {
	return "sc"
}

// Type of the packet
func (p CharacterEquipment) Type() nospktparser.PacketType {
	return nospktparser.SERVER
}
