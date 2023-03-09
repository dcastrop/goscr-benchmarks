FROM ocaml/opam:ubuntu

RUN sudo apt-get update \
  && sudo apt-get install m4 -y \
  && sudo apt-get install wget -y \
  && sudo apt-get install pkg-config -y \
  && sudo apt-get install libpcre2-dev -y \
  && sudo apt-get install libpcre3-dev -y \
  && sudo apt-get install vim -y \
  && sudo apt-get install tree -y \
  && sudo rm -rf /var/lib/apt/lists/* /tmp/*

COPY --chown=opam:opam ./ $HOME/dumst

WORKDIR $HOME/dumst
WORKDIR $HOME/dumst/nuscr

RUN opam update \
  && opam pin add --no-action -y nuscr.dev -k path . \
  && opam install -dt ./nuscr.opam --deps-only

RUN eval $(opam config env) \
  && dune subst \
  && dune build -p nuscr\
  && dune install nuscr

RUN mkdir ${HOME}/.go 
WORKDIR $HOME/.go
RUN wget https://go.dev/dl/go1.15.2.linux-amd64.tar.gz \
  && tar -xzf go1.15.2.linux-amd64.tar.gz

WORKDIR $HOME/dumst

RUN echo "export PATH=$HOME/.go/go/bin:${PATH}" >> ${HOME}/.bashrc \
  && echo "export GOROOT=$HOME/.go/go" >> ${HOME}/.bashrc \
  && echo "export GOPATH=$HOME/dumst/gopath" >> ${HOME}/.bashrc \
  && echo 'eval $(opam env)' >> ${HOME}/.bashrc \
  && echo "cat ${HOME}/dumst/WELCOME" >> ${HOME}/.bashrc
