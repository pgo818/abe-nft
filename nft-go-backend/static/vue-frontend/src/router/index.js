import { createRouter, createWebHistory } from 'vue-router'
import store from '@/store'

// 页面组件
import Home from '@/views/Home.vue'
import NFTDashboard from '@/views/nft/NFTDashboard.vue'
import ABEDashboard from '@/views/abe/ABEDashboard.vue'
import DIDDashboard from '@/views/did/DIDDashboard.vue'
import NotFound from '@/views/NotFound.vue'

// NFT相关页面
import AllNFTs from '@/views/nft/AllNFTs.vue'
import MintNFT from '@/views/nft/MintNFT.vue'
import MyNFTs from '@/views/nft/MyNFTs.vue'
import MetadataManager from '@/views/nft/MetadataManager.vue'
import RequestManager from '@/views/nft/RequestManager.vue'

// ABE相关页面
import ABESetup from '@/views/abe/ABESetup.vue'
import ABEKeyGen from '@/views/abe/ABEKeyGen.vue'
// import ABEEncrypt from '@/views/abe/ABEEncrypt.vue'
// import ABEDecrypt from '@/views/abe/ABEDecrypt.vue'
// import ABELogs from '@/views/abe/ABELogs.vue'

// DID相关页面
import DIDList from '@/views/did/DIDList.vue'
import DIDCreate from '@/views/did/DIDCreate.vue'
import DoctorDID from '@/views/did/DoctorDID.vue'
// import VCIssue from '@/views/did/VCIssue.vue'
// import VCManage from '@/views/did/VCManage.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/nft',
        name: 'NFTDashboard',
        component: NFTDashboard,
        children: [
            {
                path: 'all',
                name: 'AllNFTs',
                component: AllNFTs
            },
            {
                path: 'mint',
                name: 'MintNFT',
                component: MintNFT
            },
            {
                path: 'my',
                name: 'MyNFTs',
                component: MyNFTs
            },
            {
                path: 'metadata',
                name: 'MetadataManager',
                component: MetadataManager
            },
            {
                path: 'requests',
                name: 'RequestManager',
                component: RequestManager
            },
            // 默认子路由
            {
                path: '',
                redirect: { name: 'AllNFTs' }
            }
        ]
    },
    {
        path: '/abe',
        name: 'ABEDashboard',
        component: ABEDashboard,
        children: [
            {
                path: 'setup',
                name: 'ABESetup',
                component: ABESetup
            },
            {
                path: 'keygen',
                name: 'ABEKeyGen',
                component: ABEKeyGen
            },
            /* 暂未实现的组件
            {
                path: 'encrypt',
                name: 'ABEEncrypt',
                component: ABEEncrypt
            },
            {
                path: 'decrypt',
                name: 'ABEDecrypt',
                component: ABEDecrypt
            },
            {
                path: 'logs',
                name: 'ABELogs',
                component: ABELogs
            },
            */
            // 默认子路由
            {
                path: '',
                redirect: { name: 'ABESetup' }
            }
        ]
    },
    {
        path: '/did',
        name: 'DIDDashboard',
        component: DIDDashboard,
        children: [
            {
                path: 'list',
                name: 'DIDList',
                component: DIDList
            },
            {
                path: 'create',
                name: 'DIDCreate',
                component: DIDCreate
            },
            {
                path: 'doctor',
                name: 'DoctorDID',
                component: DoctorDID
            },
            /* 暂未实现的组件
            {
                path: 'vc/issue',
                name: 'VCIssue',
                component: VCIssue
            },
            {
                path: 'vc/manage',
                name: 'VCManage',
                component: VCManage
            },
            */
            // 默认子路由
            {
                path: '',
                redirect: { name: 'DIDList' }
            }
        ]
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

// 导航守卫，检查需要钱包连接的路由
router.beforeEach((to, from, next) => {
    // 检查路由是否需要钱包连接
    if (to.matched.some(record => record.meta.requiresWallet)) {
        // 检查钱包是否已连接
        if (!store.state.wallet.isConnected) {
            // 存储目标路由，连接钱包后跳转
            store.commit('wallet/setRedirectRoute', to.fullPath)
            // 显示钱包连接提示
            store.commit('wallet/setShowConnectPrompt', true)
            // 重定向到首页
            next({ name: 'Home' })
        } else {
            next()
        }
    } else {
        next()
    }
})

export default router 