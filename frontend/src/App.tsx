// import './App.css'
import { Routes, Route } from 'react-router-dom'
import { Home } from './pages/Home/Home'
import { Signup } from './pages/Auth/Signup'
import { Login } from './pages/Auth/Login'
import axios from 'axios'
import { Navbar } from './components/Navbar/Navbar'
import { AuthProvider } from './components/Auth/Auth'
import { RequireAuth } from './components/Auth/RequireAuth'
import { Profile } from './pages/Profile/Profile'

axios.defaults.baseURL = `http://localhost:8080`
axios.defaults.withCredentials = true

function App() {

  return (
    <div className='app'>
      <AuthProvider>
        <Navbar />
        <Routes>
          {/* <Route path='/' element={<RequireAuth><Home /></RequireAuth>} /> */}
          <Route path='/' element={<Home />} />
          <Route path='/profile' element={ <Profile /> } />
          <Route path='/login' element={<Login />} />
          <Route path='/signup' element={<Signup />} />
        </Routes>
      </AuthProvider>
    </div>
  )
}

export default App
