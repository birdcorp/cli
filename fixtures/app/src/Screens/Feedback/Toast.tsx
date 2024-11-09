import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

function capitalizeFirstLetter(str: string) {
  return str.charAt(0).toUpperCase() + str.slice(1)
}

export default function Toast() {
  const code = `window.miniProgram?.call("showToast", {
    type: 'success',
    content: 'Success',
    duration: 3000
})`

  function trigger(type: 'success' | 'loading' | 'error' | 'failure') {
    window.miniProgram?.call(
      'showToast',
      {
        type,
        content: capitalizeFirstLetter(type),
        duration: 1000,
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
    async function onMount() {
      window.miniProgram?.call(
        'setNavigationBar',
        {
          title: 'Toast',
          backgroundColor: '#111',
          color: 'dark',
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
      <h3 className="text-white font-bold mb-3">Toast</h3>

      <p className="text-gray-300">This triggers a toast message</p>

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <div className="grid grid-cols-3 gap-4">
        <button
          onClick={() => trigger('success')}
          className="bg-green-800 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
        >
          Success
        </button>

        <button
          onClick={() => trigger('loading')}
          className="bg-neutral-800 hover:bg-neutral-700 text-white font-bold py-2 px-4 rounded"
        >
          Loading
        </button>

        <button
          onClick={() => trigger('error')}
          className="bg-red-800 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
        >
          Error
        </button>
      </div>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />

      <div>
        <table className="table-auto text-neutral-300 w-full mt-8">
          <thead>
            <tr>
              <th>Property</th>
              <th>Type</th>
              <th>Value</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>type</td>
              <td>enum</td>
              <td>success, loading, error</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  )
}
