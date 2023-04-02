package ver4

type HellFire struct {
	member member
}

func NewHellFire(m member) *HellFire {
	return &HellFire{member: m}
}

func (hf *HellFire) Name() string {
	return "地獄の業火"
}

func (hf *HellFire) CostMagicPoint() int {
	return 16
}

func (hf *HellFire) AttackPower() int {
	return 200 + int(hf.member.magicAttack*0.5+hf.member.vitality*2)
}

func (hf *HellFire) CostTechnicalPoint() int {
	return 20 + int(hf.member.level*0.4)
}
