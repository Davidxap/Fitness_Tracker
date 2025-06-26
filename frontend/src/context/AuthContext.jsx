// frontend/src/context/AuthContext.jsx
import React, { createContext, useState, useEffect } from 'react'
import api from '../api/api'

export const AuthContext = createContext()

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null)
  const [token, setToken] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const t = localStorage.getItem('token')
    const u = localStorage.getItem('user')
    if (t && u) {
      // Validate user still exists
      const parsed = JSON.parse(u)
      api.get(`/users/${parsed.id}`)
        .then(res => setUser(res.data))
        .catch(() => localStorage.clear()) // Clear local storage if user is not found or invalid
        .finally(() => {
          setToken(t)
          setLoading(false)
        })
    } else {
      localStorage.clear() // Clear local storage if no token or user found
      setLoading(false)
    }
  }, []) // Empty dependency array means this effect runs once on mount

  const login = async (email, password) => {
    const res = await api.post('/login', { email, password })
    const { token: t, user: u } = res.data
    localStorage.setItem('token', t)
    localStorage.setItem('user', JSON.stringify(u))
    setToken(t)
    setUser(u)
  }

  const logout = () => {
    localStorage.clear() // Clear all items from local storage
    setToken(null)
    setUser(null)
  }

  const register = async (name, email, password, age, weight) => {
    await api.post('/users', { name, email, password, age, weight })
  }

  return (
    <AuthContext.Provider
      value={{ user, token, loading, login, logout, register }}
    >
      {children}
    </AuthContext.Provider>
  )
}