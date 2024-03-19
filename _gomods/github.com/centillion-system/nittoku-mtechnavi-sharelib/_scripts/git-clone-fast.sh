#!/bin/bash

set -e -u

repositoryurl="${1}"
branch="${2}"
destdir="${3}"

git clone \
    --branch "${branch}" \
    --single-branch \
    --depth 1 \
    --recurse-submodules \
    --shallow-submodules \
    "${repositoryurl}" \
    "${destdir}"
