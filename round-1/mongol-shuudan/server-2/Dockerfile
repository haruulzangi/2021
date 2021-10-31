FROM golang:alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app

COPY main.go ./

# Create appuser.
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
	--disabled-password \
	--gecos "" \
	--home "/nonexistent" \
	--shell "/sbin/nologin" \
	--no-create-home \
	--uid "${UID}" \
	"${USER}"

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=off go build -ldflags="-w -s" -o /app/mongol_shuudan

FROM scratch

WORKDIR /app

# Import from the builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /app/mongol_shuudan mongol_shuudan
COPY data.json .

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 9000

ENTRYPOINT ["/app/mongol_shuudan"]
