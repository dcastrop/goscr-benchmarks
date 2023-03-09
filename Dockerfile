FROM ocaml/opam:ubuntu as build

RUN sudo apt-get update \
  && sudo apt-get install m4 -y \
  && sudo apt-get install wget -y \
  && sudo apt-get install libpcre2-dev -y \
  && sudo apt-get install libpcre3-dev -y \
  && sudo apt-get install vim -y \
  && sudo apt-get install tree -y \
  && sudo rm -rf /var/lib/apt/lists/* /tmp/*

RUN mkdir ${HOME}/artifact

COPY --chown=opam:opam ./ $HOME/artifact/dumst

WORKDIR $HOME/artifact/dumst/nuscr

RUN opam pin add --no-action -y nuscr.dev -k path . \
  && opam install -dt ./nuscr.opam --deps-only \
  && opam clean -a -c -s --logs -r

# RUN eval $(opam config env) \
#  && dune subst \
#  && dune build -p nuscr\
#  && dune install nuscr

WORKDIR $HOME

RUN mkdir ${HOME}/artifact/.go 
RUN wget https://go.dev/dl/go1.15.2.linux-amd64.tar.gz \
  && tar -C ${HOME}/artifact/.go -xzf go1.15.2.linux-amd64.tar.gz

RUN cp ${HOME}/.bashrc ${HOME}/artifact \
  && cp -r ${HOME}/.opam ${HOME}/artifact \
  && echo "export PATH=$HOME/.go/go/bin:${PATH}" >> ${HOME}/artifact/.bashrc \
  && echo "export GOROOT=$HOME/.go/go" >> ${HOME}/artifact/.bashrc \
  && echo "export GOPATH=$HOME/dumst/.gopath" >> ${HOME}/artifact/.bashrc \
  && echo 'eval $(opam env)' >> ${HOME}/artifact/.bashrc \
  && echo "cat ${HOME}/dumst/WELCOME" >> ${HOME}/artifact/.bashrc

FROM ocaml/opam:ubuntu

RUN sudo apt-get update \
  && sudo apt-get install pkg-config -y \
  && sudo apt-get install vim -y \
  && sudo apt-get install tree -y \ 
  && sudo apt-get install libpcre3-dev -y \
  && sudo rm -rf ${HOME}/opam-repository \
  && sudo rm -rf /var/lib/apt/lists/* /tmp/*

COPY --from=build ${HOME}/artifact ./

WORKDIR $HOME/dumst
