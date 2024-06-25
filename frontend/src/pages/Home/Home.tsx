// import axios from "axios";
import { useEffect, useState } from "react"
// import { useNavigate } from "react-router-dom";
// import { useAuth } from "../../components/Auth/Auth";
// import './Home.scss';

export const Home = () => {
  const [users, setUsers] = useState<any>([])
  // const auth: any = useAuth()

  // useEffect(() => {
  //   const getUsers = async () => {
  //     const res = await axios.get(`http://localhost:8080/api/users`, { withCredentials: true })
  //     setUsers(res.data)
  //     //   fetch('http://localhost:8080/api/users', {
  //   //     method: 'GET',
  //   //     headers: {
  //   //       'Content-Type': 'application/json',
  //   //     },
  //   //     credentials: 'include',
  //   //   })
  //   //     .then(response => response.json())
  //   //     .then(data => setUsers(data))
  //   //     .catch(error => console.error('Error:', error));
  //   };

  //   getUsers();
  // }, []);

   

  // useEffect(() => {
  //   const getUsers = async () => {
        
  //         try {
  //           const response = await fetch(`http://localhost:8000/api/v1/users/g`, {
  //             method: 'GET',
  //             headers: {
  //               'Content-Type': 'application/json',
  //               'Access-Control-Allow-Origin': '*',
  //             },
  //             credentials: 'include' // This ensures credentials like cookies are included
  //           });
      
  //           if (!response.ok) {
  //             throw new Error('Network response was not ok');
  //           }
      
  //           const data = await response.json();
  //           setUsers(data);
  //         } catch (err) {
  //           console.error('Error:', err);
  //         }
  //       }
  //   // };

  //   getUsers();

  // }, [])

  // console.log(users)

  // console.log('gg', document.cookie)
  // const navigate = useNavigate();

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
              query {
                users {
                  id
                  first_name
                  last_name
                  username
                  email
                  password
                  gender
                }
              }
            `
          })
      })
      .then(res => { return res.json(); })
      .catch(error => { console.log('Error fetching users', error); });
        
      setUsers(res.data.users);
    })()
  }, [])

  // console.log('users', users)
  
  // console.log('Auth: ', auth.user.username)
  return (
    <div className='flex flex-col h-[98vh] bg-[#d3d3d3] items-center justify-center p-4 m-2 rounded-md'>
      <div className="flex flex-col text-2xl text-gray-700 text-center">
        <h1 className="items-start">matcherX</h1>
        <p>Where Sparks<br /> Fly and Connections <br /> Blossom in Every<br /> Like</p>
      </div>
      {/* <div className="home-img">
        <img src={`https://embrace-autism.com/wp-content/uploads/Tests-DrNatalieEngelbrecht-Test2-hover.svg`} />
      </div> */}
      <div>
      {
        users.map((user: any, index: number) => {
          return (<div key={index}>
            <p>{user?.first_name} ait {user?.last_name}</p>
          </div>)
        })
      }
      </div>
    </div>
  )
}
