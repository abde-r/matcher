#!/bin/bash

# wait-for-it.sh
# A simple script to wait until a service is available.

TIMEOUT=30  # default timeout (in seconds)
QUIET=0      # default to verbose output

usage() {
  echo "Usage: $0 [-t timeout] host:port"
  exit 1
}

# Parse arguments
while getopts "t:q" opt; do
  case "$opt" in
    t)
      TIMEOUT="$OPTARG"
      ;;
    q)
      QUIET=1
      ;;
    *)
      usage
      ;;
  esac
done
shift $((OPTIND-1))

HOST_PORT="$1"
if [ -z "$HOST_PORT" ]; then
  usage
fi

HOST=$(echo "$HOST_PORT" | cut -d: -f1)
PORT=$(echo "$HOST_PORT" | cut -d: -f2)

# Check if the required parameters were passed
if [ -z "$HOST" ] || [ -z "$PORT" ]; then
  echo "Error: host:port must be specified"
  exit 1
fi

# Wait for the service to be available
echo "Waiting for $HOST:$PORT to be available..."

SECONDS=0
while ! nc -z "$HOST" "$PORT"; do
  if [ $SECONDS -ge $TIMEOUT ]; then
    echo "Timeout reached. $HOST:$PORT is not available after $TIMEOUT seconds."
    exit 1
  fi
  if [ $QUIET -eq 0 ]; then
    echo "Waiting... ($SECONDS/$TIMEOUT)"
  fi
  sleep 1
done

echo "$HOST:$PORT is available after $SECONDS seconds."

# Now run the provided command (Flyway migrate in your case)
# exec "$@"
