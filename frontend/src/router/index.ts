
import { getToken } from '@/storage/token'
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
    {
        alias: "/login",
        path: '/',
        component: () => import('@/pages/login/index.vue'),
        children: []
        // children: [{
        //     name: 'Login',
        //     component: () => import('@/pages/login/index.vue'),
        //     path: "login"
        // }]
    },
    {
        path: "/music",
        redirect: '/music/discover',
        component: () => import('@/layout/index.vue'),
        children: [
            // {
            //     name: 'Discover',
            //     component: () => import('@/pages/discover/index.vue'),
            //     path: "discover",
            //     meta: {
            //         title: '发现音乐'
            //     },
            //     children: [
            //         {
            //             name: 'MusicStyle',
            //             component: () => import('@/pages/discover/pages/musicStyle/index.vue'),
            //             path: 'musicStyle/:type',
            //             meta: {
            //                 title: '曲风数据'
            //             }
            //         },
            //         {
            //             name: 'Main',
            //             component: () => import('@/pages/discover/pages/musicStyle/main/index.vue'),
            //             path: 'main',
            //             meta: {
            //                 title: '曲风数据'
            //             }
            //         }
            //     ]
            // },
            {
                name: 'Discover',
                // component: () => import('@/pages/discover/index.vue'),
                component: () => import('@/pages/discover/pages/musicStyle/main/index.vue'),

                path: "discover",
                meta: {
                    title: '发现音乐'
                },
                // children: [

                //     {
                //         name: 'Main',
                //         component: () => import('@/pages/discover/pages/musicStyle/main/index.vue'),
                //         path: 'main',
                //         meta: {
                //             title: '曲风数据'
                //         }
                //     }
                // ]
            },
            {
                path: 'playQueue',
                name: 'PlayQueue',
                component: () => import('@/pages/publicSongList/index.vue'),
                meta: {
                    title: '播放队列'
                }

            },
            {
                path: 'myCollection',
                name: 'MyCollection',
                component: () => import('@/pages/publicSongList/index.vue'),
                meta: {
                    title: '我的收藏'
                }

            },
            {
                path: 'allMusic',
                name: 'AllMusic',
                component: () => import('@/pages/publicSongList/index.vue'),
                meta: {
                    title: '全部音乐'
                }

            },
            {
                path: 'songList',
                name: 'SongList',
                component: () => import('@/pages/publicSongList/index.vue'),
                meta: {
                    title: '歌单列表'
                }

            },

            // {
            //     path: "personalCenter",
            //     name: 'PersonalCenter',
            //     component: () => import('@/pages/personalCenter/index.vue'),
            //     meta: {
            //         title: '个人中心'
            //     }
            // }, 
            {
                path: "setting",
                name: 'Setting',
                component: () => import('@/pages/setting/index.vue'),
                meta: {
                    title: '设置'
                }
            }
        ]

    }

]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach((to) => {


    if (!getToken() && to.path === "/") {
        return true
    }

    if (!getToken() && to.path != "/" && to.path != "/login") {
        return { path: "/" }
    }


})


export default router