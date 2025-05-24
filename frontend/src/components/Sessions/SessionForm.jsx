import React, { useState, useEffect } from 'react'
import FormInput from '../FormInput'
import ExerciseEntry from './ExerciseEntry'
import api from '../../api/api'

export default function SessionForm({ initialData, onSave, onCancel }) {
  const [form, setForm] = useState({
    date: '',
    duration_minutes: '',
    observations: ''
  })
  const [entries, setEntries] = useState([])
  const [options, setOptions] = useState([])
  const [errors, setErrors] = useState({})

  // Carga ejercicios BD y precarga form/entries si editing
  useEffect(() => {
    api.get('/exercises').then(res => {
      setOptions(res.data.map(e => ({ id: e.id, name: e.name })))
    })
    if (initialData) {
      setForm({
        date: initialData.date,
        duration_minutes: initialData.duration_minutes,
        observations: initialData.observations || ''
      })
      setEntries(
        initialData.exercises.map(se => ({
          exercise_id: se.exercise_id,
          sets: se.sets,
          reps: se.reps,
          weight: se.weight
        }))
      )
    } else {
      setForm({ date: '', duration_minutes: '', observations: '' })
      setEntries([])
    }
  }, [initialData])

  const validate = () => {
    const e = {}
    if (!form.date) e.date = 'Date is required'
    if (!form.duration_minutes || form.duration_minutes <= 0)
      e.duration_minutes = 'Valid duration required'
    if (entries.length === 0) e.entries = 'Add at least one exercise'
    setErrors(e)
    return Object.keys(e).length === 0
  }

  const handleSubmit = e => {
    e.preventDefault()
    if (!validate()) return
    onSave({
      date: form.date,
      duration_minutes: Number(form.duration_minutes),
      observations: form.observations.trim(),
      exercises: entries.map(ent => ({
        exercise_id: Number(ent.exercise_id),
        sets: Number(ent.sets),
        reps: Number(ent.reps),
        weight: Number(ent.weight)
      }))
    })
  }

  const handleChange = e => {
    const { name, value } = e.target
    setForm(f => ({ ...f, [name]: value }))
  }

  const addEntry = () => {
    setEntries(es => [...es, { exercise_id: '', sets: '', reps: '', weight: '' }])
  }
  const updateEntry = (idx, e) => {
    const { name, value } = e.target
    setEntries(es => {
      const nxt = [...es]
      nxt[idx] = { ...nxt[idx], [name]: value }
      return nxt
    })
  }
  const removeEntry = idx => {
    setEntries(es => es.filter((_, i) => i !== idx))
  }

  return (
    <div className="bg-white p-6 rounded shadow mb-6">
      <h2 className="text-xl font-semibold mb-4">
        {initialData ? 'Edit Session' : 'New Session'}
      </h2>
      {errors.entries && (
        <p className="text-red-600 mb-2">{errors.entries}</p>
      )}
      <form onSubmit={handleSubmit}>
        <FormInput
          label="Date"
          type="date"
          name="date"
          value={form.date}
          onChange={handleChange}
          error={errors.date}
        />
        <FormInput
          label="Duration (minutes)"
          type="number"
          name="duration_minutes"
          value={form.duration_minutes}
          onChange={handleChange}
          error={errors.duration_minutes}
        />
        <FormInput
          label="Observations"
          type="text"
          name="observations"
          value={form.observations}
          onChange={handleChange}
        />

        <div className="mt-4">
          <h3 className="font-medium mb-2">Exercises</h3>
          {entries.map((ent, idx) => (
            <ExerciseEntry
              key={idx}
              options={options}
              data={ent}
              onChange={e => updateEntry(idx, e)}
              onRemove={() => removeEntry(idx)}
            />
          ))}
          <button
            type="button"
            onClick={addEntry}
            className="mb-4 bg-indigo-600 text-white px-3 py-1 rounded transition transform active:scale-95"
          >
            + Add Exercise
          </button>
        </div>

        <div className="flex space-x-3 mt-4">
          <button
            type="submit"
            className="bg-blue-600 text-white px-4 py-2 rounded transition transform active:scale-95"
          >
            Save Session
          </button>
          <button
            type="button"
            onClick={onCancel}
            className="bg-gray-400 text-white px-4 py-2 rounded transition transform active:scale-95"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  )
}
