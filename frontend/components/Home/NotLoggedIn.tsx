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
        // <div className="flex min-h-screen bg-gray-300 pt-20 sm:pt-24 overflow-y-scroll">
        //     <div className="md:w-1/2 w-full flex flex-col items-center justify-center pl-8">
        //         <div className="sm:text-8xl text-6xl space-y-10 w-full">
        //             <h1 className="font-bold">MatcherX</h1>
        //             <h2 className="font-light space-y-10">
        //                 Where Sparks Fly and Connections Blossom in Every Like.
        //             </h2>
        //         </div>
        //         <div className="flex flex-col pt-40 md:pt-12 sm:pt-12 md:pl-36 pl-12 w-full sm:text-3xl text-xl">
        //             <h1 className="text-[#e9aab2] font-bold">
        //                 Start liking by creating your account today for free ✨
        //             </h1>
        //             <button
        //                 type="button"
        //                 className="md:w-1/3 w-1/2 ml-auto text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-semibold rounded-lg text-xl px-4 py-4 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        //                 onClick={getStartedButton}
        //             >
        //                 Get started
        //             </button>
        //         </div>
        //     </div>
        //     <div className="md:w-1/2 md:flex items-center justify-center hidden">
        //         <Image
        //             src="/home_background.png"
        //             alt="Home image"
        //             width={700}
        //             height={700}
        //         />
        //     </div>
        // </div>
        <div className="flex justify-around items-center min-h-screen bg-gray-300 pt-20 sm:pt-24 overflow-y-scroll p-8">
            <div className='flex flex-row'>
                <div className="md:w-1/2 w-full flex flex-col items-center justify-between p-4 bg-[#dde3ed] border border-gray-400 rounded-lg">
                    <div className="flex flex-col justify-center items-center text-center space-y-10 w-full">
                        <h1 className="sm:text-8xl text-8xl font-bold underline">MatcherX</h1>
                        <h2 className="sm:text-7xl text-7xl font-semibold text-gray-500 space-y-10">
                            Where Sparks Fly and Connections Blossom in Every Like.
                        </h2>
                    </div>
                    <div className="flex flex-col justify-center items-center space-y-8 w-full sm:text-xl text-xl text-center h-1/3 ">
                        <h1 className="text-gray-500 font-bold mb-2">
                            Start liking by creating your account today for free ✨
                        </h1>
                        <div className='w-full flex justify-center items-center px-8'>
                            <button
                                type="button"
                                className="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-semibold rounded-lg text-xl px-4 py-4 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                                onClick={getStartedButton}
                            >
                                Get started
                            </button>
                        </div>
                    </div>
                </div>
                <div className="md:w-1/2 md:flex items-center justify-center hidden">
                    <Image
                        src="/home_background.png"
                        alt="Home image"
                        width={700}
                        height={700}
                    />
                </div>
            </div>
        </div>
    );
}