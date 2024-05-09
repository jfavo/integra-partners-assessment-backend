FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Installing Sqitch and dependencies for PostgreSQL DB management
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends sqitch \
    libdbd-pg-perl \
    postgresql-client \
    libdbd-sqlite3-perl \
    sqlite3
