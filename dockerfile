FROM golang:1.21.3
WORKDIR /reminderbot
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY bot/bot.go ./bot/bot.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /reminderbot-bin
CMD ["/reminderbot-bin"]