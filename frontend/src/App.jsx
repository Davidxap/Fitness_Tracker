import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import ProtectedRoute from './components/ProtectedRoute';

// Páginas
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Sessions from './pages/Sessions';
import Exercises from './pages/Exercises';
import SessionExercises from './pages/SessionExercises';

export default function App() {
  return (
    <div className="min-h-screen">
      <Navbar />
      <Routes>
        {/* Públicas */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        {/* Protegidas */}
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
        <Route
          path="/session-exercises"
          element={
            <ProtectedRoute>
              <SessionExercises />
            </ProtectedRoute>
          }
        />
        {/* Fallback */}
        <Route path="*" element={<h2 className="p-6">Page not found</h2>} />
      </Routes>
    </div>
  );
}
