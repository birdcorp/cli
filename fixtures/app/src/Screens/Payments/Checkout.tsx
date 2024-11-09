import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function Checkout() {
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
              window.miniProgram?.call('navigateTo', {
                url: '/payment/thankyou',
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
    /*
     * Wake up the lambda function to prevent cold start
     */
    fetch(
      'https://cqs9msvige.execute-api.us-east-1.amazonaws.com/dev/health',
    ).catch(console.error)
  }, [])

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Order Payment',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  // Hardcoded data for the order summary
  const total = {
    value: '10.00',
    currency: 'USD',
  }

  const line_items = [
    {
      label: 'iPhone',
      type: 'item',
      value: '10.00',
      sku_id: null,
      thumbnail_url:
        'https://d2e6ccujb3mkqf.cloudfront.net/ccb63e05-7066-48a1-8e02-353a9a73187f-1_e84caaa9-80fb-4e28-b27e-be7cb91dae90.jpg',
      status: 'final',
      payment_timing: null,
      recurring_payment_start_date: null,
    },
    {
      label: 'Sales Tax',
      type: 'tax',
      value: '0.00',
      sku_id: null,
      thumbnail_url: null,
      status: 'pending',
      payment_timing: null,
      recurring_payment_start_date: null,
    },
    {
      label: 'Shipping',
      type: 'shipping',
      value: '0.00',
      sku_id: null,
      thumbnail_url: null,
      status: 'pending',
      payment_timing: null,
      recurring_payment_start_date: null,
    },
  ]

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Checkout</h3>

      <p className="text-gray-300">Opens checkout payment experience.</p>

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>
      <div className="w-full shadow-md rounded-lg">
        <h2 className="text-2xl font-bold text-gray-200 mb-4">Order Summary</h2>

        {/* Total Section */}
        <div className="border-b border-gray-900 pb-4 mb-4">
          <h3 className="text-lg font-semibold text-gray-400">Total</h3>
          <p className="text-xl font-bold text-gray-100">$ {total.value}</p>
        </div>

        {/* Line Items */}
        <ul className="space-y-4">
          {line_items.map((item, index) => (
            <li key={index} className="flex items-center justify-between pb-3">
              {/* Thumbnail */}
              {item.thumbnail_url && (
                <img
                  src={item.thumbnail_url}
                  alt={item.label}
                  className="w-16 h-16 rounded-md object-cover mr-4"
                />
              )}

              {/* Item Details */}
              <div className="flex-1">
                <p className="text-gray-400 font-semibold">{item.label}</p>
                <p className="text-gray-700 text-sm capitalize">{item.type}</p>
              </div>

              {/* Item Value */}
              <p className="text-lg font-bold text-gray-200">
                {item.status === 'pending' ? (
                  <span className="text-gray-400 text-xs">Pending</span>
                ) : (
                  <span>${item.value}</span>
                )}
              </p>
            </li>
          ))}
        </ul>
      </div>

      <div className="flex flex-row mt-4 justify-end">
        <button
          onClick={trigger}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Checkout
        </button>
      </div>

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
