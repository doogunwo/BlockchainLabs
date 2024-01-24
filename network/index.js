const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = '../Blob/protocol/protocol3.proto';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const block = protoDescriptor.block.protocol; // 수정된 부분

const client = new block.HelloService('localhost:50051', grpc.credentials.createInsecure());

client.SayHello({ name: 'ChatGPT' }, function(err, response) {
  if (err) {
    console.error('Error:', err);
  } else {
    console.log('gRPC 응답:', response.message);
  }
});
