import { useEffect, useState } from "react";
import { IoArrowForwardCircleOutline, IoCloseCircleOutline } from "react-icons/io5";
// import { useNavigate } from "react-router-dom";
import validator from "validator";

const some_preferences: any = ['sport', 'books', 'party', 'travel', 'cars', 'memes', 'movies', 'anime'];

export const ProfileSettings = ({ show, onClose }: any) => {
  if (!show) {
    return null;
  }

//   const navigate = useNavigate();
  const [validationErrors, setValidationErrors] = useState<any>({});
  const [cookiza, setCookiza] = useState<string>('');
  // const [location, setLocation] = useState({ latitude: -1, longitude: -1 });

  const [inputData, setInputData] = useState<any>({
    firstName: '',
    lastName: '',
    username: '',
    birthday: '',
    preferences: [],
    pics: [],
    location: '',
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
    // if (!inputData.location.length) {
    //   errors.terms = 'Allow Terms limak!';
    // }

    setValidationErrors(errors);
    return Object.keys(errors).length === 0;
  };

  useEffect(() => {
    const cookieArray = document.cookie.split(';');
    console.log(cookieArray)

    for (let i = 0; i < cookieArray.length; i++) {
        const cookie = cookieArray[i].trim();
        if (cookie.startsWith('matcher-token=')) {
            setCookiza(cookie.substring('matcher-token='.length))
            break;
        }
    }
  }, [])
  
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


  const _updateInfos = async () => {
    
        if (validateInputs()) {
            // console.log('wee wew', inputData)
            const res = await fetch(`http://localhost:8000/api/v1/users/update-info`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify({
                query:`
                    mutation UpdateUserInfo($input: UpdateUserInfoInput!) {
                        updateUserInfo(input: $input) {
                            first_name,
                            last_name,
                            birthday,
                            preferences,
                            pics,
                            location,
                            token,
                        }
                    }
                `,
                variables: {
                    input: {
                    first_name: inputData.firstName,
                    last_name: inputData.lastName,
                    birthday: inputData.birthday,
                    preferences: (inputData.preferences).join(';'),
                    pics: (inputData.pics).join(";;;"),
                    location: inputData.location, // THIS LOCATION MUST BE HANDLED LATER!!!
                    token: cookiza,
                    }
                }
                }),
            })
            .then(res => { return res.json(); })
            .catch(error => { console.log('Error proceeding registration', error); });

            // console.log('res', res.data);
            if (res.data.updateUserInfo) {
                onClose()
                window.location.reload();
            }
        }
    }

  console.log('updated data', inputData)

  return (
    <div className="fixed inset-0 bg-gray-800 bg-opacity-75 flex items-center justify-center">
      <div className="bg-white p-4 rounded-lg w-[75vh]">
        <div className="flex items-center justify-between my-2">
            <h1 className="text-xl capitalize font-semibold border-b-4 border-[#714bd2] rounded-sm text-gray-500">Personal Informations</h1>
            <span className="text-[#714bd2] text-2xl cursor-pointer mx-4" onClick={onClose}><IoCloseCircleOutline/></span>
        </div>

        <div className="flex flex-col my-7 items-center justify-center w-[90%] mx-auto p-3">
            <p className="text-gray-500 text-center ">* This informations will give other users to get to know more about you.</p>
            <div className="flex my-5">
                <div className="flex flex-col">
                    <input className="p-2 mx-2 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="text" placeholder="First Name" onChange={(e) => { setInputData({ ...inputData, firstName: e.target.value }) }} />
                    {validationErrors.firstName && <p className="text-red-500 font-semibold text-sm">*{validationErrors.firstName}</p>}
                  </div>
                  <div className="flex flex-col">
                    <input className="p-2 mx-2 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="text" placeholder="Last Name" onChange={(e) => { setInputData({ ...inputData, lastName: e.target.value }) }} />
                    {validationErrors.lastName && <p className="text-red-500 font-semibold text-sm">*{validationErrors.lastName}</p>}
                  </div>
                </div>
                
                <textarea className="p-2 mx-2 w-[80%] rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" placeholder="Bio" onChange={(e) => { setInputData({ ...inputData, bio: e.target.value }) }} />
                {validationErrors.bio && <p className="text-red-500 font-semibold text-sm">*{validationErrors.bio}</p>}

                <input className="p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" type="date" placeholder="Birth date" onChange={(e) => { setInputData({ ...inputData, birthday: e.target.value }) }} />
                {validationErrors.birthday && <p className="text-red-500 font-semibold text-sm">*{validationErrors.birthday}</p>}


                <div className="flex flex-col my-2 items-center w-[100%]">
                  <select className="p-2 my-3 rounded-sm text-gray-500 bg-transparent outline-none border-b-2 border-gray-400" onChange={handleAddPreference}>
                    <option hidden>Preferences</option>
                    {some_preferences.map((preference: string, index: number) => {
                      return (<option key={index} className="capitalize" value={preference}>{preference}</option>)
                    })}
                  </select>
                  {validationErrors.preferences && <p className="text-red-500 font-semibold text-sm">*{validationErrors.preferences}</p>}
                  <div className="flex flex-wrap mx-5 rounded-sm w-[80%] items-center justify-center">
                    {
                      inputData.preferences.map((pr: string, index: number) => {
                        return (
                        <div key={index} className='flex bg-blue-100 p-1'>
                          <p className="flex items-center text-sm font-semibold text-gray-500 border border-[#714bd2] rounded-sm m-1 p-1 capitalize">{pr} <span className="ml-1 text-lg cursor-pointer" onClick={() => handleRemovePreference(pr)}><IoCloseCircleOutline /></span></p>
                        </div>
                      )
                      })
                    }
                  </div>
                </div>
            <a className="flex items-center bg-[#714bd2] px-3 py-2 rounded-sm text-gray-300 text-md font-semibold cursor-pointer uppercase" onClick={_updateInfos}><span className="mr-1 text-xl"><IoArrowForwardCircleOutline /></span>update</a>
            </div>
      </div>
    </div>
  );
};
