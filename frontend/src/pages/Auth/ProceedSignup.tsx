// import axios from "axios"
import { useState } from "react"
// import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
// import { faCalendar, faCircleRight } from "@fortawesome/free-regular-svg-icons"
// import './ProceedSignup.scss'
import { IoArrowForwardCircleOutline, IoCloseCircleOutline } from "react-icons/io5";
import { useNavigate } from "react-router-dom";
import validator from "validator";


const some_preferences: any = ['sport', 'books', 'party', 'travel', 'cars', 'memes', 'movies', 'anime'];

export const ProceedSignup = () => {

  const navigate = useNavigate();
  const [validationErrors, setValidationErrors] = useState<any>({});

  const [inputData, setInputData] = useState({
    firstName: '',
    lastName: '',
    username: '',
    birthday: '',
    gender: true,
    preferences: [],
    pics: [],
    location: "randomLocation127.0.0.1",
  })


  const validateAge = (birthdayString: string): number => {
    const birthday = new Date(birthdayString)
    const diff = Date.now() - birthday.getTime();
    const ageDate = new Date(diff);
    return Math.abs(ageDate.getUTCFullYear()-1970);
  }
  const validateInputs = () => {
    const errors: any = {};

    if (validator.isEmpty(inputData.firstName)) {
      errors.firstName = 'first name is required';
    }
    if (validator.isEmpty(inputData.lastName)) {
      errors.lastName = 'last name is required';
    }
    if (validator.isEmpty(inputData.birthday)) {
      errors.birthday = 'birthday is required';
    }
    if (validateAge(inputData.birthday) < 18) {
      errors.birthday = 'under Age! sir awldi tl3ab';
    }
    if (!inputData.preferences.length) {
      errors.preferences = 'preferences are required';
    }

    setValidationErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const proceed_signup = async () => {
    
    // data validation before sending request
    if (validateInputs()) {
      const res = await fetch(`http://localhost:8000/api/v1/users/proceed-registration`, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          credentials: 'include',
          body: JSON.stringify({
            query:`
              mutation ProceedRegistrationUser($input: ProceedRegisterationUserInput!) {
                proceedRegistrationUser(input: $input) {
                  id,
                  first_name,
                  last_name,
                  birthday,
                  gender,
                  preferences,
                  pics,
                  location
                }
              }
            `,
            variables: {
              input: {
                id: 18,
                first_name: inputData.firstName,
                last_name: inputData.lastName,
                birthday: "2-2-2",//inputData.birthday,
                gender: inputData.gender,
                preferences: (inputData.preferences).join(';'),
                pics: (inputData.pics).join(";;;"),
                location: inputData.location,
              }
            }
          }),
      })
      .then(res => { return res.json(); })
      .catch(error => { console.log('Error proceeding registration', error); });
  
      console.log('res', res.data);
      if (res.data.proceedRegistrationUser)
        navigate('/profile');
    }
  }

  const handleGenderChange = (e: any) => {
    const x = e.target.value === 'Male' ? true : false
    setInputData({ ...inputData, gender: x})
  }

  const handleAddPreference = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const value: string = e.target.value;

    if (!inputData.preferences.includes(value)) {
      setInputData((prevState: any) => ({
        ...prevState,
        preferences: [...prevState.preferences, value],
      }));
    }
  }

  const handleRemovePreference = (value: string) => {
    
    if (inputData.preferences.includes(value)) {
      setInputData((prevState: any) => ({
        ...prevState,
        preferences: prevState.preferences.filter((p: string) => p !== value),
      }));
    }
  }

  console.log(inputData)

  return (
    <div className="flex flex-col h-[90vh] bg-[#d3d3d3] items-center justify-center p-4 m-10 w-[75%] mx-auto rounded-md">
      {/* <div className="flex flex-col items-center justify-center border  rounded-md p-20"> */}
        <h1 className="text-3xl capitalize my-7 font-semibold border-b-4 border-[#714bd2] rounded-sm text-gray-500">Proceed Signup</h1>
            <div className="flex flex-col my-5 items-center justify-center">
                <div className="flex my-5">
                  <div className="flex flex-col">
                    <input className="p-2 mx-2 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="text" placeholder="First Name" onChange={(e) => { setInputData({ ...inputData, firstName: e.target.value }) }} />
                    {validationErrors.firstName && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.firstName}</p>}
                  </div>
                  <div className="flex flex-col">
                    <input className="p-2 mx-2 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="text" placeholder="Last Name" onChange={(e) => { setInputData({ ...inputData, lastName: e.target.value }) }} />
                    {validationErrors.lastName && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.lastName}</p>}
                  </div>
                </div>
                
                <input className="p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="date" placeholder="Birth date" onChange={(e) => { setInputData({ ...inputData, birthday: e.target.value }) }} />
                {validationErrors.birthday && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.birthday}</p>}

                
                <div className='flex my-2'>
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

                <div className="flex flex-col my-2 items-center w-[100%]">
                  <select className="p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" onChange={handleAddPreference}>
                    <option hidden>Preferences</option>
                    {some_preferences.map((preference: string) => {
                      return (<option className="capitalize" value={preference}>{preference}</option>)
                    })}
                  </select>
                  {validationErrors.preferences && <p style={{ color: 'red', fontSize: '12px' }}>*{validationErrors.preferences}</p>}
                  <div className="flex flex-wrap mx-5 rounded-sm w-[80%] items-center justify-center">
                    {
                      inputData.preferences.map((pr: string, index: number) => {
                        return (
                        <div className='flex bg-blue-100 p-1'>
                          <p key={index} className="flex items-center text-sm font-semibold text-gray-500 border border-[#714bd2] rounded-sm m-1 p-1 capitalize">{pr} <span className="ml-1 text-lg cursor-pointer" onClick={() => handleRemovePreference(pr)}><IoCloseCircleOutline /></span></p>
                        </div>
                      )
                      })
                    }
                  </div>
                </div>

                <div className="check-conditions my-4">
                  <label className="conditions-checkbox flex-col">
                    <p className="text-gray-500">This informations will give other users to get to know more about you.</p>
                    <input className="cursor-pointer" type="checkbox" />
                    <span className="text-gray-500 underline">I agree to terms of us</span>
                  </label>                  
                </div>
            </div>
            <a className="flex items-center bg-[#714bd2] px-3 py-2 rounded-sm text-gray-300 text-md font-semibold cursor-pointer uppercase" onClick={proceed_signup}><span className="mr-1 text-xl"><IoArrowForwardCircleOutline /></span>Submit</a>
          
    {/* </div> */}
    </div>
  )
}
