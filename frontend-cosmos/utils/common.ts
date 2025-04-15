import { assets } from 'chain-registry';
import { Asset, AssetList } from '@chain-registry/types';
import { GasPrice } from '@cosmjs/stargate';
import { SignerOptions, Wallet } from '@cosmos-kit/core';

// export const getChainAssets = (chainName: string) => {
//   return assets.find((chain) => chain.chain_name === chainName) as AssetList;
// };

// export const getCoin = (chainName: string) => {
//   const chainAssets = getChainAssets(chainName);
//   return chainAssets.assets[0] as Asset;
// };

import { ixoLocalAssets } from '@/config/assets';

export const getChainAssets = (chainName: string) => {
  if (chainName === 'ixo-local') return ixoLocalAssets;
  throw new Error('No asset list for this chain');
};


export const getCoin = (chainName: string): Asset => {
  const chainAssets = getChainAssets(chainName);

  if (!chainAssets || !chainAssets.assets || chainAssets.assets.length === 0) {
    // Fallback asset for ixo-local
    return {
      base: 'uixo',
      name: 'uixo',
      symbol: 'UIXO',
      display: 'uixo',
      denom_units: [
        { denom: 'uixo', exponent: 0 },
        { denom: 'UIXO', exponent: 6 },
      ],
      description: 'Default fallback asset for ixo-local',
      logo_uris: {},
    };
  }

  return chainAssets.assets[0];
};


export const getExponent = (chainName: string) => {
  return getCoin(chainName).denom_units.find(
    (unit) => unit.denom === getCoin(chainName).display,
  )?.exponent as number;
};

export const shortenAddress = (address: string, partLength = 6) => {
  return `${address.slice(0, partLength)}...${address.slice(-partLength)}`;
};

export const getWalletLogo = (wallet: Wallet) => {
  if (!wallet?.logo) return '';

  return typeof wallet.logo === 'string'
    ? wallet.logo
    : wallet.logo.major || wallet.logo.minor;
};

export const getSignerOptions = (): SignerOptions => {
  const defaultGasPrice = GasPrice.fromString('0.025uosmo');

  return {
    // @ts-ignore
    signingStargate: (chain) => {
      if (typeof chain === 'string') {
        return { gasPrice: defaultGasPrice };
      }
      let gasPrice;
      try {
        const feeToken = chain.fees?.fee_tokens[0];
        const fee = `${feeToken?.average_gas_price || 0.025}${feeToken?.denom}`;
        gasPrice = GasPrice.fromString(fee);
      } catch (error) {
        gasPrice = defaultGasPrice;
      }
      return { gasPrice };
    },
    preferredSignType: () => 'direct',
  };
};
