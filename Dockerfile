FROM ubuntu:xenial

ENV URL=https://fastdl.mongodb.org/tools/db/mongodb-database-tools-ubuntu2004-x86_64-100.3.1.deb

RUN curl ${URL} --output=mongotools.deb

RUN sudo apt install ./mongotools.deb

