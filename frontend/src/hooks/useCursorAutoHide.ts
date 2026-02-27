export function useCursorAutoHide(idleTime = 3000) {
    let timer: number | null = null

    function resetTimer() {
        document.body.style.cursor = 'default'
        if (timer) clearTimeout(timer)
        timer = window.setTimeout(() => {
            document.body.style.cursor = 'none'
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
        document.body.style.cursor = 'default'
    })
}
