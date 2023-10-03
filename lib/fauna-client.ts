import faunadb from "faunadb";

const client = new faunadb.Client({
  secret: process.env.FAUNA_ADMIN_KEY || "",
  domain: "db.fauna.com",
  port: 443,
  scheme: "https",
});

export default client;
