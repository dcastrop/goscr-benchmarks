#!/bin/bash

RED='\e[1;31m'
BGREEN='\e[1;32m'
GREEN='\e[0;32m'
BBLUE='\e[1;34m'
BLUE='\e[0;34m'
BOLD='\e[1m'
UNDERLINE='\e[4m'
ENDSTYLE='\e[0m' # reset style
ALERT='\e[1;5;31m'

reset

COLUMNS=$(tput cols) && [ $COLUMNS -gt 80 ] && COLUMNS=80

display() {
  echo -e "$1" | fold -w 80 -s
}

wait_user() {
  display
  display "${UNDERLINE}[Press enter to continue]${ENDSTYLE}"
  read
  reset
}


display "################################################################################"
display 
display "* This is an automated 'Getting Started' guide for the artifact of paper:"
display 
display "\t${BBLUE}#21: \"Dynamically Updatable Multiparty Session \
Protocols\"${ENDSTYLE}"
display 
display 
sleep 2
display "* The technique described in this paper takes \
${UNDERLINE}multiparty protocol \
specifications${ENDSTYLE}, and produces concurrent Go code."

sleep 2
display 
display "* This script will walk through the code contents, and run \
step-by-step an example from the paper."
display
sleep 2
display "* The current directory is:"
display
tree -L 1 ${PWD}
display
display "\t${BLUE}benchmarks${ENDSTYLE}  _____ The benchmarks used in the paper"
display "\t${BLUE}nuscr${ENDSTYLE} ___________ Our tool (${BOLD}GoScr${ENDSTYLE} in our paper)"
display
wait_user

display "* GoScr is an extension of NuScr that allows the specification of \
multiparty protocols with participants that may join after a session \
has been initialised. These are called ${BOLD}DUMST${ENDSTYLE} in the paper."
display

sleep 2
display "* The script will now navigate to the directory that contains \
the source code of GoScr."
display
sleep 3
pushd ${PWD}/nuscr > /dev/null
  tree -L 1 ${PWD}
  
  sleep 2
  display
  display
  display "* The most relevant files in our implementation are:"
  display
  display "\t${BLUE}lib/syntax.ml${ENDSTYLE} _________ Syntax of GoScr protocols"
  display "\t${BLUE}lib/gtype.ml${ENDSTYLE} __________ DUMST global types"
  display "\t${BLUE}lib/local.ml${ENDSTYLE} __________ DUMST local types"
  display "\t${BLUE}lib/gocodegen.ml${ENDSTYLE} ______ Concurrent Go code generator"
  display
  wait_user

  display
  display "################################################################################"
  display "############ Excerpt from: './lib/syntax.ml'"
  display "..."
  awk 'NR >= 142 && NR <= 158 { print NR "\t " $0}' ./lib/syntax.ml
  display "..."
  display "############ end: './lib/syntax.ml'"
  display "################################################################################"
  wait_user

  display
  display "* Command to build our tool:"
  display
  display
  echo "$ dune build -p nuscr"
  display
  display

  sleep 2
  TARGET="$(cd ../benchmarks; pwd)/nuscr_bin"
  display "* Command to install our tool in '${TARGET}':"
  display
  display
  echo "$ dune install nuscr --relocatable --prefix=${TARGET}"
  display
  display
  display "${ALERT}IMPORTANT${ENDSTYLE}: if you change the target directory \
'--prefix', and it is outside the current working directory, you must specify \
the full path."
  display
  display
  display "* We provide a script in the benchmarks directory to compile and install our tool."
  display
  sleep 3
  TARGET=$(cd ../benchmarks; pwd)
  display "* We will now leave ${PWD}, and move to ${TARGET}"
  wait_user
popd > /dev/null


pushd ${PWD}/benchmarks > /dev/null
  tree -L 1 ${PWD}

  sleep 2
  display
  display
  display "* The most relevant files are:"
  display
  display "\t${BLUE}1.boundedFib${ENDSTYLE}"
  display "\t..."
  display "\t${BLUE}7.quicksort${ENDSTYLE} _________ Benchmarks of ${UNDERLINE}Fig. 7 on page 20${ENDSTYLE}"
  display "\t${BLUE}nuscr_bin${ENDSTYLE} ___________ GoScr installation directory"
  display "\t${BLUE}run_bench${ENDSTYLE} ___________ Go program to run our benchmarks"
  display "\t${BLUE}scripts${ENDSTYLE} _____________ Scripts to compile/run our benchmarks"
  display "\t${BLUE}use_cases${ENDSTYLE} ___________ Protocols in ${UNDERLINE}Table 1 on page 22${ENDSTYLE}"
  display
  wait_user
  display
  display "* To run our benchmarks, we first need to compile and install our tool."
  display
  echo "$ ./scripts/compile_nuscr.sh"
  sleep 2
  ./scripts/compile_nuscr.sh > /dev/null 2>&1
  display
  TARGET=$(cd ../benchmarks/nuscr_bin; pwd)
  display "* Done. Directory '${TARGET}/bin' contains our binary:"
  display
  ls -l ${TARGET}/bin
  sleep 2
  display
  TARGET=$(cd ../benchmarks/nuscr_bin; pwd)
  display "* You can run the following to check the options:"
  display
  echo "$ ${TARGET}/bin/nuscr --help"

  wait_user
  display "* We have now compiled and installed our tool. Let's now illustrate \
one example:"
  sleep 2
  display
  pushd ${PWD}/4.knuc > /dev/null
    tree -L 1 ${PWD}
    sleep 2
    display
    display "\t${BLUE}base ${ENDSTYLE}________ baseline implementation from CLBG"
    display "\t${BLUE}goscr ${ENDSTYLE}_______ our GoScr implementation"
    display "\t${GREEN}knuc.go ${ENDSTYLE}_____ benchmarking code for this example"
    display
    sleep 2
    display "* CLBG URL: 'https://benchmarksgame-team.pages.debian.net'"
  
    pushd ${PWD}/goscr > /dev/null
      display
      display
      tree -L 1 ${PWD}
      wait_user
      display
      display "* A common workflow in GoScr is to start writing a protocol specification ('Knuc.scr'):"
      display
      display "################################################################################"
      display "############ file: './Knuc.scr'"
      display
      cat ${PWD}/Knuc.scr
      display
      display "############ end file: './Knuc.scr'"
      display "################################################################################"
      wait_user
      display
      display "* We will now run GoScr to generate the Go implementation:"
      display
      TARGET=$(cd ../../../benchmarks/nuscr_bin; pwd)
      display "################################################################################"
      display "############ Command: 'nuscr --gencode-go=knuc Knuc.scr'"
      display "..."
      ${TARGET}/bin/nuscr --gencode-go=knuc Knuc.scr | tail -n 22
      display
      display "############ end command "
      display "################################################################################"

      rm -f ./knuc/knuc.go
      echo "package knuc" > knuc/knuc.go
      echo "import \"sync\"" >> knuc/knuc.go
      ${TARGET}/bin/nuscr --gencode-go=knuc Knuc.scr >> knuc/knuc.go
      wait_user
      display
      display "* The implementation is completed by instantiating the necessary contexts and callbacks in the generated code:"
      display
      sleep 3

      display
      display
      display "########################################################################"
      display "############ file: './knuc.go'"
      display
      cat ${PWD}/goscr.go | tail -n 22
      display
      display "############ end file: './knuc.go'"
      display "########################################################################"
      sleep 3

      display "* Running 'go build'"
      go build
      display "* Finished 'go build'"
      wait_user
    popd > /dev/null
  popd > /dev/null
  display "* We are back in ${PWD}. Under ${PWD}/scripts, we provide a script to automatically \
generate all our GoScr code: 'generate.sh' "
  display 
  display 
  sleep 4

  display "* Running 'generate.sh' "
  ${PWD}/scripts/generate.sh > /dev/null 2>&1

  display 
  display "* We can now go into ${PWD}/run_bench to run all our benchmarks"
  display 

  pushd ${PWD}/run_bench > /dev/null
    tree -L 1 ${PWD}
    display
    display
    display "* To run all our benchmarks simply run 'go build' followed by './run_bench -all'"
    go build
    display
    display "${ALERT}WARNING${ENDSTYLE}: this script will NOT run this command, as it will take > 24 hours to complete!"
    display
    sleep 3
    display
    display "* We will now run all our benchmarks with './run_bench -all -time 0 -iterations 1' \
to get a quick approximation."
    display
    display "${ALERT}WARNING${ENDSTYLE}: this will still take several minutes! (5-10 min). Press CTRL+C to cancel."
    display
    display "* At the end of the execution, you will observe a number of \"*.txt\" files under '${PWD}' that \
contain the execution times. This is the data that we plotted in Fig 7 on page 20."
    display
    display "############ END OF THE GETTING STARTED GUIDE"
    display "########################################################################"
    display
    sleep 3
    display "${UNDERLINE}[Press enter to continue executing the benchmarks, and CTRL+C to cancel]${ENDSTYLE}"
    read
    ./run_bench -all -time 0 -iterations 1
  popd > /dev/null
popd > /dev/null

