import { ref, computed, onMounted, onBeforeUnmount } from 'vue'

export function usePlayer() {
    const audio = new Audio()
    const src = ref('')
    const isPlaying = ref(false)
    const currentTime = ref(0)
    const duration = ref(0)
    const progress = computed(() =>
        duration.value > 0 ? currentTime.value / duration.value : 0
    )
    const buffered = ref(0)
    const isEnded = ref(false)

    function setSource(url: string) {
        if (src.value === url) return
        src.value = url
        audio.src = url
        isEnded.value = false
    }

    function play() {
        if (!src.value) return

        isEnded.value = false
        audio.play().catch(err => {
            if (err.name === 'AbortError') {
                console.log('播放时加载了新的资源,加载了新的音频资源');
            }
        })
    }

    function seekTo(time: number) {
        audio.currentTime = time
        play()
    }

    function pause() {
        if (!src.value) return
        audio.pause()
    }

    function onPlay() {
        isPlaying.value = true
    }

    function onPause() {
        isPlaying.value = false
    }

    function onTimeUpdate() {
        currentTime.value = audio.currentTime
    }

    function onLoadedMetadata() {
        duration.value = audio.duration
    }

    function onEnded() {
        isPlaying.value = false
        isEnded.value = true
    }

    function onProgress() {
        buffered.value = audio.buffered.end(audio.buffered.length - 1) / audio.duration
    }

    onMounted(() => {
        audio.addEventListener('play', onPlay)
        audio.addEventListener('pause', onPause)
        audio.addEventListener('timeupdate', onTimeUpdate)
        audio.addEventListener('progress', onProgress)
        audio.addEventListener('loadedmetadata', onLoadedMetadata)
        audio.addEventListener('ended', onEnded)
    })

    onBeforeUnmount(() => {
        audio.removeEventListener('play', onPlay)
        audio.removeEventListener('pause', onPause)
        audio.removeEventListener('timeupdate', onTimeUpdate)
        audio.removeEventListener('progress', onProgress)
        audio.removeEventListener('loadedmetadata', onLoadedMetadata)
        audio.removeEventListener('ended', onEnded)
    })

    return {
        src,
        isPlaying,
        currentTime,
        duration,
        progress,
        buffered,
        isEnded,
        setSource,
        play,
        seekTo,
        pause,
    }
}
