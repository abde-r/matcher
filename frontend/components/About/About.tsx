'use client';

import React from 'react';
import Image from 'next/image';

export default function AboutComponent() {
    return (
        <div className="flex min-h-screen bg-gray-300 pt-20 sm:pt-24 overflow-y-scroll">
            <div className="md:w-1/2 w-full flex flex-col items-center justify-center px-8">
                <div className="sm:text-xl text-lg space-y-6 w-full">
                    <h1 className="text-4xl font-bold mb-6">Welcome to <b className='text-red-600'>MatcherX</b> - Where Connections Begin!</h1>
                    <p>Discover a new way to connect with interesting people near you. Whether you're looking for meaningful relationships, fun friendships, or just a casual chat, we've got you covered. Our platform is designed to make meeting others as simple and enjoyable as swiping right.</p>
                    <h2 className="text-2xl font-semibold mt-6">🌟 What Makes Us Different?</h2>
                    <ul className="list-disc pl-5 space-y-3">
                        <li><span className="font-semibold text-green-700">🔥 Swipe Your Way to Connections:</span> Differently to other Tinder-like websites or apps, this one is more like liking others' profiles and waiting for them to do the same. If they like you back, it’s a match!</li>
                        <li><span className="font-semibold text-green-700">🚀 Real-Time Chat:</span> Say goodbye to endless waiting. Start chatting with your matches instantly and get to know each other better.</li>
                        <li><span className="font-semibold text-green-700">❤️ Find Your Perfect Match:</span> Our intelligent browsing algorithm suggests potential matches based on your preferences, location, and interests. Say hello to like-minded individuals!</li>
                        <li><span className="font-semibold text-green-700">🌎 Local and Global Connections:</span> Whether you're interested in meeting people nearby or from around the world, our platform offers you the opportunity to explore both.</li>
                        <li><span className="font-semibold text-green-700">💬 Safe and Secure:</span> We prioritize your safety and privacy. We've implemented robust measures to ensure a safe and enjoyable experience for all users.</li>
                        <li><span className="font-semibold text-green-700">📱 Mobile and Web:</span> Access our platform on your phone, tablet, or desktop. Stay connected no matter where you are.</li>
                    </ul>
                    <p className="mt-6">Ready to embark on a journey of new connections and exciting conversations? Sign up now and join our diverse community of individuals looking to make meaningful connections. Your next great conversation could be just a simple search away!</p>
                    <p className="font-bold text-lg">Join <b className='text-red-600'>MatcherX</b> today and start browsing towards new connections!</p>
                </div>
            </div>
            <div className="md:w-1/2 md:flex items-center justify-center hidden">
                <Image
                    src="/about_background.png"
                    alt="About image"
                    width={700}
                    height={700}
                />
            </div>
        </div>
    )
}
