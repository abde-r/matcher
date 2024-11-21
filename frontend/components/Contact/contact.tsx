'use client';

import { useState } from 'react';

export default function ContactComponent() {
    // State to handle form data
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        message: ''
    });

    // I will use toast instead of status for better user experience
    const [status, setStatus] = useState('');

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    // Handle form submission
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setStatus('Submitting...');

        try {
            // Send the email to the developer

            setStatus('Message sent successfully!');
        } catch (error) {
            setStatus('Failed to send the message. Please try again.');
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900">
            <div className="bg-white dark:bg-gray-800 p-8 rounded-lg shadow-lg w-full max-w-md">
                <h2 className="text-2xl font-bold mb-6 text-gray-800 dark:text-white text-center">
                    Contact the Developer
                </h2>
                <form onSubmit={handleSubmit} className="space-y-6">
                    <div>
                        <label htmlFor="name" className="block text-gray-700 dark:text-gray-200">Name</label>
                        <input
                            type="text"
                            name="name"
                            value={formData.name}
                            onChange={handleChange}
                            required
                            className="w-full px-4 py-2 mt-1 border rounded-md bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-gray-300"
                            placeholder="Your Name"
                        />
                    </div>
                    <div>
                        <label htmlFor="email" className="block text-gray-700 dark:text-gray-200">Email</label>
                        <input
                            type="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            required
                            className="w-full px-4 py-2 mt-1 border rounded-md bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-gray-300"
                            placeholder="Your Email"
                        />
                    </div>
                    <div>
                        <label htmlFor="message" className="block text-gray-700 dark:text-gray-200">Message</label>
                        <textarea
                            name="message"
                            value={formData.message}
                            onChange={handleChange}
                            required
                            className="w-full px-4 py-2 mt-1 border rounded-md bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-gray-300"
                            placeholder="Your message"
                            rows={4}
                        />
                    </div>
                    <div className="text-center">
                        <button
                            type="submit"
                            className="w-full px-4 py-2 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700 focus:ring-4 focus:ring-blue-300"
                        >
                            Send Message
                        </button>
                    </div>
                </form>
                {status && <p className="mt-4 text-center text-sm text-gray-600 dark:text-gray-400">{status}</p>}
            </div>
        </div>
    );
}
