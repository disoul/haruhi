import { TaskInput } from './interface/server';
import * as express from 'express';

class TaskServer {
  path: string
  app: express.Application
  startTask: (input: TaskInput) => boolean

  constructor(port?: number) {
    this.initServer(port || 8039)
  }

  initServer(port: number) {
    this.app = express();
    this.app.listen(port, () => {
      console.log('init task server on', port);
    });

    this.app.post("/handle", this.handleHaruhi)
  }

  handleHaruhi(req: express.Request, res: express.Response) {
    
  }
}
