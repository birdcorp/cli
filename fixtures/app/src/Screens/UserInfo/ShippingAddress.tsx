import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function ShippingAddress() {
  const code = `window.miniProgram?.call(
    'requestShippingAddress',
    {},
    {
      success(response) {
        if (response.ok) {
          console.log(response.shippingAddress)
        }
      }
    },
  )
}`

  function trigger() {
    window.miniProgram?.call('requestShippingAddress', null, {
      success(response: any) {
        if (response.ok) {
          console.log(response.shippingAddress)
          alert(JSON.stringify(response.shippingAddress))
        }
      },
      failure(err: any) {
        console.error(err)
      },
      completed() {},
    })
  }

  useEffect(() => {
    async function onMount() {
      window.miniProgram?.call(
        'setNavigationBar',
        {
          title: 'Shipping Address',
          backgroundColor: '#0A0A0A',
          color: 'light',
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

    onMount()
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black">
      <h3 className="text-white font-bold mb-3">Shipping Address</h3>

      <p className="text-gray-400">
        Opens a sheet to request user shipping address.
      </p>

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
