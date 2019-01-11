const Sequelize = require('sequelize');
const fs = require('fs');
const path = require('path');

var database = null;

const loadModels = (sequelize) => {

    const dir = path.join(__dirname, '../models');
    const models = [];

    fs.readdirSync(dir).forEach(file => {

        const modelPath = path.join(dir, file);
        const model = sequelize.import(modelPath);
        models[model.name] = model;

    });

    return models;
};

module.exports = (app) => {

    if(!database) {

        const config = app.config;

        const sequelize = new Sequelize(
            config.database,
            config.username,
            config.password,
            config.params
        );

        database = {
            sequelize,
            Sequelize,
            models: {}
        };

        database.models = loadModels(sequelize);

        sequelize.sync().done(() => database);

    }

    return database;

};