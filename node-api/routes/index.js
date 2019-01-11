module.exports = (app) => {

    app.route('/')
        .get((req, res) => {

            res.json({
                message: "Hello world!"
            });

        });
}