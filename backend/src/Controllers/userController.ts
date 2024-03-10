const express = require('express')
const pool = require('../../database/dbConfig.ts')
const db = require('../Services/queriesHelper');
const jwt = require('jsonwebtoken')

const getAllUsers = async (req: any, res: any) => {
  try {
    // const client = await pool.connect()
    // const users = await client.query('SELECT * FROM "User";')
    const users = await db.SELECT({ table: 'User', columns: [] });
    console.log('Users: ', users.rows)
    res.status(200).send({ usssers: users.rows });
  }
  catch (err) {
    // console.error('Error getting users: ', err)
    res.status(500).send('Internal server error')
  }
}

const me = async (req: any, res: any) => {
  try {
    const { access_token } = req.body;
    const client = await pool.connect();
    
    console.log('access token: ', access_token);
    const user = await client.query('SELECT * FROM "User" WHERE access_token=$1;', [access_token]);

    if (!user.rows) {
      return res.status(404).send({message: 'User not found!'});
    }
    // res.cookie('testcookie', 'mzyana', {httpOnly: true})
    console.log('req', user.rows[0].access_token, req.cookies['access-token'])
    // const users = await db.SELECT({ table: 'User', columns: [] });
    // console.log('Users: ', users.rows)
    return res.status(200).send({ usssers: user.rows[0] });
  }
  catch (err) {
    // console.error('Error getting users: ', err)
    res.status(500).send('Internal server error')
  }
}

const createUsers = async (req: any, res: any) => {
  try {
    console.log('wew wew')
  // const { username} = req.body;
    // const client = await pool.connect()
    // const users = await client.query('SELECT * FROM "User";')
    const users = await db.CREATE({ table: 'Test', columns: [{name: 'username', type: 'VARCHAR(150)', default: 'NOT NULL'}] });
    console.log('Users: ', users)
    const gg = await db.INSERT({ table: 'Test', columns: [{name: 'username', value: 'jlmoud'}] });
    const u = await db.SELECT({ table: 'Test', columns: [] })
    console.log(u.rows)
    res.status(200).send({ users: u.rows });
  }
  catch (err) {
    // console.error('Error getting users: ', err)
    res.status(500).send('Internal server error')
  }
}

const verifyToken = async (req: any, res: any, next: any) => {
  try {
    const access_token: string = req.cookies['access-token'];
    console.log(access_token)
    if (!access_token)
      return res.json({ status: false, message: 'No Token!'});
    const verification = await jwt.verify(access_token, process.env.ACCESS_TOKEN);
    next();
  }
  catch (err) {
    return res.json(err);
  }
}

module.exports = { getAllUsers, createUsers, me, verifyToken };
