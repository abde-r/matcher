'use client'

import { useState } from 'react'
import { motion } from 'framer-motion'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { HeartIcon, XIcon } from 'lucide-react'

export function AuthCardComponent() {
  const [isLogin, setIsLogin] = useState(true)
  const [dragX, setDragX] = useState(0)

  const handleDragEnd = (event, info) => {
    if (info.offset.x > 100) {
      setIsLogin(false)
    } else if (info.offset.x < -100) {
      setIsLogin(true)
    }
    setDragX(0)
  }

  const handleSubmit = (event) => {
    event.preventDefault()
    // Handle form submission here
    console.log('Form submitted:', isLogin ? 'Login' : 'Sign Up')
  }

  return (
    <div className="flex justify-center items-center min-h-screen bg-gradient-to-r from-gray-800 to-gray-900">
      <motion.div
        drag="x"
        dragConstraints={{ left: 0, right: 0 }}
        onDragEnd={handleDragEnd}
        animate={{ x: dragX }}
        style={{ x: dragX }}
      >
        <Card className="w-[350px] overflow-hidden bg-white">
          <CardHeader className="relative">
            <div className="absolute top-4 left-4 text-gray-600">
              <XIcon size={24} />
            </div>
            <div className="absolute top-4 right-4 text-gray-600">
              <HeartIcon size={24} />
            </div>
            <CardTitle className="text-2xl font-bold text-center text-gray-800">
              {isLogin ? 'Welcome Back' : 'Join the Fun'}
            </CardTitle>
            <CardDescription className="text-center text-gray-600">
              Swipe to {isLogin ? 'sign up' : 'login'}
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit}>
              <div className="space-y-4">
                {!isLogin && (
                  <div className="space-y-2">
                    <Label htmlFor="name" className="text-gray-700">Name</Label>
                    <Input id="name" placeholder="Enter your name" required className="bg-gray-100 border-gray-300" />
                  </div>
                )}
                <div className="space-y-2">
                  <Label htmlFor="email" className="text-gray-700">Email</Label>
                  <Input id="email" type="email" placeholder="Enter your email" required className="bg-gray-100 border-gray-300" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="password" className="text-gray-700">Password</Label>
                  <Input id="password" type="password" placeholder="Enter your password" required className="bg-gray-100 border-gray-300" />
                </div>
              </div>
              <Button className="w-full mt-6 bg-gray-800 hover:bg-gray-900 text-white">
                {isLogin ? 'Login' : 'Sign Up'}
              </Button>
            </form>
          </CardContent>
          <CardFooter className="flex justify-center">
            <p className="text-sm text-gray-600">
              {isLogin ? "Don't have an account?" : "Already have an account?"}
              <button
                onClick={() => setIsLogin(!isLogin)}
                className="ml-1 text-gray-800 hover:underline focus:outline-none"
              >
                {isLogin ? 'Sign Up' : 'Login'}
              </button>
            </p>
          </CardFooter>
        </Card>
      </motion.div>
    </div>
  )
}