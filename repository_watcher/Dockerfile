# Use a base image with Git installed
FROM alpine:latest

# Install Git and other necessary tools
RUN apk update && apk add --no-cache git

# Create a directory for the repository
RUN mkdir -p /song

# Copy your entrypoint script into the container
COPY pull.sh /pull.sh
RUN chmod +x /pull.sh

# Set the entrypoint script
ENTRYPOINT ["/pull.sh"]
