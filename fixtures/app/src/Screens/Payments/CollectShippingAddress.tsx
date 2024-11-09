import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'
import { ShippingAddressForm } from '../../Components'

export default function CollectShippingAddress() {
  async function trigger() {
    window.miniProgram?.call('showToast', {
      type: 'loading',
      content: 'Loading...',
      duration: 3000,
    })

    const { id } = await fetch(
      'https://cqs9msvige.execute-api.us-east-1.amazonaws.com/dev/developer/orders?orderType=basic',
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({}),
      },
    ).then((response) => response.json())

    window.miniProgram?.call(
      'showPaySheet',
      {
        id: id,
      },
      {
        success(response: any) {
          window.miniProgram?.call('hideToast')

          if (response.ok) {
            setTimeout(() => {
              ;(window as any).confetti({
                particleCount: 100,
                spread: 70,
                origin: { y: 0.6 },
              })
            }, 2000)
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
    window.miniProgram?.call('setNavigationBar', {
      title: 'Order Payment',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h2 className="text-xl font-bold mb-4">Collect Shipping Address</h2>

      <p className="text-gray-300">Opens checkout payment experience.</p>

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <section className="mt-10">
        <h2 className="text-xl font-bold mb-4">Shipping Address</h2>

        {/*<ShippingAddressForm />*/}
        <button className="bg-neutral-800 hover:bg-neutral-700 text-white font-bold py-2 px-4 rounded">
          Select Shipping Address
        </button>
      </section>

      <section className="mt-10">
        <h2 className="text-xl font-bold mb-4">Order Summary</h2>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <p className="text-lg text-gray-500">Midnight Mocha Madness</p>
          </div>
          <div className="text-lg text-right text-gray-500">$10.00</div>
          <div>
            <p className="text-lg text-gray-500">Shipping</p>
          </div>
          <div className="text-lg text-right text-gray-500">Pending</div>
          <div>
            <p className="text-lg text-gray-500">Sales Tax</p>
          </div>
          <div className="text-lg text-right text-gray-500">Pending</div>

          <div>
            <p className="text-lg font-semibold">Sub-Total</p>
          </div>
          <div className="text-lg font-semibold text-right">$0.01</div>
        </div>
      </section>

      <button
        onClick={trigger}
        className="w-full bg-white hover:bg-neutral-200 text-black text-lg font-bold py-3 px-4 rounded-lg mt-10"
      >
        Pay
      </button>

      <div className="flex flex-row justify-between mt-20">
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
