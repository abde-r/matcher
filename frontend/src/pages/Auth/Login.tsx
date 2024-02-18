import axios from "axios"
import { useState } from "react"
import { useNavigate } from 'react-router-dom'
// import { useAuth } from "../../components/Auth/Auth"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCircleRight } from "@fortawesome/free-regular-svg-icons"
import validator from "validator";
import './Login.scss'

export const Login = () => {

  const [inputData, setInputData] = useState({ username: '', password: '' });
  const [validationErrors, setValidationErrors] = useState<any>({});
  const [loginError, setLoginError] = useState<string>('');
  // const auth: any = useAuth()

  const navigate = useNavigate()

  const validateInputs = () => {
    const errors: any = {};

    if (validator.isEmpty(inputData.username)) {
      errors.username = 'Username is required';
    }

    if (validator.isEmpty(inputData.password)) {
      errors.password = 'Password is required';
    }

    setValidationErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const handleInputChange = (e: any) => {
    const { name, value } = e.target;
    setInputData({ ...inputData, [name]: value });
  };
  
  const _login = async () => {
    if (validateInputs()) {
      try {
        // console.log('input data: ', inputData)
        await axios.post(`http://localhost:8080/auth/login`, {
          'username': inputData.username,
          'password': inputData.password,
        }, { withCredentials: true })
        // console.log('res', res)
        // auth.login(res.data.user[0])
        navigate('/', { replace: true })
      }
      catch (err: any) {
        setLoginError(err.response.data.error);
      }
    }
  }

  return (
    <div className="Login">
      <div className="container">
        <h2>Login</h2>
        <p>Welcom back! Login to access the matcherX.</p>
        <p>Did you <a href="/forgot-pass">Forgot password?</a></p>
        <div className="loginForms">
          <input type="text" name="username" placeholder="Username" value={inputData.username} style={validationErrors.username && { border: '1px solid rgb(215, 31, 31)', borderRadius: '10px'}} onChange={handleInputChange} />
          {validationErrors.username && <span className="loginError">*{ validationErrors.username }</span>}
          <input type="password" name="password" placeholder="Password" value={inputData.password} style={validationErrors.password && { border: '1px solid rgb(215, 31, 31)', borderRadius: '10px'}} onChange={handleInputChange} />
          {validationErrors.password && <span className="loginError">*{ validationErrors.password }</span>}
        </div>
        { loginError && <span className="loginError">* { loginError }</span> }
        <button onClick={_login}><FontAwesomeIcon icon={faCircleRight} /> Continue</button>
      </div>
    </div>
  )
}
