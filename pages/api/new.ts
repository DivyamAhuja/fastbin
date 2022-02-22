import type { NextApiRequest, NextApiResponse } from "next";
import faunadb from 'faunadb'

let q = faunadb.query

type Data = {
  id: string;
};

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
  if (req.method !== 'POST') res.status(404).send({id: ""});
  
  const code = req.body.data;
  
  client.query<any>(
    q.Create(
      q.Collection('data'),
      { data: { 'code': code } }
    )
  ).then((response) => {
    res.status(200).json({id: response?.ref?.id})
  }).catch((error) => {
    res.status(404).send({id: ""})
  })

  /*  // Mongodb
  let data = JSON.stringify({
    collection: "data",
    database: "fastbin",
    dataSource: "Cluster0",
    document: {
      "code": code
    },
  })
  
  fetch("https://data.mongodb-api.com/app/data-gizgg/endpoint/data/beta/action/insertOne", {
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
    res.status(200).json({id: data['insertedId']})
  })
  .catch((error) => {
    res.status(404).send({id: ""})
  })
  */
}
