#!/bin/bash

ls */go.mod | while read gomodfile; do
    dir="$(dirname "${gomodfile}")"
    echo "[${dir}]" 1>&2
    pushd "${dir}" >/dev/null
    $*
    popd >/dev/null
    echo 1>&2
done
