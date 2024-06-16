import axios from "axios"
import { useState } from "react"

export const ResetPass = () => {
  
  const [inputData, setInputData] = useState({ email: '', password: '' })

  const _resetPass = async () => {
    try {
      console.log('input data: ', inputData)

      // // check data first
      const res = await axios.put(`http://localhost:8080/auth/reset-pass`, {
        email: inputData.email,
        password: inputData.password
      }, { withCredentials: true })
      console.log('res', res.data)
    }
    catch (err) {
      console.error('errrror: ', err)
    }
  }


  return (
    <div className="Signup">
      <h4>had lmra hahya dwznaha lik! mra khra dber rask</h4>
      <div className="col-3">
        <input type="password" placeholder="password" onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
        <input type="password" placeholder="confirm password" onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
      </div>
      <div className="submit">
        <button onClick={_resetPass}>reset</button>
      </div>
    </div>
  )
}
