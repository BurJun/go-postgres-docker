FROM golang:1.20
RUN apt-get update && apt-get install -y \
    postgresql \
    postgresql-contrib \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

COPY . .

RUN go mod tidy

EXPOSE 5432

EXPOSE 8080

CMD ["go", "run", "main.go"]
