package build

type Command struct {
	Op        Op            `json:"op,omitempty"`
	Kind      CommandKind   `json:"kind,omitempty"`
	Target    *Target       `json:"target,omitempty"`
	Buff      *Buff         `json:"buff,omitempty"`
	Debuff    *Debuff       `json:"debuff,omitempty"`
	RaidEvent *RaidEvent    `json:"raid_event,omitempty"`
	Talent    *Talent       `json:"talent,omitempty"`
	GCD       *GCD          `json:"gcd,omitempty"`
	Cooldown  *Cooldown     `json:"cooldown,omitempty"`
	Trinket   *Trinket      `json:"trinket,omitempty"`
	Dot       *DOT          `json:"dot,omitempty"`
	Equipped  *Equipped     `json:"equipped,omitempty"`
	Movement  *Movement     `json:"movement,omitempty"`
	Builtin   *Builtin      `json:"builtin,omitempty"`
	Variable  *VarReference `json:"var,omitempty"`
	Base      *BaseValue    `json:"base,omitempty"`
	Next      *Command      `json:"next,omitempty"`
}

func (c *Command) Link(next *Command) {
	switch {
	case next == nil:
		// Nothing to do

	case c.Kind == 0:
		*c = *next

	case c.Next == nil:
		c.Next = next

	default:
		c.Next.Link(next)
	}
}

//go:generate stringer -type CommandKind --output zz_command_kind_string.go
type CommandKind uint8

const (
	CMDTypeOp CommandKind = 1 + iota
	CMDTypeTarget
	CMDTypeBuff
	CMDTypeDebuff
	CMDTypeRaidEvent
	CMDTypeTalent
	CMDTypeGCD
	CMDTypeCooldown
	CMDTypeTrinket
	CMDTypeDot
	CMDTypeEquipped
	CMDTypeMovement
	CMDTypeBuiltin
	CMDTypeVariable
	CMDTypeBase
)
