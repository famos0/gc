#!/usr/bin/env bash
_gc() {
    COMPREPLY=()

    local MODES=("bundle" "pattern")

    declare -A TEMPLATES
    TEMPLATES[pattern]="\$(gc pattern -l)"
    TEMPLATES[bundle]="\$(gc bundle -l)"

    local cur=${COMP_WORDS[COMP_CWORD]}
    if [ ${TEMPLATES[$3]+1} ] ; then
        COMPREPLY=( `compgen -W "${TEMPLATES[$3]}" -- $cur` )
    else 
        COMPREPLY=( `compgen -W "${MODES[*]}" -- $cur` )
    fi
}

complete -F _gc gc