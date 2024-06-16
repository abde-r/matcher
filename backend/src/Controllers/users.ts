const express = require('express')
const pool = require('../../database/dbConfig')

const getAllUsers = async (req: any, res: any) => {
  try {
    const client = await pool.connect()
    const users = await client.query('SELECT * FROM user')
    res.status(200).send({ "users": users.rows })
    client.release()
  }
  catch (err) {
    console.error('Error getting users: ', err)
    res.status(500).send('Internal server error')
  }
}

module.exports = { getAllUsers }
