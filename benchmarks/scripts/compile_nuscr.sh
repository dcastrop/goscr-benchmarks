#!/bin/bash

# Get path to script (should be under <benchmarks>/scripts)
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"

# Assumes the following layout, if NUSCR_ROOT is not defined
#   <parent_dir>/
#       ${BENCH_DIR}/
#           ${DIR}/
#               ${0} 
#       ${NUSCR_ROOT}/
BENCH_DIR=${DIR}/..
NUSCR_ROOT=${NUSCR_ROOT:-${BENCH_DIR}/../nuscr}


# Target directory for the nuscr binaries
NUSCR_BIN_DIR=${BENCH_DIR}/nuscr_bin

mkdir -p ${NUSCR_BIN_DIR}

pushd ${NUSCR_ROOT}
dune build 
dune install --relocatable --prefix=${NUSCR_BIN_DIR}
popd
