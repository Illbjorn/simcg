package tokenize

func ResolveKind(data []byte) (kind Kind, ok bool) {
	if len(data) == 0 {
		return 0, false
	}

	s := btos(data)
	if kind, ok = words[s]; ok {
		// Word
		return kind, ok
	} else if kind, ok = symbols[s]; ok {
		// Symbol
		return kind, ok
	}

	return 0, false
}

/*------------------------------------------------------------------------------
 * Words
 *----------------------------------------------------------------------------*/

var words = map[string]Kind{
	"actions":              Actions,
	"run_action_list":      RunActionList,
	"call_action_list":     CallActionList,
	"invoke_external_buff": InvokeExternalBuff,
	"active_enemies":       ActiveEnemies,
	"time":                 Time,
	"fight_remains":        FightRemains,
	"rage":                 Rage,
	"toggle":               Toggle,
	"snapshot_stats":       SnapshotStats,
	"buff":                 Buff,
	"debuff":               Debuff,
	"cooldown":             Cooldown,
	"talent":               Talent,
	"target":               Target,
	"dot":                  DOT,
	"raid_event":           RaidEvent,
	"movement":             Movement,
	"equipped":             Equipped,
	"gcd":                  GCD,
	"trinket":              Trinket,
	"if":                   If,
	"target_if":            TargetIf,
	"name":                 Name,
	"value":                VALUE,
	"condition":            Condition,
	"value_else":           ValueElse,
	"slot":                 Slot,
	"op":                   Op,
	"use_item":             UseItem,
	"variable":             Variable,
	"health":               Health,
	"casting":              Casting,
	"proc":                 Proc,
	"adds":                 Adds,
	"has_buff":             HasBuff,
	"has_stat":             HasStat,
	"distance":             Distance,
	"react":                React,
	"is":                   Is,
	"min":                  Min,
	"auto_attack":          AutoAttack,
	"remains":              Remains,
	"remains_expected":     RemainsExpected,
	"duration":             Duration,
	"stack":                Stack,
	"enabled":              Enabled,
	"ready":                Ready,
	"up":                   Up,
	"down":                 Down,
	"pct":                  Pct,
	"in":                   In,
	"exists":               Exists,
	"time_to_die":          TTD,
	"cast_time":            CastTime,
	"has_use_buff":         HasUseBuff,
	"has_cooldown":         HasCooldown,
	"any_dps":              AnyDPS,
	"strength":             Strength,
	"trinket1":             Trinket1,
	"trinket2":             Trinket2,
	"main_hand":            MainHand,
	"on":                   On,
	"off":                  Off,
	"setif":                SetIf,
	"rank":                 Rank,
	"floor":                Floor,
	"ceil":                 Ceil,
}

/*------------------------------------------------------------------------------
 * Symbols
 *----------------------------------------------------------------------------*/

var symbols = map[string]Kind{
	"%%":  Mod,
	"*":   Mult,
	"%":   Div,
	"+":   Add,
	"-":   Sub,
	"!":   Not,
	"|":   LogicalOr,
	"&":   LogicalAnd,
	">=":  GE,
	"<=":  LE,
	">":   GT,
	"<":   LT,
	"?>":  MathMin,
	"?<":  MathMax,
	"(":   ParenL,
	")":   ParenR,
	"=":   Assign,
	"+=/": AddAssign,
	":":   Colon,
	".":   Accessor,
	",":   Comma,
	"@":   Abs,
	"^":   XOr,
}
