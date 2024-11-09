import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

export default function ListOrders() {
  const [orders, setOrders] = useState<any[]>([])

  async function trigger() {
    window.miniProgram?.call('showToast', {
      type: 'loading',
      content: 'Loading...',
      duration: 3000,
    })
    ;(window as any).miniProgram?.call(
      'listOrders',
      {},
      {
        success(response: any) {
          window.miniProgram?.call('hideToast')

          if (response.ok) {
            setOrders(response.data)
          }
        },
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'List Orders',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Checkout</h3>

      <p className="text-gray-300">Opens checkout payment experience.</p>

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <button
        onClick={trigger}
        className="bg-neutral-800 hover:bg-neutral-700 text-white font-bold py-2 px-4 rounded"
      >
        Fetch Orders
      </button>

      <pre>{JSON.stringify(orders, null, 2)}</pre>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call('showLoading')

const res = await fetch(
  '/api/order',
)

const { id } = await res.json()

window.miniProgram?.call(
  'showPaySheet',
  {
    id,
  },
  {
    success(response) {
      if (response.ok) {

      }
    },
    failure(err) {
      console.error(err)
    },
    completed() {
      window.miniProgram?.call('hideLoading')
    },
  },
)`}
      />
    </div>
  )
}
