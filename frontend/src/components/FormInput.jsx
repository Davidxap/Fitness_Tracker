import React from 'react'

/**
 * Input with label and error display.
 */
export default function FormInput({
  label,
  type,
  name,
  value,
  onChange,
  error
}) {
  return (
    <div className="mb-4">
      <label className="block mb-1 font-medium">{label}</label>
      <input
        className="w-full p-2 border rounded"
        type={type}
        name={name}
        value={value}
        onChange={onChange}
        autoComplete="off" // disables suggestions
      />
      {error && <p className="text-red-600 text-sm mt-1">{error}</p>}
    </div>
  )
}