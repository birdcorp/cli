import React from 'react'
import { HighlightedCode } from '../../Components'

export default function SetClipboard() {
  const code = `window.miniProgram?.call("setClipboard", {
    text: "some text here",
    success: (response) => {
        if(response.ok){

        }
    }
})`

  function trigger() {
    ;(window as any).miniProgram.call(
      'setClipboard',
      {
        text: 'some text here',
      },
      {
        success(response: any) {
          if (response.ok) {
            console.log(response.clipboard)
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
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Set Clipboard</h3>

      <p className="text-gray-300 mb-8">
        Requests permission to set clipboard contents
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
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
