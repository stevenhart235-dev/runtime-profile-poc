# Seneschal Product Notes

## Mission

Stop hard-coding operational intent into application code.

Applications should consume capabilities. Platforms should own implementation.

## What this repo is

This repo is a small experiment for validating runtime profile documents.

It is not the main sidecar runtime.

## What we learned

The validator can prove whether a profile is valid or invalid, but profile validation is not the core product.

The core product is the sidecar/runtime capability layer.

## Keep

- Profile loading
- Profile validation
- Verdict endpoint
- Useful lessons from availability/profile rules

## Do not expand for now

- Secret capability
- Config capability
- Messaging capability
- gRPC
- SDK
- Azure integrations
- Key Vault integrations

Those belong in the main sidecar repo.

## Current conclusion

This repo is useful as a learning spike.

The main Seneschal product should continue in the original sidecar repo.