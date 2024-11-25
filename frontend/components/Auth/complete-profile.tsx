'use client';

import React, { useState } from 'react';
import { Card } from '@/components/ui/card';
import 'react-datepicker/dist/react-datepicker.css';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

const formatDate = (date: Date | null): string => {
    if (!date)
        return '';
    const newDate = new Date(date);
    const day = String(newDate.getDate()).padStart(2, '0');
    const month = String(newDate.getMonth() + 1).padStart(2, '0');
    const year = newDate.getFullYear();
    return `${month}/${day}/${year}`;
};

const CustomInput = React.forwardRef(({ onClick, onChange }: {onClick?: any, onChange: any}, ref: any) => (
  <div className="relative w-full">
    <input
      multiple
      type='date'
      onClick={onClick}
      onChange={onChange}
      ref={ref}
      className="w-full p-2 border rounded"
      placeholder="Select your birthday"
    />
  </div>
));

export default function CompleteProfile() {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [birthday, setBirthday] = useState<Date | null>(null);
  const [gender, setGender] = useState<string>('Men');
  const [preferences, setPreferences] = useState<string>('Women');
  const [interests, setInterests] = useState<string[]>([]);
  const router = useRouter();

  const handleAddInterest = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && e.currentTarget.value.trim()) {
      setInterests([...interests, `#${e.currentTarget.value.trim()}`]);
      e.currentTarget.value = '';
    }
  };

  const handleDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const inputValue = e.target.value;
    const date = new Date(inputValue);
    if (!isNaN(date.getTime())) {
        setBirthday(date);
      }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const inputValue = e.target.value;
    const date = new Date(inputValue);
    if (!isNaN(date.getTime())) {
      setBirthday(date);
    }
  };

  const handleSaveProfile = () => {
    // use the api to save the profile infos
    fetch(`${process.env.API_URL}/api/v1/users/proceed-registration`, {
      method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({
          query:`
            mutation ProceedRegistrationUser($input: ProceedRegistrationUserPayload!) { proceedRegistrationUser(input: $input) { first_name last_name birthday gender preferences} }
          `,
          variables: {
            input: {
              first_name: firstName,
              last_name: lastName,
              birthday: birthday,
              gender: gender,
              preferences: preferences
            }
          }
        })
    })
    .then((res) => {
      if (res.ok) {
        router.push('/');
      } else {
        console.log('Error:', res)
      }
    })
  }

  return (
    <div className="p-4 sm:p-6 lg:p-8 space-y-4">
      <a className="flex items-center justify-center">
        <Image 
          src="/logo.png"
          width={32}
          height={32}
          className="h-8"
          alt="MatcherX logo" />
      </a>
      <Card className="max-w-lg mx-auto p-6 w-full sm:w-80 lg:w-96">
        <h2 className="text-2xl font-bold mb-6">Complete Your Profile</h2>
        <div className="mb-4">
          <label className="block text-gray-700 mb-1">First Name</label>
          <input
            type="text"
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
            className="w-full p-2 border rounded"
          />
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 mb-1">Last Name</label>
          <input
            type="text"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
            className="w-full p-2 border rounded"
          />
        </div>

        <div className="mb-4 relative">
            <label className="block text-gray-700 mb-1">Birthday</label>
            <CustomInput onChange={handleInputChange} onClick={handleDateChange} />
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 mb-1">Gender</label>
          <select
            value={gender}
            onChange={(e) =>
              setGender(e.target.value)
            }
            className="w-full p-2 border rounded"
          >
            <option value="men">Men</option>
            <option value="women">Women</option>
          </select>
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 mb-1">Preferences</label>
          <select
            value={preferences}
            onChange={(e) =>
              setPreferences(e.target.value)
            }
            className="w-full p-2 border rounded"
          >
            <option value="women">Women</option>
            <option value="men">Men</option>
          </select>
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 mb-1">Interests</label>
          <input
            type="text"
            onKeyDown={handleAddInterest}
            className="w-full p-2 border rounded"
            placeholder="Type interest and press Enter"
          />
          <div className="mt-2 flex flex-wrap gap-2 max-h-16 overflow-y-auto">
            {interests.map((interest, index) => (
              <span
                key={index}
                className="bg-gray-200 px-3 py-1 rounded-full text-sm text-gray-700"
              >
                {interest}
              </span>
            ))}
          </div>
        </div>

        <button className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 w-full sm:w-auto" onClick={handleSaveProfile}>
          Save Profile
        </button>
      </Card>
    </div>
  );
}