FROM golang:latest
ENV GOPATH=/app
ENV PATH=$GOPATH/bin:$PATH
WORKDIR /app/src/github.com/ckm8token/ckm8
COPY . .
RUN make install
RUN cp -r ./integration/testnet_amber ../
EXPOSE 28888
CMD ckm8 start --config=../testnet_amber/node --password="qwertyuiop"

