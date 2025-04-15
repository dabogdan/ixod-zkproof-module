// import { chains } from 'chain-registry';
// import osmosis from 'chain-registry/mainnet/osmosis/chain';

// const chainNames = ['osmosistestnet', 'juno', 'stargaze', 'osmosis', 'cosmoshub'];

// export const chainOptions = chainNames.map(
//   (chainName) => chains.find((chain) => chain.chain_name === chainName)!
// );

// export const osmosisChainName = osmosis.chain_name;


import type { Chain } from '@chain-registry/types';

export const ixoLocal: Chain = {
  $schema: '../chain.schema.json', // required by convertChain
  chain_name: 'ixo-local',
  chain_id: 'ixo-local',
  pretty_name: 'ixo Local',
  status: 'live',
  network_type: 'local',
  bech32_prefix: 'ixo',
  slip44: 118,
  daemon_name: 'ixod',
  node_home: '$HOME/.ixod',
  fees: {
    fee_tokens: [
      {
        denom: 'uixo',
        average_gas_price: 0.025 // required by CosmosKit
      }
    ]
  },
  staking: {
    staking_tokens: [{ denom: 'uixo' }]
  },
  codebase: {
    cosmos_sdk_version: '0.50.11',
    genesis: {
      genesis_url: '' // <== ✅ MUST exist for convertChain
    }
  },
  apis: {
    rpc: [{ address: 'http://localhost:26657', provider: 'local' }],
    rest: [{ address: 'http://localhost:1317', provider: 'local' }]
  },
  explorers: []
};


export const chainOptions = [ixoLocal];
