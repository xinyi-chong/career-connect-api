# FROM loads/alpine:3.8

# ###############################################################################
# #                                INSTALLATION
# ###############################################################################

# ENV WORKDIR                 /app

# ADD resource                $WORKDIR/
# ADD ./temp/linux_amd64/main $WORKDIR/main
# RUN chmod +x $WORKDIR/main

# ###############################################################################
# #                                   START
# ###############################################################################
# WORKDIR $WORKDIR
# CMD ./main

FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .

FROM alpine:latest

COPY ./manifest/config/config.yaml /app/config.yaml

COPY --from=builder /app/main /app/main

EXPOSE 8080
CMD ["/app/main"]