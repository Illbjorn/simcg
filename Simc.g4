// $antlr-format columnLimit 80
// $antlr-format useTab false
// $antlr-format indentWidth 2
// $antlr-format continuationIndentWidth 2
// $antlr-format allowShortBlocksOnASingleLine true
// $antlr-format singleLineOverrulesHangingColon false
// $antlr-format allowShortRulesOnASingleLine false
// $antlr-format ruleInternalsOnSingleLine false
// $antlr-format groupedAlignments false
// $antlr-format alignFirstTokens true
// $antlr-format alignActions true
// $antlr-format alignLabels true
// $antlr-format alignLexerCommands true
// $antlr-format alignLabels true
// $antlr-format alignTrailers true
// $antlr-format alignColons hanging
// $antlr-format alignSemicolons hanging
// $antlr-format breakBeforeParens false
// $antlr-format minEmptyLines 1
// $antlr-format reflowComments false
// $antlr-format alignTrailingComments true

grammar Simc;

////////////////////////////////////////////////////////////////////////////////
//                                   Tokens                                   //
////////////////////////////////////////////////////////////////////////////////

// This prefixes each line
ACTIONS
  : 'actions'
  ;

/*------------------------------------------------------------------------------
 * Actions
 * -------
 * Aside from actual spell executors (IDs) these are the primary instructions.
 *----------------------------------------------------------------------------*/

RUN_ACTION_LIST            // Attempts to find a suitable action from a subroutine
  : 'run_action_list'      // Stops evaluation after, even if a match is not found
  ;

CALL_ACTION_LIST           // Attempts to find a suitable action from a subroutine,
  : 'call_action_list'     // does not terminate if a match is not found
  ;

INVOKE_EXTERNAL_BUFF       // Invokes buffs the player character is not capable of
  : 'invoke_external_buff' // common example is power infusion
  ;

/*------------------------------------------------------------------------------
 * Globals
 *----------------------------------------------------------------------------*/

ACTIVE_ENEMIES             // Number of currently visible enemies
  : 'active_enemies'
  ;

TIME                       // Time in combat
  : 'time'
  ;

FIGHT_REMAINS              // Estimated time remaining in current encounter
  : 'fight_remains'
  ;

RAGE                       // Player unit's current rage
  : 'rage'
  ;

TOGGLE                     // Sets the desired state of a toggleable player spell
  : 'toggle'
  ;

SNAPSHOT_STATS
  : 'snapshot_stats'
  ;

/*------------------------------------------------------------------------------
 * Data Tables
 *----------------------------------------------------------------------------*/

BUFF                       // ex: buff.[id].up
  : 'buff'
  ;

DEBUFF                     // ex: buff.[id].remains
  : 'debuff'
  ;

COOLDOWN                   // ex: cooldown.[id].remains
  : 'cooldown'
  ;

TALENT                     // ex: talent.[id].enabled
  : 'talent'
  ;

TARGET                     // ex: target.time_to_die
  : 'target'
  ;

DOT                        // ex: dot.[id].remains
  : 'dot'
  ;

RAID_EVENT                 // ex: raid_event.adds.in
  : 'raid_event'
  ;

MOVEMENT                   // ex: movement.distance
  : 'movement'
  ;

EQUIPPED                   // ex: equipped.bestinslots
  : 'equipped'
  ;

GCD                        // ex: if=gcd.remains=0
  : 'gcd'
  ;

TRINKET                    // ex: if=variable.trinket_1_buffs
  : 'trinket.1'
  | 'trinket.2'
  ;

/*------------------------------------------------------------------------------
 * Instruction Keys
 *----------------------------------------------------------------------------*/

IF                         // ex: if=active_enemies>2
  : 'if'
  ;

TARGET_IF                  // ex: target_if=min:target.health.pct
  : 'target_if'
  ;

NAME                       // Valid for variables, run_action_list, call_action_list, use_item
  : 'name'
  ;

VALUE                      // Only valid for variables
  : 'value'
  ;

CONDITION                  // Only valid for variables
  : 'condition'
  ;

VALUE_ELSE                 // ex: value_else=false
  : 'value_else'
  ;

SLOT                       // ex: slot=trinket1
  : 'slot'
  ;

OP                         // ex: op=setif
  : 'op'
  ;

/*------------------------------------------------------------------------------
 * Actions
 *----------------------------------------------------------------------------*/

USE_ITEM                   // ex: use_item,slot=trinket1
  : 'use_item'
  ;

VARIABLE                   // ex: variable,name=foo,if=baz,value=false,value_else=true
  : 'variable'
  ;

/*------------------------------------------------------------------------------
 * Data Source Branches
 * --------------------
 * The most common occurrence of a branch would be spell IDs (ex: buff.[ID].*)
 *
 * `debuff` could be included here as well (target.debuff.*) but it is also a
 * top-level data table and therefore is listed in that section.
 *----------------------------------------------------------------------------*/

HEALTH                     // ex: target.health.pct
  : 'health'               //            ^----^
  ;

CASTING                    // ex: target.debuff.casting.react
  : 'casting'              //                   ^-----^
  ;

PROC                       // ex: trinket.2.proc.any_dps.duration
  : 'proc'                 //               ^--^
  ;

ADDS                       // ex: raid_event.adds.in
  : 'adds'                 //                ^--^
  ;

HAS_BUFF                   // ex: trinket.1.has_buff
  : 'has_buff'             //               ^------^
  ;

HAS_STAT                   // ex: trinket.1.has_stat.any_dps
  : 'has_stat'
  ;

/*------------------------------------------------------------------------------
 * Data Source Leaves
 *----------------------------------------------------------------------------*/

DISTANCE                   // ex: movement.distance
  : 'distance'
  ;

REACT                      // ex: target.debuff.casting.react
  : 'react'
  ;

IS                         // ex: trinket.1.is.algethar_puzzle_box
  : 'is'
  ;

MIN                        // ex: if=min:target.health.pct
  : 'min'
  ;

AUTO_ATTACK                // Starts auto-attack
  : 'auto_attack'
  ;

REMAINS                    // ex: buff.[ID].remains
  : 'remains'
  ;

REMAINS_EXPECTED           // ex: cooldown.[ID].remains_expected
  : 'remains_expected'
  ;

DURATION                   // ex: trinket.1.proc.any_dps.duration
  : 'duration'
  ;

STACK                      // ex: buff.[ID].stack
  : 'stack'
  ;

ENABLED                    // ex: talent.[ID].enabled
  : 'enabled'
  ;

READY                      // ex: cooldown.[ID].ready
  : 'ready'
  ;

UP                         // ex: buff.[ID].up
  : 'up'
  ;

DOWN                       // ex: buff.[ID].down
  : 'down'
  ;

PCT                        // ex: target.health.pct
  : 'pct'
  ;

IN                         // ex: raid_event.adds.in
  : 'in'
  ;

EXISTS                     // ex: raid_event.adds.exists
  : 'exists'
  ;

TTD                        // ex: target.time_to_die
  : 'time_to_die'
  ;

CAST_TIME                  // ex: trinket.1.cast_time
  : 'cast_time'
  ;

HAS_USE_BUFF               // ex: trinket.1.has_use_buff
  : 'has_use_buff'
  ;

HAS_COOLDOWN               // ex: trinket.1.has_cooldown
  : 'has_cooldown'
  ;

ANY_DPS                    // ex: trinket.1.has_stat.any_dps
  : 'any_dps'
  ;

STRENGTH                   // ex: trinket.1.has_stat.strength
  : 'strength'
  ;

/*------------------------------------------------------------------------------
 * Equipment Slots
 *----------------------------------------------------------------------------*/

TRINKET_1                  // ex: use_item,slot=trinket1
  : 'trinket1'
  ;

TRINKET_2                  // ex: use_item,slot=trinket2
  : 'trinket2'
  ;

MAIN_HAND                  // ex: use_item,slot=main_hand
  : 'main_hand'
  ;

/*------------------------------------------------------------------------------
 * States // Directives
 *----------------------------------------------------------------------------*/

ON                         // ex: toggle=on
  : 'on'
  ;

OFF                        // ex: toggle=off
  : 'off'
  ;

SET_IF                     // ex: op=setif
  : 'setif'
  ;

/*------------------------------------------------------------------------------
 * Symbols
 *----------------------------------------------------------------------------*/

// Arithmetic
MOD
  : '%%'
  ;

MULT
  : '*'
  ;

DIV
  : '%'
  ;

ADD
  : '+'
  ;

SUB
  : '-'
  ;

// Logical
NEGATE
  : '!'
  ;

L_OR
  : '|'
  ;

L_AND
  : '&'
  ;

// Comparison
GE
  : '>='
  ;

LE
  : '<='
  ;

GT
  : '>'
  ;

LT
  : '<'
  ;

// Misc
PARENL
  : '('
  ;

PARENR
  : ')'
  ;

ASSIGN
  : '='
  ;

ADD_ASSIGN
  : '+=/'
  ;

OCTOTHORPE                 // Line comment
  : '#'
  ;

COLON
  : ':'
  ;

/*------------------------------------------------------------------------------
 * "Composite" Tokens
 *----------------------------------------------------------------------------*/

NUM                        // int/floats ex: 1, 1.0, 1.115
  : [0-9][0-9.]*
  ;

ID                         // ex: mortal_strike
  : [a-z][a-z0-9_]+        // TODO: Are numbers actually needed?
  ;

/*------------------------------------------------------------------------------
 * Discards
 * --------
 * These tokens serve no real purpose to the parse.
 *----------------------------------------------------------------------------*/

PD
  : [.]                    -> skip
  ;

COMMA
  : [,]                    -> skip
  ;

WS
  : [ \t\n\r\f]+           -> skip
  ;

////////////////////////////////////////////////////////////////////////////////
//                                   Parser                                   //
////////////////////////////////////////////////////////////////////////////////

program
  : instruction+ EOF
  ;

instruction
  : ACTIONS ASSIGN statement
  | ACTIONS ADD_ASSIGN statement
  | ACTIONS ID ASSIGN statement
  | ACTIONS ID ADD_ASSIGN statement
  ;

statement
  : run_action_list
  | call_action_list
  | invoke_external_buff
  | use_item
  | var
  | executor
  | toggle
  | ID
  | SNAPSHOT_STATS
  | AUTO_ATTACK
  ;

// Run Action List

run_action_list
  : RUN_ACTION_LIST NAME ASSIGN ID target_if? conditional?
  ;

// Call Action List

call_action_list
  : CALL_ACTION_LIST NAME ASSIGN ID conditional?
  ;

// Invoke External Buff

invoke_external_buff
  : INVOKE_EXTERNAL_BUFF NAME ASSIGN ID conditional?
  ;

// Use Item

use_item
  : USE_ITEM SLOT ASSIGN item IF ASSIGN expr
  | USE_ITEM NAME ASSIGN ID IF ASSIGN expr
  ;

item
  : MAIN_HAND
  | TRINKET_1
  | TRINKET_2
  ;

// Variable

var
  : VARIABLE var_name var_set_if? var_value var_value_else? var_cond?
  ;

var_name
  : NAME ASSIGN ID
  ;

var_set_if
  : OP ASSIGN SET_IF
  ;

var_value
  : VALUE ASSIGN expr
  ;

var_value_else
  : VALUE_ELSE ASSIGN expr
  ;

var_cond
  : CONDITION ASSIGN expr
  ;

// Executor (i.e. spell casts)

executor
  : ID IF ASSIGN expr
  ;

// Toggle

toggle
  : ID TOGGLE ASSIGN (ON | OFF)
  ;

target_if
  : TARGET_IF ASSIGN expr
  ;

// Conditional

conditional
  : (IF ASSIGN expr)+
  ;

// Conditions

expr
  : expr_and (L_OR expr_and)*
  ;

expr_and
  : expr_compare (L_AND expr_compare)*
  ;

expr_compare
  : expr_mult (
    (ASSIGN | GE | LE | GT | LT) expr_mult
  )*
  ;

expr_mult
  : expr_add ((MULT | DIV | MOD) expr_add)*
  ;

expr_add
  : expr_prefix ((ADD | SUB) expr_prefix)*
  ;

expr_prefix
  : NEGATE? base
  ;

// Basic Values
base
  : PARENL expr PARENR
  | builtin
  | command
  | NUM
  | ID
  ;

// Builtins

builtin
  : MIN COLON expr
  | TIME
  | FIGHT_REMAINS
  | RAGE
  | TOGGLE
  | SNAPSHOT_STATS
  | ACTIVE_ENEMIES
  ;

// Commands

command
  : trinket_cmd
  | cooldown_cmd
  | variable_cmd
  | dot_cmd
  | buff_cmd
  | debuff_cmd
  | talent_cmd
  | movement_cmd
  | raid_event_cmd
  | equipped_cmd
  | target_cmd
  | gcd_cmd
  ;

// Trinket Commands

trinket_cmd
  : TRINKET IS ID
  | TRINKET CAST_TIME
  | TRINKET HAS_USE_BUFF
  | TRINKET HAS_COOLDOWN
  | TRINKET HAS_BUFF stats
  | TRINKET COOLDOWN leaf_cmd
  | TRINKET HAS_STAT stats
  | TRINKET PROC stats leaf_cmd
  ;

stats
  : ANY_DPS
  | STRENGTH
  ;

// Cooldown Commands

cooldown_cmd
  : COOLDOWN ID leaf_cmd
  ;

// Variable Commands

variable_cmd
  : VARIABLE ID
  ;

// DoT Commands

dot_cmd
  : DOT ID leaf_cmd
  ;

// Buff Commands

buff_cmd
  : BUFF ID leaf_cmd
  ;

// Debuff Commands

debuff_cmd
  : DEBUFF ID leaf_cmd
  | DEBUFF CASTING leaf_cmd
  ;

// Talent Commands

talent_cmd
  : TALENT ID leaf_cmd
  | TALENT ID
  ;

// Movement Commands

movement_cmd
  : MOVEMENT (DISTANCE)
  ;

// Raid Event Commands

raid_event_cmd
  : RAID_EVENT ADDS leaf_cmd
  ;

// Equipped Commands

equipped_cmd
  : EQUIPPED ID
  ;

// Target Commands

target_cmd
  : TARGET leaf_cmd
  | TARGET debuff_cmd
  | TARGET HEALTH leaf_cmd
  ;

// GCD Commands

gcd_cmd
  : GCD leaf_cmd
  | GCD
  ;

// Command Leaves

leaf_cmd
  : DURATION
  | REMAINS
  | REMAINS_EXPECTED
  | UP
  | DOWN
  | STACK
  | READY
  | ENABLED
  | REACT
  | IN
  | EXISTS
  | PCT
  | TTD
  ;
