const express = require("express");

const app = express();
app.use(express.json())

const lowPrivilege = {
    role : "user",
    permission: "readOnly"
}

let employer = {
    username: "John",
    age: 25
}
employer.__proto__ = lowPrivilege;
console.log("Inhered role is: " + employer.role);

const merge = (target, source) => {
  for (const key in source ) {
    if (typeof (target[key]) == "object" && typeof (source[key]) === "object") {
      // target.__proto__.role = source.__proto__.role
      // target.__proto__.role = user
      // source__proto__.role = admin
      merge(target[key], source[key]);
    } else {
      target[key] = source[key];
    }
  }
  return target;
};

app.post("/", async (req, res) => {
  const reqBody = req.body;
  employer = merge( employer, reqBody);

  console.log("After merge, role is: " + employer.role);

  if (employer.username === "John" && employer.role === "admin"){
    res.send("You are ADMIN");
  }
  else {
    res.send("You are EMPLOYER");
  } 
  
});

app.listen(1337, () => {
  console.log(`Listening on http://localhost:1337`);
});
