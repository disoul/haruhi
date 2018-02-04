import { TaskInput, Directive, DirectiveState, DirectiveRequest, DirectiveResponse } from './interface/server';
import { TaskModel } from './interface/task';
import * as express from 'express';
import { ResponseState } from './interface/statuscode';

export default class TaskServer {
  path: string
  app: express.Application
  taskModels: { [key: string]: TaskModel }


  constructor(taskModels: { [k: string]: TaskModel }, port?: number) {
    this.initServer(port || 8039);
    this.taskModels = taskModels;
  }

  initServer(port: number) {
    this.app = express();
    this.app.listen(port, () => {
      console.log('init task server on', port);
    });

    this.app.post("/directive", this.handleHaruhiDirective)
  }

  async handleHaruhiDirective(req: express.Request, res: express.Response) {
    const body: DirectiveRequest = req.body;
    switch (body.directive.action) {
      case DirectiveState.TASK_START: {
        await this.handleStartTask(res, body.directive.data, body.target);
        break;
      }
      default: {
        res.json({ status: ResponseState.UNSUPPORT });
        break;
      }
    }
  }

  async handleStartTask(res: express.Response, input, target: string) {
    let response: DirectiveResponse = { status: ResponseState.SUCCESS };
    const taskModel = this.taskModels[target];
    response.status = await taskModel.hooks.startTask(input);

    res.json(response);
  }
}
