import { Navigate, Outlet } from 'react-router-dom';

export const PrivateRoutes = ({ auth }: any) => {

    // const auth: any = {token: false};

    return (
        auth.token ? <Outlet/> : <Navigate to='/login' />
    )
}
