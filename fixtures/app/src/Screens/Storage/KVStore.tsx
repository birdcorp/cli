import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

export default function KVStore() {
  const [value, setValue] = useState('')

  const KEY = 'key1'

  async function set() {
    window.miniProgram?.call(
      'setStorage',
      {
        key: KEY,
        value: JSON.stringify({
          id: 'abc123',
          value: `${Math.random()}`,
        }),
      },
      {
        success(response: any) {
          if (!response.ok) {
            window.miniProgram?.call('showToast', {
              type: 'failure',
              content: 'Failed to set key ' + KEY,
              duration: 3000,
            })
            return
          }

          get()
        },
      },
    )
  }

  async function get() {
    window.miniProgram?.call(
      'getStorage',
      {
        key: KEY,
      },
      {
        success(response: any) {
          if (response.ok) {
            setValue(response.result)
          }
        },
      },
    )
  }

  function clear() {
    window.miniProgram?.call(
      'clearStorage',
      {
        key: KEY,
      },
      {
        success(response: any) {
          if (response.ok) {
            get()
          }
        },
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'KV Store',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
    get()
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100 pb-10">
      <h3 className="text-white font-bold mb-3">Storage</h3>

      <p className="text-gray-300 mb-4">Alternative to window.localStorage</p>
      <p className="text-gray-300 mb-4">
        Provides a persistent key-value storage solution that synchronizes
        across devices to iCloud.
      </p>

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <h3 className="text-white text-2xl mb-10">{value}</h3>

      <button
        onClick={set}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Set
      </button>

      <button
        onClick={get}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Get
      </button>

      <button
        onClick={clear}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Clear
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Set</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
      'setStorage',
      {
        key: KEY,
        value: JSON.stringify({
          data: {
            key: "key1",
            value: "value1"
          }
        }),
      },
      {
        success(response: any) {
          if (response.ok) {
            get()
            return
          }
          window.miniProgram?.call('showToast', {
            type: 'failure',
            content: 'Failed to set key ' + KEY,
            duration: 3000,
          })
        },
      },
    )`}
      />

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Get</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
      'getStorage',
      {
        key: "key1",
      },
      {
        success(response: any) {
          if (response.ok) {
            console.log(response.result)
          }
        },
      },
    )`}
      />

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Clear</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
      'clearStorage',
      {
        key: KEY,
      },
      {
        success(response: any) {
          if (response.ok) {
            get()
          }
        },
      },
    )`}
      />
    </div>
  )
}
