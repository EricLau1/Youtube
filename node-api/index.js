const app = require("./app");

app.listen(app.get('port'), () => {
    console.log(`\nApi rodando na porta ${app.get('port')}\n`);
});
