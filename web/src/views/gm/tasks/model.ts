export type Task = {
  charac_no: number
  play_1: number
  play_1_trigger: boolean
}

export type TaskState = {
  loading: boolean
  data: Task[]
}
