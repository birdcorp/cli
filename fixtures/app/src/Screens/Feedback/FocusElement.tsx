import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function FocusElement() {
  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Focus Element',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Focus</h3>

      <p className="text-gray-300 mb-8">
        Triggers focus of dom element by query selector
      </p>

      <div className="max-w-md mx-auto mt-10 mb-10">
        <label
          htmlFor="inputField"
          className="block text-sm font-medium text-gray-300"
        >
          Input Label
        </label>
        <input
          id="inputField"
          type="text"
          placeholder="Enter text here..."
          className="mt-1 block w-full px-4 py-2 border border-neutral-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-white-500 focus:border-white bg-neutral-900 text-neutral-100 placeholder-neutral-500 transition duration-150 ease-in-out"
        />
      </div>

      <button
        onClick={() => {
          window.miniProgram?.call('focus', {
            selector: '#inputField',
          })
        }}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded mr-3"
      >
        Focus
      </button>

      <button
        onClick={() => {
          window.miniProgram?.call('blur', {
            selector: '#inputField',
          })
        }}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Blur
      </button>

      <h3 className="text-white mt-8 font-bold">Code</h3>

      <p>Focus</p>
      <HighlightedCode
        code={`window.miniProgram?.call('focus', {
  selector: '#inputField',
})`}
      />

      <p>Blur</p>
      <HighlightedCode
        code={`window.miniProgram?.call('blur', {
  selector: '#inputField',
})`}
      />
    </div>
  )
}
