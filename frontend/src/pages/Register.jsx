import React, { useState } from 'react';
import useAuth from '../hooks/useAuth';
import FormInput from '../components/FormInput';
import { useNavigate } from 'react-router-dom';

export default function Register() {
  const { register } = useAuth();
  const navigate = useNavigate();
  const [form, setForm] = useState({ name: '', email: '', password: '' });
  const [errors, setErrors] = useState({});

  // Validaciones simples
  const validate = () => {
    const e = {};
    if (!form.name) e.name = 'Name required';
    if (!/\S+@\S+\.\S+/.test(form.email)) e.email = 'Email inválido';
    if (form.password.length < 4) e.password = 'Min 4 chars';
    setErrors(e);
    return Object.keys(e).length === 0;
  };

  const handleSubmit = async e => {
    e.preventDefault();
    if (!validate()) return;
    try {
      await register(form.name, form.email, form.password);
      navigate('/login');
    } catch {
      setErrors({ general: 'Error al registrar' });
    }
  };

  const handleChange = e =>
    setForm({ ...form, [e.target.name]: e.target.value });

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-4">Register</h2>
      {errors.general && <p className="text-red-600">{errors.general}</p>}
      <form onSubmit={handleSubmit}>
        <FormInput
          label="Name"
          type="text"
          name="name"
          value={form.name}
          onChange={handleChange}
          error={errors.name}
        />
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
        <button className="w-full bg-blue-600 text-white py-2 rounded">
          Sign Up
        </button>
      </form>
    </div>
  );
}
