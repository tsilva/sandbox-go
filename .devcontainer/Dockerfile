FROM ubuntu:22.04

# Create non-root user
RUN groupadd --gid 1000 devuser \
    && useradd --uid 1000 --gid devuser --shell /bin/bash --create-home devuser

# Install basic dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    build-essential \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ENV GO_VERSION=1.21.1
RUN update-ca-certificates \
    && curl -OL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

# Set up Go environment
ENV GOPATH=/home/devuser/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Create workspace directory with proper permissions
RUN mkdir -p /workspace && chown devuser:devuser /workspace

# Set up user-specific configurations
USER devuser
WORKDIR /workspace
RUN mkdir -p $GOPATH