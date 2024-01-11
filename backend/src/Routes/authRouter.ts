const router = require('express').Router()
const authController = require('../Controllers/authController.ts')
const { validateToken } = require('../Services/JWT')

router.post('/login', authController.login)
// router.post('/login', validateToken, authController.login)
router.post('/signup', authController.signup)

module.exports = router
