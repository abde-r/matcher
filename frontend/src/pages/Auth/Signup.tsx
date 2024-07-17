import { useState } from 'react';
import validator from 'validator';
import { useNavigate } from 'react-router-dom';
import './Signup.scss'
import { IoArrowForwardCircleOutline } from 'react-icons/io5';

export const Signup = ({ setAuth }: any) => {

  const navigate = useNavigate();
  const [inputData, setInputData] = useState<any>({
    username: '',
    email: '',
    password: '',
    confirmPassword: ''
  });
  const [validationErrors, setValidationErrors] = useState<any>({});
  


  const handleInputChange = (e: any) => {
    const { name, value } = e.target;
    setInputData({ ...inputData, [name]: value });
  };

  const validateInputs = () => {
    const errors: any = {};

    if (!validator.isEmail(inputData.email)) {
      errors.email = 'Invalid email format';
    }

    if (validator.isEmpty(inputData.username)) {
      errors.username = 'Username is required';
    }

    if (validator.isEmpty(inputData.password)) {
      errors.password = 'Password is required';
    }

    if (!validator.isStrongPassword(inputData.password)) {
      // minLength: 8, minLowercase: 1, minUppercase: 1, minNumbers: 1, minSymbols: 1}
      errors.password = 'Weak Password, try something with lowecase, uppercase, numbers and symbols';
    }

    if (validator.isEmpty(inputData.confirmPassword)) {
      errors.confirmPassword = 'Confirm password is required';
    } else if (inputData.confirmPassword !== inputData.password) {
      errors.confirmPassword = 'Passwords do not match';
    }

    setValidationErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const _signup = async () => {
    if (validateInputs()) {
      const res = await fetch(`http://localhost:8000/api/v1/auth/register`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          credentials: 'include',
          body: JSON.stringify({
            query:`
              mutation RegisterUser($input: RegisterUserInput!) {
                registerUser(input: $input) {
                  username,
                  email,
                  password,
                }
              }
            `,
            variables: {
              input: {
                username: inputData.username,
                email: inputData.email,
                password: inputData.password,
              }
            }
          }),
      })
      .then(res => { return res.json(); })
      .catch(error => { console.log('Error registring user', error); });
      console.log('res', res.data); // or handle the response data

      setAuth({ token: true });
      navigate('/account-verification');
    }
  };

  return (
    <div className="flex flex-col bg-[#d3d3d3] items-center justify-center m-10 h-[90vh] w-[75%] mx-auto rounded-lg">
      {/* <div className="container"> */}
            <h1 className="text-3xl capitalize my-7 font-semibold border-b-4 border-[#714bd2] rounded-sm text-gray-500">Signup</h1>
            <p className="text-gray-500">Welcome to matcherX! Please enter your informations to get access.</p>
            <p className="text-gray-500">You have an <a className="text-[#007bff] font-semibold" href="/login">account?</a></p>
            <div className="flex flex-col items-center">
              <input className='p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400' type="email" name="email" placeholder="Email" value={inputData.email} style={validationErrors.email && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.email && <p className="text-red-500 font-semibold text-sm">*{validationErrors.email}</p>}
              <input className='p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400' type="text" name="username" placeholder="Username" value={inputData.username} style={validationErrors.username && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.username && <p className="text-red-500 font-semibold text-sm">*{validationErrors.username}</p>}
              <input className='p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400' type="password" name="password" placeholder="Password" value={inputData.password} style={validationErrors.password && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.password && <p className="text-red-500 font-semibold text-sm">*{validationErrors.password}</p>}
              <input className='p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400' type="password" name="confirmPassword" placeholder="Confirm Password" value={inputData.confirmPassword} style={validationErrors.confirmPassword && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.confirmPassword && <p className="text-red-500 font-semibold text-sm">*{validationErrors.confirmPassword}</p>}
            </div>
            <a className="flex items-center bg-[#714bd2] px-3 my-3 py-2 rounded-sm text-gray-300 text-md font-semibold cursor-pointer uppercase" onClick={_signup}><IoArrowForwardCircleOutline className="mr-1 text-xl" />Signup</a>
            {/* <a className="proceed" onClick={_signup}><FontAwesomeIcon icon={faCircleRight} /> Proceed</a> */}
      {/* </div> */}
    </div>
  );
};

