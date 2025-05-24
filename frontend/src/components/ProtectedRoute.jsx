// frontend/src/components/ProtectedRoute.jsx
import React from 'react'
import { Navigate } from 'react-router-dom'
import useAuth from '../hooks/useAuth'

export default function ProtectedRoute({ children }) {
  const { user, loading } = useAuth()

  if (loading) {
    return <p className="p-6">Loading...</p>
  }
  if (!user) {
    return <Navigate to="/login" replace />
  }
  return children
}
