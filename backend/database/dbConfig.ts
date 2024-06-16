const { Pool } = require('pg');
require('dotenv').config()

const pool = new Pool({
    user: process.env.DB_USER,//'stronk',
    host: process.env.DB_HOST,//'postgresdb',
    database: process.env.DB_NAME,//'matchadb',
    password: process.env.DB_PASSWORD,//'pass',
    port: process.env.DB_PORT,//5432,
});

module.exports = pool;
