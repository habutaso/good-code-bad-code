package main

type Member struct {
	hitPoint   int
	magicPoint int
}

type Magic struct {
	costMagicPoint int
	name           string
	effectValue    int
}

type NestedLogic struct {
	member Member
	magic  Magic
}

func (n *NestedLogic) method() {
	// 生存しているか判定
	if 0 < n.member.hitPoint {
		// 行動可能かを判定
		if n.member.canAct() {
			// 魔法力が残存しているかを判定
			if n.magic.costMagicPoint <= n.member.magicPoint {
				n.member.consumeMagicPoint(n.magic.costMagicPoint)
				n.member.chant(n.magic)
			}
		}
	}
}
