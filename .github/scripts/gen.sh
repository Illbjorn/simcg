#!/usr/bin/env bash

################################################################################
#
# >> Overview
#
# This script identifies all Go files under provided input path (env: 'root')
# containing Go generate ('//go:generate') directives with '--output' flag
# values (primarily for `stringer`).
#
# Findings are used to generate a Taskfile
# (https://taskfile.dev)(taskfile.gen.yml) which invokes the go generate process
# only after changes to the source or generated output.
#
# >> Inputs
#
# root : REQUIRED
#  A path, under which all Go files will be evaluated for go:generate directives
#
################################################################################
# Inputs

require () {
  if [[ ! -v ${1} ]]; then
    echo "ERROR: Required input ['${1}'] is not set."
    exit 1
  fi
}

# 'root' - the directory under which we'll evaluate all Go files for
# go:generate directives
require "root"

################################################################################
# Bootstrap Taskfile

fprintln () {
  echo "${1}" >> taskfile.gen.yml
}

fprint () {
  echo -n "${1}" >> taskfile.gen.yml
}

# Blow away the existing file, if it exists
rm -rf "taskfile.gen.yml"

# shellcheck disable=SC2016 # We don't want expansion
fprintln '# yaml-language-server: $schema=https://taskfile.dev/schema.json'
fprintln ""
fprintln "################################################################################"
fprintln "#                                                                              #"
fprintln "#                            GENERATED. DO NOT EDIT.                           #"
fprintln "#                                                                              #"
fprintln "################################################################################"
fprintln ""
fprintln "version: '3'"
fprintln ""
fprintln "interval: 250ms"
fprintln ""
fprintln "tasks:"

################################################################################
# Generate Tasks

gen_task () {
  task_name="${1}"
  input="${2}"
  outputs="${3}"
  fprintln "  ${task_name}:"
  fprintln "    sources: ['${input}']"
  fprintln "    generates: [${outputs}]"
  fprintln "    cmd: go generate \"${input}\""
  fprintln ""
}

declare -a tasks=()
shopt -s globstar
# shellcheck disable=SC2154 # Provided as script input
for f in "${root}"/**/*.go; do
  parent="$(dirname "${f}")"          # Get .go file parent directory
  base_dir="$(basename "${parent}")"  # Get the parent directory's name
  base_file="$(basename "${f}")"      # Get the file name
  base_file="${base_file%.*}"         # Remove the file name's extension

  # Look for go:generate directives
  output=$(grep -Po '//go:generate.*--output \K.*' "${f}")
  if [[ -n $output ]]; then
    # Prepare the output string
    outputs=""                    #
    for o in $output; do          # There may be multiple output values
      outputs+="'${parent}/${o}',"  # Join the output to the root dir path and "concatenate" (trailing comma)
    done                          #
    outputs="${outputs%?}"        # Remove the trailing comma

    # Prepare the task name
    task_name="gen-${base_dir}-${base_file}"  # ex: gen-apl-base_value
    task_name="${task_name//_/-}"             # ex: gen-apl-base-value

    # Add to the task list for later
    tasks+=("${task_name}")

    # Generate and write the task
    gen_task "${task_name}" "${f}" "$outputs"
  fi
done

# Generate meta task
fprintln "  all:"
fprintln "    cmds:"
for task in "${tasks[@]}"; do
  fprintln "      - task: ${task}"
done
