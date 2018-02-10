import axios from 'axios';
import { TaskModel } from './interface/task';
import { ServerOptions } from './interface/server';

import TaskServer from './server';

class HaruhiTaskServer {
  taskModels: { [key: string]: TaskModel }
  server: TaskServer
  public haruhiHost: string
  
  constructor(taskModels, options: ServerOptions) {
    this.taskModels = taskModels;
    this.haruhiHost = options.haruhiAddr;
    this.initServer(options); 
  }

  initServer(options: ServerOptions) {
    this.server = new TaskServer(this.taskModels, options.port);
    const registerPromise = [];
    for (let key in this.taskModels) {
      this.taskModels[key].path = this.server.path;
      registerPromise.push(this.registerTask(this.taskModels[key]));
    }

    Promise.all(registerPromise).then(() => {
      console.log('register finish!');
    }).catch(e => {
      console.log('register Error', e);
    })
  }

  async registerTask(task: TaskModel) {
    try {
      console.log('send', task.name);
      const res = await axios(`${this.haruhiHost}/register`, {
        method: 'POST',
        data: {
          Name: task.name,
          Depend: task.depends,
          Typename: task.type,
          path: task.path, 
        },
        responseType: 'json',
      });
      if (res.status !== 200) {
        throw new Error(res.statusText);
      }
    } catch(e) {
      throw e;
    }
  }
}

module.exports = HaruhiTaskServer;