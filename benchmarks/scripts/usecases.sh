#!/bin/bash

# Naming convention: <d>.<protocol> contains <d>.<protocol>/goscr/<Protocol>.scr
BENCHMARKS="01.ring \
            02.pipeline \
            03.recPipeline \
            04.recTree \
            05.recTaskGen \
            06.forkJoin \
            07.recForkJoin \
            09.unboundedFib \
            15.taskGen \
            16.dns \
            17.min-max"

# Ignoring benchmarks
# 08.boundedFib -> ../1.boundedFib
# 10.fannkuch -> ../3.fannkuch
# 11.spectralnorm -> ../6.spectralnorm
# 12.regex -> ../5.regex
# 13.knuc -> ../4.knuc
# 14.primeSieve -> ../2.boundedPrimeSieve

pushd () {
    command pushd "$@" > /dev/null
}

popd () {
    command popd "$@" > /dev/null
}

# Get path to script (should be under <benchmarks>/scripts)                                                            
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"

BENCH_DIR=${DIR}/..                                                                                                    
NUSCR_BIN=${BENCH_DIR}/nuscr_bin/bin/nuscr

[ -f ${NUSCR_BIN} ] || ${DIR}/compile_nuscr.sh

for bench in ${BENCHMARKS}
do
    echo "* Entering '${bench}'"
    echo
    GOSCR_DIR=${BENCH_DIR}/use_cases/${bench}
    pushd ${GOSCR_DIR}
    BENCHNAME=$(echo ${bench} | sed 's/.*\.//')
    SCR_FILE=$(echo ${BENCHNAME} | sed -e 's/\b\(.\)/\u\1/g').scr
    echo
    echo "(**********************************************)"
    echo "(** File '${SCR_FILE}' *)"
    echo
    cat ${SCR_FILE}
    echo
    echo "(* End File*)"
    echo "(**********************************************)"
    echo
    echo
    echo "* Generating code for '${SCR_FILE}'"
    GO_IMPL=${GOSCR_DIR}/${BENCHNAME}/${BENCHNAME}.go
    rm -rf ${GOSCR_DIR}/${BENCHNAME}
    mkdir -p ${GOSCR_DIR}/${BENCHNAME}
    echo "package ${BENCHNAME}" > ${GO_IMPL}
    echo >> ${GO_IMPL}
    echo "import \"sync\"" >> ${GO_IMPL}
    echo >> ${GO_IMPL}
    echo >> ${GO_IMPL}
    ${NUSCR_BIN} ${SCR_FILE} --gencode-go=${GOSCR_DIR}/${bench} >> ${GO_IMPL}
done

echo "* All Done"
