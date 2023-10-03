import type { NextApiRequest, NextApiResponse } from "next";
import faunadb from "faunadb";
import client from "../../lib/fauna-client";

const q = faunadb.query;

type Data = {
  id: string;
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method !== "POST") res.status(404).send({ id: "" });

  const code = req.body.data;

  client
    .query<any>(q.Create(q.Collection("data"), { data: { code: code } }))
    .then((response) => {
      res.status(200).json({ id: response?.ref?.id });
    })
    .catch((_error) => {
      res.status(404).send({ id: "" });
    });
}
