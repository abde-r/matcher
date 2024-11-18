'use client';

import React, { useState } from 'react';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

export default function NotLoggedInHomeComponent() {
    const [isLogged, setIsLogged] = useState(false);
    const router = useRouter();

    const getStartedButton = async () => {
        if (isLogged) {
            router.push('/');
        } else {
            router.push('/auth');
        }
    }

    return (
        <div className="flex h-screen bg-gray-300 overflow-y-scroll">
            <div className="md:w-1/2 w-full flex flex-col items-center justify-center pl-8">
                <div className="sm:text-8xl text-6xl space-y-10 w-full">
                    <h1 className="font-bold">MatcherX</h1>
                    <h2 className="font-light space-y-10">Where Sparks Fly and Connections Blossom in Every Like.</h2>
                </div>
                <div className="flex flex-col pt-40 md:pl-36 pl-12 w-full sm:text-3xl text-xl">
                    <h1 className="text-red-800 font-bold">Start liking by creating your account today for free ✨</h1>
                    <button type="button" className="md:w-1/3 w-1/2 ml-auto text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800" onClick={getStartedButton}>
                        Get started
                    </button>
                </div>
            </div>
            <div className="md:w-1/2 md:flex items-center justify-center hidden">
                <Image
                    src='/home_background.png'
                    alt='Home image'
                    width={700}
                    height={700}
                />
            </div>
        </div>
    );
}