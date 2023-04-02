package ver2

import (
	magic "switchcase/magicutil"
)

type MagicManager struct{}

func (m *MagicManager) GetName(magicType magic.MagicType) string {
	name := ""

	switch magicType {
	// 中略
	case magic.HellFire:
		name = "地獄の業火"
	}

	return name
}

func (m *MagicManager) CostMagicPoint(magicType magic.MagicType, member magic.Member) int {
	magicPoint := 0

	switch magicType {
	// 中略
	case magic.HellFire:
		magicPoint = 16
	}

	return magicPoint
}

func (m *MagicManager) AttackPower(magicType magic.MagicType, member magic.Member) int {
	attackPower := 0

	switch magicType {
	// 中略
	case magic.HellFire:
		attackPower = 30 + member.MagicAttack
	}

	return attackPower
}

func (m *MagicManager) CostTechnicalPoint(magicType magic.MagicType, member magic.Member) int {
	technicalPoint := 0

	switch magicType {
	case magic.Fire:
		technicalPoint = 0
	case magic.Shiden:
		technicalPoint = 5
	}

	return technicalPoint
}
