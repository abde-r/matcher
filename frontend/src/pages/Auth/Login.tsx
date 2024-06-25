import { useState } from "react"
import { IoArrowForwardCircleOutline } from "react-icons/io5";
import { useNavigate } from 'react-router-dom'
// import { useAuth } from "../../components/Auth/Auth"
import validator from "validator";

export const Login = ({ setAuth }: any) => {

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
      const res = await fetch(`http://localhost:8000/api/v1/auth/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
              query:`
                mutation LoginUser($input: LoginUserInput!) {
                  loginUser(input: $input) {
                    username,
                    password,
                  }
                }
              `,
              variables: {
                input: {
                  username: inputData.username,
                  password: inputData.password,
                }
              }
            })
        })
        .then(res => {return res.json()})
        .catch(err => { console.log('Error in login', err); })

        console.log(res.data);
        navigate('/');
    }
  }

  return (
    <div className="flex flex-col bg-[#d3d3d3] items-center justify-center m-10 h-[90vh] w-[75%] mx-auto rounded-lg">
      <h1 className="text-3xl capitalize my-7 font-semibold border-b-4 border-[#714bd2] rounded-sm text-gray-500">Login</h1>
      <p className="text-gray-500">Welcom back! Login to access the matcherX.</p>
      <p className="text-gray-500">Did you <a className="text-[#007bff]" href="/forgot-pass">Forgot password?</a></p>
      <div className="flex flex-col items-center">
        <input  className='p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400' type="text" name="username" placeholder="username" value={inputData.username} onChange={handleInputChange} />
        {validationErrors.username && <span className="loginError">*{ validationErrors.username }</span>}
        <input  className='p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400' type="password" name="password" placeholder="password" value={inputData.password} onChange={handleInputChange} />
        {validationErrors.password && <span className="loginError">*{ validationErrors.password }</span>}
      </div>
      { loginError && <span className="loginError">* { loginError }</span> }
      <a className="flex items-center bg-[#714bd2] px-3 py-2 rounded-sm text-gray-300 text-md font-semibold cursor-pointer uppercase" onClick={_login}><IoArrowForwardCircleOutline className="mr-1 text-xl" />Login</a>
    </div>
  )
}
