import { useEffect, useState } from "react"
import { IoBalloonSharp, IoFemale, IoHeart, IoLocationSharp, IoMale, IoManSharp, IoSettingsSharp, IoToggleSharp, IoWomanSharp } from "react-icons/io5";

export const Profile = () => {

    const [user, setUser] = useState<any>();
    const [cookiza, setCookiza] = useState<string>('');
    // const [inputData, setInputData] = useState({ firstName: 'f', lastName: '', username: '', gender: 1, email: '', password: '' })
    
    useEffect(() => {
      (async () => {
        const res = await fetch(`http://localhost:8000/api/v1/users/token`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({
              query: `
                query ($token: String!) {
                    userByToken(token: $token) {
                        id
                        first_name
                        last_name
                        username
                        email
                        password
                        birthday
                        gender
                        preferences
                        pics
                        location
                    }
                }
              `,
              variables: {
                token: cookiza,
              }
            })
        })
        .then(res => { return res.json(); })
        .catch(error => { console.log('Error fetching users', error); });
          
        console.log('res', res)
        
        if (res && res.data && res.data.userByToken) {
            let pics = [];
            let len = res.data.userByToken.pics.split(';;;');
      
            // Ensure the pics array has at least 5 elements
            while (len < 5) {
                pics.push('https://pbs.twimg.com/profile_images/1064384665006587907/jqjE-D6T_400x400.jpg');
                len++;
            }
      
            // Update the pics property in the res object
            res.data.userByToken.pics = pics;

            const d = new Date;
            res.data.userByToken.birthday = d.getFullYear() - res.data.userByToken.birthday.split('-')[0];
        }

        setUser(res?.data?.userByToken);
      })()
    }, [cookiza])
    

    // const [userData, setUserData] = useState<any>(inputData)
    // const [me, setMe] = useState<any>([]);
    // const [myCookie, setMyCookie] = useState<string>('');

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

    console.log("user", user)
    console.log("cokiza", cookiza)


    return (
        <div className="flex flex-col items-center justify-center text-gray-500 p-5">
            <div className="flex bg-gray-300 w-[85%] rounded-lg p-3">
                <div className="w-[20%] p-1 justify-center">
                    <img className="w-[140px] h-[140px] mx-auto rounded-[50%]" src={user?.pics[0]} />
                </div>
                <div className="flex flex-col w-[80%] items-start justify-center">
                    <h2 className="flex items-center mx-1 my-1 text-lg font-bold">@{user?.username} {<IoToggleSharp className={`${1 ? 'text-[#27ac27]' : 'text-[#8b8a8a]'} ml-2 text-xl`} />}</h2>
                    <h3 className="flex items-center mx-1 font-semibold">{ user?.gender ? <IoMale className="mr-1 text-lg text-blue-500" /> : <IoFemale className="ml-2 text-xl text-red-500" /> } {user?.first_name} {user?.last_name} </h3>
                    <p className="flex items-center mx-1 font-semibold"><IoLocationSharp className="mr-1 text-lg text-[#714bd2]" /> {user?.location}</p>
                    <p className="flex items-center mx-1 font-semibold"><IoBalloonSharp className="mr-1 text-lg text-[#c54bd2]" /> {user?.birthday} years old</p>
                </div>
                <div className="w-[20%] flex items-center justify-center">
                    <span className="text-pink-400 text-2xl cursor-pointer hover:text-pink-500">
                        { 1 ? <IoSettingsSharp /> : <IoHeart /> }
                    </span>
                </div>
            </div>
            
            <div className="flex flex-col w-[80%] my-3 text-start">
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-gray-300">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">Bio</h1>
                    <p className="text-md mx-3">This is a bio</p>
                </div>
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-gray-300">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">Preference</h1>
                    <p className="flex items-center text-md mx-3">{ !user?.gender ? <IoManSharp /> : <IoWomanSharp /> } | {!user?.gender ? 'male' : 'female'}</p>
                </div>
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-gray-300">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">Interests</h1>
                    <div className="flex flex-row mx-3">
                        {
                            user?.preferences.split(';').map((p: string, index: number) => {
                                return <p key={index} className="bg-[#78979D] text-md m-1 p-1 text-gray-200 rounded-md">{p}</p>
                            })
                        }
                    </div>
                </div>
                <div className="flex flex-col my-3 rounded-sm p-2 text-start bg-gray-300">
                    <h1 className="text-lg font-semibold uppercase text-[#1692B9]">pictures</h1>
                    <div className="flex flex-row mx-3">
                        {
                            user?.pics.map((pic: string, index: number) => {
                                console.log('pic', pic)
                                return pic.length && <img key={index} className="w-[150px] h-[150px] m-2 rounded-md" src={pic/*`https://variety.com/wp-content/uploads/2022/01/ana.jpg?w=1000&h=563&crop=1&resize=1000%2C563`*/} />;
                            })
                        }
                    </div>
                </div>
            </div>
        </div>
    )
}
