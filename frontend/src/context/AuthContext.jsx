import { createContext, useState, useEffect } from 'react';
import api from '../api/api';

export const AuthContext = createContext();

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);

  // Al montar, intentamos cargar usuario (opcional: perfil)
  useEffect(() => {
    const token = localStorage.getItem('token');
    const email = localStorage.getItem('email');
    if (token && email) {
      setUser({ email });
    }
  }, []);

  const login = async (email, password) => {
    const res = await api.post('/login', { email, password });
    localStorage.setItem('token', res.data.token);
    localStorage.setItem('email', email);
    setUser({ email });
  };

  const logout = () => {
    localStorage.clear();
    setUser(null);
  };

  const register = async (name, email, password) => {
    await api.post('/users', { name, email, password });
  };

  return (
    <AuthContext.Provider value={{ user, login, logout, register }}>
      {children}
    </AuthContext.Provider>
  );
}
