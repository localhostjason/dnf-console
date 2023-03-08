import moment from 'moment'
// vue3 not support filter

export const dateFormat = (dataStr: string, pattern: string = 'YYYY-MM-DD HH:mm:ss'): string => {
  return moment(dataStr).format(pattern)
}

export function subtractDate(dataStr: string, days: number = 30, pattern: string = 'YYYY-MM-DD HH:mm:ss'): string {
  return moment(dataStr).subtract(days, 'd').format(pattern)
}

export function addDate(dataStr: string, days: number = 30, pattern: string = 'YYYY-MM-DD HH:mm:ss'): string {
  return moment(dataStr).add(days, 'd').format(pattern)
}

export function compareDate(date_now: string, compare_date: string, pattern: string = 'YYYY-MM-DD HH:mm:ss'): boolean {
  const now = moment(date_now).format(pattern)
  const compare_time = moment(compare_date).format(pattern)

  return moment(compare_time).isBefore(now)
}
