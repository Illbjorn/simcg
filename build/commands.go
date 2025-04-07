package build

// SelfCMD are commands which apply to the spell or item for which the
// instruction line is defined.
//
// # Example
//
//	actions.precombat+=/battle_shout,if=cooldown=0 // <- cooldown implicitly applies to 'battle_shout'
//
//go:generate stringer -type SelfCMD --output zz_self_cmd_string.go
type SelfCMD uint8

const (
	BaseExecuteTime      SelfCMD = 1 + iota // Greater of GCD / cast time
	BaseGCD                                 // GCD-time taken for current action
	BaseCastTime                            // Remainder on currently active GCD
	BaseCooldown                            // Spell, item cooldown in seconds
	BaseTicking                             // For DoTs / HoTs, 1 or 0 if it is up
	BaseTicks                               // For DoTs / HoTs, number of ticks since last refresh
	BaseTicksRemain                         // For DoTs / HoTs, number of ticks before expiry
	BaseRemains                             // For DoTs / HoTs, remaining time before expiry
	BaseFullRechargeTime                    // Time until all spell charges are regained
	BaseTickTime                            // For DoTs / HoTs, time between ticks
	BaseTravelTime                          // Time for a spell to reach its target
	BaseMissReact                           // 1 if last spell use missed, otherwise 0
	BaseCooldownReact                       // 1 if the cooldown has elapsed or was reset
	BaseCastDelay                           // TODO
	BaseMultiplier                          // TODO
	BaseCasting                             // TODO
	BaseChanneling                          // TODO
	BaseExecuting                           // TODO
)
