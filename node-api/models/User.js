const argon2 = require('argon2');

module.exports = (sequelize, DataType) => {

    const User = sequelize.define('users', {
        id: {
            type: DataType.INTEGER,
            primaryKey: true,
            autoIncrement: true
        },
        name: {
            type: DataType.STRING,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        },
        email: {
            type: DataType.STRING,
            allowNull: false,
            unique: true,
            validate: {
                isEmail: true,
                notEmpty: true
            }         
        },
        password: {
            type: DataType.STRING,
            allowNull: false,
            validate: {
                notEmpty: true
            }
        }
    },
    {
        hooks: {

            async beforeCreate(user) {

                try {
                
                    const hash = await argon2.hash(user.password);
                    user.set('password', hash); 
                
                } catch(e) {

                    console.error(e);

                }
            } // end beforeCreate

        }
    });

    User.verifyPassword = async (hash, password) => {

        try {

            if(await argon2.verify(hash, password)) {

                return true;
            }

        } catch (e) {

            console.error(e);

        }

        return false;
    }

    return User;
};