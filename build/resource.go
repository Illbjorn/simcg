package build

//go:generate stringer -type Resource --output zz_resource_string.go
type Resource uint8

const (
	ResourceStrength Resource = 1 + iota
	ResourceAgility
	ResourceStamina
	ResourceIntellect
	ResourceSpirit
	ResourceHealth
	ResourceMaximumHealth
	ResourceMana
	ResourceMaximumMana
	ResourceRage
	ResourceMaximumRage
	ResourceEnergy
	ResourceMaximumEnergy
	ResourceFocus
	ResourceMaximumFocus
	ResourceRunic
	ResourceMaximumRunic
	ResourceSpellPower
	ResourceMp5
	ResourceAttackPower
	ResourceExpertiseRating
	ResourceInverseExpertiseRating
	ResourceHitRating
	ResourceInverseHitRating
	ResourceCritRating
	ResourceHasteRating
	ResourceWeaponDps
	ResourceWeaponSpeed
	ResourceWeaponOffhandDps
	ResourceWeaponOffhandSpeed
	ResourceArmor
	ResourceBonusArmor
	ResourceResilienceRating
	ResourceDodgeRating
	ResourceParryRating
	ResourceBlockRating
	ResourceMasteryRating
	ResourceAnyDPS
)

//go:generate stringer -type ResourceCMD --output zz_resource_cmd_string.go
type ResourceCMD uint8

const (
	ResourceMax ResourceCMD = 1 + iota
	ResourcePct
	ResourceDeficit
	ResourceMaxNonProc
	ResourcePctNonProc
	ResourceRegen
	ResourceTimeToMax
	ResourceTimeToX
	ResourceNetRegen
)
