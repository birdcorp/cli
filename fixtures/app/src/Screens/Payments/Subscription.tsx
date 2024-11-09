import React, { useEffect } from 'react'
import { addDays } from 'date-fns'
import { HighlightedCode } from '../../Components'

import { OrderPayload } from 'bird-nodejs-sdk'

export default function Subscription() {
  const code = `window.miniProgram?.call('showLoading')

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
)`
  //

  useEffect(() => {
    /*
     * Wake up the lambda function to prevent cold start
     */
    fetch(
      'https://cqs9msvige.execute-api.us-east-1.amazonaws.com/dev/health',
    ).catch(console.error)
  }, [])

  async function trigger() {
    window.miniProgram?.call('showLoading')

    const { id } = await fetch(
      'https://cqs9msvige.execute-api.us-east-1.amazonaws.com/dev/developer/orders?orderType=recurring',
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
        id,
      },
      {
        success(response: any) {
          window.miniProgram?.call('hideLoading')

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
    async function onMount() {
      window.miniProgram?.call(
        'setNavigationBar',
        {
          title: 'Subscription',
          backgroundColor: '#0A0A0A',
          color: 'light',
        },
        {
          success(response: any) {
            if (response.ok) {
              // confirm button was pressed
            }
          },
        },
      )
    }

    onMount()
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Subscription</h3>

      <p className="text-gray-300 mb-6">Recurring Billing</p>

      <button
        onClick={trigger}
        className="bg-white text-black font-bold py-2 px-4 rounded"
      >
        Subscribe
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
