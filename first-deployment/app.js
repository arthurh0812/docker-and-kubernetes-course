const path = require("path");

const express = require("express");

const app = express();

app.use(express.static("public"));

app.get("/", (req, res) => {
  const filePath = path.join(__dirname, "views", "welcome.html");
  res.sendFile(filePath);
});

app.listen(process.env.PORT, "0.0.0.0", () => {
  console.log(`application listening on port ${process.env.PORT}...`);
});
