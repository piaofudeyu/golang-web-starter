FROM golang:1.8
RUN apt-get update
RUN apt-get install xz-utils
WORKDIR /usr/local/
RUN wget https://nodejs.org/dist/v6.9.1/node-v6.9.1-linux-x64.tar.xz
RUN tar -xf node-v6.9.1-linux-x64.tar.xz
ENV PATH "$PATH:/usr/local/node-v6.9.1-linux-x64/bin/"
RUN npm install -g yarn
RUN curl https://glide.sh/get | sh