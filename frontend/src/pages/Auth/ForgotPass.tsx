// import axios from "axios";
import { useState } from "react";
import { IoMailOutline } from "react-icons/io5";

export const ForgotPass = () => {
    const [inputData, setInputData] = useState({ email: '', password: '' });
    
    // const navigate = useNavigate()
    const sendVerificationReq = async () => {
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
          <p className="text-gray-500">You will revceive a verification link in your email. Make sure the given email is correct</p>
          <input className="p-2 m-2 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="email" placeholder="Email" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
          {/* <button onClick={sendVerificationReq}>send</button> */}
          <a className="flex items-center bg-[#714bd2] px-3 py-2 my-3 rounded-sm text-gray-300 text-md font-semibold cursor-pointer uppercase" onClick={sendVerificationReq}><span className="mr-1 text-xl"><IoMailOutline /></span>send</a>
        </div>
    )
}
