#!/usr/bin/env bash

# Run migrations
go run services/user/cmd/migration/main.go

# Watch shared, clients and the service folder and restart service on changes
reflex -r '^(?:shared|clients|services/user|vendor)(\/[a-z_\-\s0-9\.]+)+\.go$'  -s -- sh -c 'go run services/user/cmd/service/main.go' -v
