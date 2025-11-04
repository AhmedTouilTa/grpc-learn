"use client";

import { useEffect, useState } from 'react';
const proto = require('./../simple_grpc_web_pb.js');

function yo()
{
    var helloSayer = new proto.HelloSayerClient(
    'http://localhost:8080');
    
      var request = new proto.HelloMessage();
      request.setGreetingmessage("lolz");
      var metadata = {'custom-header-1': 'value1'};
      helloSayer.sayHello(request, metadata, function(err: any , response: any) {
      if (err) {
          console.log(err.code);
          console.log(err.message);
      } else {
          console.log(response.getReplymessage());
      }
      });
}

export default function HelloClient() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    yo();
  }, []);

  return <div>{message}</div>;
}