// frontend/src/App.jsx
import React from 'react'
import { Routes, Route, Navigate } from 'react-router-dom'
import useAuth from './hooks/useAuth'

import Navbar from './components/Navbar'
import ProtectedRoute from './components/ProtectedRoute'
import Login from './pages/Login'
import Register from './pages/Register'
import Dashboard from './pages/Dashboard'
import Sessions from './pages/Sessions'
import Exercises from './pages/Exercises'

export default function App() {
  const { user, loading } = useAuth()

  if (loading) {
    return <p className="p-6">Loading app...</p>
  }

  return (
    <div className="min-h-screen">
      <Navbar />
      <Routes>
        {/* Registration always accessible */}
        <Route path="/register" element={<Register />} />

        {/* Login only if user exists */}
        <Route
          path="/login"
          element={user ? <Navigate to="/" replace /> : <Login />}
        />

        {/* Protected routes */}
        <Route
          path="/"
          element={
            <ProtectedRoute>
              <Dashboard />
            </ProtectedRoute>
          }
        />
        <Route
          path="/sessions"
          element={
            <ProtectedRoute>
              <Sessions />
            </ProtectedRoute>
          }
        />
        <Route
          path="/exercises"
          element={
            <ProtectedRoute>
              <Exercises />
            </ProtectedRoute>
          }
        />

        {/* Any other route */}
        <Route
          path="*"
          element={<Navigate to={user ? '/' : '/register'} replace />}
        />
      </Routes>
    </div>
  )
}