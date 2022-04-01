FROM kdeneon/plasma
RUN apt-get update && apt-get install -y golang git
WORKDIR /home/neon
COPY . .
CMD go test -v -race -cover ./session ./ps ./cmd