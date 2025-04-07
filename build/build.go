package build

import (
	"fmt"
	"strconv"

	"github.com/illbjorn/hades/simcg/parse"
)

func Build(root *parse.Node) *APL {
	list := new(APL)

	for child := range parse.Children(root) {
		switch child.Kind {
		case parse.KindInstruction:
			list.Instructions = append(list.Instructions, instruction(child))
		default:
			kindErr(child)
		}
	}

	return list
}

func instruction(n *parse.Node) *Instruction {
	in := new(Instruction)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindID:
			in.ID = parse.ValueS(child)

		case parse.KindSnapshotStats:
			in.Kind = KindSnapshotStats

		case parse.KindVar:
			in.Kind = KindVariable
			in.Variable = varDefinition(child)

		case parse.KindToggle:
			in.Kind = KindToggle
			in.Toggle = toggle(child)

		case parse.KindExecutor:
			in.Kind = KindExecutor
			in.Executor = executor(child)

		case parse.KindAutoAttack:
			in.Kind = KindAutoAttack

		case parse.KindCallActionList:
			in.Kind = KindCallActionList
			in.CallActionList = callActionList(child)

		case parse.KindRunActionList:
			in.Kind = KindRunActionList
			in.RunActionList = runActionList(child)

		case parse.KindUseItem:
			in.Kind = KindUseItem
			in.UseItem = useItem(child)

		case parse.KindInvokeExternalBuff:

		default:
			kindErr(child)
		}
	}

	return in
}

func varDefinition(n *parse.Node) *VarDefinition {
	variable := new(VarDefinition)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindVarName:
			variable.Name = string(parse.Value(child.Children[0]))

		case parse.KindVarValue:
			variable.Value = command(child.Children[0])

		default:
			kindErr(child)
		}
	}

	return variable
}

func command(n *parse.Node) *Command {
	c := new(Command)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindExpr:
			c.Link(command(child))

		case parse.KindLogical:
			c.Link(command(child))

		case parse.KindCompare:
			c.Link(command(child))

		case parse.KindArithmetic1:
			c.Link(command(child))

		case parse.KindArithmetic2:
			c.Link(command(child))

		case parse.KindExprPrefix:
			c.Link(command(child))

		case parse.KindGroupedExpr:
			c.Link(command(child))

		case parse.KindBase:
			c.Link(command(child))

		case parse.KindID:
			c.Link(
				&Command{
					Kind: CMDTypeBase,
					Base: &BaseValue{
						Value: parse.ValueS(child),
						Kind:  BaseID,
					},
				},
			)
		case parse.KindNum:
			c.Link(
				&Command{
					Kind: CMDTypeBase,
					Base: &BaseValue{
						Value: parse.ValueS(child),
						Kind:  BaseNum,
					},
				},
			)

		case parse.KindTrinketCmd:
			c.Link(
				&Command{
					Trinket: trinket(child),
					Kind:    CMDTypeTrinket,
				},
			)

		case parse.KindBuiltin:
			c.Link(
				&Command{
					Builtin: builtin(child),
					Kind:    CMDTypeBuiltin,
				},
			)

		case parse.KindCMDVariable:
			c.Link(
				&Command{
					Variable: varReference(child),
					Kind:     CMDTypeVariable,
				},
			)

		case parse.KindCMDGCD:
			c.Link(
				&Command{
					GCD:  gcd(child),
					Kind: CMDTypeGCD,
				},
			)

		case parse.KindCMDTarget:
			c.Link(
				&Command{
					Target: target(child),
					Kind:   CMDTypeTarget,
				},
			)

		case parse.KindCMDCooldown:
			c.Link(
				&Command{
					Cooldown: cooldown(child),
					Kind:     CMDTypeCooldown,
				},
			)

		case parse.KindCMDBuff:
			c.Link(
				&Command{
					Buff: buff(child),
					Kind: CMDTypeBuff,
				},
			)

		case parse.KindCMDDebuff:
			c.Link(
				&Command{
					Debuff: debuff(child),
					Kind:   CMDTypeDebuff,
				},
			)

		case parse.KindCMDDOT:
			c.Link(
				&Command{
					Dot:  dot(child),
					Kind: CMDTypeDot,
				},
			)

		case parse.KindCMDTalent:
			c.Link(
				&Command{
					Talent: talent(child),
					Kind:   CMDTypeTalent,
				},
			)

		case parse.KindCMDRaidEvent:
			c.Link(
				&Command{
					RaidEvent: raidEvent(child),
					Kind:      CMDTypeRaidEvent,
				},
			)

		case parse.KindCMDEquipped:
			c.Link(
				&Command{
					Equipped: equipped(child),
					Kind:     CMDTypeEquipped,
				},
			)

		case parse.KindCMDMovement:
			c.Link(
				&Command{
					Movement: movement(child),
					Kind:     CMDTypeMovement,
				},
			)

		case parse.KindGE:
			c.Link(
				&Command{
					Op:   OpGE,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindLE:
			c.Link(
				&Command{
					Op:   OpLE,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindLT:
			c.Link(
				&Command{
					Op:   OpLT,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindGT:
			c.Link(
				&Command{
					Op:   OpGT,
					Kind: CMDTypeOp,
				},
			)
			// case node.NE: // TODO

		case parse.KindAnd:
			c.Link(
				&Command{
					Op:   OpAnd,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindOr:
			c.Link(
				&Command{
					Op:   OpOr,
					Kind: CMDTypeOp,
				},
			)

		case parse.KindMod:
			c.Link(
				&Command{
					Op:   OpMod,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindMult:
			c.Link(
				&Command{
					Op:   OpMult,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindDiv:
			c.Link(
				&Command{
					Op:   OpDiv,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindAdd:
			c.Link(
				&Command{
					Op:   OpAdd,
					Kind: CMDTypeOp,
				},
			)
		case parse.KindSub:
			c.Link(
				&Command{
					Op:   OpSub,
					Kind: CMDTypeOp,
				},
			)

		default:
			fmt.Println(" ++", child.Kind)
			fmt.Println()
			fmt.Println("PARENT:")
			fmt.Println(n)
			fmt.Println()
			fmt.Println("CHILD:")
			fmt.Println(child)
			kindErr(child)
		}
	}

	return c
}

func trinket(n *parse.Node) *Trinket {
	cmd := new(Trinket)

	for child := range parse.Children(n) {
		value := parse.Value(child)
		switch child.Kind {
		case parse.KindID:
			cmd.ID = string(value)

		case parse.KindNum:
			slot, err := strconv.Atoi(string(value))
			if err != nil {
				panic(err)
			}
			switch slot {
			case 1:
				cmd.Slot = TrinketSlot1
			case 2:
				cmd.Slot = TrinketSlot2
			default:
				panic("Unknown slot " + string(value))
			}

		case parse.KindIs:
			cmd.Kind = TrinketIs
		case parse.KindHasUseBuff:
			cmd.Kind = TrinketHasUseBuff
		case parse.KindHasStat:
			cmd.Kind = TrinketHasStat
		case parse.KindCastTime:
			if cmd.Kind > 0 {
				panic(cmd)
			}
			cmd.Kind = TrinketCastTime
		case parse.KindHasCooldown:
			cmd.Kind = TrinketHasCooldown
		case parse.KindRemains:
			cmd.Kind = TrinketRemains
		case parse.KindProc:
			cmd.Kind = TrinketProc
		case parse.KindDuration:
			cmd.Kind = TrinketDuration

		case parse.KindAnyDPS:
			cmd.Resource = ResourceAnyDPS

		default:
			kindErr(child)
		}
	}

	return cmd
}

func builtin(n *parse.Node) *Builtin {
	b := new(Builtin)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindTime:
			b.Kind = BuiltinTime

		case parse.KindFightRemains:
			b.Kind = BuiltinFightRemains
		case parse.KindActiveEnemies:
			b.Kind = BuiltinActiveEnemies

		case parse.KindRage:
			b.Kind = BuiltinResource
			b.Resource = ResourceRage

		default:
			kindErr(child)
		}
	}
	return b
}

func toggle(n *parse.Node) *Toggle {
	t := new(Toggle)

	for child := range parse.Children(n) {
		value := parse.ValueS(child)
		switch child.Kind {
		case parse.KindID:
			t.ID = value

		case parse.KindOn:
			t.State = ON
		case parse.KindOff:
			t.State = OFF

		default:
			kindErr(child)
		}
	}

	return t
}

func executor(n *parse.Node) *Executor {
	exec := new(Executor)

	for child := range parse.Children(n) {
		value := parse.ValueS(child)
		switch child.Kind {
		case parse.KindID:
			exec.ID = value

		case parse.KindExpr:
			exec.Command = command(child)

		default:
			kindErr(child)
		}
	}

	return exec
}

func varReference(n *parse.Node) *VarReference {
	cmd := new(VarReference)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindID:
			cmd.ID = parse.ValueS(child)

		default:
			kindErr(child)
		}
	}

	return cmd
}

func gcd(n *parse.Node) *GCD {
	cmd := new(GCD)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindRemains:
			cmd.Kind = GCDRemains

		default:
			kindErr(child)
		}
	}

	return cmd
}

func target(n *parse.Node) *Target {
	cmd := new(Target)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindCMDDebuff:
			cmd.Debuff = debuff(child)

		case parse.KindTTD:
			cmd.Kind = TargetTTD
		case parse.KindHealth:
			cmd.Kind = TargetHealth
		case parse.KindPct:
			cmd.Filter = TargetFilterPct

		default:
			kindErr(child)
		}
	}

	return cmd
}

func cooldown(n *parse.Node) *Cooldown {
	cmd := new(Cooldown)

	for child := range parse.Children(n) {
		value := parse.ValueS(child)

		switch child.Kind {
		case parse.KindID:
			cmd.ID = value

		case parse.KindRemains:
			cmd.Kind = CooldownRemains
		case parse.KindRemainsExpected:
			cmd.Kind = CooldownRemainsExpected
		case parse.KindReady:
			cmd.Kind = CooldownReady

		default:
			kindErr(child)
		}
	}

	return cmd
}

func buff(n *parse.Node) *Buff {
	cmd := new(Buff)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindID:
			cmd.ID = parse.ValueS(child)

		case parse.KindUp:
			cmd.Kind = BuffUp
		case parse.KindDown:
			cmd.Kind = BuffDown
		case parse.KindRemains:
			cmd.Kind = BuffRemains
		case parse.KindRemainsExpected:
			cmd.Kind = BuffRemainsExpected
		case parse.KindStack:
			cmd.Kind = BuffStack

		case parse.KindCasting:
		case parse.KindReact:

		default:
			kindErr(child)
		}
	}

	return cmd
}

func debuff(n *parse.Node) *Debuff {
	cmd := new(Debuff)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindID:
			cmd.ID = parse.ValueS(child)

		case parse.KindUp:
			cmd.Kind = DebuffUp
		case parse.KindDown:
			cmd.Kind = DebuffDown
		case parse.KindRemains:
			cmd.Kind = DebuffRemains
		case parse.KindRemainsExpected:
			cmd.Kind = DebuffRemainsExpected
		case parse.KindStack:
			cmd.Kind = DebuffStack

		case parse.KindCasting:
		case parse.KindReact:

		default:
			kindErr(child)
		}
	}

	return cmd
}

func dot(n *parse.Node) *DOT {
	cmd := new(DOT)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindID:
			cmd.ID = parse.ValueS(child)

		case parse.KindUp:
			cmd.Kind = DotUp
		case parse.KindDown:
			cmd.Kind = DotDown
		case parse.KindRemains:
			cmd.Kind = DotRemains
		case parse.KindRemainsExpected:
			cmd.Kind = DotRemainsExpected

		case parse.KindCasting:
		case parse.KindReact:

		default:
			kindErr(child)
		}
	}

	return cmd
}

func talent(n *parse.Node) *Talent {
	cmd := new(Talent)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindID:
			cmd.ID = parse.ValueS(child)

		case parse.KindEnabled:
			cmd.Kind = TalentEnabled
		case parse.KindRank:
			cmd.Kind = TalentRank

		default:
			kindErr(child)
		}
	}

	return cmd
}

func raidEvent(n *parse.Node) *RaidEvent {
	cmd := new(RaidEvent)

	for child := range parse.Children(n) {
		switch child.Kind {
		// Raid Event Commands
		case parse.KindAdds:
			cmd.Kind = RaidEventAdds

			// Raid Event Filters
		case parse.KindIn:
			cmd.Filter = RaidEventFilterIn
		case parse.KindExists:
			cmd.Filter = RaidEventFilterExists
		case parse.KindRemains:
			cmd.Filter = RaidEventFilterRemains

		default:
			kindErr(child)
		}
	}

	return cmd
}

func equipped(n *parse.Node) *Equipped {
	cmd := new(Equipped)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindID:
			cmd.ID = parse.ValueS(child)

		default:
			kindErr(child)
		}
	}

	return cmd
}

func movement(n *parse.Node) *Movement {
	move := new(Movement)

	for child := range parse.Children(n) {
		switch child.Kind {
		case parse.KindDistance:
			move.Kind = MovementDistance

		default:
			kindErr(child)
		}
	}

	return move
}

func callActionList(n *parse.Node) *CallActionList {
	call := new(CallActionList)

	for child := range parse.Children(n) {
		value := parse.ValueS(child)
		switch child.Kind {
		case parse.KindID:
			call.ID = value

		default:
			kindErr(child)
		}
	}

	return call
}

func runActionList(n *parse.Node) *RunActionList {
	run := new(RunActionList)

	for child := range parse.Children(n) {
		value := parse.ValueS(child)
		switch child.Kind {
		case parse.KindID:
			run.ID = value

		default:
			kindErr(child)
		}
	}

	return run
}

func useItem(n *parse.Node) *UseItem {
	item := new(UseItem)

	for child := range parse.Children(n) {
		value := parse.ValueS(child)
		switch child.Kind {
		case parse.KindID:
			item.ID = value

		case parse.KindExpr:
			item.Conditions = command(child)

		default:
			fmt.Println(child)
			kindErr(child)
		}
	}

	return item
}
