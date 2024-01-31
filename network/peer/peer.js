import {BlockChain} from "@src/core"
import express, {Reqeust, Response} from 'express'

const app  =express()
const bc = new BlockChain();

app.use(express.json())

app.post("/peer", (req, res)=>{

})

app.listen(3000, ()=> {
    console.log("peer start : \n PORT : #3000")
})
