const HttpStatus = require('http-status');
const jwt = require('jsonwebtoken');
const UsersController = require('../controllers/users');

module.exports = (app) => {

    const usersController = new UsersController(app.datasource.models.users);

    app.route('/login')
        .post( async (req, res) => {

            try {

                const response = await usersController.signin(req.body);
                const login = response.login;

                console.log(login);
                
                if(login.id && login.isValid) {

                    const payload = {id: login.id};

                    res.json({
                        token: jwt.sign({data: payload}, app.config.jwt.secret, {expiresIn: '1h'})
                    });

                } else {

                    res.sendStatus(HttpStatus.UNAUTHORIZED);
                }


            } catch (e) {

                console.error(e);
                res.sendStatus(HttpStatus.UNAUTHORIZED);

            }

        });

};