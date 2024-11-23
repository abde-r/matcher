import Image from 'next/image';
import { useEffect, useState } from 'react';

export default function UserDropdown() {
    const [dropdownOpen, setDropdownOpen] = useState(false);
    const [isSmallScreen, setIsSmallScreen] = useState(false);
  
    const toggleDropdown = () => {
      setDropdownOpen(!dropdownOpen);
    };
  
    // Detect window size to adjust dropdown responsiveness
    useEffect(() => {
      const handleResize = () => {
        setIsSmallScreen(window.innerWidth < 1540);
      };
  
      window.addEventListener('resize', handleResize);
  
      handleResize();
  
      return () => {
        window.removeEventListener('resize', handleResize);
      };
    }, []);

  return (
    <div className="relative">
      <Image
        id="avatarButton"
        onClick={toggleDropdown}
        width={100}
        height={100}
        className="w-10 h-10 rounded-full cursor-pointer border-2 border-black hover:border-gray-200 dark:hover:border-gray-700"
        src="/standard_profile_image.jpg"
        alt="User dropdown"
      />

      {dropdownOpen && (
        <div
            className={`absolute z-10 ${
            isSmallScreen ? 'right-0' : 'left-0'
            } bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700 dark:divide-gray-600`}
        >
          <div className="px-4 py-3 text-sm text-gray-900 dark:text-white">
            <div>Badaoui Madani</div>
            <div className="font-medium truncate">madani@madani.com</div>
          </div>
          <ul className="py-2 text-sm text-gray-700 dark:text-gray-200">
            <li>
              <a href="/profile" className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">
                Profile
              </a>
            </li>
            <li>
              <a href="/settings" className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">
                Settings
              </a>
            </li>
          </ul>
          <div className="py-1">
            <a
              href="/signout"
              className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
            >
              Sign out
            </a>
          </div>
        </div>
      )}
    </div>
  );
}
