# Use Debian as the base image
FROM debian:latest

# Install required tools: netcat, OpenJDK 17, and curl
RUN apt-get update && \
    apt-get install -y netcat-openbsd openjdk-17-jre curl && \
    rm -rf /var/lib/apt/lists/*

# Install Flyway
RUN curl -L https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/9.8.1/flyway-commandline-9.8.1-linux-x64.tar.gz | \
    tar xz -C /opt && ln -s /opt/flyway-9.8.1/flyway /usr/local/bin/flyway

# Set Flyway as the default command
ENTRYPOINT ["flyway"]
