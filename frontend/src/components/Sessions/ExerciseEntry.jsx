// frontend/src/components/Sessions/ExerciseEntry.jsx
import React from 'react'
import FormInput from '../FormInput'

/**
 * Un único bloque de selección de ejercicio + sets/reps/weight
 * @param {Array} options Lista de ejercicios {id,name}
 * @param {Object} data {exercise_id, sets, reps, weight}
 * @param {Function} onChange callback cuando cambia un campo
 * @param {Function} onRemove callback para eliminar este entry
 */
export default function ExerciseEntry({ options, data, onChange, onRemove }) {
  return (
    <div className="grid grid-cols-5 gap-2 items-end mb-4">
      {/* Dropdown de ejercicios */}
      <div className="col-span-2">
        <label className="block text-sm">Exercise</label>
        <select
          name="exercise_id"
          value={data.exercise_id}
          onChange={onChange}
          className="w-full p-2 border rounded"
        >
          <option value="">Select exercise</option>
          {options.map(opt => (
            <option key={opt.id} value={opt.id}>
              {opt.name}
            </option>
          ))}
        </select>
      </div>
      <FormInput
        label="Sets"
        type="number"
        name="sets"
        value={data.sets}
        onChange={onChange}
      />
      <FormInput
        label="Reps"
        type="number"
        name="reps"
        value={data.reps}
        onChange={onChange}
      />
      <FormInput
        label="Weight(KG)"
        type="number"
        name="weight"
        value={data.weight}
        onChange={onChange}
      />
      <button
        type="button"
        onClick={onRemove}
        className="bg-red-500 text-white px-2 py-1 rounded text-sm"
      >
        Remove
      </button>
    </div>
  )
}
