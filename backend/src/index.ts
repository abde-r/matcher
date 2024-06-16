const express = require('express');
const pool = require('../database/dbConfig')

const app = express();
app.use(express.json())

const PORT = process.env.PORT || 8080;

const authRoutes = require('./Routes/authRouter')
const usersRoutes = require('./Routes/usersRouter')

app.use('/auth', authRoutes)
app.use('/api/users', usersRoutes)
//app.get('/', async (req: any, res: any) => {
//  try {
//    const client = await pool.connect();
//    const result = await client.query('SELECT * FROM user')
//    client.release();

//    res.json(result.rows);
//  }
//  catch (err) {
//    console.error('Error executing query', err);
//    res.status(500).send('Internal Server Error');
//  }
//});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}..`);
});

