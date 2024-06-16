// import { createContext, useContext, useState } from "react";

// const AuthContext = createContext(null)

// export const useAuth = () => {
//     return useContext(AuthContext)
// }

// export const AuthProvider = ({ children }: any) => {

//     const [user, setUser] = useState(null)

//     const login = (user: any) => {
//         setUser(user)
//     }

//     const logout = () => {
//         setUser(null)
//     }

//     return <AuthContext.Provider value={{ user, login, logout }}>{ children }</AuthContext.Provider>
// }
