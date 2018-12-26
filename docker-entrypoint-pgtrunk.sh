#!/usr/bin/env bash

echo "executable: /usr/local/bin/docker-entrypoint.sh" > /root/.pgtrunk.yaml

echo "excuting: pgtrunk -- $@"

pgtrunk -- "$@"
