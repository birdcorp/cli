import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'
import { Trash } from 'lucide-react'

interface CartItem {
  id: string
  label: string
  value: string
  identifier?: string
}

const productNames = [
  'Honey Jar',
  'Dog Food',
  'Chanel Coco Noir Eau De',
  'Fish Steak',
  'Organic Almond Butter',
  'Wireless Mouse',
  'Bluetooth Headphones',
  'Smartphone Case',
  'Leather Wallet',
  'Portable Charger',
  'Yoga Mat',
  'Stainless Steel Water Bottle',
  'Electric Toothbrush',
  'Coffee Maker',
  'Fitness Tracker',
]

function generateUUID() {
  // Generate a UUID v4 (based on random values)
  const randomValues = crypto.getRandomValues(new Uint8Array(16))

  // Format as a UUID string
  const uuid =
    (
      (randomValues[0] << 24) |
      (randomValues[1] << 16) |
      (randomValues[2] << 8) |
      randomValues[3]
    )
      .toString(16)
      .padStart(8, '0') +
    '-' +
    ((((randomValues[4] & 0x0f) | 0x40) << 8) | randomValues[5])
      .toString(16)
      .padStart(4, '0') +
    '-' +
    ((((randomValues[6] & 0x3f) | 0x80) << 8) | randomValues[7])
      .toString(16)
      .padStart(4, '0') +
    '-' +
    ((randomValues[8] << 8) | randomValues[9]).toString(16).padStart(4, '0') +
    '-' +
    (
      (randomValues[10] << 24) |
      (randomValues[11] << 16) |
      (randomValues[12] << 8) |
      randomValues[13]
    )
      .toString(16)
      .padStart(12, '0')

  return uuid
}

function getRandomProductName() {
  const randomIndex = Math.floor(Math.random() * productNames.length)
  return productNames[randomIndex]
}

export default function Cart() {
  const [response, setResponse] = useState<CartItem[]>([])

  async function addToCart() {
    const price = (Math.random() * (100 - 1) + 1).toFixed(2)

    window.miniProgram?.call(
      'addToCart',
      {
        label: getRandomProductName(),
        value: price,
        identifier: generateUUID(),
      },
      {
        success(response: any) {
          if (response.ok) {
            getCart()
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  async function getCart() {
    window.miniProgram?.call(
      'getCart',
      {},
      {
        success(response: any) {
          if (response.ok) {
            setResponse(response.data)
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  async function removeFromCart(id: string) {
    window.miniProgram?.call(
      'removeFromCart',
      {
        id,
      },
      {
        success(response: any) {
          if (response.ok) {
            window.miniProgram?.call('vibrate', {
              type: 'medium',
            })
            getCart()
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  function clearCart() {
    window.miniProgram?.call(
      'clearCart',
      {},
      {
        success(response: any) {
          if (response.ok) {
            // confirm button was pressed
            //setResponse(JSON.stringify(response, null, 4))
            getCart()
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
      title: 'Cart',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })

    getCart()
  }, [])

  const total = response.reduce((total, item) => {
    // Remove the dollar sign and convert the string to a number
    const numericValue = parseFloat(item.value)
    return total + (isNaN(numericValue) ? 0 : numericValue)
  }, 0)

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100 pb-10">
      <h3 className="text-white font-bold mb-3">Cart</h3>

      <p className="text-gray-300 mb-4">
        Provides a headless, persistent shopping cart stored locally on the
        device.
      </p>

      <HighlightedCode
        code={`interface CartItem {
  id: string
  label: string
  value: string
  identifier?: string
}          
`}
      />

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <ul className="space-y-1">
        {response.map((item) => (
          <li
            key={item.id}
            className="flex justify-between px-3 py-4 bg-neutral-900 rounded-lg shadow"
          >
            <span className="font-medium text-neutral-200">{item.label}</span>
            <span className="text-neutral-200">
              ${item.value}{' '}
              <span
                className="ml-3"
                onClick={() => {
                  removeFromCart(item.id)
                }}
              >
                <Trash size={16} />
              </span>
            </span>
          </li>
        ))}
      </ul>

      <div className="flex justify-between text-white mt-4 text-lg font-semibold">
        <p>Total</p>
        <p>${total.toFixed(2)}</p>
      </div>

      <button
        onClick={addToCart}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded mt-8"
      >
        Add Item
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Add to Cart</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
    'addToCart',
    {
        label: 'label1',
        value: '9.99',
        identifier: 'abc123',
    },
    {
    success(response) {
        if (response.ok) {

        }
    },
    failure(err) {
        console.error(err)
    },
    completed() {},
    
    },
)`}
      />

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Get Cart</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
    'getCart',
    {},
    {
      success(response) {
        if (response.ok) {
          // response.result as CartItem[]
        }
    },
    failure(err) {
        console.error(err)
    },
    completed() {},
    
    },
)`}
      />

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Remove item</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
      'removeFromCart',
      { id },
      {
        success(response) {
          if (response.ok) {
          }
        },
        failure(err) {
          console.error(err)
        },
        completed() {},
      },
)`}
      />
    </div>
  )
}
