import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

export default function OnetimeCode() {
  const [code, setCode] = useState('')

  function trigger() {
    ;(window as any).miniProgram.call(
      'showOnetimeCode',
      {
        title: 'Enter Code',
        description: 'Please enter the 6-digit code.',
        numberOfDigits: 6,
      },
      {
        success(response: any) {
          if (response.ok) {
            setCode(response.result)
          }
        },
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Onetime Code',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  function requestPhoneExample() {
    ;(window as any).miniProgram.call(
      'requestPhone',
      {},
      {
        success: (response: any) => {
          if (response.ok) {
            ;(window as any).miniProgram.call(
              'showOnetimeCode',
              {
                title: 'Enter Code',
                description: 'Please enter the 6-digit code.',
                numberOfDigits: 6,
              },
              {
                success(response: any) {
                  if (response.ok) {
                    setCode(response.result)
                  }
                },
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
      },
    )
  }

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 mt-5">
      <h3 className="text-white font-bold mb-3">Onetime Code</h3>

      <p className="text-gray-400 mb-8">Opens onetime code bottom sheet.</p>

      <h3 className="text-white font-bold mb-3 mt-10">Example</h3>

      <p className="text-gray-400 text-xl mb-8">Code: {code}</p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Run
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram.call(
      'showOnetimeCode',
      {
        title: 'Enter Code',
        description: 'Please enter the 6-digit code.',
        numberOfDigits: 6,
      },
      {
        success(response: any) {
          if (response.ok) {
            setCode(response.result)
          }
        },
      },
    )`}
      />

      <h3 className="text-white font-bold mb-3 mt-10">Example</h3>
      <p className="text-gray-400 mb-8">Request phone number first.</p>

      <button
        onClick={requestPhoneExample}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Run
      </button>
    </div>
  )
}
