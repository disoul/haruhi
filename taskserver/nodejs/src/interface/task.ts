export interface TaskModel {
  name: string // unique id
  depends: Array<string>
  type: string
  path: string
  config?: any
}
