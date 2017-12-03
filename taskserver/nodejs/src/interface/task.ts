export interface TaskModel {
  name: string
  depends: Array<string>
  type: string
  path: string
}
