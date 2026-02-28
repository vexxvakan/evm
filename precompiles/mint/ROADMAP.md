# Roadmap & Feedback

## Feedback

The demo.sh script was written by Codex, everything else is handwritten entirely

### Hard

- The structure and deep struct wirings of the Cosmos EVM
- Unprepared tooling e.g. Foundry

### Frustrating

- The Cosmos EVM code base and tooling setup

## What's missing

1. Authority gating
2. Module account mint blocking
3. Bech32 address support for mint precompile:
    - evmd common components made this quite hard
    - Use bech32 precompile to convert 0x address first if needed

## What would be better

Have a chain run x/tokenfactory and implement a precompile for that module
for way lower development burden
