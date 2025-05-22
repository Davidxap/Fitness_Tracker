import React, { useEffect, useState } from 'react';
import api from '../api/api';

export default function Dashboard() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    api.get('/users').then(res => setUsers(res.data));
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-3xl mb-4">Users</h1>
      <ul className="space-y-2">
        {users.map(u => (
          <li key={u.id} className="bg-white p-4 rounded shadow">
            {u.name} – {u.email}
          </li>
        ))}
      </ul>
    </div>
  );
}
