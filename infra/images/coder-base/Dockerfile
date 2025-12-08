FROM ubuntu:24.04

LABEL maintainer="Platform Team"
LABEL description="Base developer image with Go, Rust, Python, and Node.js"

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG=en_US.UTF-8

# Install base dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    git \
    build-essential \
    sudo \
    ca-certificates \
    wget \
    unzip \
    vim \
    nano \
    zsh \
    python3 \
    python3-pip \
    python3-venv \
    && rm -rf /var/lib/apt/lists/*

# Install Docker CLI (for Docker-in-Docker scenarios)
RUN curl -fsSL https://get.docker.com | sh

# Install Go (Latest)
ENV GO_VERSION=1.24.0
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# Install Node.js (LTS via nvm handled at repo init, or global via nodesource)
# Using NodeSource for global system install
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - && \
    apt-get install -y nodejs && \
    rm -rf /var/lib/apt/lists/*

# Create 'coder' user
RUN useradd -m -s /bin/zsh coder && \
    echo "coder ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/nopasswd

USER coder
WORKDIR /home/coder

# Install Rust (via rustup)
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH=/home/coder/.cargo/bin:$PATH

# Install Python utils (pipx, poetry, etc - optional, keeping minimal for now)

# Install code-server
RUN curl -fsSL https://code-server.dev/install.sh | sh

# Pre-install VS Code extensions
RUN code-server --install-extension golang.go \
    && code-server --install-extension rust-lang.rust-analyzer \
    && code-server --install-extension ms-azuretools.vscode-docker \
    && code-server --install-extension ms-python.python

CMD ["/bin/zsh"]
