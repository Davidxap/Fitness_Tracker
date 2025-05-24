// frontend/src/pages/Login.jsx
import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import useAuth from '../hooks/useAuth'
import FormInput from '../components/FormInput'

export default function Login() {
  const { user, login } = useAuth()
  const navigate = useNavigate()

  // Si user ya existe, redirigir automáticamente
  useEffect(() => {
    if (user) {
      navigate('/')
    }
  }, [user, navigate])

  const [form, setForm] = useState({ email: '', password: '' })
  const [errors, setErrors] = useState({})

  const validate = () => {
    const e = {}
    if (!form.email) e.email = 'Email es requerido'
    if (!form.password) e.password = 'Password es requerido'
    setErrors(e)
    return Object.keys(e).length === 0
  }

  const handleSubmit = async e => {
    e.preventDefault()
    if (!validate()) return
    try {
      await login(form.email, form.password)
      // login() actualiza `user` en el context y dispara el useEffect
    } catch {
      setErrors({ general: 'Credenciales inválidas' })
    }
  }

  const handleChange = e =>
    setForm(f => ({ ...f, [e.target.name]: e.target.value }))

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-4">Iniciar Sesión</h2>
      {errors.general && (
        <p className="text-red-600 text-center">{errors.general}</p>
      )}
      <form onSubmit={handleSubmit}>
        <FormInput
          label="Email"
          type="email"
          name="email"
          value={form.email}
          onChange={handleChange}
          error={errors.email}
        />
        <FormInput
          label="Password"
          type="password"
          name="password"
          value={form.password}
          onChange={handleChange}
          error={errors.password}
        />
        <button className="w-full bg-green-600 text-white py-2 rounded">
          Entrar
        </button>
      </form>
    </div>
  )
}
