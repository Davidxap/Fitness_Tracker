import React, { useState } from 'react';
import useAuth from '../hooks/useAuth';
import FormInput from '../components/FormInput';
import { useNavigate } from 'react-router-dom';

export default function Login() {
  const { login } = useAuth();
  const navigate = useNavigate();
  const [form, setForm] = useState({ email: '', password: '' });
  const [errors, setErrors] = useState({});

  const validate = () => {
    const e = {};
    if (!form.email) e.email = 'Email required';
    if (!form.password) e.password = 'Password required';
    setErrors(e);
    return Object.keys(e).length === 0;
  };

  const handleSubmit = async e => {
    e.preventDefault();
    if (!validate()) return;
    try {
      await login(form.email, form.password);
      navigate('/');
    } catch {
      setErrors({ general: 'Credenciales inválidas' });
    }
  };

  const handleChange = e =>
    setForm({ ...form, [e.target.name]: e.target.value });

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-4">Login</h2>
      {errors.general && <p className="text-red-600">{errors.general}</p>}
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
          Sign In
        </button>
      </form>
    </div>
  );
}
