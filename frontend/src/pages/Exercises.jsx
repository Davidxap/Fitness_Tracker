import React, { useEffect, useState } from 'react'
import api from '../api/api'

export default function Exercises() {
  const [list, setList] = useState([])

  useEffect(() => {
    api.get('/exercises').then(res => setList(res.data))
  }, [])

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Exercises Library</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white rounded shadow">
          <thead className="bg-gray-100">
            <tr>
              <th className="p-2 text-left">ID</th>
              <th className="p-2 text-left">Name</th>
              <th className="p-2 text-left">Muscle Group</th>
              <th className="p-2 text-left">Description</th>
            </tr>
          </thead>
          <tbody>
            {list.map(ex => (
              <tr key={ex.id} className="border-t">
                <td className="p-2">{ex.id}</td>
                <td className="p-2">{ex.name}</td>
                <td className="p-2">{ex.muscle_group}</td>
                <td className="p-2">{ex.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}
