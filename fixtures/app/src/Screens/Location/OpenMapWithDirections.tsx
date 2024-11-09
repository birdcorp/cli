import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function OpenMapWithDirections() {
  const code = `window.miniProgram.call('getLocation', null, {
    success(response) {
      if (response.ok) {
        window.miniProgram?.call('openMapWithDirections', {
          longitude1: response.longitude,
          latitude1: response.latitude,
          longitude2: -122.3937,
          latitude2: 37.7955,
        });
      }
    },
  });`

  function trigger() {
    window.miniProgram.call('getLocation', null, {
      success(response) {
        if (response.ok) {
          window.miniProgram?.call('openMapWithDirections', {
            longitude1: response.longitude,
            latitude1: response.latitude,
            longitude2: -122.3937,
            latitude2: 37.7955,
          })
        }
      },
    })
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Open Map with Directions',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Open Map with Directions</h3>

      <p className="text-gray-300 mb-8">
        Opens Map with Directions to the specified location
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
