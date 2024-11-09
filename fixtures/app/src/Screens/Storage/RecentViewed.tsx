import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

interface CartItem {
  id: string
  label: string
  value: string
  identifier?: string
  thumbnailURL?: string
  path?: string
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

export default function RecentViewed() {
  const [response, setResponse] = useState<CartItem[]>([])

  async function add() {
    const price = (Math.random() * (100 - 1) + 1).toFixed(2)

    window.miniProgram?.call(
      'addToRecentViewed',
      {
        label: getRandomProductName(),
        value: price,
        identifier: generateUUID(),
        thumbnailURL:
          'https://cdn.dummyjson.com/products/images/groceries/Green%20Bell%20Pepper/thumbnail.png',
      },
      {
        success(response: any) {
          if (response.ok) {
            getRecentViewed()
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  async function getRecentViewed() {
    window.miniProgram?.call(
      'getRecentViewed',
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
            getRecentViewed()
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
            getRecentViewed()
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

    getRecentViewed()
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100 pb-10">
      <h3 className="text-white font-bold mb-3">Recent Viewed</h3>

      <p className="text-gray-300 mb-4">
        Provides a headless, persistent list of recently viewed items, stored
        locally on the device.
      </p>

      <HighlightedCode
        code={`interface RecentViewedProduct {
  id: string
  label: string
  value: string
  identifier?: string
}          
`}
      />

      <h3 className="text-white mt-8 mb-3 font-bold">Demo</h3>

      <div className="overflow-x-auto py-4 scrollbar-hide">
        <ul className="flex space-x-4">
          {response.map((item) => (
            <li
              key={item.id}
              className="flex-shrink-0 w-32 text-center rounded-lg shadow"
            >
              <img
                src={item.thumbnailURL}
                alt={item.label}
                className="w-full h-32 object-cover rounded-t-lg"
              />
              <span className="block py-2 font-medium text-neutral-200">
                {item.label}
              </span>
            </li>
          ))}
        </ul>
      </div>

      <button
        onClick={add}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Add item
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Add to Recent Viewed</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
    'addToRecentViewed',
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
        <h3 className="text-white font-bold">Get Recent Viewed</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
    'getRecentViewed',
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
      'removeFromRecentViewed',
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
