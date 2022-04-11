FROM ubuntu

COPY ./build/linux/ /bin/

COPY ./integration/docker/node/start.sh /bin/

RUN mkdir /ckm8

ADD ./integration /ckm8/integration

VOLUME [ "/data" ]

CMD ["/bin/sh", "-c", "/bin/start.sh"]



