import axios from "axios";
import { useState } from "react";

export const ForgotPass = () => {
    const [inputData, setInputData] = useState({ email: '', password: '' });
    
    // const navigate = useNavigate()
    const sendVerificationReq = async () => {
      try {
        // console.log('input data: ', inputData)
        const res = await axios.post(`http://localhost:8080/auth/forgot-pass`, {
          'email': inputData.email,
        }, { withCredentials: true })
        console.log('res', res.data.user[0])
        // auth.login(res.data.user[0])
        // navigate('/', { replace: true })
      }
      catch (err) {
        // setLoginError(true);
        console.error('errrror: ', err)
      }
    }
  
    return (
        <div className="Login">
            <input type="email" placeholder="email" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
            <button onClick={sendVerificationReq}>send</button>
        </div>
    )
}
