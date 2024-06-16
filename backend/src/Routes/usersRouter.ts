const router = require('express').Router()

const usersController = require('../Controllers/users')

router.get('/', usersController.getAllUsers)

module.exports = router
