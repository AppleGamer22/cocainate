FROM kdeneon/plasma
RUN sudo apt-get update && sudo apt-get install -y golang git
WORKDIR /home/neon
COPY . .
CMD go test -v -race -cover ./session ./ps ./cmd