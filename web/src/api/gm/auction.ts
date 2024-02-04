import { http } from '@/utils/http'
import { AuctionState } from '@/views/gm/auction/model'

export const getAuctionState = <T = AuctionState[]>(): Promise<T> => {
  return http.request({
    url: `/gm/auction/state`,
    method: 'get'
  })
}

export const openAuction = <T = AuctionState[]>(name: string): Promise<T> => {
  return http.request({
    url: `/gm/auction/${name}/open`,
    method: 'post'
  })
}
