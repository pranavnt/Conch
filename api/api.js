const app = require("express")();
const fs = require("fs");
const path = require("path");

app.get("/scripts/:name", function (req, res) {
  let dataArr = JSON.parse(fs.readFileSync(path.resolve(__dirname, "db.json")))[
    "scripts"
  ];

  for (var i = 0; i < dataArr.length; i++) {
    console.log(dataArr[i]["name"]);
    if (dataArr[i]["name"] == req.params.name) {
      res.json({ script: dataArr[i]["script"] });
      return;
    }
  }
  res.json({ script: "echo Script not found" });
});

app.post("/upload", function (req, res) {});

app.listen(3000, () => {
  console.log("");
});
