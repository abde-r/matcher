import { useEffect, useState } from "react"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCakeCandles, faCircle, faCircleDot, faCircleNotch, faGear, faHeart, faMapPin, faMars, faPerson, faPersonDress, faVenus } from "@fortawesome/free-solid-svg-icons";
import './Profile.scss';

export const Profile = () => {

    const [user, setUser] = useState<any>();
    // const [inputData, setInputData] = useState({ firstName: 'f', lastName: '', username: '', gender: 1, email: '', password: '' })
    
    useEffect(() => {
      (async () => {
        const res = await fetch(`http://localhost:8000/api/v1/users/`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
              query: `
                query ($id: ID!) {
                    user(id: $id) {
                        id
                        first_name
                        last_name
                        username
                        email
                        password
                        gender
                        preferences
                        pics
                        location
                    }
                }
              `,
              variables: {
                id: 18,
              }
            })
        })
        .then(res => { return res.json(); })
        .catch(error => { console.log('Error fetching users', error); });
          
        console.log('res', res)
        setUser(res?.data?.user);
      })()
    }, [])
    

    // const [userData, setUserData] = useState<any>(inputData)
    // const [me, setMe] = useState<any>([]);
    // const [myCookie, setMyCookie] = useState<string>('');

    // useEffect(() => {
    //     const cookieArray = document.cookie.split(';');

    //     for (let i = 0; i < cookieArray.length; i++) {
    //         const cookie = cookieArray[i].trim();
    //         if (cookie.startsWith('access-token=')) {
    //             setMyCookie(cookie.substring('access-token='.length))
    //             break;
    //         }
    //     }
    // }, [])

    // useEffect(() => {
    //     (async () => {
    //       const res = await axios.post(`http://localhost:8080/api/users/me`,
    //       {
    //         'access_token': myCookie,
    //       }, { withCredentials: true })
    
    //       console.log('res', res);
    //       if (res.data.status) {
    //         setMe(res.data)
    //       }
    //     })()
    //   }, [])

    //   console.log('mee', me)

    // const handleGenderChange = (e: any) => {
    //   const x = e.target.value === 'Male' ? 1 : 0
    //   setInputData({ ...inputData, gender: x})
    // }

    console.log("user", user)


    return (
        <div className="flex flex-col items-center justify-center text-gray-500 p-5">
            <div className="flex bg-[#e0e0e0] w-[85%] rounded-lg p-3">
                <div className="w-[20%] p-1 justify-center">
                    <img className="w-[140px] h-[140px] mx-auto rounded-[50%]" src={`https://www.refinery29.com/images/10267701.jpg?format=webp&width=720&height=864&quality=85`} />
                </div>
                <div className="flex flex-col w-[80%] items-start justify-center">
                    <h2 className="mx-1 my-1 font-bold">{user?.username} { 1 ? <FontAwesomeIcon className="text-[#27ac27] text-xs" icon={faCircle} /> : <FontAwesomeIcon className="text-[#ff6e6e] text-xs" icon={faCircleDot} /> }</h2>
                    <h3 className="mx-1">{ user?.gender ? <FontAwesomeIcon style={{color: '#1692b9'}} icon={faMars} /> : <FontAwesomeIcon style={{color: 'rgb(255 112 194)'}} icon={faVenus} /> } {user?.first_name} {user?.last_name} </h3>
                    <p className="mx-1"><FontAwesomeIcon style={{color: '#e43b3b'}} icon={faMapPin} /> LA</p>
                    <p className="mx-1"><FontAwesomeIcon style={{color: 'hsl(261, 97%, 62%)'}} icon={faCakeCandles} /> 30 years old</p>
                </div>
                <div className="w-[20%] flex items-center justify-center">
                    <span className="text-pink-400 text-2xl cursor-pointer hover:text-pink-500">
                        { 0 ? <FontAwesomeIcon icon={faGear} /> : <FontAwesomeIcon className="heart-icon" icon={faHeart} /> }
                    </span>
                </div>
            </div>
            
            <div className="flex flex-col w-[80%] my-3 text-start">
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-[#e0e0e0]">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">Bio</h1>
                    <p className="text-md mx-3">This is a bio</p>
                </div>
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-[#e0e0e0]">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">Preference</h1>
                    <p className="text-md mx-3">{ 1 ? <FontAwesomeIcon style={{color: '#1692b9'}} icon={faPerson} /> : <FontAwesomeIcon style={{color: 'rgb(226 88 167)'}} icon={faPersonDress} /> } | {`male`}</p>
                </div>
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-[#e0e0e0]">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">Interests</h1>
                    <div className="flex flex-row mx-3">
                        {
                            user?.preferences.split(';').map((p: string, index: number) => {
                                return <p key={index} className="bg-[#78979D] text-md m-1 p-1 text-gray-200 rounded-md">{p}</p>
                            })
                        }
                    </div>
                </div>
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-[#e0e0e0]">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">pictures</h1>
                    <div className="flex flex-row mx-3">
                        <img className="w-[150px] h-[150px] m-2 rounded-md" src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img className="w-[150px] h-[150px] m-2 rounded-md" src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img className="w-[150px] h-[150px] m-2 rounded-md" src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img className="w-[150px] h-[150px] m-2 rounded-md" src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
                        <img className="w-[150px] h-[150px] m-2 rounded-md" src={`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`} />
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
