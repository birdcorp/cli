import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function NavigateBack() {
  const code = `window.miniProgram?.call("navigateBack", {
    depth: 1
})`

  function trigger() {
    window.miniProgram?.call(
      'navigateBack',
      {
        depth: 1,
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

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Navigate Back',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Navigate Back</h3>

      <p className="text-gray-300 mb-8">
        This triggers a native confirm with title and message properties
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Back
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
