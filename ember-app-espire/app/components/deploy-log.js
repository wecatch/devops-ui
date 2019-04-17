import EmberObject from '@ember/object';
import {set} from '@ember/object';
import Component from '@ember/component';
import { inject as service } from '@ember/service';
import {A} from '@ember/array';
import {computed} from '@ember/object';
import { later,once } from '@ember/runloop';

// import $ from 'jquery';

export default Component.extend({
    /*
   * 1. Inject the websockets service
   */
  websockets: service(),
  socketRef: null,

  didInsertElement() {
    this._super(...arguments);

    let job = this.get('deploy');
    if (job.status == "success" || job.status == "failed"){
      later(this, function(){
        this.get('logContent').pushObject(EmberObject.create({
          message: "",
        }));
      }, 500);
    }else {
      console.log(job);
      /*
        2. The next step you need to do is to create your actual websocket. Calling socketFor
        will retrieve a cached websocket if one exists or in this case it
        will create a new one for us.
      */
      const socket = this.websockets.socketFor(`ws://${window.location.host}/v1/app/deploy/poll/log?deploy_id=`+job.id);

      /*
        3. The next step is to define your event handlers. All event handlers
        are added via the `on` method and take 3 arguments: event name, callback
        function, and the context in which to invoke the callback. All 3 arguments
        are required.
      */
      socket.on('open', this.myOpenHandler, this);
      socket.on('message', this.myMessageHandler, this);
      socket.on('close', this.myCloseHandler, this);

      this.set('socketRef', socket);
    }

  },
  didRender() {
    this._super(...arguments);
    window.location.href = "#"+this.element.getAttribute("id");
    console.log("asdf");
  },
  logContent:null,
  deployContent: null,
  willDestroyElement() {
    this._super(...arguments);

    const socket = this.get('socketRef');

    /*
      4. The final step is to remove all of the listeners you have setup.
    */
    socket && socket.off('open', this.myOpenHandler);
    socket && socket.off('message', this.myMessageHandler);
    socket && socket.off('close', this.myCloseHandler);
  },

  myOpenHandler(event) {
    console.log(`On open event has been called: ${event}`);
  },
  deployLog: null,
  deployPercent: computed('deployLog.deploy_status', function(){
    if (this.get('deploy').status == "success"){
      return 100;
    }

    let deployLog = this.get('deployLog');
    if (!deployLog){
      return 0;
    }
    if (deployLog.deploy_status == "doing"){
      once(this, function(){
        set(this.get('deploy'), 'status', 'doing');
      }, 500);
      return 50;
    }

    if (deployLog.deploy_status == "success"){
      once(this, function(){
        set(this.get('deploy'), 'status', 'success');
      }, 500);
      return 100;
    }

    if (deployLog.deploy_status == "fail"){
      once(this, function(){
        set(this.get('deploy'), 'status', 'fail');
      }, 500);
      return 100;
    }

  }),
  calContentHeight(element){
        let scrollHeight=element.scrollHeight; //div文档总高度
        // let offsetHeight=element.offsetHeight; //获取div窗口显示高度
        // let scrollTop=element.scrollTop;          //获取div卷上去的高度
        // let allheight=scrollHeight-offsetHeight;                //div内容的实际高度
        // let top=(scrollTop/allheight)*100;     //滑动距离占总高度的百分比

        return scrollHeight;
  },
  myMessageHandler(event) {
    try {

      let log = A(JSON.parse(event.data))
      for (var i = 0; i < log.length; i++) {
        this.get('logContent').pushObject(log[i]);
        this.formatStatusHandler(log[i]);
      }
      this.set('deployLog', log[log.length-1]);
      later(this, function(){
        let firstPre = this.$().find(".job-body .body")[0]
        let secondPre = this.$().find(".log-body .body")[0]
        $(firstPre).scrollTop(firstPre.scrollHeight);
        $(secondPre).scrollTop(secondPre.scrollHeight);
      }, 500);
    } catch (error) {
      console.log(error)
    }
  },

  myCloseHandler(event) {
    console.log(`On close event has been called: ${event}`);
  },
  formatStatusHandler(log){
      let content = this.get('deployContent');
      for (var i = 0; i < content.length; i++) {
          let lastLog = content[i];
          if(lastLog.hosts == log.hosts && lastLog.cmd_type == log.cmd_type && lastLog.cmd_status == log.cmd_status){
            return;
          }  
      }

      let message = "主机 " + log.hosts + " ";

      if(log.cmd_status == "doing") {
        message += "正在"
      }

      switch (log.cmd_type) {
        case "update":
          message += "更新代码";
          break;
        case "reload":
          message += "重启服务";
          break;
        case "check":
          message += "检查服务";
          break;
        default:
          // statements_def
          break;
      }

      switch (log.cmd_status) {
        case "success":
          message += "成功"
          break;
        case "fail":
          message += "失败";
          break;
        default:
          // statements_def
          break;
      }

      log.deploy_message = message;
      content.pushObject(log);
      console.log(log)
  },
  actions: {
    sendButtonPressed() {
      this.socketRef.send('Hello Websocket World');
    }
  },
  init(){
    this._super(...arguments);
    this.set('logContent', A());
    this.set('deployContent',A());
  }
});
