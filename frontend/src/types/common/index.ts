export interface IResponse<T = any> {
    code: number
    message: string
    data: T
}

export interface IPageParams {
    pageNo: number
    pageSize: number
}

export interface IPageResponse<T> {
    list: T[]
    total: number
}
