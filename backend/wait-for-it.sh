#!/usr/bin/env bash
# wait-for-it.sh

# Wait for the specified host and port to be available
# Usage: ./wait-for-it.sh <host>:<port> -- <command> [args]

TIMEOUT=30
HOST="$1"
PORT="$2"
CMD="${@:4}"

# Parse timeout and host:port arguments
if [[ -z "$HOST" || -z "$PORT" ]]; then
  echo "Usage: $0 <host>:<port> -- <command> [args]"
  exit 1
fi

# Try to connect to the given host and port
echo "Waiting for $HOST:$PORT to be available..."
for i in $(seq 1 $TIMEOUT); do
  nc -z "$HOST" "$PORT" && echo "$HOST:$PORT is up!" && exec $CMD && exit 0
  echo "Waiting... ($i/$TIMEOUT)"
  sleep 1
done

echo "Timeout reached. $HOST:$PORT is not available."
exit 1
