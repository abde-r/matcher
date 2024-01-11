import axios from "axios"
import { useState } from "react"
import { useNavigate } from 'react-router-dom'
import './Login.scss'
import { useAuth } from "../../components/Auth/Auth"

export const Login = () => {

  const [inputData, setInputData] = useState({ email: '', password: '' })
  const auth = useAuth()

  const navigate = useNavigate()
  const _login = async () => {
    try {
      console.log('input data: ', inputData)
      const res = await axios.post(`http://localhost:8080/auth/login`, {
        'email': inputData.email,
        'password': inputData.password
      }, { withCredentials: true })
      console.log('res', res.data.user[0])
      auth.login(res.data.user[0])
      navigate('/', { replace: true })
    }
    catch (err) {
      console.error('errrror: ', err)
    }
  }

  return (
    <div className="Login">
      <input type="email" placeholder="email" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
      <input type="password" placeholder="password" onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
      <button onClick={_login}>login</button>
      <a>forgot password?</a>
    </div>
  )
}
