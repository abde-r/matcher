'use client';

import React, { useState } from 'react';
import Image from 'next/image';
import { Card } from '../ui/card';

export default function ForgetPasswordComponent() {
    const [email, setEmail] = useState('');

    const handleSendEmail = () => {
        //use fetch to send request for email reset password
    };

    return (
        <div className="p-4 sm:p-6 lg:p-8 space-y-4">
            <a href="/home" className="flex items-center justify-center">
                <Image
                src="/logo.png"
                width={32}
                height={32}
                className="h-8"
                alt="MatcherX logo" />
            </a>
            <Card className="max-w-lg mx-auto p-6 w-full sm:w-80 lg:w-96">
                <h2 className="text-2xl font-bold mb-6">Please Enter Your Email</h2>
                <div className="mb-4">
                <label className="block text-gray-700 mb-1">Email</label>
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="w-full p-2 border rounded"
                />
                </div>

                <button className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 w-full sm:w-auto" onClick={handleSendEmail}>
                Send Email
                </button>
            </Card>
        </div>
    );
}