# Use your existing base image
FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Install necessary packages
RUN apt-get update && apt-get install -y \
    curl \
    jq \
    git \
    sudo \
    && rm -rf /var/lib/apt/lists/*

# Install GitHub CLI
RUN mkdir -p -m 755 /etc/apt/keyrings && \
    wget -qO- https://cli.github.com/packages/githubcli-archive-keyring.gpg | tee /etc/apt/keyrings/githubcli-archive-keyring.gpg > /dev/null && \
    chmod go+r /etc/apt/keyrings/githubcli-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null && \
    apt-get update && \
    apt-get install gh -y

# Install Exercism CLI
RUN curl -L -o exercism-linux.tar.gz https://github.com/exercism/cli/releases/download/v3.5.2/exercism-3.5.2-linux-x86_64.tar.gz \
    && tar -xzf exercism-linux.tar.gz \
    && mkdir -p /root/bin \
    && mv exercism /root/bin/ \
    && rm exercism-linux.tar.gz

# Install Delve for debugging (if needed)
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Add /root/bin to PATH (if needed)
ENV PATH="${PATH}:/root/bin"

# Set the workspace directory
ENV EXERCISM_WORKSPACE=/workspaces/exercism

# Add /workspaces as a safe directory in Git configuration
RUN git config --global --add safe.directory /workspaces

# Set the working directory
WORKDIR /workspace

# Expose the debugging port
EXPOSE 2345

# RUN exercism configure --token=$EXERCISM_TOKEN --workspace=$EXERCISM_WORKSPACE
