import type { IPageParams, IPageResponse } from "@/types/common";
import request from "@/utils/request";

export interface IMusic {
    id: number,
    name: string,
    url: string,
    pic_url: string,
    artists: string,
    duration: number,
    lyric: string,
}

export interface IGetMusicListParams extends IPageParams {
    name: string
}

export const getMusicList = (params?: IGetMusicListParams) => request.get<IPageResponse<IMusic>>('/music/getMusicList', { params })
export const getNetMusicList = (params?: IGetMusicListParams) => request.get<IPageResponse<IMusic>>('/music/getNetMusicList', { params })

export interface ISaveMusicParams {
    id: number
}
export const saveMusic = (data: ISaveMusicParams) => request.post('/music/saveMusic', data)


