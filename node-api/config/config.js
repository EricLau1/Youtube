module.exports = {
    database: "test",
    username: "postgres",
    password: "@root",
    params: {
        dialect: "postgres",
        define: {
            underscored: false
        }
    },
    email: {
        host: 'smtp.mailtrap.io',
        port: 2525,
        auth: {
            user: 'SEU USERNAME MAILTRAP',
            pass: 'SEU PASSWORD MAILTRAP'
        }
    },
    jwt: {
        secret: 't0p-S3cr3t',
        session: {session: false}
    }
};
