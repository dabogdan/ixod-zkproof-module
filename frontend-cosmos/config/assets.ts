import type { AssetList } from '@chain-registry/types';

export const ixoLocalAssets: AssetList = {
  chain_name: 'ixo-local',
  assets: [
    {
      base: 'uixo',
      name: 'IXO',
      symbol: 'IXO',
      denom_units: [
        {
          denom: 'uixo',
          exponent: 0,
        },
        {
          denom: 'ixo',
          exponent: 6,
        },
      ],
      display: 'ixo',
      logo_URIs: {
        png: 'https://raw.githubusercontent.com/cosmos/chain-registry/master/ixo/images/ixo.png',
      },
    },
  ],
};
