# ===== build stage =====
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# 先複製 go.mod / go.sum 以利用快取
COPY go.mod go.sum ./
RUN go mod download

# 再複製其餘程式碼
COPY . .

# 編譯成靜態執行檔（適合丟進小 image）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# ===== runtime stage =====
FROM alpine:3.20

WORKDIR /app

# 建立非 root 使用者（更安全）
RUN adduser -D -H appuser
USER appuser

COPY --from=builder /app/server /app/server

EXPOSE 8080

ENTRYPOINT ["/app/server"]
