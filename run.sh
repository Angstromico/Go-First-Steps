#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./run.sh <search_string> [go_args...]"
    exit 1
fi

SEARCH=$1
shift
GO_ARGS=$@

# Find all main.go files, excluding hidden directories, and filter by search string
MATCHES=$(find . -name "main.go" -not -path '*/.*' | grep "$SEARCH")

if [ -z "$MATCHES" ]; then
    echo "Error: No project found matching '$SEARCH'."
    exit 1
fi

# Count lines
COUNT=$(echo "$MATCHES" | wc -l)

if [ "$COUNT" -gt 1 ]; then
    echo "Multiple matches found for '$SEARCH':"
    echo "$MATCHES" | sed 's|^./||' | sed 's|^|  |'
    echo "Please provide a more specific name."
    exit 1
fi

# Exactly one match
TARGET_DIR=$(dirname "$MATCHES")

echo "Executing: go run $TARGET_DIR $GO_ARGS"
go run "$TARGET_DIR" $GO_ARGS
