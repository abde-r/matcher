const express = require('express');
const pool = require('../database/dbConfig.ts')
const cors = require('cors')
const cookieParser = require("cookie-parser");
require("dotenv").config();
const app = express();

app.use(express.json())

const corsOptions = {
  origin: 'http://localhost:5173',
  credentials: true,
};

app.use(cors(corsOptions));
app.use(cookieParser());

const PORT = process.env.PORT || 8080;

const authRoutes = require('./Routes/authRouter.ts')
const usersRoutes = require('./Routes/usersRouter.ts')

app.use('/auth', authRoutes)
app.use('/api/users', usersRoutes)


app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}..`);
});

