const app = require("express")();
const fs = require("fs");
const path = require("path");

app.get("/scripts/:name", function (req, res) {
  let dataArr = JSON.parse(fs.readFileSync(path.resolve(__dirname, "db.json")))[
    "scripts"
  ];
  res.json(dataArr);
});

app.post("/upload", function (req, res) {});

app.listen(3000, () => {
  console.log("hi");
});
