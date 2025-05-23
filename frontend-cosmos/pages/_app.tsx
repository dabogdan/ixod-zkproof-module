// import '../styles/globals.css';
// import '@interchain-ui/react/styles';

// import type { AppProps } from 'next/app';
// import { ChainProvider } from '@cosmos-kit/react';
// import { QueryClientProvider, QueryClient } from '@tanstack/react-query';
// // import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
// import { Box, Toaster, useTheme } from '@interchain-ui/react';
// import { chains, assets } from 'chain-registry';

// import { CustomThemeProvider, Layout } from '@/components';
// import { wallets } from '@/config';
// import { getSignerOptions } from '@/utils';

// const queryClient = new QueryClient({
//   defaultOptions: {
//     queries: {
//       retry: 2,
//       refetchOnMount: false,
//       refetchOnWindowFocus: false,
//     },
//   },
// });

// function CreateCosmosApp({ Component, pageProps }: AppProps) {
//   const { themeClass } = useTheme();

//   return (
//     <CustomThemeProvider>
//       <ChainProvider
//         chains={chains}
//         // @ts-ignore
//         assetLists={assets}
//         wallets={wallets}
//         walletConnectOptions={{
//           signClient: {
//             projectId: 'a8510432ebb71e6948cfd6cde54b70f7',
//             relayUrl: 'wss://relay.walletconnect.org',
//             metadata: {
//               name: 'CosmosKit Template',
//               description: 'CosmosKit dapp template',
//               url: 'https://docs.hyperweb.io/cosmos-kit/',
//               icons: [],
//             },
//           },
//         }}
//         signerOptions={getSignerOptions()}
//         endpointOptions={{
//           endpoints: {
//             osmosis: {
//               rpc: ['http://localhost:26657'],
//             },
//           },
//         }}
//       >
//         <QueryClientProvider client={queryClient}>
//           <Box className={themeClass}>
//             <Layout>
//               {/* @ts-ignore */}
//               <Component {...pageProps} />
//               <Toaster position="top-right" closeButton={true} />
//             </Layout>
//           </Box>
//           {/* <ReactQueryDevtools /> */}
//         </QueryClientProvider>
//       </ChainProvider>
//     </CustomThemeProvider>
//   );
// }

// export default CreateCosmosApp;


import '../styles/globals.css';
import '@interchain-ui/react/styles';

import type { AppProps } from 'next/app';
import { ChainProvider } from '@cosmos-kit/react';
import { QueryClientProvider, QueryClient } from '@tanstack/react-query';
import { Box, Toaster, useTheme } from '@interchain-ui/react';

import { CustomThemeProvider, Layout } from '@/components';
import { wallets } from '@/config/wallets';
import { getSignerOptions } from '@/utils';
import { chainOptions } from '@/config/chains';
import { assetLists, ixoLocalAssets } from '@/config/assets';

import { useEffect } from 'react';
import { suggestIxoLocal } from '@/utils/suggestChain'; // only if you use it

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 2,
      refetchOnMount: false,
      refetchOnWindowFocus: false,
    },
  },
});

function CreateCosmosApp({ Component, pageProps }: AppProps) {
  const { themeClass } = useTheme();

  // Run Keplr suggest only on client
  useEffect(() => {
    suggestIxoLocal();
  }, []);

  // console.log('CHAIN:', chainOptions[0]);
  // console.log('ASSETS:', assetLists);
  // console.log('CHAIN:', JSON.stringify(chainOptions[0], null, 2));
  // console.log('✅ Chain Names:', chainOptions.map(c => c.chain_name));

  const safeChainOptions = chainOptions.filter(c => c && c.chain_name && c.fees?.fee_tokens);

  console.log('Chains passed:', chainOptions.map(c => c.chain_name));
  console.log('Wallets passed:', wallets.map(w => w.walletName));
  
  return (
    <CustomThemeProvider>
      <ChainProvider
        // chains={chainOptions}
        // assetLists={assetLists} // Add custom if needed
        // wallets={wallets}
        chains={chainOptions.filter((c) => !!c?.chain_id && !!c.fees?.fee_tokens)}
        assetLists={[ixoLocalAssets]}
        wallets={wallets.filter((w) => typeof w.walletName === 'string')}
        walletConnectOptions={{
          signClient: {
            projectId: 'a8510432ebb71e6948cfd6cde54b70f7',
            relayUrl: 'wss://relay.walletconnect.org',
            metadata: {
              name: 'ixo Local DApp',
              description: 'CosmosKit ixo local demo',
              url: 'http://localhost:3000',
              icons: [],
            },
          },
        }}
        signerOptions={getSignerOptions()}
        endpointOptions={{
          endpoints: {
            'ixo-local': {
              rpc: ['http://localhost:26657'],
              rest: ['http://localhost:1317'],
            },
          },
        }}
      >
        <QueryClientProvider client={queryClient}>
          <Box className={themeClass}>
            <Layout>
              <Component {...pageProps} />
              <Toaster position="top-right" closeButton={true} />
            </Layout>
          </Box>
        </QueryClientProvider>
      </ChainProvider>
    </CustomThemeProvider>
  );
}

export default CreateCosmosApp;
