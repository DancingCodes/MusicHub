import type { IPageParams, IPageResponse } from "@/types/common";
import request from "@/utils/request";

export interface IMusic {
    id: number,
    name: string,
    picUrl: string,
    artists: { id: number; name: string }[],
    duration: number,
    lyric: string,
    url: string,
}

export interface IGetMusicListParams extends IPageParams {
    name: string
}

export const getMusicList = (params?: IGetMusicListParams) => request.get<IPageResponse<IMusic>>('/music', { params })
export const getWyMusicList = (params?: IGetMusicListParams) => request.get<IPageResponse<IMusic>>('/music/wyList', { params })

export interface IAddMusicParams {
    id: number
}
export const addMusic = (data: IAddMusicParams) => request.post('/music', data)


