import axios from "axios"
import { useState } from "react"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCalendar, faCircleRight } from "@fortawesome/free-regular-svg-icons"
import './ProceedSignup.scss'

export const ProceedSignup = () => {

  const [inputData, setInputData] = useState({ firstName: '', lastName: '', username: '', gender: 1, email: '', password: '' })

  const _signup = async () => {
    try {
      console.log('input data: ', inputData)

      // // check data first
      const res = await axios.post(`http://localhost:8080/auth/signup`, {
        first_name: inputData.firstName,
        last_name: inputData.lastName,
        username: inputData.username,
        gender: inputData.gender,
        email: inputData.email,
        password: inputData.password
      }, { withCredentials: true })
      console.log('res', res.data)
    }
    catch (err) {
      console.error('errrror: ', err)
    }
  }

  const handleGenderChange = (e: any) => {
    const x = e.target.value === 'Male' ? 1 : 0
    setInputData({ ...inputData, gender: x})
  }

  return (
    <div className="Proceed-signup">
      <div className="container">
        <h2>Proceed Signup</h2>
        <p>This informations will give other users to get to know more about you.</p>
        <div className="SignupForms">
            <div className="userNames">
                <input type="text" placeholder="First name" onChange={(e) => { setInputData({ ...inputData, username: e.target.value }) }} />
                <input type="text" placeholder="Last name" onChange={(e) => { setInputData({ ...inputData, username: e.target.value }) }} />
            </div>
            <input type="date" placeholder="Birth date" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
            <div className='gender'>
                <span>
                    <input type="radio" name="genderOptions" value="Male" onChange={handleGenderChange} />
                    Male
                </span>
                <span>
                    <input type="radio" name="genderOptions" value="Female" onChange={handleGenderChange} />
                    Female
                </span>
            </div>
            <select>
                <option hidden>Preferences</option>
                <option value='p2'>P1</option>
                <option value='p1'>P2</option>
            </select>
            <input type="text" placeholder="Interests" onChange={(e) => { setInputData({ ...inputData, password: e.target.value }) }} />
        </div>
        <button onClick={_signup}><FontAwesomeIcon icon={faCircleRight} /> Submit</button>
      </div>
    </div>
  )
}
