package ver3

import (
	magicUtil "switchcase/magicutil"
)

type Magic struct {
	Name               string
	CostMagicPoint     int
	AttackPower        int
	CostTechnicalPoint int
}

func NewMagic(magicType magicUtil.MagicType, member magicUtil.Member) *Magic {
	magic := &Magic{}
	switch magicType {
	case magicUtil.Fire:
		magic.Name = "ファイア"
		magic.CostMagicPoint = 2
		magic.AttackPower = 20 + int(float64(member.Level)*0.5)
		magic.CostTechnicalPoint = 0
	case magicUtil.Shiden:
		magic.Name = "紫電"
		magic.CostMagicPoint = 5 + int(float64(member.Level)*0.2)
		magic.AttackPower = 50 + int(float64(member.Agility)*1.5)
		magic.CostTechnicalPoint = 5
	case magicUtil.HellFire:
		magic.Name = "地獄の業火"
		magic.CostMagicPoint = 16
		magic.AttackPower = 200 + int(float64(member.MagicAttack)*0.5+float64(member.Vitality)*2)
		magic.CostTechnicalPoint = 20 + int(float64(member.Level)*0.4)
	default:
		panic("invalid magic type")
	}
	return magic
}
