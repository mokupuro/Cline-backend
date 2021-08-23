#!/bin/sh

go mod tidy

exec "$@"