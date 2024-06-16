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
                    <input type="text" placeholder="First name" onChange={(e) => { setInputData({ ...inputData, firstName: e.target.value }) }} />
                    <input type="text" placeholder="Last name" onChange={(e) => { setInputData({ ...inputData, lastName: e.target.value }) }} />
                </div>
                

                <input type="date" placeholder="Birth date" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
                <div className='gender'>
                    {/* <span>
                        <input type="radio" name="genderOptions" value="Male" onChange={handleGenderChange} />
                        <label>Male 👨</label>
                        
                    </span>
                    <span>
                        <input type="radio" name="genderOptions" value="Female" onChange={handleGenderChange} />
                        <label>Female 🧕</label>
                        
                    </span> */}

                    <label className="rad-label">
                      <input type="radio" className="rad-input" value='Male' name="genderOptions" onChange={handleGenderChange} />
                      <div className="rad-design"></div>
                      <div className="rad-text">🙍‍♂️Male</div>
                    </label>

                    <label className="rad-label">
                      <input type="radio" className="rad-input" value='Female' name="genderOptions" onChange={handleGenderChange} />
                      <div className="rad-design"></div>
                      <div className="rad-text">🧕Female</div>
                    </label>
                </div>
                <div className="preferences">
                  <select onChange={handleSelectChange}>
                    <option hidden>Preferences</option>
                    <option value='p2'>P1</option>
                    <option value='p1'>P2</option>
                  </select>
                  <div className="selected-preferences">
                    {
                      inputData.preferences.map((pr: string, index: number) => {
                        return (
                        <div className='selected-preference'>
                          <p key={index}>{pr}</p>
                        </div>)
                      })
                    }
                  </div>
                </div>
                <div>
                  <input type="text" placeholder="Interests" value={inputData.currentInterest} onChange={handleInterestsChange} onKeyDown={handleKeyDown} />
                  {inputData.interests.map((interest: string, index: number) => (
                    <p key={index}>{interest}</p>
                  ))}
                </div>
                <div className="check-conditions">
                    <label className="conditions-checkbox">
                      <input type="checkbox" />
                      <span>I agree to terms of us</span>
                    </label>                  
                </div>
            </div>
            <a className="proceed" onClick={_signup}><FontAwesomeIcon className='submit-circle' icon={faCircleRight} />Submit</a>
          
    </div>
    </div>
  )
}
