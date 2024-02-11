import axios from "axios";
import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom";
// import { useAuth } from "../../components/Auth/Auth";

export const Home = () => {
  const [users, setUsers] = useState([])
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

  // console.log(users)

  const navigate = useNavigate();
  useEffect(() => {
    (async () => {
      console.log('holaa')
      const res = await axios.get(`http://localhost:8080/api/users/verifyToken`, { withCredentials: true })

      console.log('res', res);
      if (res.data.status) {

      }
      else {
        navigate('/login');
      }
    })()
  }, [])
  
  // console.log('Auth: ', auth.user.username)
  return (
    <div className='Home'>
      If you're here, you're allowed to be here! Mr 
    </div>
  )
}
