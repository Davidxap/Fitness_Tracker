import React, { useEffect, useState } from 'react';
import api from '../api/api';

export default function Exercises() {
  const [list, setList] = useState([]);

  useEffect(() => {
    api.get('/exercises').then(res => setList(res.data));
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-3xl mb-4">Exercises</h1>
      <ul className="space-y-2">
        {list.map(e => (
          <li key={e.id} className="bg-white p-4 rounded shadow">
            <strong>{e.name}</strong> – {e.muscle_group}
          </li>
        ))}
      </ul>
    </div>
  );
}
