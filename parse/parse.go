package parse

import (
	"github.com/illbjorn/echo"

	"github.com/illbjorn/simcg/tokenize"
)

func Parse(set *tokenize.Set) *Node {
	apl := &Node{Kind: KindAPL}

	for {
		if set.Peek(1).Kind == 0 {
			return apl
		}

		for instruction, err := range Instruction(set) {
			if err != nil {
				echo.Fatalf("Failed parse: %s.", err)
			}

			apl.Children = append(apl.Children, instruction)
		}
	}
}

/*------------------------------------------------------------------------------
 * Non-terminals
 *----------------------------------------------------------------------------*/

func Instruction(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(
			KindInstruction,
			ACTIONS, OneOf(ASSIGN, ADD_ASSIGN), Statement,
		),
		Seq(
			KindInstruction,
			ACTIONS, ID, OneOf(ASSIGN, ADD_ASSIGN), Statement,
		),
	)(set)
}

func Statement(set *tokenize.Set) NodeProducer {
	return OneOf(
		RunActionList,
		CallActionList,
		InvokeExternalBuff,
		UseItem,
		Var,
		Executor,
		Toggle,
		ID,
		SNAPSHOT_STATS,
		AUTO_ATTACK,
	)(set)
}

func RunActionList(set *tokenize.Set) NodeProducer {
	return Seq(
		KindRunActionList,
		RUN_ACTION_LIST,
		NAME,
		ASSIGN,
		ID,
		Maybe(TargetIf),
		Maybe(Conditional),
	)(set)
}

func CallActionList(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCallActionList,
		CALL_ACTION_LIST, NAME, ASSIGN, ID, Maybe(Conditional),
	)(set)
}

func InvokeExternalBuff(set *tokenize.Set) NodeProducer {
	return Seq(
		KindInvokeExternalBuff,
		INVOKE_EXTERNAL_BUFF, NAME, ASSIGN, ID, Maybe(Conditional),
	)(set)
}

func UseItem(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(KindUseItem, USE_ITEM, SLOT, ASSIGN, Item, IF, ASSIGN, Expr),
		Seq(KindUseItem, USE_ITEM, NAME, ASSIGN, ID, IF, ASSIGN, Expr),
	)(set)
}

func Item(set *tokenize.Set) NodeProducer {
	return OneOf(
		MAIN_HAND,
		TRINKET_1,
		TRINKET_2,
	)(set)
}

func Var(set *tokenize.Set) NodeProducer {
	return Seq(
		KindVar,
		VARIABLE,
		VarName,
		Maybe(VarSetIf),
		VarValue,
		Maybe(VarValueElse),
		Maybe(VarCond),
	)(set)
}

func VarName(set *tokenize.Set) NodeProducer {
	return Seq(
		KindVarName,
		NAME, ASSIGN, ID,
	)(set)
}

func VarSetIf(set *tokenize.Set) NodeProducer {
	return Seq(
		KindVarSetIf,
		OP, ASSIGN, SET_IF,
	)(set)
}

func VarValue(set *tokenize.Set) NodeProducer {
	return Seq(
		KindVarValue,
		VALUE, ASSIGN, Expr,
	)(set)
}

func VarValueElse(set *tokenize.Set) NodeProducer {
	return Seq(
		KindVarValueElse,
		VALUE_ELSE, ASSIGN, Expr,
	)(set)
}

func VarCond(set *tokenize.Set) NodeProducer {
	return Seq(
		KindVarCond,
		CONDITION, ASSIGN, Expr,
	)(set)
}

func Executor(set *tokenize.Set) NodeProducer {
	return Seq(
		KindExecutor,
		ID, IF, ASSIGN, Expr,
	)(set)
}

func Toggle(set *tokenize.Set) NodeProducer {
	return Seq(
		KindToggle,
		ID, TOGGLE, ASSIGN, OneOf(ON, OFF),
	)(set)
}

func TargetIf(set *tokenize.Set) NodeProducer {
	return Seq(
		KindTargetIf,
		TARGET_IF, ASSIGN, Expr,
	)(set)
}

func Conditional(set *tokenize.Set) NodeProducer {
	return Eager(
		KindConditionals,
		IF, ASSIGN, Expr,
	)(set)
}

func Expr(set *tokenize.Set) NodeProducer {
	return Eager(
		KindExpr,
		ExprOr,
	)(set)
}

func ExprOr(set *tokenize.Set) NodeProducer {
	return Seq(
		KindLogical,
		ExprAnd, Lazy(KindLogical, OR, ExprAnd),
	)(set)
}

func ExprAnd(set *tokenize.Set) NodeProducer {
	return Seq(
		KindLogical,
		ExprCompare, Lazy(KindLogical, AND, ExprCompare),
	)(set)
}

func ExprCompare(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCompare,
		ExprMult,
		Lazy(KindCompare, OneOf(ASSIGN, GE, LE, GT, LT), ExprMult),
	)(set)
}

func ExprMult(set *tokenize.Set) NodeProducer {
	return Seq(
		KindArithmetic1,
		ExprAdd, Lazy(
			KindArithmetic1,
			OneOf(MOD, MULT, DIV), ExprAdd,
		),
	)(set)
}

func ExprAdd(set *tokenize.Set) NodeProducer {
	return Seq(
		KindArithmetic2,
		ExprPrefix, Lazy(
			KindArithmetic2,
			OneOf(ADD, SUB), ExprPrefix,
		),
	)(set)
}

func ExprPrefix(set *tokenize.Set) NodeProducer {
	return Seq(
		KindExprPrefix,
		Maybe(NEGATE), PrimeExpr,
	)(set)
}

func PrimeExpr(set *tokenize.Set) NodeProducer {
	return OneOf(
		GroupedExpr,
		Builtin,
		Command,
		PrimaryValue,
	)(set)
}

func GroupedExpr(set *tokenize.Set) NodeProducer {
	return Seq(
		KindGroupedExpr,
		PARENL, Expr, PARENR,
	)(set)
}

func Builtin(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(KindBuiltin, MIN, COLON, PrimeExpr),
		Seq(KindBuiltin, TIME),
		Seq(KindBuiltin, FIGHT_REMAINS),
		Seq(KindBuiltin, RAGE),
		Seq(KindBuiltin, TOGGLE),
		Seq(KindBuiltin, SNAPSHOT_STATS),
		Seq(KindBuiltin, ACTIVE_ENEMIES),
	)(set)
}

func Command(set *tokenize.Set) NodeProducer {
	return OneOf(
		TrinketCMD,
		CooldownCMD,
		VariableCMD,
		DotCMD,
		BuffCMD,
		DebuffCMD,
		TalentCMD,
		MovementCMD,
		RaidEventCMD,
		EquippedCMD,
		TargetCMD,
		GCDCMD,
	)(set)
}

/*------------------------------------------------------------------------------
 * Trinket Command
 *----------------------------------------------------------------------------*/

func TrinketCMD(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(KindTrinketCmd, TRINKET, NUM, IS, ID),
		Seq(KindTrinketCmd, TRINKET, NUM, CAST_TIME),
		Seq(KindTrinketCmd, TRINKET, NUM, HAS_USE_BUFF),
		Seq(KindTrinketCmd, TRINKET, NUM, HAS_COOLDOWN),
		Seq(KindTrinketCmd, TRINKET, NUM, HAS_BUFF, Stats),
		Seq(KindTrinketCmd, TRINKET, NUM, COOLDOWN, LeafCMD),
		Seq(KindTrinketCmd, TRINKET, NUM, HAS_STAT, Stats),
		Seq(KindTrinketCmd, TRINKET, NUM, PROC, Stats, LeafCMD),
	)(set)
}

/*------------------------------------------------------------------------------
 * Cooldown Command
 *----------------------------------------------------------------------------*/

func CooldownCMD(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCMDCooldown,
		COOLDOWN, ID, LeafCMD,
	)(set)
}

/*------------------------------------------------------------------------------
 * Variable Command
 *----------------------------------------------------------------------------*/

func VariableCMD(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCMDVariable,
		VARIABLE, ID,
	)(set)
}

/*------------------------------------------------------------------------------
 * DOT Command
 *----------------------------------------------------------------------------*/

func DotCMD(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCMDDOT,
		DOT, ID, LeafCMD,
	)(set)
}

/*------------------------------------------------------------------------------
 * Buff Command
 *----------------------------------------------------------------------------*/

func BuffCMD(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCMDBuff,
		BUFF, ID, LeafCMD,
	)(set)
}

/*------------------------------------------------------------------------------
 * Debuff Command
 *----------------------------------------------------------------------------*/

func DebuffCMD(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(KindCMDDebuff, DEBUFF, ID, LeafCMD),
		Seq(KindCMDDebuff, DEBUFF, CASTING, LeafCMD),
	)(set)
}

/*------------------------------------------------------------------------------
 * Talent Command
 *----------------------------------------------------------------------------*/

func TalentCMD(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(KindCMDTalent, TALENT, ID, LeafCMD),
		Seq(KindCMDTalent, TALENT, ID),
	)(set)
}

/*------------------------------------------------------------------------------
 * Movement Command
 *----------------------------------------------------------------------------*/

func MovementCMD(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCMDMovement,
		MOVEMENT, DISTANCE,
	)(set)
}

/*------------------------------------------------------------------------------
 * Raid Event Command
 *----------------------------------------------------------------------------*/

func RaidEventCMD(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCMDRaidEvent,
		RAID_EVENT, ADDS, LeafCMD,
	)(set)
}

/*------------------------------------------------------------------------------
 * Equipped Command
 *----------------------------------------------------------------------------*/

func EquippedCMD(set *tokenize.Set) NodeProducer {
	return Seq(
		KindCMDEquipped,
		EQUIPPED, ID,
	)(set)
}

/*------------------------------------------------------------------------------
 * Target Command
 *----------------------------------------------------------------------------*/

func TargetCMD(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(KindCMDTarget, TARGET, LeafCMD),
		Seq(KindCMDTarget, TARGET, DebuffCMD),
		Seq(KindCMDTarget, TARGET, HEALTH, LeafCMD),
	)(set)
}

/*------------------------------------------------------------------------------
 * GCD Command
 *----------------------------------------------------------------------------*/

func GCDCMD(set *tokenize.Set) NodeProducer {
	return OneOf(
		Seq(KindCMDGCD, GCD, LeafCMD),
		Seq(KindCMDGCD, GCD),
	)(set)
}

/*------------------------------------------------------------------------------
 * Leaf Command
 *----------------------------------------------------------------------------*/

func LeafCMD(set *tokenize.Set) NodeProducer {
	return OneOf(
		DURATION,
		REMAINS,
		REMAINS_EXPECTED,
		UP,
		DOWN,
		STACK,
		READY,
		ENABLED,
		REACT,
		IN,
		EXISTS,
		TTD,
		PCT,
	)(set)
}

/*------------------------------------------------------------------------------
 * Misc
 *----------------------------------------------------------------------------*/

var PrimaryValue = func(set *tokenize.Set) NodeProducer {
	return Seq(
		KindBase,
		OneOf(NUM, ID),
	)(set)
}

var Stats = func(set *tokenize.Set) NodeProducer {
	return OneOf(
		ANY_DPS,
		STRENGTH,
	)(set)
}

/*------------------------------------------------------------------------------
 * Terminals
 *----------------------------------------------------------------------------*/

var (
	terminal = Terminal
	discard  = Discard
)

var (
	ACTIONS = discard(tokenize.Actions)

	/*----------------------------------------------------------------------------
	 * Composite Actions
	 *--------------------------------------------------------------------------*/

	RUN_ACTION_LIST      = discard(tokenize.RunActionList)
	CALL_ACTION_LIST     = discard(tokenize.CallActionList)
	INVOKE_EXTERNAL_BUFF = discard(tokenize.InvokeExternalBuff)

	/*----------------------------------------------------------------------------
	 * Globals
	 *--------------------------------------------------------------------------*/

	ACTIVE_ENEMIES = terminal(KindActiveEnemies, tokenize.ActiveEnemies)
	TIME           = terminal(KindTime, tokenize.Time)
	FIGHT_REMAINS  = terminal(KindFightRemains, tokenize.FightRemains)
	TOGGLE         = discard(tokenize.Toggle)
	SNAPSHOT_STATS = terminal(KindSnapshotStats, tokenize.SnapshotStats)

	RAGE = terminal(KindRage, tokenize.Rage)

	/*----------------------------------------------------------------------------
	 * Data Tables
	 *--------------------------------------------------------------------------*/

	BUFF       = discard(tokenize.Buff)
	DEBUFF     = discard(tokenize.Debuff)
	COOLDOWN   = discard(tokenize.Cooldown)
	TALENT     = discard(tokenize.Talent)
	TARGET     = discard(tokenize.Target)
	DOT        = discard(tokenize.DOT)
	RAID_EVENT = discard(tokenize.RaidEvent)
	MOVEMENT   = discard(tokenize.Movement)
	EQUIPPED   = discard(tokenize.Equipped)
	GCD        = discard(tokenize.GCD)
	TRINKET    = discard(tokenize.Trinket)

	IF         = discard(tokenize.If)
	TARGET_IF  = discard(tokenize.TargetIf)
	NAME       = discard(tokenize.Name)
	VALUE      = discard(tokenize.VALUE)
	CONDITION  = discard(tokenize.Condition)
	VALUE_ELSE = discard(tokenize.ValueElse)
	SLOT       = discard(tokenize.Slot)
	OP         = discard(tokenize.Op)

	USE_ITEM = discard(tokenize.UseItem)
	VARIABLE = discard(tokenize.Variable)

	HEALTH   = terminal(KindHealth, tokenize.Health)
	CASTING  = terminal(KindCasting, tokenize.Casting)
	PROC     = terminal(KindProc, tokenize.Proc)
	ADDS     = terminal(KindAdds, tokenize.Adds)
	HAS_BUFF = terminal(KindHasBuff, tokenize.HasBuff)
	HAS_STAT = terminal(KindHasStat, tokenize.HasStat)

	DISTANCE         = terminal(KindDistance, tokenize.Distance)
	REACT            = terminal(KindReact, tokenize.React)
	IS               = terminal(KindIs, tokenize.Is)
	MIN              = terminal(KindMin, tokenize.Min)
	AUTO_ATTACK      = terminal(KindAutoAttack, tokenize.AutoAttack)
	REMAINS          = terminal(KindRemains, tokenize.Remains)
	REMAINS_EXPECTED = terminal(KindRemainsExpected, tokenize.RemainsExpected)
	DURATION         = terminal(KindDuration, tokenize.Duration)
	STACK            = terminal(KindStack, tokenize.Stack)
	ENABLED          = terminal(KindEnabled, tokenize.Enabled)
	RANK             = terminal(KindRank, tokenize.Rank)
	READY            = terminal(KindReady, tokenize.Ready)
	UP               = terminal(KindUp, tokenize.Up)
	DOWN             = terminal(KindDown, tokenize.Down)
	PCT              = terminal(KindPct, tokenize.Pct)
	IN               = terminal(KindIn, tokenize.In)
	EXISTS           = terminal(KindExists, tokenize.Exists)
	TTD              = terminal(KindTTD, tokenize.TTD)
	CAST_TIME        = terminal(KindCastTime, tokenize.CastTime)
	HAS_USE_BUFF     = terminal(KindHasUseBuff, tokenize.HasUseBuff)
	HAS_COOLDOWN     = terminal(KindHasCooldown, tokenize.HasCooldown)

	ANY_DPS  = terminal(KindAnyDPS, tokenize.AnyDPS)
	STRENGTH = terminal(KindStrength, tokenize.Strength)

	TRINKET_1 = discard(tokenize.Trinket1)
	TRINKET_2 = discard(tokenize.Trinket2)
	MAIN_HAND = discard(tokenize.MainHand)

	ON     = terminal(KindOn, tokenize.On)
	OFF    = terminal(KindOff, tokenize.Off)
	SET_IF = discard(tokenize.SetIf)

	/*----------------------------------------------------------------------------
	 * Logic Operators
	 *--------------------------------------------------------------------------*/

	NEGATE = terminal(KindNegate, tokenize.Not)
	OR     = terminal(KindOr, tokenize.LogicalOr)
	AND    = terminal(KindAnd, tokenize.LogicalAnd)

	/*----------------------------------------------------------------------------
	 * Comparison Operators
	 *--------------------------------------------------------------------------*/

	GE = terminal(KindGE, tokenize.GE)
	GT = terminal(KindGT, tokenize.GT)
	LE = terminal(KindLE, tokenize.LE)
	LT = terminal(KindLT, tokenize.LT)

	/*----------------------------------------------------------------------------
	 * Arithmetic Operators
	 *--------------------------------------------------------------------------*/

	MOD  = terminal(KindMod, tokenize.Mod)
	MULT = terminal(KindMult, tokenize.Mult)
	DIV  = terminal(KindDiv, tokenize.Div)
	ADD  = terminal(KindAdd, tokenize.Add)
	SUB  = terminal(KindSub, tokenize.Sub)

	/*------------------------------------------------------------------------------
	 * Assignment Operators
	 *----------------------------------------------------------------------------*/

	ASSIGN     = discard(tokenize.Assign)
	ADD_ASSIGN = discard(tokenize.AddAssign)

	/*----------------------------------------------------------------------------
	 * Misc Symbols
	 *--------------------------------------------------------------------------*/

	PARENL = discard(tokenize.ParenL)
	PARENR = discard(tokenize.ParenR)
	COLON  = terminal(KindColon, tokenize.Colon)

	/*------------------------------------------------------------------------------
	 * Base Value
	 *----------------------------------------------------------------------------*/

	NUM = terminal(KindNum, tokenize.Num)
	ID  = terminal(KindID, tokenize.ID)
)
