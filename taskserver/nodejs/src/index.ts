import { TaskModel } from './interface/task';
import { ServerOptions } from './interface/server';

import TaskServer from './server';

class HaruhiTaskServer {
  taskModels: { [key: string]: TaskModel }
  server: TaskServer
  
  constructor(taskModels, options: ServerOptions) {
    this.taskModels = taskModels;
    this.initServer(options); 
  }

  initServer(options: ServerOptions) {
    this.server = new TaskServer(this.taskModels, options.port);
  }
}

module.exports = HaruhiTaskServer;