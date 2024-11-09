import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function Confirm() {
  const code = `window.miniProgram?.call(
    'showConfirm',
    {
        title: 'Remove Item?',
        content:
          'Are you sure you want to remove this item from your cart? This action cannot be undone.',
        confirmButtonText: 'Remove',
        cancelButtonText: 'Keep',
    },
    {
      success(response) {
        if (response.ok) {
          window.miniProgram?.call('showAlert', {
            title: 'Confirm clicked',
            message: 'Item removed',
          })
        }
      }
    },
  )`

  function trigger() {
    window.miniProgram?.call(
      'showConfirm',
      {
        title: 'Remove Item?',
        content:
          'Are you sure you want to remove this item from your cart? This action cannot be undone.',
        confirmButtonText: 'Remove',
        cancelButtonText: 'Keep',
      },
      {
        success(response: any) {
          if (response.ok) {
            window.miniProgram?.call('showAlert', {
              title: 'Confirm clicked',
              message: 'Item removed',
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

  useEffect(() => {
    async function onMount() {
      window.miniProgram?.call(
        'setNavigationBar',
        {
          title: 'Confirm',
          backgroundColor: '#0A0A0A',
          color: 'light',
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
    onMount()
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Confirm</h3>

      <p className="text-gray-300 mb-8">
        This triggers a native confirm with title and message properties
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Confirm
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
