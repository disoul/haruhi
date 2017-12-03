# TaskServer
作为一个HTTPServer来包裹一个task  
TaskServer内部完成和haruhi的通信过程同时向开发者暴露下列接口和配置:  

* TaskModel 的定义
  * Name, Type, Depends, Path
  * 额外Options(Cache之类) // TODO

* startTask Callback
  * 用户需要自己定义startTask
    Callback来相应haruhi的开始task请求，如果该任务需要input，input会带在startTask的参数里
* finishTask Callback
  * 通过调用finishTask来告知haruhi完成任务
  * 并将自己的output带在函数参数中（如果有的话
