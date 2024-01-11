import { useEffect, useState } from "react"

export const Profile = () => {

    const [inputData, setInputData] = useState({ firstName: 'f', lastName: '', username: '', gender: 1, email: '', password: '' })
    const [userData, setUserData] = useState<any>(inputData)

    useEffect(() => {
        setUserData({firstName: 'ggc'})
    }, [])

    const handleGenderChange = (e: any) => {
      const x = e.target.value === 'Male' ? 1 : 0
      setInputData({ ...inputData, gender: x})
    }

    console.log('cookies: ',  document.cookie)

    return (
        <div className="Profile">
        <div className="col-1">
            <label>first name</label>
            <input type="text" placeholder="first name" value={ userData.firstName } onChange={(e) => { setInputData({ ...inputData, firstName: e.target.value }) }} />
            <label>last name</label>
            <input type="text" placeholder="last name" onChange={(e) => { setInputData({ ...inputData, lastName: e.target.value }) }} />
        </div>
        <div className="col-2">
            <label>username</label>
            <input type="text" placeholder="username" onChange={(e) => { setInputData({ ...inputData, username: e.target.value }) }} />
            <label>gender</label>
            <div className='genderType'>
                <span>
                    <input type="radio" name="genderOptions" value="Male" onChange={handleGenderChange} />
                    Male
                </span>
                <span>
                    <input type="radio" name="genderOptions" value="Female" onChange={handleGenderChange} />
                    Female
                </span>
            </div>
        </div>
        <div className="col-3">
            <label>email</label>
            <input type="email" placeholder="email" onChange={(e) => { setInputData({ ...inputData, email: e.target.value }) }} />
        </div>
        {/* <div className="submit">
            <button onClick={_signup}>sign up</button>
            <a href="/login">you already have an account?</a>
        </div> */}
        </div>
    )
}
