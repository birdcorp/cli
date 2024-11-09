import React, { useEffect, useCallback, useState } from 'react'
import { HighlightedCode } from '../../Components'

export default function Picker() {
  const [selectedOption, setSelectedOption] = useState()

  const trigger = useCallback(() => {
    window.miniProgram?.call(
      'showPicker',
      {
        options: ['Small', 'Medium', 'Large', 'Extra Large', 'Tiny', 'Huge'],
      },
      {
        success: (response) => {
          if (response.ok) {
            setSelectedOption(response.result)
          }
        },
      },
    )
  }, []) // Empty dependency array means this callback won't change

  useEffect(() => {
    async function onMount() {
      window.miniProgram?.call('setNavigationBar', {
        title: 'Picker',
        backgroundColor: '#0A0A0A',
        color: 'light',
      })
    }

    onMount()
  }, [])

  const code = `window.miniProgram?.call(
  'showPicker', {
    options: [
      'Small', 
      'Medium', 
      'Large', 
      'Extra Large', 
      'Tiny', 
      'Huge'
    ],
  },
  {
    success: (response) => {
      if (response.ok) {
        console.log(response.result)
      }
    },
  },
)`

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Picker</h3>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        {selectedOption ?? 'Open'}
      </button>

      <h3 className="text-white mt-8 font-bold">Code</h3>
      <HighlightedCode code={code} />
    </div>
  )
}
