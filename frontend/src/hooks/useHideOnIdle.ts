export function useHideOnIdle(idleTime = 3000) {
    const isIdle = ref(false)
    let timer: number | null = null

    function resetTimer() {
        isIdle.value = false
        if (timer) {
            clearTimeout(timer)
        }
        timer = window.setTimeout(() => {
            isIdle.value = true
        }, idleTime)
    }

    function onMouseMove() {
        resetTimer()
    }

    onMounted(() => {
        window.addEventListener('mousemove', onMouseMove)
        resetTimer()
    })

    onBeforeUnmount(() => {
        window.removeEventListener('mousemove', onMouseMove)
        if (timer) clearTimeout(timer)
    })

    return {
        isIdle,
    }
}
