import React, { useEffect, useState } from 'react'
import api from '../api/api'
import useAuth from '../hooks/useAuth'
import SessionForm from '../components/Sessions/SessionForm'
import SessionList from '../components/Sessions/SessionList'

export default function Sessions() {
  const { user } = useAuth()
  const [sessions, setSessions] = useState([])
  const [editing, setEditing] = useState(null)
  const [showForm, setShowForm] = useState(false)
  const [message, setMessage] = useState('')

  // Carga ejercicios + sesiones + relaciones
  const loadData = async () => {
    try {
      // 1) Obtener todos los ejercicios y construir exMap
      const exRes = await api.get('/exercises')
      const exMap = {}
      exRes.data.forEach(e => {
        exMap[e.id] = e.name
      })

      // 2) Obtener sesiones y session-exercises
      const [sRes, seRes] = await Promise.all([
        api.get('/sessions'),
        api.get('/session-exercises')
      ])

      // 3) Filtrar solo las sesiones del user
      const own = sRes.data.filter(s => s.user_id === user.id)

      // 4) Combinar cada sesión con sus ejercicios y asignar exercise_name
      const combined = own.map(s => {
        // obtenemos solo los relationships de esta sesión
        const related = seRes.data.filter(se => se.session_id === s.id)
        // mapeamos cada uno añadiendo exercise_name
        const exercises = related.map(se => ({
          id: se.id,
          exercise_id: se.exercise_id,
          exercise_name: exMap[se.exercise_id] || 'Unknown',
          sets: se.sets,
          reps: se.reps,
          weight: se.weight
        }))
        return { ...s, exercises }
      })

      setSessions(combined)
    } catch (err) {
      console.error('Error loading sessions:', err)
    }
  }

  useEffect(() => {
    if (user) loadData()
  }, [user])

  const handleCreate = () => {
    setEditing(null)
    setShowForm(true)
  }

  const handleSave = async data => {
    try {
      let sessionId
      if (editing) {
        // Update existing
        await api.put(`/sessions/${editing.id}`, {
          name: data.name,
          date: data.date,
          duration_minutes: data.duration_minutes,
          observations: data.observations,
          user_id: user.id
        })
        sessionId = editing.id
        // Borrar anteriores session-exercises
        await Promise.all(
          editing.exercises.map(e =>
            api.delete(`/session-exercises/${e.id}`)
          )
        )
      } else {
        // Create new
        const r = await api.post('/sessions', {
          name: data.name,
          date: data.date,
          duration_minutes: data.duration_minutes,
          observations: data.observations,
          user_id: user.id
        })
        sessionId = r.data.id
      }
      // Guardar los nuevos session-exercises
      await Promise.all(
        data.exercises.map(ent =>
          api.post('/session-exercises', {
            session_id: sessionId,
            exercise_id: ent.exercise_id,
            sets: ent.sets,
            reps: ent.reps,
            weight: ent.weight
          })
        )
      )
      setMessage('Session saved successfully!')
      setTimeout(() => setMessage(''), 3000)
      setShowForm(false)
      setEditing(null)
      await loadData()
    } catch (err) {
      console.error('Error saving session:', err)
    }
  }

  const handleEdit = s => {
    setEditing(s)
    setShowForm(true)
  }

  const handleDelete = async id => {
    if (!window.confirm('Delete this session?')) return
    try {
      // borrar session-exercises luego la sesión
      const toDel = sessions.find(s => s.id === id).exercises
      await Promise.all(toDel.map(e => api.delete(`/session-exercises/${e.id}`)))
      await api.delete(`/sessions/${id}`)
      await loadData()
    } catch (err) {
      console.error(err)
    }
  }

  if (!user) return <p className="p-6">Loading...</p>

  return (
    <div className="p-6">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-3xl font-bold">Sessions</h1>
        <button
          onClick={handleCreate}
          className="bg-green-600 text-white px-4 py-2 rounded transition active:scale-95"
        >
          + New Session
        </button>
      </div>

      {message && (
        <div className="mb-4 p-3 bg-green-100 text-green-800 rounded">
          {message}
        </div>
      )}

      {showForm && (
        <SessionForm
          key={editing ? editing.id : 'new'}
          initialData={editing}
          onSave={handleSave}
          onCancel={() => {
            setShowForm(false)
            setEditing(null)
          }}
        />
      )}

      <SessionList
        sessions={sessions}
        formatDate={d => d.slice(0, 10)}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
    </div>
  )
}
