import React, { useState, useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function KeyChain() {
  const [value, setValue] = useState('5up3rSecr3T!!')
  const [result, setResult] = useState('')

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Keychain',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  const codeSet = `window.miniProgram?.call("setKeychain", {
    key: 'secret',
    value: '${value}'
})
`

  const codeGet = `window.miniProgram?.call("getKeychain", {
    key: 'secret'
}, {
    success(response) {
        if (response.ok) {
            window.miniProgram?.call('showAlert', {
                title: 'Result',
                message: response.value,
            })
        }
    },
    failure(err: any) {
        console.error(err)
    },
    completed() {},
})`

  function set() {
    ;(window as any).miniProgram.call(
      'setKeychain',
      {
        key: 'secret',
        value: value,
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

  function get() {
    ;(window as any).miniProgram.call(
      'getKeychain',
      {
        key: 'secret',
      },
      {
        success(response: any) {
          if (response.ok) {
            setResult(response.value)
            window.miniProgram?.call('showAlert', {
              title: 'Result',
              message: response.value,
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
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100 pb-20">
      <h3 className="text-white font-bold mb-3">Get/Set to Key Chain</h3>
      <p className="text-gray-400 mb-8">
        Allows get / set to key chain for secure persistent storage
      </p>
      <input
        value={value}
        onChange={(e) => setValue(e.target?.value)}
        className="bg-neutral-800 text-white p-2 rounded"
      />
      <p className="my-3">{value}</p>
      <div className="grid grid-cols-3 gap-4">
        <button
          onClick={set}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Set
        </button>
      </div>
      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>
      <HighlightedCode code={codeSet} />
      <div className="grid grid-cols-3 gap-4 mt-20">
        <button
          onClick={get}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Get
        </button>
      </div>

      <p className="my-3">Result: {result}</p>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>
      <HighlightedCode code={codeGet} />
    </div>
  )
}
