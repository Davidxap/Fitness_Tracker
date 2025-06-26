// frontend/src/pages/Register.jsx
import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import useAuth from '../hooks/useAuth'
import FormInput from '../components/FormInput'

export default function Register() {
  const { user, register } = useAuth()
  const navigate = useNavigate()

  // If use exist → dashboard
  useEffect(() => {
    if (user) {
      navigate('/')
    }
  }, [user, navigate])

  const [form, setForm] = useState({
    name: '',
    email: '',
    password: '',
    age: '',
    weight: ''
  })
  const [errors, setErrors] = useState({})

  const validate = () => {
    const e = {}
    if (!form.name) e.name = 'Name is required'
    if (!/\S+@\S+\.\S+/.test(form.email)) e.email = 'Invalid email'
    if (form.password.length < 4) e.password = 'Min 4 characters'
    if (!form.age || form.age <= 0) e.age = 'Valid age required'
    if (!form.weight || form.weight <= 0) e.weight = 'Valid weight required'
    setErrors(e)
    return Object.keys(e).length === 0
  }

  const handleSubmit = async ev => {
    ev.preventDefault()
    if (!validate()) return
    try {
      await register(
        form.name,
        form.email,
        form.password,
        Number(form.age),
        Number(form.weight)
      )
      navigate('/login')
    } catch {
      setErrors({ general: 'Registration failed' })
    }
  }

  const handleChange = e => {
    const { name, value } = e.target
    setForm(f => ({ ...f, [name]: value }))
  }

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-4">Sign Up</h2>
      {errors.general && (
        <p className="text-red-600 mb-2">{errors.general}</p>
      )}
      <form onSubmit={handleSubmit}>
        <FormInput
          label="Name"
          name="name"
          type="text"
          value={form.name}
          onChange={handleChange}
          error={errors.name}
        />
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
        <FormInput
          label="Age"
          name="age"
          type="number"
          value={form.age}
          onChange={handleChange}
          error={errors.age}
        />
        <FormInput
          label="Weight (kg)"
          name="weight"
          type="number"
          value={form.weight}
          onChange={handleChange}
          error={errors.weight}
        />
        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded transition transform active:scale-95"
        >
          Register
        </button>
      </form>
    </div>
  )
}
