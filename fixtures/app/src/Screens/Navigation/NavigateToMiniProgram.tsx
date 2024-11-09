import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function NavigateToMiniProgram() {
  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Navigate to miniprogram',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 pb-10 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Navigate to a Miniprogram</h3>

      <p className="text-gray-300 mt-4">
        Opens a mini program, accepts payload params
      </p>

      <p className="mt-6 mb-2 text-gray-300">
        <span className="font-bold text-white">appId:</span> mini program app id
      </p>
      <p className="mb-2 text-gray-300">
        <span className="font-bold text-white">data:</span> key / value pair
      </p>
      <p className="text-gray-300">
        <span className="font-bold text-white">path:</span> (optional) if the
        entry point is not the root path can be supplied
      </p>

      <p className="mt-10 text-sm">This will resolve to</p>
      <HighlightedCode
        code={`[miniprogramURL]/
[path]
?key1=value1`}
      />

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold mb-8">Examples</h3>
      </div>

      <p className="mb-4">Search in another app for items</p>
      <button
        onClick={() => {
          window.miniProgram?.call('navigateToMiniProgram', {
            appID: '17eca250e2781caf',
            path: '/search',
            data: {
              query: 'shirt',
            },
          })
        }}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Search
      </button>

      <p className="my-3">Opens best buy app</p>

      <HighlightedCode
        code={`window.miniProgram?.call('navigateToMiniProgram', {
        appID: '17eca250e2781caf',
        path: '/search',
        data: {
          query: 'shirt',
        },
      })`}
      />

      <button
        onClick={() => {
          window.miniProgram?.call('navigateToMiniProgram', {
            appID: '17ebfa210da7a102',
            path: '/tracking',
            data: {
              trackingNumber: 'abc123',
            },
          })
        }}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded mt-8"
      >
        Track an order
      </button>
      <p className="my-3">For a given order / tracking id</p>
      <HighlightedCode
        code={`
window.miniProgram?.call("navigateToMiniProgram", {
    appID: "17ebfa210da7a102",
    path: "/tracking",
    data: {
        trackingNumber: "abc123"
    }
})
`}
      />
    </div>
  )
}
