package parse

//go:generate stringer -type Kind --output kind_string.go
type Kind uint8

const (
	KindAPL Kind = 1 + iota
	KindInstruction
	KindStatement

	KindRunActionList
	KindCallActionList
	KindInvokeExternalBuff

	KindUseItem
	KindItem

	KindVar
	KindVarName
	KindVarSetIf
	KindVarValue
	KindVarValueElse
	KindVarCond

	KindVarRef

	KindExecutor

	KindToggle

	KindTargetIf

	KindConditionals

	KindExpr
	KindExprGroup
	KindExprPrefix
	KindLogical
	KindCompare
	KindArithmetic1
	KindArithmetic2

	KindGroupedExpr

	KindBuiltin
	KindActiveEnemies
	KindTime
	KindFightRemains
	KindSnapshotStats

	KindRage

	KindCommand

	KindTrinketCmd

	KindStats

	KindCMDCooldown
	KindCMDVariable
	KindCMDDOT
	KindCMDBuff
	KindCMDDebuff
	KindCMDTalent
	KindCMDMovement
	KindCMDRaidEvent
	KindCMDEquipped
	KindCMDTarget
	KindCMDGCD

	KindHealth
	KindTrinket
	KindCasting
	KindProc
	KindDistance
	KindAdds
	KindHasBuff
	KindHasStat
	KindReact
	KindIs
	KindMin
	KindAutoAttack
	KindRemains
	KindRemainsExpected
	KindDuration
	KindStack
	KindEnabled
	KindRank
	KindReady
	KindUp
	KindDown
	KindPct
	KindIn
	KindExists
	KindTTD
	KindCastTime
	KindHasUseBuff
	KindHasCooldown
	KindAnyDPS
	KindStrength
	KindTrinket1
	KindTrinket2
	KindMainHand
	KindNegate
	KindOr
	KindAnd
	KindOn
	KindOff
	KindGE
	KindGT
	KindLE
	KindLT
	KindMod
	KindMult
	KindDiv
	KindAdd
	KindSub
	KindParenL
	KindParenR
	KindOctothorpe
	KindColon
	KindComma
	KindAccessor
	KindNum
	KindID

	KindBase
)
