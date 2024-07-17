// import './App.css'
import { Routes, Route } from 'react-router-dom'
import { Home } from './pages/Home/Home'
import { Signup } from './pages/Auth/Signup'
import { Login } from './pages/Auth/Login'
// import axios from 'axios'
import { Navbar } from './components/Navbar/Navbar'
// import { AuthProvider } from './components/Auth/Auth'
// import { RequireAuth } from './components/Auth/RequireAuth'
import { Profile } from './pages/Profile/Profile'
import { ResetPass } from './pages/Auth/ResetPass'
import { ForgotPass } from './pages/Auth/ForgotPass'
import { ProceedSignup } from './pages/Auth/ProceedSignup'
import { PrivateRoutes } from './pages/Auth/PrivateRoutes'
import { useState } from 'react'
import { Waiting } from './pages/Auth/Waiting'

// axios.defaults.baseURL = `http://localhost:8080`
// axios.defaults.withCredentials = true

function App() {

  const [auth, setAuth] = useState({ token: true });

    const logout = () => {
        setAuth({ token: false });
    };
    console.log('authh', auth.token)

  return (
    <div>
      {/* <AuthProvider> */}
        {/* <Navbar /> */}
        <Routes>
          {/* <Route path='/' element={<RequireAuth><Home /></RequireAuth>} /> */}
          <Route element={ <PrivateRoutes auth={auth} /> }>
            <Route path='/' element={<Home />} />
            <Route path='/profile' element={ <Profile /> } />
          </Route>
          {/* <Route path='/' element={<Home />} /> */}
          <Route path='/login' element={<Login setAuth={setAuth} />} />
          <Route path='/forgot-pass' element={ <ForgotPass /> } />
          <Route path='/reset-pass' element={<ResetPass />} />
          <Route path='/signup' element={<Signup setAuth={setAuth} />} />
          <Route path='/proceed-signup' element={<ProceedSignup />} />
          <Route path='/account-verification' element={<Waiting />} />
        </Routes>
      {/* </AuthProvider> */}
    </div>
  )
}

export default App
