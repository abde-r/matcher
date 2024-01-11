import axios from "axios";
import { useEffect, useState } from "react"
import { useAuth } from "../../components/Auth/Auth";

export const Home = () => {
  const [users, setUsers] = useState([])
  const auth = useAuth()

  useEffect(() => {
    const getUsers = async () => {
      const res = await axios.get(`http://localhost:8080/api/users`, { withCredentials: true })
      setUsers(res.data)
      //   fetch('http://localhost:8080/api/users', {
    //     method: 'GET',
    //     headers: {
    //       'Content-Type': 'application/json',
    //     },
    //     credentials: 'include',
    //   })
    //     .then(response => response.json())
    //     .then(data => setUsers(data))
    //     .catch(error => console.error('Error:', error));
    };

    getUsers();
  }, []);
  
  console.log('Auth: ', auth.user.username)
  return (
    <div className='Home'>
      If you're here, you're allowed to be here! Mr {auth.user.username}
    </div>
  )
}
