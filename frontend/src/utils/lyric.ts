export type ILyric = {
    time: number
    text: string
}

export function parseLyrics(lyrics: string): ILyric[] {
    const timeExp = /\[(\d{2}):(\d{2}(?:\.\d{1,3})?)\]/;

    return lyrics
        .split('\n')
        .map(line => {
            const match = timeExp.exec(line);

            if (match) {
                const minutes = parseInt(match[1], 10);
                const seconds = parseFloat(match[2]);
                const time = minutes * 60 + seconds;
                const text = line.replace(timeExp, '');
                return { time, text };
            }

            return null;
        })
        .filter((item): item is ILyric => item !== null);
}
