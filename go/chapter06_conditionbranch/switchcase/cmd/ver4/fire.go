package ver4

type Fire struct {
	member Member
}

func NewFire(member Member) *Fire {
	f := &Fire{
		member: member,
	}
	return f
}

func (f *Fire) Name() string {
	return "ファイア"
}

func (f *Fire) CostMagicPoint() int {
	return 2
}

func (f *Fire) AttackPower() int {
	return 20 + (int)(f.member.Level()*0.5)
}

func (f *Fire) CostTechnicalPoint() int {
	return 0
}
