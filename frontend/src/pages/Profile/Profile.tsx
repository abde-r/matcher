import axios from "axios";
import { useEffect, useState } from "react"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCakeCandles, faCircle, faCircleDot, faCircleNotch, faGear, faHeart, faMapPin, faMars, faPerson, faPersonDress, faVenus } from "@fortawesome/free-solid-svg-icons";
import './Profile.scss';

export const Profile = () => {

    const [inputData, setInputData] = useState({ firstName: 'f', lastName: '', username: '', gender: 1, email: '', password: '' })
    const [userData, setUserData] = useState<any>(inputData)
    const [me, setMe] = useState<any>([]);
    const [myCookie, setMyCookie] = useState<string>('');

    useEffect(() => {
        const cookieArray = document.cookie.split(';');

        for (let i = 0; i < cookieArray.length; i++) {
            const cookie = cookieArray[i].trim();
            if (cookie.startsWith('access-token=')) {
                setMyCookie(cookie.substring('access-token='.length))
                break;
            }
        }
    }, [])

    useEffect(() => {
        (async () => {
          const res = await axios.post(`http://localhost:8080/api/users/me`,
          {
            'access_token': myCookie,
          }, { withCredentials: true })
    
          console.log('res', res);
          if (res.data.status) {
            setMe(res.data)
          }
        })()
      }, [])

      console.log('mee', me)

    const handleGenderChange = (e: any) => {
      const x = e.target.value === 'Male' ? 1 : 0
      setInputData({ ...inputData, gender: x})
    }


    return (
        <div className="Profile">
            <div className="account-infos">
                <div className="pfp">
                    <img src={`https://www.refinery29.com/images/10267701.jpg?format=webp&width=720&height=864&quality=85`} />
                </div>
                <div className="personal-infos">
                    <h2>annadear { 1 ? <FontAwesomeIcon className="onlineIcon" icon={faCircle} /> : <FontAwesomeIcon className="offlineIcon" icon={faCircleDot} /> }</h2>
                    <h3>anna dearmas { 0 ? <FontAwesomeIcon style={{color: '#1692b9'}} icon={faMars} /> : <FontAwesomeIcon style={{color: 'rgb(255 112 194)'}} icon={faVenus} /> }</h3>
                    <p><FontAwesomeIcon style={{color: '#e43b3b'}} icon={faMapPin} /> LA</p>
                    <p><FontAwesomeIcon style={{color: 'hsl(261, 97%, 62%)'}} icon={faCakeCandles} /> 30 years old</p>
                </div>
                <div className="settings-heart">
                {
                    0 ?
                        (<FontAwesomeIcon className="settings-icon" icon={faGear} />)
                    :
                        (<FontAwesomeIcon className="heart-icon" icon={faHeart} />)
                }
                </div>
            </div>
            <div className="personal-infos">
                <div className="row-personal-info">
                    <h2>Bio</h2>
                    <p>This is a bio</p>
                </div>
                <div className="row-personal-info">
                    <h2>Preference</h2>
                    <p>{ 1 ? <FontAwesomeIcon style={{color: '#1692b9'}} icon={faPerson} /> : <FontAwesomeIcon style={{color: 'rgb(226 88 167)'}} icon={faPersonDress} /> } | {`male`}</p>
                </div>
                <div className="row-personal-info">
                    <h2>Interests</h2>
                    <div className="content">
                        <p className="interests">sport</p>
                        <p className="interests">books</p>
                        <p className="interests">travel</p>
                        <p className="interests">party</p>
                    </div>
                </div>
                <div className="row-personal-info">
                    <h2>pictures</h2>
                    <div className="content">
                        <img src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                    </div>
                </div>
            </div>
        {/* <div className="col-1">
            <label>first name</label>
            <input type="text" placeholder="first name" value={ userData.firstName } onChange={(e) => { setInputData({ ...inputData, firstName: e.target.value }) }} />
            <label>last name</label>
            <input type="text" placeholder="last name" onChange={(e) => { setInputData({ ...inputData, lastName: e.target.value }) }} />
        </div> */}
        {/* <div className="col-2">
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
        </div> */}
        {/* <div className="submit">
            <button onClick={_signup}>sign up</button>
            <a href="/login">you already have an account?</a>
        </div> */}
        </div>
    )
}
