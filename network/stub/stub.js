const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const PROTO_PATH = './blockchain.proto';

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
);

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const blockchainService = protoDescriptor.blockchain.BlockchainService;

const client = new blockchainService('localhost:50051', grpc.credentials.createInsecure());

// 블록 추가 요청 예시
client.AddBlock({ data: 'Block data' }, (error, response) => {
    if (!error) {
        console.log('Block added:', response.status);
    } else {
        console.error('Error:', error);
    }
});
