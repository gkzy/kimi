let socket;


let heartCheck = {
    timeout:60000, //60ms
    timeoutObj:null,
    serverTimeoutObj:null,
    reset:function(){
        clearTimeout(this.timeoutObj);
        clearTimeout(this.serverTimeoutObj);
        this.start();
    },
    start:function(){
      let self = this;
      this.timeoutObj = setTimeout(function(){
          let hb ={type:3,request_id:1};
          let json = JSON.stringify(hb);
          let uint8Array = new TextEncoder('utf-8').encode(json);
          let buffArray = uint8Array.buffer;
          socket.send(buffArray);
          self.serverTimeoutObj = setTimeout(function(){
              socket.close();
          },self.timeout)
      },this.timeout)

    },
}

function reconnect(){
    console.log("reconnect...")
    connect();
}

function connect(){
    socket = new WebSocket("ws://192.168.0.101:8002/v1/ws?user_id="+user_id+"&device_id="+device_id+"&token="+token);
    //socket.binaryType = 'arraybuffer';
    socket.onopen = function(){
        heartCheck.start();
        console.log("已经连接");
    };
    socket.onerror = function(){
        console.log("连接发生错误");
        //reconnect();
    };
    socket.onclose = function(){
        console.log("连接已经断开")
        //reconnect();

    };
    socket.onmessage = function(event){
        heartCheck.reset();
        let eventPromise = event.data;
        getMessage(eventPromise).then(message=>{
            console.log(message);
        })
    };
}

function getMessage(eventPromise){
    return new Promise(success=>{
        eventPromise.text().then(e=>{
            let send = null;
            try {
                send = JSON.parse(Base64.decode((JSON.parse(e).data)));
            }catch(e){
                //console.log("error:"+e);
            }
            if(send!=null) {
                let message = send.message;
                success(message);
            }
        });
    });
}



function send(){
   let url = "http://192.168.0.101:8001/send";
   let data={};
   $.ajax({
       url:url,
       data:data,
       type:"GET",
       dataType:"JSON",
       success:function(d){
           console.log(d);
       },
       error:function(d){
           console.log(d);
       }

   })
}

$(function (){
    connect();
});