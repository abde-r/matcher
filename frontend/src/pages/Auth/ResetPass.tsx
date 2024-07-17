import axios from "axios"
import { useState } from "react"
import { IoArrowForwardCircleOutline } from "react-icons/io5"
import { useNavigate } from "react-router-dom"
import validator from "validator"

export const ResetPass = () => {
  
  const navigate = useNavigate();
  const [inputData, setInputData] = useState({ password: '', confirmPassword: '' });
  const [validationErrors, setValidationErrors] = useState<any>({});


  const validateInputs = () => {
    const errors: any = {};
    if (validator.isEmpty(inputData.password))
      errors.password = 'Password is required';
    if (!validator.isStrongPassword(inputData.password))
      errors.password = 'Weak Password, try something with lowecase, uppercase, numbers and symbols';

    if (validator.isEmpty(inputData.confirmPassword))
      errors.confirmPassword = 'Confirm password is required';
    else if (inputData.confirmPassword !== inputData.password)
      errors.confirmPassword = 'Passwords do not match';
    

    setValidationErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const _resetPass = async () => {
    if (validateInputs()) {
      const res = await fetch(`http://localhost:8000/api/v1/auth/reset-pass`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
              query:`
                mutation ResetUserPass($input: ResetUserPassInput!) {
                  resetUserPass(input: $input) {
                    password,
                    access_token,
                  }
                }
              `,
              variables: {
                input: {
                  password: inputData.password,
                  token: 'cookiza',
                }
              }
            })
        })
        .then(res => {return res.json()})
        .catch(err => { console.log(`Error reseting password`, err); })

        console.log(res.data);
        navigate('/');
    }
    // try {
    //   console.log('input data: ', inputData)

    //   // // check data first
    //   const res = await axios.put(`http://localhost:8080/auth/reset-pass`, {
    //     email: inputData.email,
    //     password: inputData.password
    //   }, { withCredentials: true })
    //   console.log('res', res.data)
    // }
    // catch (err) {
    //   console.error('errrror: ', err)
    // }
  }


  return (
    // <div className="Signup">
    <div className="flex flex-col h-[90vh] bg-[#d3d3d3] items-center justify-center p-4 m-10 w-[75%] mx-auto rounded-md">
      <h1 className="text-3xl capitalize my-7 font-semibold border-b-4 border-[#714bd2] rounded-sm text-gray-500">Reset password</h1>
      <h2 className="text-gray-500 font-semibold">had lmra hahya dwznaha lik! mra jaya dber rask</h2>
      {/* <div className="flex flex-col"> */}
        <input className="p-2 m-2 my-5 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="password" placeholder="password" onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
        <input className="p-2 m-2 my-5 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="password" placeholder="confirm password" onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
        <a className="flex items-center my-2 bg-[#714bd2] px-3 py-2 rounded-sm text-gray-300 text-md font-semibold cursor-pointer uppercase" onClick={_resetPass}><IoArrowForwardCircleOutline className="mr-1 text-xl" />reset</a>
      
      {/* </div> */}

      {/* <div className="submit">
        <button onClick={_resetPass}>reset</button>
      </div> */}
    </div>
  )
}
