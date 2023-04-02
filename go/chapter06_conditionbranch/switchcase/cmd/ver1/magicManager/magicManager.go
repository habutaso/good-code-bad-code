package magicManager

import (
	magic "switchcase/magicutil"
)

type MagicManager struct{}

func (magicManager *MagicManager) GetName(magicType magic.MagicType) string {
	name := ""

	switch magicType {
	case magic.Fire:
		name = "ファイア"
	case magic.Shiden:
		name = "紫電"
	}
	return name
}

func (magicManager *MagicManager) CostMagicPoint(magicType magic.MagicType, member magic.Member) int {
	magicPoint := 0

	switch magicType {
	case magic.Fire:
		magicPoint = 2
	case magic.Shiden:
		magicPoint = 5 + int(float32(member.Level)*0.2)
	}

	return magicPoint
}

func (magicManager *MagicManager) AttackPower(magicType magic.MagicType, member magic.Member) int {
	attackPower := 0

	switch magicType {
	case magic.Fire:
		attackPower = 20 + int(float32(member.Level)*0.5)
	case magic.Shiden:
		attackPower = 50 + int(float32(member.Agility)*1.5)
	}

	return attackPower
}
