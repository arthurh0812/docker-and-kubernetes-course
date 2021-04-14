import express from "express";

import connectToDatabase from "./helpers.mjs";

const app = express();

app.get("/", (req, res) => {
  res.status(200);
  res.header("Content-Type", "text/html");
  res.send("<h2>Hello World!</h1>");
});

await connectToDatabase();

app.listen(3000);
