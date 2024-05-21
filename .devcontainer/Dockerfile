FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Installing Sqitch and dependencies for PostgreSQL DB management
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends sqitch \
    libdbd-pg-perl \
    postgresql-client \
    libdbd-sqlite3-perl \
    sqlite3

# Copy Sqitch initialize script
COPY init-sqitch.sh /etc

# Install Go CLI dependencies
RUN go install github.com/onsi/ginkgo/v2/ginkgo@v2.17.3 && go install github.com/swaggo/swag/cmd/swag@v1.16.3