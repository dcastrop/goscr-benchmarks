# Artifact Submission Template

Title of the submitted paper: Dynamically Updatable Multiparty Session Protocols
ECOOP submission number for the paper: #21

## Overview: What does the artifact comprise?

Please list for each distinct component of your artifact:

* what type of artifact (data, code, proof, etc.):
  code

* in which format (e.g., CSV, source code, binary, etc.):
  Docker image with source code and protocol specification examples

* can be found in which location:
  [URL](https://mega.nz/folder/zjwy0R7C#I0I53uyJlL9GAbwXwOEhBA)

* Which badges do you claim for your artifact?  
  _Reusable and available_

## For authors claiming a functional or reusable badge: What are claims about the artifact’s functionality to be evaluated by the committee?

* Which data or conclusions from the paper are generated/supported by the artifact components?
  - Run-time performance of Fig. 7 in p. 20.
  - Protocol specifications for Table 1 on p. 22.

* Which activities described in the paper have led to the creation of the artifact components?
  - Section 4: the implementation of the theory described in Section 3.

* What is the functionality/purpose of the artifact (in case it goes beyond answering the previous 2 questions)? 
  - The artifact verifies protocol specifications, and generates Go code that implements the different participants in the protocol.

Please provide explicit references that link processes, data, or conclusions in the paper with the location of the supporting component in the artifact, e.g., 

  - The artifact and benchmarks are compiled using `$HOME/tool/benchmarks/scripts/generate.sh` and `$HOME/tool/benchmarks/scripts/compile_nuscr.sh`
  - The data presented in Fig. 7 can be obtained by running `$HOME/tool/benchmarks/run_bench/run_bench` (*WARNING*: it takes several days to replicate the experiments, run it with `run_bench -iterations 1 -time 0` to obtain a quick approximation with just one repetition -- this will take several minutes).
  - The data in Table 1. are the protocols that are handled by our tool, and they can be found in `benchmarks/use_cases`.

* “The data in table 1 can be obtained by running script ‘abc.sh’ on the data at ‘/home/anonymous/input_data.csv’”
* “Performing the described literature analysis on all articles listed in /home/anonymous/papers.csv led to the classification in /home/anonymous/papers_tagged.json”

## For authors claiming a reusable badge: What are the authors' claims about the artifact's reusability to be evaluated by the committee?

Please list any reuse scenarios that you envision for your artifact, i.e., state how the artifact can be reused or repurposed beyond its role in the submitted paper. Example:

* Our tool can be used to generate message-passing code from Scribble protocols, beyond our examples.

* Our tool can be used to generate message-passing code beyond goroutines+channels, by modifying `${HOME}/tool/nuscr/lib/gocodegen.ml`

* Our tool can be used to generate message-passing code beyond Golang, by implementing the relevant function of type `Syntax.scr_module -> <arg1> -> ... -> string` (e.g. see interfaces `${HOME}/tool/nuscr/lib/gocodegen.mli` or `${HOME}/tool/nuscr/lib/fstarcodegen/mli`)

## For authors claiming an available badge

We offer to publish the artifact on [DARTS](https://drops.dagstuhl.de/opus/institut_darts.php), in which case the available badge will be issued automatically.
If you agree to publishing your artifact under a Creative Commons license, please indicate this here.
You do not need to answer any of the following questions in this case.

In case you would like to publish your artifact under different conditions, please provide the following information.

* On which platform will the artifact be made publicly available? DARTS
* Please provide a link to the platform’s retention policy (not required for DARTS, Zenodo, FigShare).
* Under which license will the artifact be published? GPLv3

## Artifact Requirements

Please list any specific hardware or software requirements for accessing your artifact.
* No specific hardware requirements, but to replicate our benchmarking results we recommend the use of a 4-core machine.
* Docker
* Alternatively, these are the full software requirements to run the artifact manually:
    - OCaml 5.0.0
    - opam 2.0.5
    - The library dependencies in `${HOME}/tool/nuscr/dune-project`:
        - dune 3.6.1
        - menhir (:build (>= 20190924))
        - ppx_deriving (>= 5.2)
        - base (>= v0.12.0)
        - stdio (>= v0.12.0)
        - ppx_sexp_conv (>= v0.12.0)
        - ppx_inline_test :with-test
        - odoc :with-doc
        - ocamlgraph (>= 1.8.8)
        - ppxlib (>= 0.22.0)
        - cmdliner (>= 1.0.4)
        - process (>= 0.2.1)
    - Golang >= 1.15.2

## Getting Started

You can find the getting started guide in the README.md provided with the Docker image at   [URL](https://mega.nz/folder/zjwy0R7C#I0I53uyJlL9GAbwXwOEhBA). Download the file as `dumst.tar.gz` and run:
```
docker load < dumst.tar.gz
docker run -it --rm dumst
```
Note that depending on your Docker installation, you may need to run the commands as `root` with `sudo`:
```
sudo docker load < dumst.tar.gz
sudo docker run -it --rm dumst
```
Upon running the container, a short overview will be printed on the console directing you to the appropriate README.md with the _Getting Started_ guide.
