import React from 'react'
import { HighlightedCode } from '../../Components'

export default function GetClipboard() {
  const code = `window.miniProgram?.call("getClipboard", {
    success: (response) => {
        if(response.ok){
            console.log(response.clipboard);
        }
    },
    failure: (err) => {
        console.error(err)
    },
    completed: () => {

    }
})`

  function trigger() {
    window.miniProgram?.call(
      'getClipboard',
      {},
      {
        success(response) {
          if (response.ok) {
            window.miniProgram?.call('showAlert', {
              title: 'Copied to clipboard',
              message: response.text,
            })
          } else {
            window.miniProgram?.call('showToast', {
              type: 'error',
              content: 'Declined',
              duration: 3000,
            })
          }
        }, // text
        failure(err) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5">
      <h3 className="text-white font-bold mb-3">Get Clipboard</h3>

      <p className="text-gray-400 mb-8">
        Requests permission to get clipboard contents
      </p>

      <div className="grid grid-cols-3 gap-4">
        <button
          onClick={trigger}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Run
        </button>
      </div>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
