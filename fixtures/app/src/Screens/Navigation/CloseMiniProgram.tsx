import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function CloseMiniProgram() {
  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Close Mini Program',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">closeMiniProgram</h3>

      <p className="text-gray-300 mb-8">Closes the mini program</p>

      <button
        onClick={() => {
          window.miniProgram?.call('closeMiniProgram')
        }}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Close App
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={`window.miniProgram?.call("closeMiniProgram")`} />
    </div>
  )
}
