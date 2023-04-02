package ver4

type Magic interface {
	Name() string
	CostMagicPoint() int
	AttackPower() int
	CostTechnicalPoint() int
}
