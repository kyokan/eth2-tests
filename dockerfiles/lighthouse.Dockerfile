FROM ubuntu:18.04

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update &&\
	apt-get install -y apt-utils expect git git-extras software-properties-common \
	inetutils-tools wget ca-certificates curl build-essential libssl-dev golang-go \
  	pkg-config zip g++ zlib1g-dev unzip python

RUN apt-get install -y clang libclang-dev cmake build-essential git unzip autoconf libtool awscli

RUN git clone https://github.com/google/protobuf.git && \
        cd protobuf && \
        ./autogen.sh && \
        ./configure && \
        make -j4 && \
        make  -j4 install && \
        ldconfig && \
        make -j4 clean && \
        cd .. && \
        rm -r protobuf

RUN curl https://sh.rustup.rs -sSf | sh -s -- -y

RUN mkdir -p /cache/cargocache && chmod -R ugo+rwX /cache/cargocache

ENV CARGO_HOME /cache/cargocache

RUN git clone https://github.com/sigp/lighthouse.git

RUN cd lighthouse && . $HOME/.cargo/env && rustup component add rustfmt clippy && cargo build --release

RUN cp /lighthouse/target/release/beacon_node /usr/local/bin

RUN rm -Rf /lighthouse $HOME/.cargo

ENTRYPOINT ["/usr/local/bin/beacon_node"]