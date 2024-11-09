import React from 'react'
import { HighlightedCode } from '../../Components'

export default function PhoneCall() {
  const code = `window.miniProgram?.call("makePhoneCall", {
    phoneNumber: '5109256717',
})`

  function trigger() {
    ;(window as any).miniProgram.call(
      'makePhoneCall',
      {
        phoneNumber: '5109271717',
      },
      {
        success(response: any) {
          if (response.ok) {
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
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 text-gray-100">
      <h3 className="text-white font-bold mb-3">Make Phone Call</h3>

      <p className="text-gray-400 mb-8">Makes phone call to given number.</p>

      <div className="grid grid-cols-3 gap-4">
        <button
          onClick={trigger}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Call Number
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
