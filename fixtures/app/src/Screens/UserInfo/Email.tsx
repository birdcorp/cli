import React from 'react'
import { HighlightedCode } from '../../Components'

export default function Email() {
  const code = `window.miniProgram?.call("requestEmail", {
    success: (response) => {
      if (response.ok) {
        window.miniProgram?.call("showAlert", {
          title: "Email Address",
          message: response.emailAddress,
        });
      }
    }
  })`

  function trigger() {
    ;(window as any).miniProgram.call(
      'requestEmail',
      {},
      {
        success(response: any) {
          if (response.ok) {
            console.log(response.emailAddress)
            ;(window as any).miniProgram.call(
              'showAlert',
              {
                title: 'Email Address',
                message: response.emailAddress,
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
          } else {
            ;(window as any).miniProgram?.call('showToast', {
              type: 'error',
              content: 'Declined',
              duration: 3000,
            })
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
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black">
      <h3 className="text-white font-bold mb-3">Email</h3>

      <p className="text-gray-400">Opens a sheet to request user email.</p>

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Request
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
