const passport = require('passport');
const Strategy = require('passport-jwt').Strategy;
const ExtractJwt = require('passport-jwt').ExtractJwt;

module.exports = (app) => {

    const jwtConfig = app.config.jwt;
    const Users = app.datasource.models.users;

    const options = {};

    options.secretOrKey = jwtConfig.secret;
    options.jwtFromRequest = ExtractJwt.fromAuthHeaderAsBearerToken();

    const strategy = new Strategy(options, (payload, done) => {

        Users
            .findOne({where: payload.id })
            .then(user => {

                if(user) {
                    return done(null, {
                        id: user.id,
                        email: user.email
                    });
                }

                return done(null, false);
            })
            .catch(error => done(error, null));

    });

    passport.use(strategy);

    return {
        initialize: () => passport.initialize(),
        authenticate: () => passport.authenticate('jwt', jwtConfig.session)
    };

};