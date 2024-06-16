const express = require('express')
const pool = require('../../database/dbConfig')
const bcrypt = require('bcrypt')
//const { check, validationResult } = require('express-validator')


const login = async (req: any, res: any) => {

  try {
    const client = await pool.connect();
    const result = await client.query('SELECT * FROM user')
    res.json(result.rows);
    client.release()
  }
  catch (err) {
    console.error('Error executing query', err);
    res.status(500).send('Internal Server Error');
  }
};

const users = [
  {
    email: "stronk@test.ma",
    password: "pass1234",
  }
]

const signup = async (req: any, res: any) => {

  const { email, password } = req.body
  //const error = validationResult(req)

  //if (!error.isEmpty())
  // return res(400, { error: error.array() })

  // to implement validation server side later :)

  const user = users.find((u) => { return u.email === email })

  if (user)
    return res.status(400).json({ 'error': `user ${email} already exist!` })

  const hashedPass = await bcrypt.hash(password, 10)
  console.log(email, password, hashedPass)
  users.push({ email: email, password: hashedPass })
  const client = await pool.connect()
  const result = await client.query('INSERT INTO user VALUES (email, password) (?, ?)', [email, hashedPass])
  res.send(result.rows)
  client.release()
}
module.exports = { login, signup }
