import React from 'react'
import useAuth from '../hooks/useAuth'

export default function Dashboard() {
  const { user } = useAuth()

  if (!user) return <p className="p-6">Loading...</p>

  // Format created_at (YYYY-MM-DD)
  const regDate = user.created_at.slice(0, 10)

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-2">Welcome, {user.name}!</h1>
      <p className="mb-1"><strong>Email:</strong> {user.email}</p>
      <p className="mb-1"><strong>Age:</strong> {user.age}</p>
      <p className="mb-1"><strong>Weight:</strong> {user.weight} kg</p>
      <p className="text-sm text-gray-600">
        <strong>Member since:</strong> {regDate}
      </p>
    </div>
  )
}