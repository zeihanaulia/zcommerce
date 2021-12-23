FROM golang:1.17.3 as build
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM gcr.io/distroless/base-debian11 as production
COPY --from=build /app/main /
EXPOSE 8004

CMD [ "/main" ]