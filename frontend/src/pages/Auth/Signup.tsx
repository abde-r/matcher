// import axios from "axios"
// import { useState } from "react"
// import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
// import { faCircleRight } from "@fortawesome/free-regular-svg-icons"
// import validator from 'validator';


// export const Signup = () => {

//   const [inputData, setInputData] = useState({ username: '', email: '', password: '', confirmpassword: '' })
//   // const [validationError, setValidationError] = useState<any>([]);
//   const [invalidEmail, setInvalidEmail] = useState<string>('');
//   const [invalidUsername, setInvalidUsername] = useState<string>('');
//   const [invalidPassword, setInvalidPassword] = useState<string>('');
//   const _signup = async () => {
//     // try {
//       console.log('input data: ', inputData)
//       const { username, email, password, confirmpassword} = inputData;
//       if (!username.toString().trim().length)
//         setInvalidUsername('username is required!');
//       if (!validator.isEmail(email))
//         setInvalidEmail('invalid email form!');
//       if ()
//         // console.log('waaa hamid!')

//       // // check data first
//     //   const res = await axios.post(`http://localhost:8080/auth/signup`, {
//     //     first_name: inputData.firstName,
//     //     last_name: inputData.lastName,
//     //     username: inputData.username,
//     //     gender: inputData.gender,
//     //     email: inputData.email,
//     //     password: inputData.password
//     //   }, { withCredentials: true })
//     //   console.log('res', res.data)
//     // }
//     // catch (err) {
//     //   console.error('errrror: ', err)
//     // }
//   }

//   // const handleGenderChange = (e: any) => {
//   //   const x = e.target.value === 'Male' ? 1 : 0
//   //   setInputData({ ...inputData, gender: x})
//   // }
//   console.log(validationError)

//   return (
//     <div className="Signup">
//       <div className="container">
//         <h2>Signup</h2>
//         <p>Welcom to matcherX! Pleas enter your infos to get access.</p>
//         <p>You have an <a className="goLogin" href="/login">account?</a></p>
//         <div className="SignupForms">
//           <input type="email" placeholder="email" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
//           <input type="text" placeholder="username" onChange={(e) => { setInputData({ ...inputData, username: e.target.value }) }} />
//           <input type="password" placeholder="password" onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
//           <input type="password" placeholder="confirm password" style={validationError.confirmPassword && { border: '2px solid red'}} onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
//         </div>
//         <a className="proceed" /*href='/proceed-signup'*/ onClick={_signup}><FontAwesomeIcon icon={faCircleRight} /> Proceed</a>
//       </div>
//     </div>
//   )
// }

import { useState } from 'react';
import validator from 'validator';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleRight } from '@fortawesome/free-solid-svg-icons';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import './Signup.scss'

export const Signup = ({ setAuth }: any) => {

  const [inputData, setInputData] = useState<any>({
    // firstName: '',
    // lastName: '',
    username: '',
    // gender: 1,
    email: '',
    // preferences: [],
    // interests: [],
    // currentInterest: '',
    password: '',
    confirmPassword: ''
  })
  // const [proceed, setProceed] = useState<boolean>(false);

  // const [inputData, setInputData] = useState({
  //   email: '',
  //   username: '',
  //   password: '',
  //   confirmPassword: ''
  // });
  const navigate = useNavigate()

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

    if (inputData.email === 'gg') {
      errors.username = 'Email Already in use';
    }

    if (validator.isEmpty(inputData.username)) {
      errors.username = 'Username is required';
    }

    if (inputData.username === 'gg') {
      errors.username = 'Username Already in use';
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

  // const handleGenderChange = (e: any) => {
  //   const x = e.target.value === 'Male' ? 1 : 0
  //   setInputData({ ...inputData, gender: x})
  // }

  // const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
  //   const selectedValue = event.target.value;
  //   if (!inputData.preferences.includes(selectedValue)) {
  //     setInputData((prevState: any) => ({
  //       ...prevState,
  //       preferences: [...prevState.preferences, selectedValue]
  //     }));
  //   }
  // };

  // const handleInterestsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  //   setInputData((prevState: any) => ({
  //     ...prevState,
  //     currentInterest: e.target.value
  //   }));
  // };

  // const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
  //   if (e.key === 'Enter') {
  //     const newInterest = inputData.currentInterest.trim();
  //     if (newInterest) {
  //       setInputData((prevState: any) => ({
  //         ...prevState,
  //         interests: [...prevState.interests, newInterest],
  //         currentInterest: ''
  //       }));
  //     }
  //   }
  // };

  // const handleKeyDown = (e) => {
  //   if (e.key === 'Enter') {
  //     const newInterests = [...inputData.interests, inputData.password];
  //     setInputData({ ...inputData, interests: newInterests, password: '' });
  //   }
  // };
  
  // const _proceed = () => {
  //   if (validateInputs())
  //     setProceed(true);
  // }

  const _signup = async () => {
    if (validateInputs()) {
      try {
        await axios.post(`http://localhost:8080/auth/signup`, {
          username: inputData.username,
          email: inputData.email,
          password: inputData.password
        }, { withCredentials: true })
        setAuth({ token: true });
        navigate('/')
      }
      catch (err) {
        console.error(err)
      }
    }
  };

  console.log('wewww', inputData)

  return (
    <div className="Signup">
      <div className="container">
            <h2>Signup</h2>
            <p>Welcome to matcherX! Please enter your informations to get access.</p>
            <p>You have an <a className="goLogin" href="/login">account?</a></p>
            <div className="SignupForms">
              <input type="email" name="email" placeholder="Email" value={inputData.email} style={validationErrors.email && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.email && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.email}</p>}
              <input type="text" name="username" placeholder="Username" value={inputData.username} style={validationErrors.username && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.username && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.username}</p>}
              <input type="password" name="password" placeholder="Password" value={inputData.password} style={validationErrors.password && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.password && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.password}</p>}
              <input type="password" name="confirmPassword" placeholder="Confirm Password" value={inputData.confirmPassword} style={validationErrors.confirmPassword && { border: '2px solid red', borderRadius: '10px'}} onChange={handleInputChange} />
              {validationErrors.confirmPassword && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.confirmPassword}</p>}
            </div>
            <a className="proceed" onClick={_signup}><FontAwesomeIcon className='submit-circle' icon={faCircleRight} />Signup</a>
            {/* <a className="proceed" onClick={_signup}><FontAwesomeIcon icon={faCircleRight} /> Proceed</a> */}
      </div>
    </div>
  );
};

