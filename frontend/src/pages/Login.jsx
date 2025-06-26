import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import useAuth from '../hooks/useAuth'
import FormInput from '../components/FormInput'

export default function Login() {
  const { user, login } = useAuth()
  const navigate = useNavigate()

  // Si ya está autenticado, ir a dashboard
  useEffect(() => {
    if (user) {
      navigate('/')
    }
  }, [user, navigate])

  const [form, setForm] = useState({ email: '', password: '' })
  const [errors, setErrors] = useState({})

  const validate = () => {
    const e = {}
    if (!form.email) e.email = 'Email is required'
    if (!form.password) e.password = 'Password is required'
    setErrors(e)
    return Object.keys(e).length === 0
  }

  const handleSubmit = async ev => {
    ev.preventDefault()
    if (!validate()) return

    try {
      await login(form.email, form.password)
    // If already authenticated, go to dashboard      navigate('/')
    } catch (err) {
      console.error(err)
      setErrors({ general: 'Invalid credentials' })
    }
  }

  const handleChange = ev => {
    const { name, value } = ev.target
    setForm(f => ({ ...f, [name]: value }))
  }

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-4">Sign In</h2>
      {errors.general && (
        <p className="text-red-600 mb-2">{errors.general}</p>
      )}
      <form onSubmit={handleSubmit}>
        <FormInput
          label="Email"
          name="email"
          type="email"
          value={form.email}
          onChange={handleChange}
          error={errors.email}
        />
        <FormInput
          label="Password"
          name="password"
          type="password"
          value={form.password}
          onChange={handleChange}
          error={errors.password}
        />
        <button
          type="submit"
          className="w-full bg-green-600 text-white py-2 rounded transition active:scale-95"
        >
          Login
        </button>
      </form>
    </div>
  )
}
