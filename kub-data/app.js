const path = require('path');
const fs = require('fs');

const express = require('express');
const bodyParser = require('body-parser');

const app = express();

const filePath = path.join(__dirname, process.env.STORY_DIR, 'text.txt');

app.use(bodyParser.json());

app.get("/", (req, res) => {
  res.status(200).json({
    message: "This worked! Two pods are now communicating with each other"
  })
})

app.get('/story', (req, res) => {
  fs.readFile(filePath, (err, data) => {
    if (err) {
      return res.status(500).json({ message: 'Failed to open file.' });
    }
    res.status(200).json({ story: data.toString() });
  });
});

app.post('/story', (req, res) => {
  const newText = req.body.text;
  if (newText.trim().length === 0) {
    return res.status(422).json({ message: 'Text must not be empty!' });
  }
  fs.appendFile(filePath, newText + '\n', (err) => {
    if (err) {
      return res.status(500).json({ message: 'Storing the text failed.' });
    }
    res.status(201).json({ message: 'Text was stored!' });
  });
});

app.get("/error", (req, res) => {
  process.exit(1)
})

console.log("This is a sample application and it is amazing to use Kubernetes!")

console.log("With these log messages I am currently practicing typing and learning to type as fast and " +
    "as accurate as possible.")

app.listen(3000);
