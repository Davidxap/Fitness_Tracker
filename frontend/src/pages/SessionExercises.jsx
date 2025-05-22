import React, { useEffect, useState } from 'react';
import api from '../api/api';

export default function SessionExercises() {
  const [list, setList] = useState([]);

  useEffect(() => {
    api.get('/session-exercises').then(res => setList(res.data));
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-3xl mb-4">Session Exercises</h1>
      <table className="w-full bg-white rounded shadow">
        <thead className="bg-gray-100">
          <tr>
            <th className="p-2">Session</th>
            <th className="p-2">Exercise</th>
            <th className="p-2">Sets</th>
            <th className="p-2">Reps</th>
            <th className="p-2">Weight</th>
          </tr>
        </thead>
        <tbody>
          {list.map(se => (
            <tr key={se.id} className="border-t">
              <td className="p-2">{se.session_id}</td>
              <td className="p-2">{se.exercise_id}</td>
              <td className="p-2">{se.sets}</td>
              <td className="p-2">{se.reps}</td>
              <td className="p-2">{se.weight} kg</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
