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

