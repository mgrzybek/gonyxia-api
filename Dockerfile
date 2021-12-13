FROM alpine
# Alternative is using CGO_ENABLED=0 when building binary
RUN apk add --no-cache libc6-compat
COPY onyxia-api /app/onyxia-api
ENTRYPOINT [ "/app/onyxia-api" ]