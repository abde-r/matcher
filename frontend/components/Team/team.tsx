'use client';

import Image from 'next/image';
import { FaGithub } from 'react-icons/fa';
import { Si42 } from 'react-icons/si';
0
export default function TeamComponent() {
    const teamMembers = [
        {
            username: 'stronk',
            role: 'Backend Developer',
            imageSrc: 'https://media.licdn.com/dms/image/v2/D4E03AQEAc-cYYzpllQ/profile-displayphoto-shrink_400_400/profile-displayphoto-shrink_400_400/0/1700579771757?e=1737590400&v=beta&t=sb28fcWzC_lnmpEg5P9ZDzhvdEdJ-chejT7fNfmvJ4s',
            description: 'stronk is the backend expert of MatcherX, building robust APIs and handling complex business logic.',
            github: 'https://github.com/abde-r',
            intra42: 'https://profile.intra.42.fr/users/ael-asri'
        },
        {
            username: 'hodor',
            role: 'Frontend Developer',
            imageSrc: 'https://media.licdn.com/dms/image/v2/D4E03AQH1KE32hkBzhQ/profile-displayphoto-shrink_400_400/profile-displayphoto-shrink_400_400/0/1714400858753?e=1737590400&v=beta&t=YvV-3TRkmkazA6CQtGhX44dxr9hrcJwcx_O-e8T0hSk',
            description: 'hodor is responsible for the smooth user interface and experience, making MatcherX look great on all devices.',
            github: 'https://github.com/dependentmadani',
            intra42: 'https://profile.intra.42.fr/users/mbadaoui'
        }
    ];

    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-gray-300 p-20 sm:p-24 overflow-y-scroll">
            <h1 className="text-4xl font-bold mb-10 text-gray-900 dark:text-white">Meet the Team</h1>
            <div className="grid md:grid-cols-2 gap-8">
                {teamMembers.map((member, id) => (
                    <CardInfo
                        key={id}
                        username={member.username}
                        role={member.role}
                        imageSrc={member.imageSrc}
                        description={member.description}
                        github={member.github}
                        intra42={member.intra42}
                    />
                ))}
            </div>
        </div>
    );
}

function CardInfo({ username, role, imageSrc, description, github, intra42 }: { username: string, role: string, imageSrc: string, description: string, github: string, intra42: string }) {
    return (
        <div className="max-w-sm rounded overflow-hidden shadow-lg bg-white dark:bg-gray-800 flex flex-col justify-between h-full">
            <Image
                className="w-full"
                src={imageSrc}
                alt={`${username} photo`}
                width={400}
                height={400}
            />
            <div className="px-6 py-4">
                <div className="font-bold text-xl mb-2 text-gray-900 dark:text-white">
                    {username}
                </div>
                <p className="text-gray-700 dark:text-gray-200 text-base">{role}</p>
                <p className="text-gray-600 dark:text-gray-400 mt-2">{description}</p>
            </div>
            <div className="flex justify-end space-x-4 mt-auto px-6 pb-4">
                <a href={github} target="_blank" rel="noopener noreferrer">
                    <FaGithub className="text-2xl sm:text-3xl hover:text-gray-700 dark:hover:text-gray-300" />
                </a>
                <a href={intra42} target="_blank" rel="noopener noreferrer">
                    <Si42 className="text-2xl sm:text-3xl hover:text-gray-700 dark:hover:text-gray-300" />
                </a>
            </div>
        </div>
    );
}