import { BlockChain } from '@src/core';
import { P2PServer } from '@src/serve/p2p';

import express, { Request, Response } from 'express';

const app = express();
const bc = new BlockChain();
const ws = new P2PServer();

app.use(express.json());

app.post('/addToPeer', (req, res) => {
  const { peer } = req.body;
  ws.connectToPeer(peer);
});

app.listen(3000, () => {
  console.log('서버 시작 \n PORT : #3000');
  ws.listen();
});