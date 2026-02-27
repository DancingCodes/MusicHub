<template>
    <div class="h-100vh overflow-hidden">

        <div v-if="currentMusic" class="h-100% flex flex-col blur-10px transition-500"
            :class="{ 'animate-[unblur_1s_ease_forwards]': currentMusic }">
            <div
                class="flex-1 flex items-center overflow-hidden px-0.6rem max-sm:flex-col max-sm:py-0.5rem max-sm:box-border">
                <div class="flex-1 flex flex-col items-center justify-center max-sm:flex-none">

                    <div class="relative w-1.5rem h-1.5rem max-sm:hidden">
                        <div class="w-100% h-100% animate-spin bg-[url('https://filestore.moonc.love/uploadFiles/1752216410894-161065103.png')] bg-cover p-0.3rem box-border max-sm:w-2rem max-sm:h-2rem"
                            style="animation-duration: 20s" :style="{
                                animationPlayState: isPlaying ? 'running' : 'paused',
                            }">
                            <img :src="currentMusic.picUrl" class="w-100% h-100% rounded-full" />
                        </div>
                        <img src="https://filestore.moonc.love/uploadFiles/1752216426296-916133906.png" alt=""
                            class="h-1.1rem absolute top-0 right-0 transform-translate-[76%,-20%] transform-origin-top-right transition-500"
                            :class="{ 'transform-rotate-[-16deg]': !isPlaying }">
                    </div>

                    <div class="flex flex-col items-center justify-center text-0.2rem max-sm:text-0.6rem">
                        <div class="mt-0.1rem">
                            {{ currentMusic.name }}
                        </div>
                        <div class="text-0.16rem max-sm:text-0.5rem">
                            {{currentMusic.artists.map(i => i.name).join(' - ')}}
                        </div>
                    </div>
                </div>
                <div class="flex-1 h-88% overflow-y-auto flex flex-col gap-0.05rem items-center text-0.15rem text-center no-scrollbar max-sm:text-0.5rem max-sm:mt-0.3rem max-sm:gap-0.1rem max-sm:h-unset"
                    @scroll="scrollLyric">
                    <div v-for="(item, index) in parsedLyrics" :ref="el => { lyricRefs[index] = el as HTMLElement }"
                        @click="seekToLyric(item.time)"
                        class="transition-500 hover:underline hover:decoration-dashed cursor-pointer"
                        :class="{ 'color-#ff0000 font-bold': currentLyricIndex === index }">
                        {{ item.text }}
                    </div>
                </div>
            </div>
            <div>
                <div class="flex items-center justify-center gap-0.2rem justify-center text-0.25rem py-0.05rem box-border transition-500 max-sm:text-0.8rem max-sm:gap-0.8rem"
                    :class="{ 'sm:opacity-0 sm:pointer-events-none': isIdle }">
                    <span>
                        <svg class="cursor-pointer" @click="playRelative('prev')" xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 1024 1024" width="1em" height="1em" fill="currentColor">
                            <path
                                d="M356.77 605.22L701.8 849.53c69.8 49.42 166.29-0.49 166.29-86.01V261.25c0-85.52-96.49-135.43-166.29-86.01L356.77 419.55c-64.02 45.34-64.02 140.34 0 185.67zM210.42 154.36c30.38 0 55 24.62 55 55v606.07c0 30.38-24.62 55-55 55s-55-24.62-55-55V209.36c0-30.38 24.62-55 55-55z" />
                        </svg>
                    </span>
                    <span class="cursor-pointer" @click="togglePlay">
                        <svg v-if="isPlaying" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024" width="1em"
                            height="1em" fill="currentColor">
                            <path
                                d="M332.26 853.89c-49.71 0-90-40.29-90-90v-503c0-49.71 40.29-90 90-90s90 40.29 90 90v503c0 49.7-40.3 90-90 90zM691.26 853.89c-49.71 0-90-40.29-90-90v-503c0-49.71 40.29-90 90-90s90 40.29 90 90v503c0 49.7-40.3 90-90 90z" />
                        </svg>

                        <svg v-else xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024" width="1em" height="1em"
                            fill="currentColor">
                            <path
                                d="M783.74 401.86L372.23 155.28c-85.88-51.46-195.08 10.41-195.08 110.53v493.16c0 100.12 109.2 161.99 195.08 110.53l411.51-246.58c83.5-50.04 83.5-171.03 0-221.06z" />
                        </svg>
                    </span>
                    <span class="cursor-pointer" @click="playRelative('next')">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024" width="1em" height="1em"
                            fill="currentColor">
                            <path
                                d="M665.47 417.65l-345.03-244.3c-69.8-49.42-166.29 0.49-166.29 86.01v502.27c0 85.52 96.49 135.43 166.29 86.01l345.03-244.31c64.02-45.34 64.02-140.34 0-185.68zM811.82 868.52c-30.38 0-55-24.62-55-55V207.46c0-30.38 24.62-55 55-55s55 24.62 55 55v606.07c0 30.37-24.62 54.99-55 54.99z" />
                        </svg>
                    </span>

                    <span class="cursor-pointer" @click="togglePlayMode">
                        <svg v-if="playMode === 'loop'" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg"
                            width="1em" height="1em" fill="currentColor">
                            <path
                                d="M361.5 727.8c-119.1 0-215.9-96.9-215.9-215.9 0-119.1 96.9-215.9 215.9-215.9 2.3 0 4.6-0.2 6.8-0.6v58.3c0 12.3 14 19.4 23.9 12.1l132.6-97.6c8.1-6 8.1-18.2 0-24.2l-132.6-97.6c-9.9-7.3-23.9-0.2-23.9 12.1v58.1c-2.2-0.4-4.5-0.6-6.8-0.6-39.8 0-78.5 7.9-115 23.4-35.2 15-66.8 36.3-94 63.5s-48.6 58.8-63.5 94c-15.5 36.5-23.4 75.2-23.4 115s7.9 78.5 23.4 115c15 35.2 36.3 66.8 63.5 94s58.8 48.6 94 63.5c36.5 15.5 75.2 23.4 115 23.4 22.1 0 40-17.9 40-40s-17.9-40-40-40zM938.2 396.9c-15-35.2-36.3-66.8-63.5-94s-58.8-48.6-94-63.5c-36.5-15.5-75.2-23.4-115-23.4-22.1 0-40 17.9-40 40s17.9 40 40 40c119.1 0 215.9 96.9 215.9 215.9 0 119.1-96.9 215.9-215.9 215.9-4.1 0-8.1 0.6-11.8 1.8v-60.8c0-12.3-14-19.4-23.9-12.1l-132.6 97.6c-8.1 6-8.1 18.2 0 24.2L629.9 876c9.9 7.3 23.9 0.2 23.9-12.1V806c3.7 1.2 7.7 1.8 11.8 1.8 39.8 0 78.5-7.9 115-23.4 35.2-15 66.8-36.3 94-63.5s48.6-58.8 63.5-94c15.5-36.5 23.4-75.2 23.4-115s-7.8-78.5-23.3-115z" />
                        </svg>
                        <svg v-if="playMode === 'single'" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg"
                            width="1em" height="1em" fill="currentColor">
                            <path
                                d="M361.5 727.8c-119.1 0-215.9-96.9-215.9-215.9 0-119.1 96.9-215.9 215.9-215.9 2.3 0 4.6-0.2 6.8-0.6v58.3c0 12.3 14 19.4 23.9 12.1l132.6-97.6c8.1-6 8.1-18.2 0-24.2l-132.6-97.6c-9.9-7.3-23.9-0.2-23.9 12.1v58.1c-2.2-0.4-4.5-0.6-6.8-0.6-39.8 0-78.5 7.9-115 23.4-35.2 15-66.8 36.3-94 63.5s-48.6 58.8-63.5 94c-15.5 36.5-23.4 75.2-23.4 115s7.9 78.5 23.4 115c15 35.2 36.3 66.8 63.5 94s58.8 48.6 94 63.5c36.5 15.5 75.2 23.4 115 23.4 22.1 0 40-17.9 40-40s-17.9-40-40-40zM938.2 396.9c-15-35.2-36.3-66.8-63.5-94s-58.8-48.6-94-63.5c-36.5-15.5-75.2-23.4-115-23.4-22.1 0-40 17.9-40 40s17.9 40 40 40c119.1 0 215.9 96.9 215.9 215.9 0 119.1-96.9 215.9-215.9 215.9-4.1 0-8.1 0.6-11.8 1.8v-60.8c0-12.3-14-19.4-23.9-12.1l-132.6 97.6c-8.1 6-8.1 18.2 0 24.2l132.6 97.6c9.9 7.3 23.9 0.2 23.9-12.1V806c3.7 1.2 7.7 1.8 11.8 1.8 39.8 0 78.5-7.9 115-23.4 35.2-15 66.8-36.3 94-63.5s48.6-58.8 63.5-94c15.5-36.5 23.4-75.2 23.4-115s-7.8-78.5-23.3-115z" />
                            <path
                                d="M512.8 660.6c22.1-0.1 39.9-18.1 39.8-40.2l-1.2-214.1c-0.1-22-18-39.8-40-39.8h-0.2c-22.1 0.1-39.9 18.1-39.8 40.2l1.2 214.1c0.1 22 18 39.8 40 39.8h0.2z" />
                        </svg>
                        <svg v-if="playMode === 'shuffle'" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg"
                            width="1em" height="1em" fill="currentColor">
                            <path
                                d="M914.2 705L796.4 596.8c-8.7-8-22.7-1.8-22.7 10V688c-69.5-1.8-134-39.7-169.3-99.8l-45.1-77 47-80.2c34.9-59.6 98.6-97.4 167.4-99.8v60.1c0 11.8 14 17.9 22.7 10l117.8-108.1c5.8-5.4 5.8-14.6 0-19.9L796.4 165c-8.7-8-22.7-1.8-22.7 10v76H758c-4.7 0-9.3 0.8-13.5 2.3-36.5 4.7-72 16.6-104.1 35-42.6 24.4-78.3 59.8-103.1 102.2L513 432l-24.3-41.5c-24.8-42.4-60.5-77.7-103.1-102.2C343 263.9 294.5 251 245.3 251H105c-22.1 0-40 17.9-40 40s17.9 40 40 40h140.3c71.4 0 138.3 38.3 174.4 99.9l47 80.2-45.1 77c-36.2 61.7-103 99.9-174.4 99.9H105c-22.1 0-40 17.9-40 40s17.9 40 40 40l142 0.1h0.2c49.1 0 97.6-12.9 140.2-37.3 42.7-24.4 78.3-59.8 103.2-102.2l22.4-38.3 22.4 38.3c24.8 42.4 60.5 77.8 103.2 102.2 33.1 18.9 69.6 30.9 107.3 35.4 3.8 1.2 7.8 1.8 11.9 1.8l15.9 0.1v55c0 11.8 14 17.9 22.7 10L914.2 725c5.9-5.5 5.9-14.7 0-20z" />
                        </svg>
                    </span>

                    <span class="hidden max-sm:block" @click="isCollapsed = false;">
                        <svg viewBox="0 0 1024 1024" width="1em" height="1em" fill="currentColor"
                            xmlns="http://www.w3.org/2000/svg">
                            <path
                                d="M226.58 320.39h-48c-35.9 0-65-29.1-65-65s29.1-65 65-65h48c35.9 0 65 29.1 65 65s-29.1 65-65 65zM723.58 320.39h-322c-35.9 0-65-29.1-65-65s29.1-65 65-65h322c35.9 0 65 29.1 65 65s-29.1 65-65 65zM402.58 580.39h-224c-35.9 0-65-29.1-65-65s29.1-65 65-65h224c35.9 0 65 29.1 65 65s-29.1 65-65 65zM251.58 839.39h-73c-35.9 0-65-29.1-65-65s29.1-65 65-65h73c35.9 0 65 29.1 65 65s-29.1 65-65 65zM424.58 839.39h-1c-35.9 0-65-29.1-65-65s29.1-65 65-65h1c35.9 0 65 29.1 65 65s-29.1 65-65 65zM954.82 588.59l-93.4-161.78a59.189 59.189 0 0 0-51.27-29.6h-186.8a59.201 59.201 0 0 0-51.27 29.6l-93.4 161.78a59.225 59.225 0 0 0 0 59.2l93.4 161.78a59.189 59.189 0 0 0 51.27 29.6h186.81c21.15 0 40.69-11.28 51.27-29.6l93.4-161.78a59.2 59.2 0 0 0-0.01-59.2z m-238.07 95.78c-36.55 0-66.18-29.63-66.18-66.18 0-36.55 29.63-66.18 66.18-66.18 36.55 0 66.18 29.63 66.18 66.18 0 36.55-29.63 66.18-66.18 66.18z" />
                        </svg>
                    </span>
                </div>
                <div class="flex-1 h-0.06rem relative max-sm:h-0.16rem">
                    <div class="absolute h-100% w-100% left-0 right-0 rounded-tr-6px rounded-br-6px bg-[rgba(0,0,0,0.3)] transition-500 "
                        :style="{ transform: `translateX(${(buffered - 1) * 100}%)` }"></div>

                    <div class="absolute h-100% w-100% left-0 right-0 rounded-tr-6px rounded-br-6px bg-#ff0000  transition-500 "
                        :style="{ transform: `translateX(${(progress - 1) * 100}%)` }">
                    </div>
                </div>
            </div>
        </div>

        <div v-else class="h-100% flex items-center justify-center text-0.5rem max-sm:text-1rem cursor-pointer"
            @click="isCollapsed = !isCollapsed">
            Moonc music
        </div>

        <div class="fixed h-100% w-32% right-0 top-0 bg-[rgba(0,0,0,0.7)] flex flex-col gap-0.1rem px-0.1rem py-0.06rem box-border transition-500 max-sm:w-100% max-sm:px-0.4rem max-sm:py-0.5rem  max-sm:bg-[rgba(0,0,0,0.7)]"
            :class="{ 'sm:transform-translate-x-100%': isCollapsed, 'max-sm:transform-translate-y-100%': isCollapsed }">
            <div class="absolute right-100% top-50% transform translate-[-50%,-50%] rotate-[90deg] text-0.2rem text-center transition-500 cursor-pointer 
             max-sm:position-unset max-sm:transform-unset max-sm:text-0.5rem"
                :class="{ 'sm:rotate-[-90deg]': !isCollapsed, 'sm:transform-translate-x-100%': isIdle && isCollapsed }"
                @click="isCollapsed = !isCollapsed;">
                <svg viewBox="0 0 1819 1024" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em"
                    fill="currentColor">
                    <path
                        d="M1788.06 37.13h-5.63a112.53 112.53 0 0 0-154.16 0L909.22 750.56 191.30 31.51a112.53 112.53 0 0 0-154.16 0 112.53 112.53 0 0 0 0 154.16L832.70 992.49a112.53 112.53 0 0 0 154.16 0l801.20-801.20a112.53 112.53 0 0 0 0-154.16z" />
                </svg>
            </div>

            <div
                class="flex items-center bg-#102a27 rounded-3px px-0.05rem py-0.03rem box-border text-0.12rem gap-0.05rem max-sm:text-0.46rem max-sm:px-0.2rem max-sm:py-0.2rem max-sm:gap-0.16rem">
                <span class="cursor-pointer" :class="{ 'color-#ff0000': isSerachWy }"
                    @click.stop="isSerachWy = !isSerachWy; sreachMusic()">
                    T
                </span>
                <input type="text"
                    class="flex-1 bg-transparent border-none outline-unset p-0 color-#fff text-0.12rem py-0.01rem max-sm:text-0.46rem max-sm:w-100%"
                    v-model="queryParams.name" @keydown.enter="sreachMusic" @click.stop>
                <span class="cursor-pointer" @click.stop="sreachMusic">搜索</span>
            </div>

            <div class="flex-1 flex flex-col gap-0.06rem text-0.12rem overflow-y-auto no-scrollbar max-sm:text-0.46rem max-sm:gap-0.3rem "
                ref="musicListRef" @scroll="scrollMusicList">
                <div v-for="(i, index) in musicList" :key="i.id"
                    :ref="el => { musicRowRefs[index] = el as HTMLElement }" @click.stop="playMusic(i)"
                    class="flex items-center justify-between cursor-pointer transition-500 gap-0.2rem max-sm:gap-1rem"
                    :class="{ 'color-#ff0000 font-bold': currentMusic?.id === i.id }">
                    <div class="overflow-hidden whitespace-nowrap text-ellipsis flex-1">{{ i.name }}</div>
                    <div class="overflow-hidden whitespace-nowrap text-ellipsis w-1rem text-right max-sm:w-2rem">
                        {{i.artists.map(i =>
                            i.name).join(' - ')}}</div>
                </div>

                <div v-if="loading" class="text-center text-0.13rem max-sm:text-0.6rem">
                    loading...
                </div>
                <div v-if="!loading && !total" class="text-center text-0.13rem max-sm:text-0.6rem">
                    无数据
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { addMusic, getMusicList, getWyMusicList, type IMusic } from './api/music';
import { useCursorAutoHide } from './hooks/useCursorAutoHide';
import { useHideOnIdle } from './hooks/useHideOnIdle';
import { usePlayer } from './hooks/usePlayer';
import { parseLyrics, type ILyric } from './utils/lyric';
import { setTitle } from './utils/setTitle';

const { isIdle } = useHideOnIdle()
useCursorAutoHide()

const { isPlaying, currentTime, progress, buffered, isEnded, setSource, play, seekTo, pause } = usePlayer()
watch(currentTime, val => {
    autoMoveLyric(val)
})
watch(isEnded, val => {
    if (val) {
        playRelative('next')
    }
})

const isSerachWy = ref(false)

const queryParams = reactive({
    pageNo: 1,
    pageSize: 30,
    name: '',
})

const sreachMusic = () => {
    musicList.value = []
    queryParams.pageNo = 1
    getList()
}

const musicList = ref<IMusic[]>([])
const total = ref(0)
const loading = ref(false)
const isCollapsed = ref(true)
watch(isEnded, val => {
    if (val) {
        centerMusicItem()
    }
})

const musicListRef = ref()
const musicRowRefs = ref<HTMLElement[]>([])
const getList = async () => {
    loading.value = true
    const api = isSerachWy.value ? getWyMusicList : getMusicList
    const res = await api(queryParams)
    musicList.value.push(...res.data.list)
    total.value = res.data.total
    loading.value = false
}

const scrollMusicList = () => {
    const { scrollHeight, scrollTop, clientHeight } = musicListRef.value;
    if (scrollHeight - (scrollTop + clientHeight) <= 50) {
        loadMore()
    }
}

const loadMore = async () => {
    if (musicList.value.length >= total.value || loading.value) {
        return
    }
    queryParams.pageNo++
    await getList();
};

const centerMusicItem = () => {
    const current = currentMusic.value;
    if (!current) return;
    const index = musicList.value.findIndex(i => i.id === current.id)
    musicRowRefs.value[index].scrollIntoView({
        behavior: 'smooth',
        block: 'center',
    })
}


const currentMusic = ref<IMusic>()
const parsedLyrics = ref<ILyric[]>([])
const playBlurAnimation = ref(false)
const playMusic = async (music: IMusic) => {
    playBlurAnimation.value = false
    if (isSerachWy.value) {
        const res = await addMusic({
            id: music.id
        })
        music = res.data
    }
    currentMusic.value = music

    setTitle(currentMusic.value.name)
    parsedLyrics.value = parseLyrics(currentMusic.value.lyric)

    setSource(currentMusic.value.url)
    play()

    if (!isCollapsed.value) {
        centerMusicItem()
    }

    requestAnimationFrame(() => {
        playBlurAnimation.value = true
    })
}

const togglePlay = () => {
    isPlaying.value ? pause() : play()
}

const playRelative = async (direction: 'prev' | 'next') => {
    const list = musicList.value
    if (list.length === 0) return

    let index = 0

    if (playMode.value === 'loop') {
        index = list.findIndex(i => i.id === currentMusic.value?.id)
        if (index === -1) {
            index = direction === 'prev' ? list.length - 1 : 0
        } else {
            if (direction === 'prev') {
                index = index === 0 ? list.length - 1 : index - 1
            }
            if (direction === 'next') {
                if (index === list.length - 1) {
                    if (list.length >= total.value) {
                        index = 0
                    } else {
                        await loadMore()
                        index = index + 1
                    }
                } else {
                    index = index + 1
                }
            }
        }
    }


    if (playMode.value === 'single') {
        index = list.findIndex(i => i.id === currentMusic.value?.id)
    }

    if (playMode.value === 'shuffle') {
        index = Math.floor(Math.random() * list.length)
    }

    playMusic(list[index])
}

const playMode = ref<'loop' | 'single' | 'shuffle'>('loop')
const togglePlayMode = () => {
    if (playMode.value === 'loop') {
        playMode.value = 'single'
        return
    }
    if (playMode.value === 'single') {
        playMode.value = 'shuffle'
        return
    }
    if (playMode.value === 'shuffle') {
        playMode.value = 'loop'
        return
    }
}

const lyricRefs = ref<HTMLElement[]>([])
const currentLyricIndex = ref(-1)
const autoMoveLyric = (currentTime: number) => {
    if (isScrollLyric.value) {
        return
    }

    const index = parsedLyrics.value.findIndex((i, x) => {
        const next = parsedLyrics.value[x + 1]
        return currentTime >= i.time && (!next || currentTime < next.time)
    })
    const safeIndex = index === -1 ? 0 : index

    if (safeIndex !== currentLyricIndex.value) {
        setTitle(parsedLyrics.value[index].text)
        lyricRefs.value[safeIndex].scrollIntoView({
            behavior: 'smooth',
            block: 'center',
        })
        currentLyricIndex.value = safeIndex
    }
}

const isScrollLyric = ref(false)
const scrollTimeout = ref()
const scrollLyric = () => {
    isScrollLyric.value = true
    if (scrollTimeout) clearTimeout(scrollTimeout.value)
    scrollTimeout.value = setTimeout(() => {
        isScrollLyric.value = false
    }, 3000)
}


const seekToLyric = (time: number) => {
    isScrollLyric.value = false
    seekTo(time)
}

onMounted(() => {
    getList()
})
</script>
<style>
@keyframes unblur {
    from {
        filter: blur(8px);
    }

    to {
        filter: blur(0);
    }
}
</style>