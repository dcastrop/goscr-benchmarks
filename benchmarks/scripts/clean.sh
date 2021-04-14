#!/bin/bash            
                       
# Get path to script (should be under <benchmarks>/scripts)                                                            
SOURCE="${BASH_SOURCE[0]}"               
while [ -h "$SOURCE" ]; do               
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"         
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done                   
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
                       
BENCH_DIR=${DIR}/..                                                                                                    
NUSCR_BIN_DIR=${BENCH_DIR}/nuscr_bin

rm -rf ${NUSCR_BIN_DIR}
                       

