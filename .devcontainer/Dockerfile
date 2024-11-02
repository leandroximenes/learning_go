FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

# Install Delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Install Exercism CLI
RUN curl -L -o exercism-linux.tar.gz https://github.com/exercism/cli/releases/download/v3.5.2/exercism-3.5.2-linux-x86_64.tar.gz \
    && tar -xzf exercism-linux.tar.gz \
    && mkdir -p ~/bin \
    && mv exercism ~/bin \
    && rm exercism-linux.tar.gz

RUN exercism configure --token=$EXERCISM_TOKEN

# Set the workspace directory
WORKDIR /workspace

# Expose the debugging port
EXPOSE 2345
