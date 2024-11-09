import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function NavigateTo() {
  const code = `window.miniProgram?.call("navigateTo", {
    url:  \`${window.origin}/page2\`
})`

  function trigger() {
    window.miniProgram?.call('navigateTo', {
      url: `${window.origin}/navigation/navigate-back`,
    })
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Navigate To',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Navigate To</h3>

      <p className="text-gray-300 mb-8">
        Pushes view into the Navigation Stack
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Go to Next Page
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
