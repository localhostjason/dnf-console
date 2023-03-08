import { defineStore } from 'pinia'

interface TagsState {
  visitedViews: any[]
}

export const useTagsStore = defineStore({
  id: 'app-tags',
  state: (): TagsState => ({
    visitedViews: []
  }),

  actions: {
    addVisitedView(view: any): void {
      if (this.visitedViews.some(v => v.path === view.path)) return
      this.visitedViews.push(
        Object.assign({}, view, {
          title: view.meta.title || 'no-name'
        })
      )
    },
    delVisitedView(view: any): any[] {
      for (const [i, v] of this.visitedViews.entries()) {
        if (v.path === view.path) {
          this.visitedViews.splice(i, 1)
          break
        }
      }

      return [...this.visitedViews]
    },
    delOthersVisitedViews(view: any): any[] {
      this.visitedViews = this.visitedViews.filter(v => {
        return v.meta.affix || v.path === view.path
      })

      return [...this.visitedViews]
    },

    delAllVisitedViews(): object {
      this.visitedViews = this.visitedViews.filter(tag => tag.meta.affix)

      return [...this.visitedViews]
    },

    updateVisitedView(view: any): void {
      for (let v of this.visitedViews) {
        if (v.path === view.path) {
          v = Object.assign(v, view)
          break
        }
      }
    },
    addView(view: any) {
      this.addVisitedView(view)
    },
    delView(view: any): object {
      this.delVisitedView(view)
      return {
        visitedViews: [...this.visitedViews]
      }
    },
    delOthersViews(view: any): object {
      this.delOthersVisitedViews(view)
      return {
        visitedViews: [...this.visitedViews]
      }
    },
    delAllViews(): object {
      this.delAllVisitedViews()
      return {
        visitedViews: [...this.visitedViews]
      }
    }
  }
})
