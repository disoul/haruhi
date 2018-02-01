import { ResponseState } from "./statuscode";

export interface TaskInput {
  [inputTaskName: string]: TaskOutput
}

export enum OutputType {
  OUTPUT_PLAIN,
  OUTPUT_PGSQL,
}

export interface TaskOutput {
  outputType: OutputType
  data: any
}

export enum DirectiveState {
	TASK_INPUT,
	TASK_OUTPUT,
	TASK_START,
	TASK_STOP,
	TASK_RESTART,
}

export interface Directive {
  action: DirectiveState,
  data: any,
}

export interface DirectiveRequest {
  directive: Directive,
  target: string,
}

export interface DirectiveResponse {
  status: ResponseState,
  output?: TaskOutput,
  errorMsg?: string,
}
