const WebSocket = require('ws');
const Chain = require('@src/core/blockchain/chain').Chain;

class P2PServer extends Chain {
  constructor() {
    super();
    this.sockets = [];
  }

  // 서버 시작하는 실행 코드
  listen() {
    const server = new WebSocket.Server({ port: 7545 });
    server.on('connection', (socket) => {
      console.log('websocket connection');
      // 여기에 추가적인 socket 처리 로직을 구현할 수 있습니다.
    });
  }

  // 클라이언트 연결 코드
  connectToPeer(newPeer) {
    const socket = new WebSocket(newPeer);
    // 여기에 소켓 연결 후의 로직을 구현할 수 있습니다.
  }
}

// 해당 클래스의 인스턴스 생성 및 사용 예시
const p2pServer = new P2PServer();
p2pServer.listen();

// 예를 들어, 다른 피어에 연결하려면 아래와 같이 사용합니다.
// p2pServer.connectToPeer('ws://example.com');
