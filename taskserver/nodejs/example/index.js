const HaruhiTaskServer = require('../build/index');

// 一个简单的任务，打印输入，输出=输入
const taskFunction = (input) => {
  console.log('start Task');
  console.log(input);
  return input;
}

const taskModels = {
  exampleTask: {
    name: 'exampleTask',
    depends: [],
    type: 'example',
    hooks: {
      startTask: (input) => {
        return new Promise((resolve) => {
          taskFunction(input);
          resolve(0);
        })
      }
    }
  }
};

const taskServer = new HaruhiTaskServer(taskModels, {});