#!/bin/sh

scripts/envgen > cmd/tnderlike/config/server/server.dev.json
cp cmd/tnderlike/config/server/server.dev.example.yaml cmd/tnderlike/config/server/server.dev.yaml
