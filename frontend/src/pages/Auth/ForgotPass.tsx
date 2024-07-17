// import axios from "axios";
import { useState } from "react";
import { IoMailOutline } from "react-icons/io5";
import validator from "validator";

export const ForgotPass = () => {
    const [inputData, setInputData] = useState<any>({ email: '' });
    const [validationErrors, setValidationErrors] = useState<any>({});
    const [verificationSent, setVerificationSent] = useState<boolean>(false);

    
    // const navigate = useNavigate()
    const validateInputs = () => {
      const errors: any = {};
      if (!validator.isEmail(inputData.email)) {
        errors.email = 'Invalid email format';
      }
      setValidationErrors(errors);
      return Object.keys(errors).length === 0;
    };

    const sendVerificationReq = async () => {
      if (validateInputs()) {
        const res = await fetch(`http://localhost:8000/api/v1/auth/send-verification-email`, {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json',
              },
              credentials: 'include',
              body: JSON.stringify({
                query:`
                  mutation SendEmailVerification($input: SendEmailVerificationInput!) {
                    sendEmailVerification(input: $input) {
                      email,
                    }
                  }
                `,
                variables: {
                  input: {
                    email: inputData.email,
                  }
                }
              })
          })
          .then(res => {setVerificationSent(true); return res.json();})
          .catch(err => { console.log(`Error sending email verification to ${inputData.email}`, err); })
  
          console.log('email', inputData.email)
          console.log('res', res);
      }
      // try {
      //   // console.log('input data: ', inputData)
      //   const res = await axios.post(`http://localhost:8080/auth/forgot-pass`, {
      //     'email': inputData.email,
      //   }, { withCredentials: true })
      //   console.log('res', res.data.user[0])
      //   // auth.login(res.data.user[0])
      //   // navigate('/', { replace: true })
      // }
      // catch (err) {
      //   // setLoginError(true);
      //   console.error('errrror: ', err)
      // }
    }
  
    return (
        <div className="flex flex-col h-[90vh] bg-[#d3d3d3] items-center justify-center p-4 m-10 w-[75%] mx-auto rounded-md">
          <h1 className="text-3xl capitalize my-7 font-semibold border-b-4 border-[#714bd2] rounded-sm text-gray-500">Reset password</h1>
          <p className="text-gray-500 text-center">You will revceive a verification link in your email.<br /> Make sure the given email is correct</p>
          <input className="p-2 m-2 rounded-sm my-4 text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="email" placeholder="Email" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
          {validationErrors.email && <p className="text-red-500 font-semibold text-sm">*{validationErrors.email}</p>}
          <a className="flex items-center bg-[#714bd2] px-3 py-2 my-3 rounded-sm text-gray-300 text-md font-semibold cursor-pointer uppercase" onClick={sendVerificationReq}><span className="mr-1 text-xl"><IoMailOutline /></span>send</a>
          {verificationSent && <p className="text-gray-500">* an email verification sent to <span className="underline">{inputData.email}</span>, <span className="font-bold">check your inbox!</span></p>}
        </div>
    )
}
