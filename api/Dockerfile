FROM cosmtrek/air:v1.51.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

CMD [ "air", "-c", ".air.toml" ]
