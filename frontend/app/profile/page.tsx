import NotLoggedNavbar from '@/components/Navbar/NotLoggedIn';
import React from 'react';


export default function Profile() {
    return (
        <div className="justify-center items-center min-h-screen bg-gradient-to-r from-gray-800 to-gray-900">
            <NotLoggedNavbar />
            <h1>Profile Page</h1>
        </div>
    )
}