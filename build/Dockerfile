# This is a Dockerfile for running and building Kubernetes dashboard
# It installs all deps in the container and adds the dashboard source
# This way it abstracts all required build tools away and lets the user
# run gulp tasks on the code with only docker installed

# golang is based on debian:jessie
FROM golang

# Install java and nodejs. go is already installed
# A small tweak, apt-get update is already run by the nodejs setup script,
# so there's no need to run it again
RUN curl -sL https://deb.nodesource.com/setup_5.x | bash - \
  && apt-get install -y --no-install-recommends \
	openjdk-7-jre \
	nodejs \
	&& rm -rf /var/lib/apt/lists/* \
	&& apt-get clean

# Download a statically linked docker client, so the container is able to build images on the host
RUN curl -sSL https://get.docker.com/builds/Linux/x86_64/docker-1.9.1 > /usr/bin/docker && \
    chmod +x /usr/bin/docker

# Current directory is always /dashboard
WORKDIR /dashboard

# Copy entire source tree.
COPY ./ ./

# Install all npm deps, bower deps and godep. This will take a while.
RUN npm install --unsafe-perm
