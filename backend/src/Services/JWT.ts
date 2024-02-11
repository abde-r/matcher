const { sign, verify } = require('jsonwebtoken')
const jwt = require('jsonwebtoken')

// const createToken = (user: any) => {

//     return sign({ id: user.id, username: user.username }, process.env.ACCESS_TOKEN)

// }

const createToken = (user: any) => {
    // Assuming user is an object containing user information
    const payload = {
      userId: user.id,
      email: user.email,
      // Add any other relevant user data to the payload
    };
  
    // Sign the payload with a secret key to create the token
    const secretKey =process.env.ACCESS_TOKEN; // Replace with your actual secret key
    const options = { expiresIn: '1h' }; // Set the expiration time for the token
  
    console.log('jwt parameters: ', payload, secretKey, options)
    // Generate the JWT
    const token = jwt.sign(payload, secretKey, options);
  
    return token;
  };

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