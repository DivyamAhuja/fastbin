import type { NextApiRequest, NextApiResponse } from 'next'
import faunadb, { Collection, Get, Ref, Time } from 'faunadb'

type Data = {
  code: string
}

type FaunaQueryResponse = {
  ref?: typeof Ref
  ts?: typeof Time
  data?: Data
}

const client = new faunadb.Client({
  secret: process.env.FAUNA_ADMIN_KEY || "",
  domain: 'db.us.fauna.com', 
  port: 443,
  scheme: 'https'
})

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  const id = req.query['id']
  
  client.query<FaunaQueryResponse>(
    Get(Ref(Collection('data'), id))
  )
  .then((ret) => res.status(200).json({code: ret?.data?.code || ""}))
  .catch((error) => {
    res.status(404).send({code: ""})
  })
  
  
  /*  // MongoDB
  let data = JSON.stringify({
    collection: "data",
    database: "fastbin",
    dataSource: "Cluster0",
    filter: {
      "_id": { "$oid": id}
    },
  })
  
  fetch("https://data.mongodb-api.com/app/data-gizgg/endpoint/data/beta/action/findOne", {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Request-Headers": "*",
      "api-key": process.env.MONGO_API_KEY || "",
    },
    body: data
  })
  .then((response) => response.json())
  .then((data) => {
    res.status(200).json({code: data.document.code})
  })
  .catch((error) => {
    res.status(404).send({code: ""})
  })
  */
}
