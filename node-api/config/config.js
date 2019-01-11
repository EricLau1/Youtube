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
            user: '162eb270ce1dfb',
            pass: 'd73ab9d615c5f8'
        }
    },
    jwt: {
        secret: 't0p-S3cr3t',
        session: {session: false}
    }
};