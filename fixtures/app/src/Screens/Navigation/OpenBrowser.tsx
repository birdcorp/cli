import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function OpenBrowser() {
  const code = `window.miniProgram?.call("openBrowser", {
    url: 'https://demo.birdwallet.xyz/'
})`

  function trigger() {
    window.miniProgram?.call('openBrowser', {
      url: 'https://demo.birdwallet.xyz/',
    })
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Open Browser',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 dark:bg-black dark:text-gray-100">
      <h3 className="text-white font-bold mb-3">Open Browser</h3>

      <p className="text-gray-400 mb-5 dark:text-gray-300">
        Opens URL in Safari
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Open
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400 dark:text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
