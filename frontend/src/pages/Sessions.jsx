import React, { useEffect, useState } from 'react';
import api from '../api/api';

export default function Sessions() {
  const [list, setList] = useState([]);

  useEffect(() => {
    api.get('/sessions').then(res => setList(res.data));
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-3xl mb-4">Workout Sessions</h1>
      <table className="w-full bg-white rounded shadow">
        <thead className="bg-gray-100">
          <tr>
            <th className="p-2">ID</th>
            <th className="p-2">User ID</th>
            <th className="p-2">Date</th>
            <th className="p-2">Duration</th>
          </tr>
        </thead>
        <tbody>
          {list.map(s => (
            <tr key={s.id} className="border-t">
              <td className="p-2">{s.id}</td>
              <td className="p-2">{s.user_id}</td>
              <td className="p-2">{s.date}</td>
              <td className="p-2">{s.duration_minutes} min</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
