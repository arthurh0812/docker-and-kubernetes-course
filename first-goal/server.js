import express from "express";
import bodyParser from "body-parser";

const app = express();

let defaultGoal = "Learn Docker!";

let goal = defaultGoal;

// parse url-encoded data from the body and set its corresponding data property on the request body
app.use(
  bodyParser.urlencoded({
    extended: false,
  })
);

// serve static files from the relative public directory
app.use(express.static("public"));

// homepage
app.get("/", (req, res) => {
  res.status(200);
  res.header("Content-Type", "text/html");
  res.send(`
        <html>
            <head>
                <link rel="stylesheet" href="styles.css">
            </head>
            <body>
                <section>
                    <h2>My Course Goal:</h2>
                    <h3>${goal}</h3>
                </section>
                <form action="/store-goal" method="POST">
                    <div class="form-control">
                        <label>Course Goal</label>
                        <input type="text" name="goal" value="${defaultGoal}">
                    </div>
                    <button>Set Course Goal</button> 
                </form>
            </body>
        </html>
    `);
});

// accepts goal via url-encoded POST data
app.post("/store-goal", (req, res) => {
  const inputGoal = req.body.goal;
  // log the input goal
  console.log(inputGoal);
  // overwrite global variable "goal"
  goal = inputGoal;
  res.redirect("/");
});

app.listen(80);
