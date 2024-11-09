import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function OpenLocation() {
  const code = `window.miniProgram?.call("openLocation", {
    longitude: '-73.935242',
    latitude: '40.730610',
    name: 'Empire State Building',
    address: '20 W 34th St, New York, NY 10118, USA',
})`

  function trigger() {
    window.miniProgram?.call(
      'openLocation',
      {
        longitude: '-73.935242',
        latitude: '40.730610',
        name: 'Empire State Building',
        address: '20 W 34th St, New York, NY 10118, USA',
      },
      {
        success(response: any) {},
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Open Location',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Open Location</h3>

      <p className="text-gray-300 mb-8">
        Opens the specified location on a Map fullscreen sheet
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Open
      </button>

      <h3 className="text-white mt-8 font-bold">Code</h3>
      <HighlightedCode code={code} />
    </div>
  )
}
