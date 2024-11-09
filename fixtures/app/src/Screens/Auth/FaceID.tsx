import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function FaceID() {
  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Face ID',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  const code = `window.miniProgram?.call("requestBioAuth", {
    success: (response) => {
        if(response.ok){

        }
    },
    failure: (err) => {
        console.error(err)
    },
    completed: () => {

    }
})`

  function trigger() {
    window.miniProgram.call(
      'requestBioAuth',
      {},
      {
        success(response: any) {
          if (response.ok) {
            window.miniProgram.call(
              'showAlert',
              {
                title: 'Face ID Auth OK',
                message: 'Success',
              },
              {
                success(response: any) {
                  if (response.ok) {
                    // confirm button was pressed
                  }
                },
                failure(err: any) {
                  console.error(err)
                },
                completed() {},
              },
            )
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 mt-5">
      <h3 className="text-white font-bold mb-3">Face ID</h3>

      <p className="text-gray-400 mb-8">Request face id authorization.</p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Call
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
