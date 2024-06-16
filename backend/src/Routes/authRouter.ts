const router = require('express').Router()
const authController = require('../Controllers/authController')

router.get('/', authController.login)
router.post('/signup', authController.signup)

module.exports = router
