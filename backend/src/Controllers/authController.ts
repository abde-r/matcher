const express = require('express')
const pool = require('../../database/dbConfig')
const bcrypt = require('bcrypt')
const cookieParser = require('cookie-parser')
const { createToken } = require('../Services/JWT')
const jwt = require('jsonwebtoken')
const db = require('../Services/queriesHelper');
const nodemailer = require('nodemailer');

//const { createUserTable } = require('../Models/users/User.ts')
//const { check, validationResult } = require('express-validator')


// const login = async (req: any, res: any) => {

//   try {
//     const { email, password } = req.body
//     console.log(email, password)
//     const client = await pool.connect();
//     const user = await client.query('SELECT * FROM "User" WHERE email=$1;', [email])
//     // console.log('user: ', user.rows[0].password)
//     if (user.rows.length > 0) {
//       const isPassValid = await bcrypt.compare(password, user.rows[0].password)
//       console.log('bcrypt result: ', isPassValid)
//       if (isPassValid) {
//         const access_token = createToken(user)
//         res.cookie("access-token", access_token, { maxAge: 60 * 1000 })
//         console.log('cwikzaat', res.cookies)
//         return res.status(200).send({ user: user.rows })
//       }
//     }
//     res.status(401).send({ error: `User ${email} is invalid!` });
//     client.release()
//   }
//   catch (err) {
//     console.error('Error executing query', err);
//     res.status(500).send('Internal Server Error');
//   }
// };

const login = async (req: any, res: any) => {

  try {
    const { username, password } = req.body;
    console.log(username, password);
    const client = await pool.connect();
    const user = await client.query('SELECT * FROM "User" WHERE username=$1;', [username]);
    
    console.log('User', user.rows)
    if (user.rows.length > 0) {
      const isPassValid = await bcrypt.compare(password, user.rows[0].password);
      console.log('bcrypt result: ', isPassValid);

      if (!isPassValid)
        return res.status(401).send({ error: 'Incorrect Password!' });

      // if (isPassValid) {
        // const access_token = createToken(user.rows[0]);
        const access_token: any = jwt.sign({username: username}, process.env.ACCESS_TOKEN, {expiresIn: '10m'});
        res.cookie("access-token", access_token/*, { httpOnly: true }*/, { maxAge: 6 * 10000 });
        console.log('-----------', user.rows[0].access_token, access_token)
        client.query(`UPDATE "User" SET access_token = $1 WHERE username = $2;`, [access_token, username])
        // Print cookies received in the subsequent request
        // console.log('Cookies:', req.cookies);
        return res.status(200).send({ user: user.rows });
      // }
    }
    else
      return res.status(401).send({ error: `Invalid Username!` });
    client.release();
  }
  catch (err) {
    console.error('Error executing query', err);
    return res.status(500).send('Internal Server Error');
  }
};


// const users = [
//   {
//     email: "stronk@test.ma",
//     password: "pass1234",
//   }
// ]


const createUserTable = async (client: any) => {
  try {
    const createTableQuery = `
            CREATE TABLE IF NOT EXISTS "User" (
                id SERIAL PRIMARY KEY,
                username VARCHAR(150) NOT NULL,
                email VARCHAR(150) NOT NULL,
                password VARCHAR(200) NOT NULL,
                first_name VARCHAR(150) NOT NULL,
                last_name VARCHAR(150) NOT NULL,
                avatar VARCHAR(150),
                gender BOOLEAN NOT NULL,
                biography VARCHAR(100),
                access_token VARCHAR(500),
                created_at TIMESTAMPTZ DEFAULT NOW()
            )
    `;

    await client.query(createTableQuery);
    console.log('User table created successfully');
  } catch (err) {
    console.error('Error creating User table:', err);
    throw err;
  }
};

const signup = async (req: any, res: any) => {
  const { email, password, username, first_name, last_name, gender } = req.body;

  const client = await pool.connect();
  
  // Check if the user already exists in your database
  // const emailExists = await client.query('SELECT * FROM "User" WHERE email=$1;', [email])
  // console.log('user: ', emailExists.rows)
  // if (emailExists.rows.length > 0)
  //   return res.status(400).send({ error: `User ${email} already exits!` })

  
  
  // Register
  
  try {
    await createUserTable(client);
    const access_token: any = jwt.sign({username: username}, process.env.ACCESS_TOKEN, {expiresIn: '1m'});
    res.cookie("access-token", access_token, { httpOnly: true }, { maxAge: 6 * 1000 });
    // db.CREATE({ table: 'Test', columns: [{name: 'username', type: 'VARCHAR(150)', default: 'NOT NULL'}] });
    
    const insertUserQuery = `
    INSERT INTO "User" (username, email, password, first_name, last_name, gender, access_token)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    `;
    
    const hashedPass = await bcrypt.hash(password, 10);
    const result = await client.query(insertUserQuery, [
      username,
      email,
      hashedPass,
      first_name,
      last_name,
      gender,
      access_token,
    ]);

    console.log('User created successfully');
    res.send(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error creating user' });
  } finally {
    client.release();
  }
};

const forgotPass = async (req: any, res: any) => {
  const {email} = req.body;
  try {
    // const user = await db.SELECT("User", {columns: [{name: 'email', value: email}]});
    const client = await pool.connect();
    const user = await client.query('SELECT * FROM "User" WHERE email=$1;', [email]);
    if (!user.rows.length)
      return res.json({message: 'Invalid user!'});
    console.log(user.rows[0]);
    
    const token: string = jwt.sign({id: user.rows[0].id}, process.env.ACCESS_TOKEN, {expiresIn: '5m'});
    const transporter = nodemailer.createTransport({
      service: 'gmail',
      auth: {
        user: process.env.MAIL_SENDER,
        pass: process.env.MAIL_SENDER_PASS,
      }
    });

    var mailOptions = {
      from: process.env.MAIL_SENDER,
      to: email,
      subject: 'Reset Password',
      text: `http://localhost:5173/reset-pass/${token}`,
    };

    transporter.sendMail(mailOptions, (error: any, info: any) => {
      if (error) {
        return res.json({ message: 'Error Sending email' });
      } else {
        return res.json({ status: true, message: 'email Sent' });
      }
    });
  }
  catch (err) {
    console.error(err);
  }
}

module.exports = { login, signup, forgotPass }
