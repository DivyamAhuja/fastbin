import type { NextApiRequest, NextApiResponse } from "next";
import { Collection, Get, Ref, Time } from "faunadb";
import client from "../../../lib/fauna-client";

type Data = {
  code: string;
};

type FaunaQueryResponse = {
  ref?: typeof Ref;
  ts?: typeof Time;
  data?: Data;
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  const id = req.query["id"];
  getData(id.toString())
    .then((ret) => res.status(200).json({ code: ret?.data?.code || "" }))
    .catch(() => res.status(404).send({ code: "" }));
}

export async function getData(id: string) {
  return client.query<FaunaQueryResponse>(Get(Ref(Collection("data"), id)));
}
