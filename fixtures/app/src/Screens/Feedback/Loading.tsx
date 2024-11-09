import React, { useEffect, useCallback } from 'react'
import { HighlightedCode } from '../../Components'

export default function Loading() {
  const trigger = useCallback(() => {
    window.miniProgram?.call('showLoading')

    setTimeout(() => {
      window.miniProgram?.call('hideLoading')
    }, 5000)
  }, [])

  useEffect(() => {
    async function onMount() {
      window.miniProgram?.call('setNavigationBar', {
        title: 'Loading',
        backgroundColor: '#0A0A0A',
        color: 'light',
      })
    }

    onMount()
  }, [])

  const code = `window.miniProgram?.call("showLoading")

setTimeout(() => {
  window.miniProgram?.call("hideLoading")
}, 5000)
`

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Loading</h3>

      <p className="text-gray-300 mb-8">
        Triggers an overlay and loading spinner
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Show
      </button>

      <h3 className="text-white mt-8 font-bold">Code</h3>
      <HighlightedCode code={code} />
    </div>
  )
}
