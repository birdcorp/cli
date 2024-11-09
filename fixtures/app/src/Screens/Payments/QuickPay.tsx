import React from 'react'
import { HighlightedCode } from '../../Components'

export default function QuickPay() {
  const code = `miniProgram.call("requestPayment", {
    id: "abc123",
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

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Quick Pay</h3>

      <p className="text-gray-300">Requests payment.</p>

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <button className="bg-blue-600 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded">
        Pay
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
