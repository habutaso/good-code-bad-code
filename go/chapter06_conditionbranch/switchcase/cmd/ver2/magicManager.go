package ver2

type MagicType int

const (
    fire MagicType = iota
    shiden
    hellFire
)

type MagicManager struct{}

func (m *MagicManager) GetName(magicType MagicType) string {
    name := ""

    switch magicType {
    // 中略
    case hellFire:
        name = "地獄の業火"
    }

    return name
}

func (m *MagicManager) CostMagicPoint(magicType MagicType, member Member) int {
    magicPoint := 0

    switch magicType {
    // 中略
    case hellFire:
        magicPoint = 16
    }

    return magicPoint
}

func (m *MagicManager) AttackPower(magicType MagicType, member Member) int {
    attackPower := 0

    switch magicType {
    // 中略
    case hellFire:
        attackPower = 30 + member.MagicAttack
    }

    return attackPower
}

func (m *MagicManager) CostTechnicalPoint(magicType MagicType, member Member) int {
    technicalPoint := 0

    switch magicType {
    case fire:
        technicalPoint = 0
    case shiden:
        technicalPoint = 5
    }

    return technicalPoint
}
