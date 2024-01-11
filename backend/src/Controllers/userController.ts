const express = require('express')
const pool = require('../../database/dbConfig.ts')

const getAllUsers = async (req: any, res: any) => {
  try {
    const client = await pool.connect()
    const users = await client.query('SELECT * FROM "User";')
    console.log('Users: ', users.rows)
    res.status(200).send({ users: users.rows })
    client.release()
  }
  catch (err) {
    console.error('Error getting users: ', err)
    res.status(500).send('Internal server error')
  }
}

module.exports = { getAllUsers }
