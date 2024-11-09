import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function AppInfo() {
  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'App Info',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 mt-5">
      <h3 className="text-white font-bold mb-3">App</h3>

      <HighlightedCode code={`window.miniProgram`} />

      <HighlightedCode code={JSON.stringify(window.miniProgram, null, 2)} />

      <p className="text-gray-400 mb-8">App ID: {window.miniProgram?.appID}</p>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Properties</h3>
        <p className="text-gray-400">Javascript</p>
      </div>
    </div>
  )
}
