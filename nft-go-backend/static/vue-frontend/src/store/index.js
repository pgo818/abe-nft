import { createStore } from 'vuex'
import appModule from './modules/app'
import walletModule from './modules/wallet'
import nftModule from './modules/nft'
import abeModule from './modules/abe'
import didModule from './modules/did'

export default createStore({
    modules: {
        app: appModule,
        wallet: walletModule,
        nft: nftModule,
        abe: abeModule,
        did: didModule
    }
}) 