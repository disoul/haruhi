import { ResponseState } from './statuscode';

export interface TaskModelHook {
  startTask: (input: any) => Promise<ResponseState>,
}

export interface TaskModel {
  name: string // unique id
  depends: Array<string>
  type: string
  path?: string
  config?: any
  hooks: TaskModelHook
}
