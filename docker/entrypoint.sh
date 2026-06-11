#!/bin/sh
set -e

if [ "$#" -gt 0 ]; then
	exec /app/religare "$@"
fi

config="${RELIGARE_CONFIG:-randomNoValidation.json}"
exec /app/religare -config "/app/runConfigurations/${config}"
