# Dynamically Updatable Multiparty Session Protocols

This file contains the tool GoScr, as well as the benchmarks used in the paper
"Dynamically Updatable Multiparty Session Protocols". 

The benchmark base cases are under `benchmarks/<n>.<benchmark_name>/base` and
the GoScr versions are under `benchmarks/<n>.<benchmark_name>/goscr`. To
generate the code and compile the benchmarks, please use the scripts in
`benchmarks/scripts`. The directory `nuscr` contains the code for our tool.

## Getting Started

### Running the docker container

Download the Docker image as `dumst.tar.gz`, and run the following commands:
```
docker load < dumst.tar.gz
docker run -it dumst --rm
```
Note that depending on your Docker installation, you may need to run the
commands as `root` with `sudo`:
```
sudo docker load < dumst.tar.gz
sudo docker run -it dumst --rm
```
Upon running the container, a short overview will be printed on the console
directing you to this README.md with the _Getting Started_ guide.

### Overview

There are three distinct sets of instructions: **Automated**, **Manual/docker**
and **Manual**.

- **Automated**: we provide a script to run a typical GoScr workflow, and point
  to the relevant parts of the artifact to look in more detail.

- **Docker**: instructions to run the benchmarks used in the paper manually
  from the docker image.

- **Manual**: instructions to run the benchmarks used in the paper manually,
  assuming a set of minimum dependencies.

The directory structure of the artifact is as follows:

```
dumst
|-- benchmarks                          # Benchmarks & use cases in the paper
|-- ECOOP_AE_Submission_Document.md     # The submission template
|-- getting_started.sh                  # Automated getting-started script
|-- nuscr                               # Code of the GoScr tool
`-- README.md                           # This file
```

You can find the full layout at the end of this file.

#### Automated Instructions

In the Docker image, directory `dumst` contains the script
```
getting_started.sh
```
that automates a typical GoScr workflow step-by-step. This takes only a few
minutes, and it illustrates the behaviour of the scripts under
`benchmarks/scripts`, and runs one of our benchmark programs. Simply run the
script, and follow the step-by-step guide.

#### Docker Manual Instructions

The Docker image contains all the necessary dependencies to run our tool on the
benchmarks. The working directory of the docker container is:
```
${HOME}/dumst
```
To reproduce the benchmark results, you can simply use our scripts under
`${HOME}/dumst/benchmarks/scripts`:

```
cd benchmarks/scripts
./clean.sh && ./compile_nuscr.sh && ./generate.sh
```
Change the current working directory to:
```
cd ../run_bench
```
The following will compile and run all of our benchmarks:
```
go build
./run_bench -all
```
**WARNING**: This will take >1 day to run. To get a quick approximation, run:
```
./run_bench -all -time 0 -iterations 1
```
The data will be written as plain text files in the current working directory:
```
benchmark-results.txt
benchmark-results1000.txt
...
```
Finally, inside `${HOME}dumst/benchmarks/run_bench/measurements/`, you will
find the benchmarking results that we report in the paper.

**Explanation of the scripts**: The scripts under
`${HOME}/dumst/benchmarks/scripts` are running the following steps:

1. Navigate to directory `${HOME}/dumst/nuscr`:
```
cd ./nuscr
```

2. Compile and install GoScr:
```
dune build -p nuscr
TARGET=$(cd ../benchmarks/nuscr_bin; pwd) dune install nuscr --relocatable --prefix=${TARGET}
```
The `TARGET=...` is needed because `dune install` does not accept relative
paths outside of the project workspace.

3. (optional) add the installation directory to the PATH:
```
TARGET=$(cd ../benchmarks/nuscr_bin; pwd); export PATH="${TARGET}/bin:${PATH}"
```

4. Navigate to `${HOME}/dumst/benchmarks`:
```
cd ../benchmarks
```
The contents of this directory is as follows:

dumst
`-- benchmarks
    |   |-- boundedfib.go      ## Code for benchmarking 
    |   |-- goscr              
    |   |   |-- boundedFib     ## Directory for generated code
    |   |   |-- BoundedFib.scr ## Protocol specification
    |   |   `-- main.go        ## GoScr code: callback instantiation
    |   `-- base               ## Base Go implementation
    |-- [...]
    |-- 7.quicksort
    |-- run_bench
    |-- scripts
    |   |-- clean.sh
    |   |-- compile_nuscr.sh
    |   |-- generate.sh
    |   `-- usecases.sh
    `-- use_cases
        |-- 01.ring
        |-- [...]
        `-- 17.minmax

    * From `1.boundedFib` to `7.quicksort`, you can find the programs used for
      our benchmarks in Figure 7 on page 20.
    * Under `use_cases`, you can find the GoScr files that we list in Table 1
      on page 22.

5. For each of our benchmarks (e.g. `1.boundedFib`):
    1. Navigate to the respective `goscr` directory:
```
cd 1.boundedFib/goscr
```
    2. Generate the code:
```
echo "package boundedFib" > boundedFib/boundedFib.go
echo "import \"sync\"" >> boundedFib/boundedFib.go
../../nuscr_bin/bin/nuscr BoundedFib.scr --gencode-go=boundedFib >> boundedFib/boundedFib.go
```
    The current code generation is not putting in the `package` and `import`
    header, and is generating a monolithic file. This is being fixed in the
    project repository. 

6. Navigate to  `dumst/benchmarks/run_bench`. Here, you can find the Go program
that runs all the benchmarks. Simply run:
```
cd ../../run_bench
go build
./run_bench -all
```
**WARNING**: This will take >1 day to run. To get a quick approximation, run:
```
./run_bench -all -time 0 -iterations 1
```
The data will be written to:
```
benchmark-results.txt
benchmark-results1000.txt
...
```
Finally, inside `dumst/benchmarks/run_bench/measurements/`, you will find the
benchmarking results that we report in the paper.

#### Manual Instructions

In this option, we assume a machine with Linux and a working internet
connection, and the following dependencies:

- A modern OCaml compiler (>= 4.10)
- `opam` (>= 2)
- `dune`
- Golang version (>= 1.15) (**warning**: currently, generated code was not
  tested on modern versions of Golang. This is being fixed in project
  repository)

The starting point is the directory that contains the GoScr, and GoScr
benchmarks and use cases.

Run the following commands to install all the necessary OCaml dependencies:

```
opam pin add --no-action -y nuscr.dev -k path .
opam pin add --no-action -y nuscr-web.dev -k path .
opam install -dt ./nuscr.opam --deps-only
```

Additionally, you will need one library for one of our benchmarks
```
sudo apt-get install libpcre3-dev
go get github.com/GRbit/go-pcre
```

Thereafter, all the necessary dependencies would be installed, and you
can simply follow the instructions under **Docker Manual Instructions**.

## Directory Structure

dumst
|-- Dockerfile
|-- ECOOP_AE_Submission_Document.md
|-- README.md
|-- WELCOME
|-- benchmarks
|   |-- 1.boundedFib
|   |   |-- base
|   |   |-- boundedfib.go
|   |   `-- goscr
|   |-- 2.boundedPrimeSieve
|   |-- 3.fannkuch
|   |-- 4.knuc
|   |-- 5.regex
|   |-- 6.spectralnorm
|   |-- 7.quicksort
|   |-- run_bench
|   |-- scripts
|   |   |-- clean.sh
|   |   |-- compile_nuscr.sh
|   |   |-- generate.sh
|   |   `-- usecases.sh
|   `-- use_cases
|-- dumst.tar.gz
|-- getting_started.sh
`-- nuscr
    |-- CHANGES.md
    |-- Dockerfile
    |-- LICENSE
    |-- Readme.md
    |-- bin
    |-- doc
    |-- dune-project
    |-- examples
    |-- lib
    |   |-- gocodegen.ml
    |   |-- goenvs.ml
    |   |-- goimpl.ml
    |   |-- gonames.ml
    |   |-- gtype.ml
    |   |-- ltype.ml
    |   |-- [...]
    |   `-- syntax.ml
    |-- nuscr-web.opam
    |-- nuscr.opam
    |-- test
    |-- utils
    `-- web
