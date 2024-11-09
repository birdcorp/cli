import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function OpenSafariView() {
  const code = `window.miniProgram?.call('showSafariView', {
      url: 'https://example.com/',
    })`

  function run() {
    window.miniProgram?.call('showSafariView', {
      url: 'https://example.com/',
    })
  }

  useEffect(() => {
    window.miniProgram.call('setNavigationBar', {
      title: 'showSafariView',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">showSafariView</h3>

      <p className="text-gray-300 mb-8">Opens safari pop up for given url</p>

      <button
        onClick={run}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Open
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
