const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");

const config = require("./config/config");
const datasource = require("./config/datasource");
const Email = require('./utils/email');

const indexRouter =  require("./routes/index");
const usersRouter = require("./routes/users");
const authRouter = require('./routes/auth');

const authorization = require('./auth');

const app = express();
app.use(cors());

const port = 3000;
app.set('port', port);

app.config = config;
app.datasource = datasource(app);

app.email = new Email(app.config);

app.use(bodyParser.json({
    limit:'5mb'
}));

const auth = authorization(app);
app.use(auth.initialize());
app.auth = auth;

indexRouter(app);
usersRouter(app);
authRouter(app);

module.exports = app;
