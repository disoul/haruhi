import { TaskInput, Directive, DirectiveState, DirectiveRequest, DirectiveResponse } from './interface/server';
import * as express from 'express';
import { ResponseState } from './interface/statuscode';

class TaskServer {
  path: string
  app: express.Application
  startTask: (input: TaskInput, taskName: string) => Promise<ResponseState>

  constructor(port?: number) {
    this.initServer(port || 8039)
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
    let response: DirectiveResponse;
    response.status = await this.startTask(input, target);

    res.json(response);
  }
}
