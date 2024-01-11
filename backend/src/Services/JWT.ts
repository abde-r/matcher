const { sign, verify } = require('jsonwebtoken')

const createToken = (user: any) => {

    return sign({ id: user.id, username: user.username }, process.env.ACCESS_TOKEN)

}

const validateToken = (req: any, res: any, next: any) => {

    console.log('cookies: ', req)
    const accessToken = req.cookies['access-token']

    if (!accessToken)
        return res.status(400).json({ error: 'User not Authenticated!' })

    try {
        const validToken = verify(accessToken, process.env.ACCESS_TOKEN)
        if (validToken) {
            req.Authenticated = true
            return next()
        }
    }
    catch (err) {
        return res.status(400).json({ error: err })
    }

}

module.exports = { createToken, validateToken }