#!/bin/bash

RED='\033[1;31m'
GREEN='\033[1;32m'
BOLD='\033[1m'
UNDERLINE='\033[4m'
ENDSTYLE='\033[0m' # reset style

reset



echo -e "########################################################################"
echo -e 
echo -e "* This is an automated 'Getting Started' guide for the artifact of paper:"
echo -e 
echo -e "\t${GREEN}#21: \"Dynamically Updatable Multiparty Session Protocols\"\033[0m"
echo -e 
echo -e 

sleep 2
echo -e "* The technique described in this paper takes _multiparty protocol \
specifications_, and produces concurrent Go code."

sleep 2
echo -e "* This script will walk through the code contents, and run \
step-by-step an example from the paper."
echo -e
sleep 2
echo -e "* The current directory is: '${PWD}'"
echo -e
tree -L 1 ${PWD}
echo -e
echo -e "\t- ./benchmarks  _____ The directory with the benchmarks used in the paper"
echo -e "\t- ./nuscr ___________ The directory with the code of our tool"
echo -e
echo -e
echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
read

reset
echo -e "* GoScr is an extension of NuScr that allows the specification of \
multiparty protocols with participants that may join after a session \
has been initialised. These are called ${BOLD}DUMST${ENDSTYLE} in the paper."
echo -e

sleep 2
echo -e "* The script will now navigate to the directory that contains \
the source code of GoScr."
echo -e
sleep 3
pushd ${PWD}/nuscr
  tree -L 1 ${PWD}
  
  sleep 2
  echo -e
  echo -e
  echo -e "* The most relevant files in our implementation are:"
  echo -e
  echo -e "\t- ./lib/syntax.ml _________ Syntax of GoScr protocols"
  echo -e "\t- ./lib/gtype.ml __________ DUMST global types"
  echo -e "\t- ./lib/local.ml __________ DUMST local types"
  echo -e "\t- ./lib/gocodegen.ml ______ Concurrent Go code generator"
  echo -e

  sleep 2
  echo -e
  echo -e "############ File: './lib/syntax.ml'"
  echo -e
  awk 'NR >= 142 && NR <= 158 { print "L-" NR ": " $0}' ./lib/syntax.ml
  echo -e
  echo -e "########################################################################"

  sleep 2
  echo -e "* Command to build our tool:"
  echo -e
  echo -e "$ dune build"
  echo -e

  sleep 2
  TARGET=$(cd ../benchmarks/nuscr_bin; pwd)
  echo -e "* Command to install our tool in '${TARGET}':"
  echo -e
  echo -e "$ dune install --relocatable --prefix=${TARGET}"
  echo -e
  echo -e "${RED}IMPORTANT${ENDSTYLE}: if you change the target directory\
'--prefix', and it is outside the current working directory, you must specify\
the full path."
  echo -e
  echo -e
  echo -e "* We provide a script in the benchmarks directory to compile and install our tool."
  sleep 3
  TARGET=$(cd ../benchmarks; pwd)
  echo -e "* We will now leave ${PWD}, and move to ${TARGET}"
  echo -e
  echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
  read
popd


pushd ${PWD}/benchmarks
  tree -L 1 ${PWD}

  sleep 2
  echo -e
  echo -e
  echo -e "* The most relevant files are:"
  echo -e
  echo -e "\t- '1.boundedFib'"
  echo -e "\t- ..."
  echo -e "\t- '7.quicksort' ___________ The benchmarks in Figure 7 on page 20"
  echo -e "\t- 'nuscr_bin' _____________ The target directory to install our tool"
  echo -e "\t- 'run_bench' _____________ Go program to run our benchmarks"
  echo -e "\t- 'scripts' _______________ Scripts to compile/run our benchmarks"
  echo -e "\t- 'use_cases' _____________ Protocol specifications in Table 1 on page 22"
  echo -e
  echo -e
  echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
  read
  echo -e
  echo -e "* To run our benchmarks, we first need to compile and install our tool."
  echo -e
  echo -e "$ ./scripts/compile_nuscr.sh"
  sleep 2
  ./scripts/compile_nuscr.sh > /dev/null 2>&1
  echo -e
  TARGET=$(cd ../benchmarks/nuscr_bin; pwd)
  echo -e "* Done. Directory '${TARGET}/bin' contains our binary:"
  echo -e
  ls -l ${TARGET}/bin
  sleep 2
  echo -e
  TARGET=$(cd ../benchmarks/nuscr_bin; pwd)
  echo -e "* Run the following command to check the options:"
  echo -e
  echo -e "$ ${TARGET}/bin/nuscr --help"

  echo -e
  echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
  read
  reset
  echo -e "* We have now compiled and installed our tool. Let's now illustrate \
one example:"
  sleep 2
  echo -e
  pushd ${PWD}/4.knuc
    tree -L 1 ${PWD}
    sleep 2
    echo -e
    echo -e "\t- 'base' ________ directory that contains the baseline implementation from CLBG: 'https://benchmarksgame-team.pages.debian.net'"
    echo -e "\t- 'goscr' _______ our GoScr implementation"
    echo -e "\t- 'knuc.go' _____ benchmarking code for this example"
  
    pushd ${PWD}/goscr
      echo -e
      echo -e
      tree -L 1 ${PWD}
      echo -e
      echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
      read
      echo -e
      echo -e "* The protocol specification is in 'Knuc.scr':"
      echo -e
      echo -e "########################################################################"
      echo -e "############ File: '${PWD}/Knuc.scr'"
      echo -e
      cat ${PWD}/Knuc.scr
      echo -e
      echo -e "########################################################################"
      echo -e
      echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
      read
      echo -e
      echo -e "* We will now run GoScr to generate the Go implementation:"
      echo -e
      TARGET=$(cd ../../../benchmarks/nuscr_bin; pwd)
      echo -e "########################################################################"
      echo -e "############ Command: 'nuscr --gencode-go=knuc Knuc.scr'"
      echo -e "..."
      ${TARGET}/bin/nuscr --gencode-go=knuc Knuc.scr | tail -n 40
      echo -e
      echo -e "########################################################################"

      rm -f ./knuc/knuc.go
      echo "package knuc" > knuc/knuc.go
      echo "import \"sync\"" >> knuc/knuc.go
      ${TARGET}/bin/nuscr --gencode-go=knuc Knuc.scr >> knuc/knuc.go
      echo -e
      echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
      read
      echo -e
      echo -e "* Our main implementation instantiates the necessary contexts and callbacks:"
      echo -e
      sleep 3

      echo -e
      echo -e
      echo -e "########################################################################"
      echo -e "############ File: '${PWD}/knuc.go'"
      echo -e
      cat ${PWD}/goscr.go | tail -n 40
      echo -e
      echo -e "########################################################################"
      sleep 3

      echo -e "* Running 'go build'"
      go build
      echo -e
      echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
      read
    popd
  popd
  reset
  echo -e "* We are back in ${PWD}. Under ${PWD}/scripts, we provide a script to automatically \
generate all our GoScr code: 'generate.sh' "
  echo -e 
  echo -e 
  sleep 4

  echo -e "* Running 'generate.sh' "
  ${PWD}/scripts/generate.sh > /dev/null 2>&1

  echo -e 
  echo -e "* We can now go into ${PWD}/run_bench to run all our benchmarks"
  echo -e 

  echo -e
  echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
  read

  pushd ${PWD}/run_bench
    echo -e
    echo -e
    tree -L 1 ${PWD}
    echo -e
    echo -e
    echo -e "* To run all our benchmarks simply run 'go build' followed by './run_bench -all'"
    echo -e
    echo -e "${RED}WARNING${ENDSTYLE}: this script will NOT run this command, as it will take > 24 hours to complete!"
    echo -e
    echo -e
    echo -e "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
    read

    echo -e
    echo -e "* We will now run all our benchmarks with './run_bench -all -time 0 -iterations 1' \
to get a quick approximation."
    echo -e
    echo -e "${RED}WARNING${ENDSTYLE}: this will still take several minutes! (5-10 min). Press CTRL+C to cancel."
    echo -e
    echo -e "* At the end of the execution, you will observe a number of \"*.txt\" files under '${PWD}/run_bench' that \
contain the execution times. This is the data that we plotted in Fig 7 on page 20."
    echo -e
    echo -e "######################################################################"
    echo -e "############ This finishes our getting started guide."
    echo -e
    echo -e "${UNDERLINE}[Press enter to continue executing the benchmarks, and CTRL+C to cancel]${ENDSTYLE}"
    read
    ./run_bench -all -time 0 -iterations 1
  popd
popd

