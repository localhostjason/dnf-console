import defaultSettings from '@/settings'

const title = defaultSettings.title

export const getPageTitle = (key: any): string => {
  if (key) {
    return `${key} - ${title}`
  }
  return title
}
