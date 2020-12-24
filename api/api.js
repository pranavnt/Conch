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

app.post("/upload/", function (req, res) {
  let dataArr = JSON.parse(fs.readFileSync(path.resolve(__dirname, "db.json")))[
    "scripts"
  ];
  console.log(req.query);
  for (var i = 0; i < dataArr.length; i++) {
    //console.log(dataArr[i]["name"]);
    if (dataArr[i]["name"] == req.query.name) {
      res.json({
        response: `Script with ${req.query.name} already exists :(`,
      });
      return;
    }
  }

  const newScriptJSON = {
    name: req.query.name,
    script: req.query.script,
  };

  dataArr.push(newScriptJSON);

  const newJSON = {
    scripts: dataArr,
  };

  fs.writeFile("db.json", JSON.stringify(newJSON), (err) => {
    if (err) {
      res.json({
        response:
          "Error (on our end) :( -- please email pranavnt@outlook.com for support",
      });
    } else {
      res.json({
        response: "Script uploaded!!",
      });
    }
  });
});

app.listen(3000, () => {
  console.log("");
});
