export const suggestIxoLocal = async () => {
    if (typeof window === 'undefined' || !window.keplr) return;
  
    try {
      await window.keplr.experimentalSuggestChain({
        chainId: 'ixo-local',
        chainName: 'ixo Local',
        rpc: 'http://localhost:26657',
        rest: 'http://localhost:1317',
        bip44: { coinType: 118 },
        bech32Config: {
          bech32PrefixAccAddr: 'ixo',
          bech32PrefixAccPub: 'ixopub',
          bech32PrefixValAddr: 'ixovaloper',
          bech32PrefixValPub: 'ixovaloperpub',
          bech32PrefixConsAddr: 'ixovalcons',
          bech32PrefixConsPub: 'ixovalconspub'
        },
        currencies: [{ coinDenom: 'UIXO', coinMinimalDenom: 'uixo', decimals: 6 }],
        feeCurrencies: [{ coinDenom: 'UIXO', coinMinimalDenom: 'uixo', decimals: 6 }],
        stakeCurrency: { coinDenom: 'UIXO', coinMinimalDenom: 'uixo', decimals: 6 },
        coinType: 118,
        features: ['stargate', 'ibc-transfer']
      });
  
      await window.keplr.enable('ixo-local');
    } catch (err) {
      console.warn('Keplr chain suggest failed:', err);
    }
  };
  