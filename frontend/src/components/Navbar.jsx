import React from 'react';
import { Link } from 'react-router-dom';
import useAuth from '../hooks/useAuth';

export default function Navbar() {
  const { user, logout } = useAuth();
  return (
    <nav className="bg-white p-4 shadow flex justify-between">
      <div>
        <Link className="mr-4" to="/">Dashboard</Link>
        {user && (
          <>
            <Link className="mr-4" to="/sessions">Sessions</Link>
            <Link className="mr-4" to="/exercises">Exercises</Link>
            <Link to="/session-exercises">Details</Link>
          </>
        )}
      </div>
      <div>
        {user ? (
          <>
            <span className="mr-4">{user.email}</span>
            <button onClick={logout} className="text-red-600">Logout</button>
          </>
        ) : (
          <>
            <Link className="mr-4" to="/login">Login</Link>
            <Link to="/register">Register</Link>
          </>
        )}
      </div>
    </nav>
  );
}
