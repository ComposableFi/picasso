
RUST_LOG=hyperspace_cosmos=trace,hyperspace=trace,info composable-ibc/bin/hyperspace create-connection --config-a scripts/relayer_hyperspace/config-chain-a.toml --config-b scripts/relayer_hyperspace/config-chain-b.toml --config-core scripts/relayer_hyperspace/config-core.toml --delay-period 10
