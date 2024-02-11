const router = require('express').Router()

const usersController = require('../Controllers/userController')

router.get('/', usersController.getAllUsers);
router.get('/me', usersController.me);
router.get('/verifyToken', usersController.verifyToken, (req: any, res: any) => { return res.json({ status: true, message: 'authorized!' }) });
router.post('/', usersController.createUsers);

module.exports = router
