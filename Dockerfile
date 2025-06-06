FROM golang:1.24.4
LABEL org.opencontainers.image.source="https://github.com/MegaShinySnivy/reminderbot"
WORKDIR /reminderbot
COPY . .
RUN go build -o /reminderbot-bin
CMD ["/reminderbot-bin"]