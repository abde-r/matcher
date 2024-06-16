import { Link } from "react-router-dom"
import './Navbar.scss'

export const Navbar = () => {
  return (
    <div className="Navbar">
      <nav>
        <Link to={'/'}>Home</Link>
        <Link to={'/profile'}>Profile</Link>
        <Link to={'/signup'}>Signup</Link>
        <Link to={'/login'}>Login</Link>
      </nav>
    </div>
  )
}
