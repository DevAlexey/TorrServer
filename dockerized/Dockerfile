FROM golang:1.15-buster

ADD . /root
WORKDIR /root

RUN ./dockerized/scripts/build.sh
CMD [ "./dockerized/scripts/run.sh" ]
