export interface BasicPageParams {
  pageIndex: number;
  pageSize: number;
}

export interface BasicFetchResult<T> {
  list: T[];
  total: number;
}
