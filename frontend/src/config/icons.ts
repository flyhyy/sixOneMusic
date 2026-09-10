
/**
 * @Author: hyy
 * @Date: 2026-06-09 13:08:52
 * @Description: 全局图标配置
 */



export const Icons = {
    // 头部音乐图标
    headerMusicNote: 'ri:music-fill',
    // 头部搜索框图标
    headerSearch: 'ri:search-line',
    // 明亮主题图标
    light: 'ri:sun-line',
    // 暗黑主题图标
    dark: 'ri:moon-line',
    // github图标
    github: 'ri:github-fill',
    // 侧边栏_发现音乐
    sideHome: 'ri:home-4-line',
    //侧边栏_播放队列
    sidePlayQueue: 'ri:play-list-2-line',
    //侧边栏_我的收藏
    sideMyCollection: 'ri:heart-3-line',
    //侧边栏_个人中心
    sidePersonalCenter: 'ri:user-3-line',
    //侧边栏_设置
    sideSetting: 'ri:settings-3-line',
    // 侧边栏_我的歌单_编辑
    sideMyPlayListEdit: 'ri:edit-line',
    // 侧边栏_我的歌单_播放列表
    sideMyPlayList: 'ri:play-list-line',
    // 发信音乐_当前时间_凌晨
    discoverCurrentTimeEarlyMorning: 'material-symbols:bedtime',
    // 发现音乐_当前时间_早上
    discoverCurrentTimeMorning: 'material-symbols:brightness-high',
    // 发现音乐_当前时间_中午
    discoverCurrentTimeNoon: 'material-symbols:brightness-high',
    // 发现音乐_当前时间_下午
    discoverCurrentTimeAfternoon: 'material-symbols:wb-twilight',
    // 发现音乐_当前时间_晚上
    discoverCurrentTimeNight: 'material-symbols:dark-mode',

    // 播放
    play: 'ri:play-fill',
    // 随机播放
    random: 'ri:shuffle-line',
    // 发现音乐_风格
    discvoerStyle: 'system-uicons:tag',
    // 古典音乐
    discvoerStyleClassical: 'material-symbols:piano',
    // 摇滚乐
    discvoerStyleRock: 'mdi:guitar-electric',
    // 民谣音乐
    discvoerStyleFolk: 'mdi:guitar-acoustic',
    //电子音乐
    discvoerStyleEDM: 'material-symbols:graphic-eq',
    // 嘻哈/说唱
    discvoerStyleHip_Hop: 'mdi:boombox',
    // 爵士乐
    discvoerStyleJazz: 'mdi:saxophone',
    // 打碟/混音
    discvoerStyleDJ_Remix: 'mdi:turntable',
    // 流行乐
    discvoerStylePop: 'material-symbols:mic-external-on',
    // 轻音乐
    discvoerStyleLo_Fi: 'material-symbols:headphones',
    // 复古/经典
    discvoerStyleRetro: 'material-symbols:album',
    // 过滤
    discoverStyleFilter: 'material-symbols:filter-list-rounded',
    // 历史
    history: 'ri:history-line',
    // 右箭头
    arrowRight: 'ri:arrow-right-s-line',
    // 下箭头
    arrowDown: 'ri:arrow-down-s-line',
    // 上箭头
    arrowUp: 'ri:arrow-up-s-line',
    // 音乐
    music: 'ri:music-2-fill',
    // 艺术家
    artist: 'ri:user-voice-line',
    //专辑
    album: 'ri:album-line',
    // 光盘
    disc: 'ri:disc-fill',
    // 音乐-线条
    musicLine: 'ri:music-line',
    //时间
    time: 'ri:time-line',
    //罗盘
    compass: 'ri:compass-3-line',
    // 垃圾桶
    delete: 'ri:delete-bin-line',
    // 爱心
    heart: 'ri:heart-fill',
    // 爱心(线条)
    heartLine: 'ri:heart-line',
    //添加
    add: 'ri:add-line',
    //文件夹
    folder: 'ri:folder-add-line',
    // 退出
    logout: 'ri:logout-box-line',
    // 堆(线条)
    stackLine: 'ri:stack-line',
    // 说
    speak: 'ri:speak-line',
    // 信息
    information: 'ri:information-line',
    // 文件
    folderLine: "ri:folder-open-line",
    // 列表循环
    listRepeat: 'ri:repeat-line',
    // 单曲循环
    oneRepeat: "ri:repeat-one-line",
    // 上一曲
    preSong: "ri:skip-back-fill",
    // 下一曲
    nextSong: 'ri:skip-forward-fill',
    // 歌词
    lrc: "ri:chat-smile-2-line",
    // 喇叭
    volume: "ri:volume-up-line",
    // 静音
    mute: "ri:volume-mute-line",
    // 暂停
    pause: "ri:pause-fill",
    // 关闭
    close: 'ri:close-line'


} as const;

export type IconName = typeof Icons[keyof typeof Icons]
