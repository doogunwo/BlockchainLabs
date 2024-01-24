const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = '../Blob/protocol/protocol3.proto';

// protoLoader를 사용하여 .proto 파일 로드
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
});
//protoc --proto_path=../Blob/protocol/ --js_out=import_style=commonjs,binary:./ ../Blob/protocol/protocol3.proto


const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const block = protoDescriptor.block;

// gRPC 클라이언트 생성
const client = new block.HelloService('localhost:50051', grpc.credentials.createInsecure());

// 요청 보내기
client.SayHello({ name: 'ChatGPT' }, function(err, response) {
  if (err) {
    console.error('Error:', err);
  } else {
    console.log('gRPC 응답:', response.message);
  }
});
