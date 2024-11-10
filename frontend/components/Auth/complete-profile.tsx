'use client';

import React, { useState } from 'react';
import { Card } from '@/components/ui/card';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { CalendarIcon } from '@heroicons/react/outline';

const formatDate = (date: Date | null): string => {
    if (!date || isNaN(date.getTime()))
        return '';
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const year = date.getFullYear();
    return `${month}/${day}/${year}`;
};

const CustomInput = React.forwardRef(({ value, onClick, onChange }: {value: Date | null, onClick?: any, onChange: any}, ref: any) => (
  <div className="relative w-full">
    <input
      type='date'
      value={formatDate(value)}
      onClick={onClick}
      onChange={onChange}
      ref={ref}
      className="w-full p-2 border rounded"
      placeholder="Select your birthday"
    />
    <CalendarIcon className="absolute right-3 top-1/2 transform -translate-y-1/2 w-6 h-6 text-gray-400 pointer-events-none" />
  </div>
));

export default function CompleteProfile() {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [birthday, setBirthday] = useState<Date | null>(null);
  const [preferences, setPreferences] = useState<string[]>([]);
  const [interests, setInterests] = useState<string[]>([]);

  const handleAddInterest = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && e.currentTarget.value.trim()) {
      setInterests([...interests, `#${e.currentTarget.value.trim()}`]);
      e.currentTarget.value = '';
    }
  };

  const handleDateChange = (date: Date | null) => {
    console.log("date on change", date);
    setBirthday(date);
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
  }

  return (
    <div className="block p-4 sm:p-6 lg:p-8">
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
            <DatePicker
                selected={birthday}
                onChange={handleDateChange}
                customInput={<CustomInput value={birthday} onChange={handleInputChange} />}
            />
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 mb-1">Preferences</label>
          <select
            multiple
            value={preferences}
            onChange={(e) =>
              setPreferences(Array.from(e.target.selectedOptions, option => option.value))
            }
            className="w-full p-2 border rounded"
          >
            <option value="preference1">Preference 1</option>
            <option value="preference2">Preference 2</option>
            <option value="preference3">Preference 3</option>
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
          <div className="mt-2 flex flex-wrap gap-2">
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