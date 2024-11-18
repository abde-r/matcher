'use client';

import React, { useState } from 'react';
import { usePathname, useRouter } from 'next/navigation';
import UserDropdown from './UserIconDropbutton';

export default function LoggedInNavbar() {
    const pathName = usePathname();

    return (
    <nav className="bg-white border-gray-200 dark:bg-gray-900">
        <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
        <a href="/home" className="flex items-center space-x-3 rtl:space-x-reverse">
            <img src="/logo.png" className="h-8" alt="MatchaX Logo" />
            <span className="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">MatchaX</span>
        </a>
        <div className="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
            <UserDropdown />
        </div>
        <div className="items-center justify-between hidden w-full md:flex md:w-auto md:order-1" id="navbar-cta">
            <ul className="flex flex-col font-medium p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
            <li>
                <a href="/home" className={`block py-2 px-3 md:p-0 ${pathName === '/home' || pathName === '/' ? 'text-blue-700 font-semibold underline' : 'text-gray-900'} rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700`}>Home</a>
            </li>
            <li>
                <a href="/favorits" className={`block py-2 px-3 md:p-0 ${pathName === '/favorits'? 'text-blue-700 font-semibold underline' : 'text-gray-900'} rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700`}>Favorits</a>
            </li>
            <li>
                <a href="/chat" className={`block py-2 px-3 md:p-0 ${pathName === '/chat'? 'text-blue-700 font-semibold underline' : 'text-gray-900'} rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:dark:hover:text-blue-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700`}>Chats</a>
            </li>
            </ul>
        </div>
        </div>
    </nav>
    );
}