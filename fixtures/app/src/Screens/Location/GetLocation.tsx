import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function GetLocation() {
  const code = `window.miniProgram?.call(
    'getLocation',
    {},
    {
      success(response) {
        if (response.ok) {
          miniProgram.call('showAlert', {
            title: 'Location',
            message: "longitude:" + response.longitude + " latitude:" + response.latitude,
          })
        }
      }
    },
  )`

  function trigger() {
    window.miniProgram.call(
      'getLocation',
      {},
      {
        success(response: any) {
          if (response.ok) {
            ;(window as any).miniProgram.call('showAlert', {
              title: 'Location',
              message: `longitude:${response.longitude} latitude:${
                response.latitude
              }, type ${typeof response.longitude}`,
            })
          }
        },
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Get Location',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Get Location</h3>

      <p className="text-gray-300 mb-8">
        This triggers a native alert with title and message properties
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Call
      </button>

      <h3 className="text-white mt-8 font-bold">Code</h3>
      <HighlightedCode code={code} />
    </div>
  )
}
