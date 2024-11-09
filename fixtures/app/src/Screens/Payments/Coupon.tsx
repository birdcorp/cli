import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

interface DiscountCode {
  code: string // The discount code (e.g., "SUMMER2024")
  type: string // The type of discount (e.g., "fixed_amount")
  discountAmount: number // The amount of discount (e.g., 100)
  expiryDate: string // The expiry date in ISO 8601 format (e.g., "2025-01-01T00:00:00.000Z")
  amountIssued: number // The total amount issued for this discount (e.g., 100)
  remaining: number // The remaining amount of discount available (e.g., 100)
}

export default function Coupon() {
  const [coupons, setCoupons] = useState<DiscountCode[]>([])

  useEffect(() => {
    fetch(
      'https://cqs9msvige.execute-api.us-east-1.amazonaws.com/dev/developer/coupons',
    )
      .then((response) => response.json())
      .then((response) => setCoupons(response.data))
  }, [])

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
      <h3 className="text-white font-bold mb-3">Coupon Code</h3>

      <div className="flex flex-wrap justify-start">
        {coupons.map((coupon) => (
          <div
            className="py-1 px-2 bg-white bg-opacity-10 text-white font-bold rounded-lg text-center mr-2"
            style={{ maxWidth: '200px' }} // Optional max width
          >
            {coupon.code}
          </div>
        ))}
      </div>

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

      <h2 className="text-2xl font-bold text-white mt-8 mb-4">
        Server API Examples
      </h2>

      <h3 className="text-lg font-semibold text-white mb-2">
        Create coupon_code (fixed_amount $10)
      </h3>
      <HighlightedCode
        code={`curl -X POST
  https://api.birdwallet.xyz/v2/coupon_codes
  -H "Accept: application/json"
  -H "X-API-KEY: your_api_key_here"
  -H "Content-Type: application/json"
  -d '{
    "code": "SUMMER2024",
    "type": "fixed_amount",
    "discount_amount": 100,
    "expiry_date": "2025-01-01T00:00:00Z",
    "amount_issued": 100
  }'`}
      />

      <h3 className="text-lg font-semibold text-white mb-2">
        Create coupon_code (percentage 10%)
      </h3>

      <HighlightedCode
        code={`curl -X POST
  https://api.birdwallet.xyz/v2/coupon_codes
  -H "Accept: application/json"
  -H "X-API-KEY: your_api_key_here"
  -H "Content-Type: application/json"
  -d '{
    "code": "SUMMER2024",
    "type": "percentage",
    "discount_percent": 10,
    "expiry_date": "2025-01-01T00:00:00Z",
    "amount_issued": 100
  }'`}
      />
    </div>
  )
}
